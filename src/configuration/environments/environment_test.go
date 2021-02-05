package environments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	env := GetEnvironment()
	assert.NotNil(t, env.ServiceName)
}
