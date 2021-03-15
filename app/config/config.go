package config

type (
	Config struct {
		MongoDbAuth MongoDbAuthConfig
		HTTPData    HTTPConfig
		UserAuth    AuthConfig
	}

	MongoDbAuthConfig struct {
		URI      string
		User     string
		Password string
		Name     string
	}

	HTTPConfig struct {
		Host string
		Port string
	}

	AuthConfig struct {
	}
)

func getConfig(path string) {

}
