package config

type JwtConfig struct {
	Key      []byte `env:"JWT_KEY"`
	Audience string `env:"JWT_AUDIENCE"`
	Issuer   string `env:"JWT_ISSUER"`
}
