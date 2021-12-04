//go:build wireinject
//+build wireinject

package depend_inject

import (
	"github.com/google/wire"
)

// run command: `wire gen open_api/depend_inject`
// or run command `cd depend_inject && wire`
func InitializeService(isAnError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}
func InitializeDatabaseRepo() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMySQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepo, NewFooService)
var barSet = wire.NewSet(NewBarRepo, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}