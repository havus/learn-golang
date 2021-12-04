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