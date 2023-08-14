package config_parser

// Config is the root of the config file.
// It contains all the configuration for the services.
type Config struct {
	Service ServiceConfig `yaml:"service"`
}

// ServiceConfig contains the configuration for a service.
// It contains the name and the port of the service.
type ServiceConfig struct {
	ServiceName string `yaml:"name"`
	ServicePort int    `yaml:"port"`
}

