package depend_inject

type Database struct {
	Name string
}

// before creating alias for database
// func NewDatabasePostgreSQL() *Database {
// 	return &Database{Name: "PostgreSQL"}
// }
// func NewDatabaseMySQL() *Database {
// 	return &Database{Name: "MySQL"}
// }
// func NewDatabaseMongoDB() *Database {
// 	return &Database{Name: "MongoDB"}
// }

// type DatabaseRepository struct {
// 	DatabasePostgreSQL 	*Database
// 	DatabaseMySQL 			*Database
// 	DatabaseMongoDB 		*Database
// }

// after creating alias for database
type DatabasePostgreSQL Database
type DatabaseMySQL 			Database
type DatabaseMongoDB 		Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}
func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{Name: "MySQL"})
}
func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}


type DatabaseRepository struct {
	DatabasePostgreSQL 	*DatabasePostgreSQL
	DatabaseMySQL 			*DatabaseMySQL
	DatabaseMongoDB 		*DatabaseMongoDB
}

func NewDatabaseRepository(
	postgreSQL *DatabasePostgreSQL,
	mySQL *DatabaseMySQL,
	mongoDB *DatabaseMongoDB
) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMySQL: 			mySQL,
		DatabaseMongoDB: 		mongoDB,
	}
}