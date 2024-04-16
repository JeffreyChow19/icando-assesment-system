package lib

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"icando/utils/logger"
	"os"
)

type Config struct {
	ServiceHost       string `envconfig:"SERVICE_HOST" required:"true"`
	ServiceState      int    `envconfig:"SERVICE_STATE" required:"true" default:"0"`
	ServiceName       string `envconfig:"SERVICE_NAME" required:"true"`
	ClientHost        string `envconfig:"CLIENT_HOST" required:"true"`
	Cors              string `envconfig:"CORS" required:"true" default:"https://localhost:5173"`
	AssessmentWebHost string `envconfig:"ASSESMENT_WEB_HOST" required:"true" default:"http://localhost:5002"`

	// environment is either 'dev' or 'prod' or 'test'
	Environment string `envconfig:"ENVIRONMENT" required:"true" default:"dev"`

	ServicePort int    `envconfig:"SERVICE_PORT" default:"8000" required:"true"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"INFO"`

	DatabaseHost     string `envconfig:"DB_HOST" required:"true"`
	DatabasePort     string `envconfig:"DB_PORT" required:"true"`
	DatabaseUsername string `envconfig:"DB_USERNAME" required:"true"`
	DatabasePassword string `envconfig:"DB_PASSWORD" required:"true"`
	DatabaseName     string `envconfig:"DB_NAME" required:"true"`

	JwtSecret    string `envconfig:"JWT_SECRET" required:"true"`
	RedisAddress string `envconfig:"REDIS_ADDRESS" required:"true"`
	SmtpUser     string `envconfig:"SMTP_USER" required:"true"`
	SmtpEmail    string `envconfig:"SMTP_EMAIL" required:"true"`
	SmtpPassword string `envconfig:"SMTP_PASSWORD" required:"true"`
	SmtpHost     string `envconfig:"SMTP_HOST" required:"true"`
	SmtpPort     int    `envconfig:"SMTP_PORT" required:"true"`
}

func NewConfig() (*Config, error) {
	var config Config

	filename := os.Getenv("CONFIG_FILE")
	if filename == "" {
		filename = ".env"
	}

	logger.Log.Info(fmt.Sprintf("Loading env from file: %s", filename))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
		return &config, nil
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil
}
