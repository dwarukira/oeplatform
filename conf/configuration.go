package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// DBConfiguration holds all the database related configuration.
type DBConfiguration struct {
	Dialect     string
	Driver      string `required:"true"`
	URL         string `envconfig:"DATABASE_URL" required:"true"`
	Namespace   string
	Automigrate bool
}

type SMTPConfiguration struct {
	Host       string `json:"host"`
	Port       int    `json:"port" default:"587"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	AdminEmail string `json:"admin_email" split_words:"true"`
}

// JWTConfiguration holds all the JWT related configuration.
type JWTConfiguration struct {
	Secret          string `json:"secret"`
	AdminGroupName  string `json:"admin_group_name" split_words:"true"`
	SellerGroupName string `json:"seller_group_name" split_words:"true"`
}

// GraphqlConfiguration holds all the Grapgql related configuration.
type GraphqlConfiguration struct {
	Path             string `split_words:"true" default:"graphql"`
	PlaygroundPath   string `split_words:"true" default:"playground"`
	EnablePlayground bool   `split_words:"true"`
}

// GlobalConfiguration holds all the global configuration for gocommerce
type GlobalConfiguration struct {
	API struct {
		Host     string
		Port     int `envconfig:"PORT" default:"8080"`
		Endpoint string
	}
	DB                DBConfiguration
	Logging           LoggingConfig `envconfig:"LOG"`
	OperatorToken     string        `split_words:"true"`
	MultiInstanceMode bool
	SMTP              SMTPConfiguration `json:"smtp"`
	Graphql           GraphqlConfiguration

	JWT JWTConfiguration
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		// handle if .env file does not exist, this is OK
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

type Configuration struct {
	SiteURL string           `json:"site_url" split_words:"true" required:"true"`
	JWT     JWTConfiguration `json:"jwt"`

	SMTP SMTPConfiguration `json:"smtp"`
}

// LoadGlobal will construct the core config from the file
func LoadGlobal(filename string) (*GlobalConfiguration, *logrus.Entry, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, nil, err
	}

	config := new(GlobalConfiguration)
	if err := envconfig.Process("oe", config); err != nil {
		return nil, nil, err
	}
	log, err := ConfigureLogging(&config.Logging)
	if err != nil {
		return nil, nil, err
	}
	return config, log, nil
}

// LoadConfig loads the per-instance configuration from a file
func LoadConfig(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("oe", config); err != nil {
		return nil, err
	}
	config.ApplyDefaults()
	return config, nil
}

// ApplyDefaults sets defaults for a Configuration
func (config *Configuration) ApplyDefaults() {
	if config.JWT.AdminGroupName == "" {
		config.JWT.AdminGroupName = "admin"
	}
}
