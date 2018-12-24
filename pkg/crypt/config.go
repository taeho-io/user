package crypt

type Config interface {
}

type DefaultConfig struct {
	Config
}

func NewConfig() Config {
	return &DefaultConfig{}
}

func NewMockConfig() Config {
	return NewConfig()
}
