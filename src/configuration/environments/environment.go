package environments

type Environment struct {
	Port             string `env:"PORT"`
	ConnectionString string `env:"CONNECTION_STRING`
}
