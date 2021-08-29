package database_learn

import (
	"testing"
	"fmt"
	"context"
	"time"
	"database/sql"
)

func TestInsertCustomer(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// CREATE TABLE customers ( id int NOT NULL, name VARCHAR(100) NOT NULL, PRIMARY KEY (id) );
	// ALTER TABLE customers 
	// 		ADD COLUMN email VARCHAR(100),
	// 		ADD COLUMN balance int DEFAULT 0,
	// 		ADD COLUMN rating DOUBLE DEFAULT 0.0,
	// 		ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 		ADD COLUMN birth_date DATE,
	// 		ADD COLUMN is_premium BOOLEAN DEFAULT false;
	query := "INSERT INTO customers(id, name, email, balance, rating, birth_date, is_premium) "
	query += "VALUES(3, 'Third Customer', 'thirdcustomer@mail.com', 100000, 5, '1998-05-05', true)"

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert to customer is success!")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, is_premium, created_at FROM customers"
	rows, err := db.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, balance, rating int
		var name string
		var email sql.NullString
		var birth_date sql.NullTime
		var created_at time.Time
		var is_premium bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &is_premium, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("-------------------------------")
		fmt.Println("id         |", id)
		fmt.Println("name       |", name)
		if email.Valid {
			fmt.Println("email      |", email.String)
		}
		fmt.Println("balance    |", balance)
		fmt.Println("rating     |", rating)
		if birth_date.Valid {
			fmt.Println("birth_date |", birth_date.Time)
		}
		fmt.Println("is_premium |", is_premium)
		fmt.Println("created_at |", created_at)
	}
}

func TestInsertUser(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// CREATE TABLE users ( username VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL, PRIMARY KEY(username) );
	query := "INSERT INTO users(username, password) "
	query += "VALUES('admin', 'admin')"

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert to customer is success!")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	// CREATE TABLE users ( username VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL, PRIMARY KEY(username) );
	query := "SELECT username FROM users WHERE username = '"
	query += username
	query += "' AND password = '"
	query += password
	query += "' LIMIT 1"

	// query: SELECT username FROM users WHERE username = 'admin'; #' AND password = 'admin' LIMIT 1
	fmt.Println("query:", query)

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var usernameValue string

		err := rows.Scan(&usernameValue)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success login:", usernameValue)
	} else {
		fmt.Println("Gagal login:", username)
	}
}

func TestSanitizeSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var usernameValue string

		err := rows.Scan(&usernameValue)
		if err != nil {
			panic(err)
		}

		fmt.Println("Success login:", usernameValue)
	} else {
		fmt.Println("Gagal login:", username)
	}
}

func TestInserProduct(t *testing.T) {
	// Auto increment lesson
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// CREATE TABLE products ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100) NOT NULL, price DECIMAL(32, 2) );
	query := "INSERT INTO products(name) VALUES(?)"

	result, err := db.ExecContext(ctx, query, "Handphone")

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last insert id:", insertId)
}