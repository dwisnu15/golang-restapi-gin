package configs

//exported configurations
type Configurations struct {
	Server ServerConfigs
	Database DatabaseConfigs
}

//exported ServerConfigurations
type ServerConfigs struct {
	Host string
	Port int
}

//exported DatabaseConfigurations
type DatabaseConfigs struct {
	DBName string
	DBUser string
	DBPassword string
}