package initialize

import (
	"SunnyContest/global"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper := viper.New()
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	//read config file
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fail to read configuaration %w\n", err))
	}

	//read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Failed to read configuaration %w\n", err))
	}
}
