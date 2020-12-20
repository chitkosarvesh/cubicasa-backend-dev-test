package config

import (
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/cubicasa/")
	viper.AddConfigPath(".")
}
func Setup() bool {
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	os.Setenv("PORT", viper.Get("port").(string))
	return true

}
