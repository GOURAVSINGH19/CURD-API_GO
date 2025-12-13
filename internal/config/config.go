package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true" env-default:":8082"`
}
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:true env-default:"production"`
	Storagepath string     `yaml:"storage_path" env-required:true`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configpath string
	configpath = os.Getenv("CONFIG_PATH")
	if configpath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configpath = *flags

		if configpath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configpath)
	}

	var cfg Config;

	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read config: %s", err.Error())
	}

	return &cfg

}
