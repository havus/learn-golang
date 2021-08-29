package greeting

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require" immediately raise error and not execute next code
)

func TestMain(m *testing.M) {
	fmt.Println("Start to testing")
	m.Run()
	fmt.Println("Testing finished!")
}

func TestGreetingAssertion(t *testing.T) {
	t.Run("Hello World", func(t *testing.T) {
		result := Greet("World")

		assert.Equal(t, "Hello World!", result, "return value must be \"Hello World!\"")
		assert.NotEqual(t, "Hai World!", result, "return value must NOT be \"Hai World!\"")
	})

	// go test -v -run=TestGreetingAssertion/Hello_John_Doe
	t.Run("Hello John Doe", func(t *testing.T) {
		result := Greet("John Doe")
	
		assert.Equal(t, "Hello John Doe!", result, "return value must be \"Hello John Doe!\"")
		assert.NotEqual(t, "Hai John Doe!", result, "return value must NOT be \"Hai John Doe!\"")
	})
}
