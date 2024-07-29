package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	ConfPath   string `env:"CONF_PATH" env-default:"/home/gorik/go/src/ozon-like/config/config.yaml"`
	HTTPServer `yaml:"http_server"`
	Database
	AuthJWT `yaml:"auth_jwt"`
	CSRFJWT `yaml:"csrfjwt"`
	GRPC    `yaml:"grpc"`

	Environmet      string `json:"environmet" env-default:"dev" env-description:"available: local, dev, prod"`
	LogFilePath     string `env:"LOG_FILEPATH" env-default:"zuzu.log"`
	PhotosFilesPath string `env:"PHOTOS_FILESPATH" env-default:"photos/"`
}

type HTTPServer struct {
}
type Database struct {
}
type AuthJWT struct {
}
type CSRFJWT struct {
}
type GRPC struct {
}

func MustLoad() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Printf("Coudln't read env   %s\n (fix: need to put .env file inroot dir) ", err)
		os.Exit(1)
	}

	pathToCfg := cfg.ConfPath
	if _, err = os.Stat(pathToCfg); os.IsNotExist(err) {
		println("err")
		log.Printf("Config file .env %s does not exist", pathToCfg)
		os.Exit(2)
	}
	err = cleanenv.ReadConfig(pathToCfg, &cfg)
	if err != nil {
		log.Printf("Coudln't read config file .env %s, %s", pathToCfg, err)
		os.Exit(3)
	}
	return &cfg
}
