package config

type RedisConfig struct {
	Host     string `yaml:"host" env:"REDIS_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
	Password string `yaml:"password" env:"REDIS_PASSWORD" env-default:""`
	Db       int    `yaml:"db" env:"REDIS_DB" env-default:"0"`
}
