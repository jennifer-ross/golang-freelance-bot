package config

type PostgresConfig struct {
	Host     string `required:"true" env_name:"POSTGRES_HOST"`
	User     string `required:"true" env_name:"POSTGRES_USER"`
	Password string `required:"true" env_name:"POSTGRES_PASSWORD"`
	DB       string `required:"true" env_name:"POSTGRES_DB"`
	Port     int    `required:"true" env_name:"POSTGRES_PORT"`
}
