package web

// we can separate per file "{model_name}{request_method}{request/response}"
// ex: activity_create_request.go / activity_create_response.go

type ActivityCreateRequest struct {
	Name 		string 	`validate:"required,min=1,max=100"`
	Status 	int 		`validate:"oneof='todo' 'done' 'in progress'"`
}
type ActivityUpdateRequest struct {
	Id 			int			`validate:"required"`
	Name 		string 	`validate:"required,min=1,max=100"`
	Status 	int 		`validate:"oneof='todo' 'done' 'in progress'"`
}
// type ActivityCreateResponse struct {
// 	Id			int
// 	Name		string
// 	Status	int
// }
type ActivityResponse struct {
	Id			int
	Name		string
	Status	int
}
