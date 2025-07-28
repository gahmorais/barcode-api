package config

type env struct {
	DatabaseName string
	Password     string
	User         string
	Address      string
	Port         int
}

func NewEnv() *env {
	return &env{
		DatabaseName: "local",
		User:         "root",
		Password:     "example",
		Address:      "localhost",
		Port:         27017,
	}
}
