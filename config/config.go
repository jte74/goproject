package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User     string
		Password string
		Addr     string
		DBName   string
		Sslmode  string
	}

	App struct {
		Addr string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	println()

	path := RootDir()
	// viper.AddConfigPath("C:/Users/OWTPF3EZC3A/Projects/Training/Go/goProject/config")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return basepath
}
