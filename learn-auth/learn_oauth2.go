package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// { github_id: user_id }
var githubConnections map[string]string

// {"data":{"viewer":{"id":"...","name":"Hafidz Mahrus"}}}
type githubResponse struct {
	Data struct {
		Viewer struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"viewer"`
	} `json:"data"`
}

var githubConfig = &oauth2.Config{
	ClientID:     "760ba0cff79d6a9bcc7b",
	ClientSecret: "8a726ae925c716b2520fec3cff329edea5652b06",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/oauth2/github", githubHandler)
	http.HandleFunc("/oauth2/receive", receiveHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Test</title>
	</head>
	<body>
		<form action="/oauth2/github" method="post">
			<input type="submit" value="Login with Github">
		</form>
	</body>
	</html>
	`)
}

func githubHandler(w http.ResponseWriter, r *http.Request) {
	redirectUrl := githubConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
}

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:3000/oauth2/receive?code=251fb64d4da769f52806&state=0000
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "state is incorrect", http.StatusBadRequest)
		return
	}

	// func (c *Config) Exchange(ctx context.Context, code string, opts ...AuthCodeOption) (*Token, error)
	token, err := githubConfig.Exchange(r.Context(), code)

	if err != nil {
		http.Error(w, "count exchange the token", http.StatusInternalServerError)
		return
	}

	// func (c *Config) TokenSource(ctx context.Context, t *Token) TokenSource
	tokenSrc := githubConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), tokenSrc)

	requestBody := strings.NewReader(`{ "query": "query { viewer { id name } }" }`)
	res, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	defer res.Body.Close()

	if err != nil {
		http.Error(w, "Couldn't get user", http.StatusInternalServerError)
		return
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Couldn't read response", http.StatusInternalServerError)
		return
	}

	log.Println(string(bs))

	var gRes githubResponse

	err = json.NewDecoder(res.Body).Decode(&gRes)
	if err != nil {
		http.Error(w, "Github invalid response", http.StatusInternalServerError)
		return
	}

	githubID := gRes.Data.Viewer.ID
	userID, ok := githubConnections[githubID]
	if !ok {
		// new user = create account
		// maybe return, maybe not, depends
	}

	fmt.Fprintf(w, userID)

	// login to account UserID using JWT
}
