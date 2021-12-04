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

var helloSet = wire.NewSet(
	NewHelloImpl,
	// wire.Bind(kalo ada yg butuh => new(Hello), balikkan => *new(HelloImpl))
	wire.Bind(new(Hello), new(*HelloImpl)),
)

func InitializeHelloService() *HelloService {
	// expected after binding
	// hello := NewHelloImpl() // *HelloImpl
	// helloService := NewHelloService(hello)

	// wrong way
	// wire.Build(NewHelloImpl, NewHelloService)

	// correct way
	wire.Build(helloSet, NewHelloService)

	return nil
}

var fooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializeFooBar() *FooBar {
	// wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	// wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))

	wire.Build(fooBarSet, wire.Struct(new(FooBar), "*"))
	return nil
}
