package utils

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	SwaggerUrl           string        `mapstructure:"SWAGGER_URL"`
	MongoDBUri           string        `mapstructure:"MONGODB_URI"`
	RedisHost            string        `mapstructure:"REDIS_HOST"`
	RedisPort            int           `mapstructure:"REDIS_PORT"`
	RedisPassword        string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB              int           `mapstructure:"REDIS_DB"`
	RedisExpiredDuration int           `mapstructure:"REDIS_EXPIRED_DURATION"`
	GinMode              string        `mapstructure:"GIN_MODE"`
	LogFilename          string        `mapstructure:"LOG_FILENAME"`
	LogMaxSize           int           `mapstructure:"LOG_MAX_SIZE"`
	LogMaxBackups        int           `mapstructure:"LOG_MAX_BACKUPS"`
	LogMaxAge            int           `mapstructure:"LOG_MAX_AGE"`
	LogCompress          bool          `mapstructure:"LOG_COMPRESS"`
	MongoTimeout         time.Duration `mapstructure:"MONGO_TIMEOUT"`
	BcryptCost           int           `mapstructure:"BCRYPT_COST"`
	MongoDBName          string        `mapstructure:"MONGO_DB_NAME"`
	SanctumSecretKey     string        `mapstructure:"SANCTUM_SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("HTTP_SERVER_ADDRESS", ":8080")
	viper.SetDefault("SWAGGER_URL", "/swagger/docs")
	viper.SetDefault("LOG_FILENAME", "log/app.log")
	viper.SetDefault("LOG_MAX_SIZE", 10)
	viper.SetDefault("LOG_MAX_BACKUPS", 5)
	viper.SetDefault("LOG_MAX_AGE", 28)
	viper.SetDefault("LOG_COMPRESS", true)
	viper.SetDefault("MONGO_TIMEOUT", 5*time.Second)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
