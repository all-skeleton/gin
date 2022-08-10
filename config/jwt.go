package config

type JwtConfig struct {
	// 发行人
	Issuer string `default:"gin-skeleton"`
	// 用于Api
	JwtSecret string `env:"JWT_SECRET_API"`
}
