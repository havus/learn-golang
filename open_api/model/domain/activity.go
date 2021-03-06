package domain

type Activity struct {
	Id 			int
	Name 		string 	`validate:"required,min=1,max=100"`
	Status 	string 	`validate:"oneof='todo' 'done' 'in progress'"`
}
