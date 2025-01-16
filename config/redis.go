package config

type RedisConfig struct {
	Host string `required:"true" env_name:"REDIS_HOST"`
	Port int    `required:"true" env_name:"REDIS_PORT"`
}
