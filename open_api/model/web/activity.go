package web

// we can separate per file "{model_name}{request_method}{request/response}"
// ex: activity_create_request.go / activity_create_response.go

type ActivityCreateRequest struct {
	Name 		string 	`validate:"required,min=1,max=100" json:"name"`
	Status 	string	`validate:"oneof='todo' 'done' 'in progress'" json:"status"`
}
type ActivityUpdateRequest struct {
	Id 			int			`validate:"required" json:"id"`
	Name 		string 	`validate:"required,min=1,max=100" json:"name"`
	Status 	string	`validate:"oneof='todo' 'done' 'in progress'" json:"status"`
}
// type ActivityCreateResponse struct {
// 	Id			int
// 	Name		string
// 	Status	int
// }
type ActivityResponse struct {
	Id			int 		`json:"id"`
	Name		string	`json:"name"`
	Status	string	`json:"status"`
}
