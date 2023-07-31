package config

import (
	env "github.com/Netflix/go-env"
)

type Environment struct {
	Constants
	RunMode               string `env:"RUN_MODE,required=true"`
	Port                  int    `env:"PORT,required=true"`
	CorsAllowOrigins      string `env:"CORS_ALLOW_ORIGINS,required=true"`
	PostgreSqlUrl         string `env:"POSTGRESQL_URL,required=true"`
	PostgresqlHost        string
	PostgresqlPort        string
	PostgresqlUserName    string
	PostgresqlPassword    string
	PostgresqlDatabase    string
	RedisURI              string `env:"REDIS_URI,required=true"`
	RedisPassword         string `env:"REDIS_PASSWORD,required=true"`
	ElasticSearchURI      string `env:"ELASTICSEARCH_URI,required=true"`
	AwsAccessKeyID        string `env:"AWS_ACCESS_KEY_ID,required=true"`
	AwsSecretAccessKey    string `env:"AWS_SECRET_ACCESS_KEY,required=true"`
	AwsRegion             string `env:"AWS_REGION,required=true"`
	AwsS3Bucket           string `env:"AWS_S3_BUCKET,required=true"`
	MongoURI              string `env:"MONGO_URI,required=true"`
	AccessTokenSecretKey  string `env:"ACCESS_TOKEN_SECRET_KEY,required=true"`
	RefreshTokenSecretKey string `env:"REFRESH_TOKEN_SECRET_KEY,required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}

	return &environment, nil
}
