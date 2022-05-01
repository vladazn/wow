package domain

type Config struct {
	Redis  RedisConfig   `yaml:"redis" env-prefix:"REDIS_"`
	Listen ListenConfigs `yaml:"listen"`
	Hash   HashConfigs   `yaml:"hash" env-prefix:"HASH_"`
	Client ClientConfigs `yaml:"client" env-prefix:"CLIENT_"`
}

type HashConfigs struct {
	Difficulty int `yaml:"difficulty" env:"DIFFICULTY" env-default:"5"`
}

type ClientConfigs struct {
	Host string `yaml:"host" env:"HOST" env-default:"http://localhost:8080"`
}

type ListenConfigs struct {
	Api  int `json:"api" env:"API_PORT" env-default:"8080"`
	Grpc int `json:"grpc" env:"GRPC_PORT"`
}

type RedisConfig struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"PORT" env-default:"6379"`
	Db   int    `yaml:"db" env:"DB" env-default:"0"`
}
