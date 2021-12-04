package test

import (
	"open_api/depend_inject"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependInjectError(t *testing.T) {
	dependInject, err := depend_inject.InitializeService(true)

	assert.NotNil(t, err)
	assert.Nil(t, dependInject)
}

func TestDependInjectSuccess(t *testing.T) {
	dependInject, err := depend_inject.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, dependInject)
}

func TestConnection(t *testing.T) {
	conn, cleanup := depend_inject.InitializeConnection("ehem ehem")
	assert.NotNil(t, conn)

	cleanup()
}