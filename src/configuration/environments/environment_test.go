package environments

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetEnvironment(t *testing.T) {
	env := GetEnvironment()

	result := env.ServiceName

	assert.NotNil(t, result)
}

func TestShouldReloadTheEnvironment(t *testing.T) {
	env := GetEnvironment()
	nomeOriginal := env.ServiceName
	_ = os.Setenv("SERVICE_NAME", "Novo nome")

	ReloadEnvironment()

	assert.Equal(t, env.ServiceName, nomeOriginal)
}
