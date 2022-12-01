package golang

import "github.com/mesosphere/daggers/daggers"

type config struct {
	GoImageRepo string   `env:"GO_IMAGE_REPO,notEmpty" envDefault:"docker.io/golang"`
	GoImageTag  string   `env:"GO_IMAGE_TAG,notEmpty" envDefault:"1.19"`
	Args        []string `env:"GO_ARGS" envDefault:""  envSeparator:" "`

	Env map[string]string
}

// WithGoImageRepo sets whether to enable go module caching. Optional, defaults to docker.io/golang.
func WithGoImageRepo(repo string) daggers.Option[config] {
	return func(c config) config {
		c.GoImageRepo = repo
		return c
	}
}

// WithGoImageTag sets the go image tag to use for the container. Optional, defaults to 1.19.
func WithGoImageTag(tag string) daggers.Option[config] {
	return func(c config) config {
		c.GoImageTag = tag
		return c
	}
}

// WithArgs sets the arguments to pass to go.
func WithArgs(args ...string) daggers.Option[config] {
	return func(c config) config {
		c.Args = args
		return c
	}
}

// WithEnv sets the environment variables to pass to go.
func WithEnv(envMap map[string]string) daggers.Option[config] {
	return func(c config) config {
		c.Env = envMap
		return c
	}
}
