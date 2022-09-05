package config

type TermitConfig struct {
	Db struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"db"`
	TermService struct {
		TermLimit int `yaml:"term_limit"`
	} `yaml:"term_service"`
	Server struct {
		GrpcPort int `yaml:"grpc_port"`
		HTTPPort int `yaml:"http_port"`
	} `yaml:"server"`
}