package service

import "github.com/devuo/go-app/config"

var (
	// Service configuration parameters
	ParameterContainer *config.ParameterContainer

	// Services singletons
	DB     *db
	Logger *logger
	Mailer *mailer
	Theme  *theme
	Web    *web
)

// Parameters contains the list of required service parameters
func Parameters() *config.ParameterContainer {
	params := []*config.Parameter{
		{Key: "db:user", Default: "", Description: "The database username"},
		{Key: "db:pass", Default: "", Description: "The database password"},
		{Key: "db:host", Default: "", Description: "The database host"},
		{Key: "db:name", Default: "", Description: "The database name"},

		{Key: "web:host", Default: "localhost", Description: "The web server hostname"},
		{Key: "web:port", Default: "4300", Description: "The web server port"},

		{Key: "mailer:host", Default: "127.0.0.1", Description: "The mail server host"},
		{Key: "mailer:port", Default: "587", Description: "The mail server port"},
		{Key: "mailer:user", Default: "", Description: "The mail server username"},
		{Key: "mailer:pass", Default: "", Description: "The mail server password"},
	}

	return &config.ParameterContainer{List: params}
}

// Init initializes and configures the application service singletons
func Init(params *config.ParameterContainer) {

	ParameterContainer = params

	DB = &db{
		User:     params.Value("db:user"),
		Password: params.Value("db:pass"),
		Host:     params.Value("db:host"),
		Name:     params.Value("db:name"),
	}

	Web = &web{
		Host: params.Value("web:host"),
		Port: params.Value("web:port"),
	}

	Mailer = &mailer{
		Host:     params.Value("mailer:host"),
		Port:     params.Value("mailer:port"),
		User:     params.Value("mailer:user"),
		Password: params.Value("mailer:pass"),
	}

	Logger = &logger{
		Debug: params.Value("logger:debug"),
	}

	Theme = &theme{}
}
