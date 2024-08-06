package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ConfPath   string `env:"CONF_PATH" env-default:"/home/gorik/go/src/ozon-like/config/config.yaml"`
	HTTPServer `yaml:"http_server"`
	Database
	AuthJWT `yaml:"auth_jwt"`
	CSRFJWT `yaml:"csrfjwt"`
	GRPC    GRPC `yaml:"grpc"`

	Environmet      string `json:"environmet" env-default:"local" env-description:"available: local, dev, prod"`
	LogFilePath     string `env:"LOG_FILEPATH" env-default:"zuzu.log"`
	PhotosFilesPath string `env:"PHOTOS_FILESPATH" env-default:"photos/"`
}

type HTTPServer struct {
}
type Database struct {
	DBuser    string
	DBpass    string
	DBname    string
	DBhost    string
	DBport    string
	DBsslmode string
}
type AuthJWT struct {
	JwtAccess            string        `env:"AUTH_JWT_SECRET_KEY" env-required:"true"`
	AccessExpirationTime time.Duration `yaml:"accessExpirationTime" yaml-defualt:"6h"`
	Issuer               string
}
type CSRFJWT struct {
}
type GRPC struct {
	AuthPort            int    `env:"GRPC_AUTH_PORT" env-defualt:"8011" yaml:"auth_port"`
	OrderPort           int    `env:"GRPC_ORDER_PORT" env-defualt:"8012"`
	ProductsPort        int    `env:"GRPC_PRODUCTS_PORT" env-defualt:"8013"`
	AuthContainerIP     string `env:"GRPC_AUTH_CONTAINER_IP" env-defualt:"zuzu-auth"`
	OrderContainerIP    string `env:"GRPC_ORDER_CONTAINER_IP" env-defualt:"zuzu-order"`
	ProductsContainerIP string `env:"GRPC_PRODUCTS_CONTAINER_IP" env-defualt:"zuzu-products"`
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

func (a AuthJWT) GetTTL() time.Duration {
	return a.AccessExpirationTime
}
func (a AuthJWT) GetSecret() string {
	return a.JwtAccess
}
func (a AuthJWT) GetIssuer() string {
	return "auth"
}
