/*
Package setting the config of service
*/
package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var v  *viper.Viper
func init() {
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("conf")

	// Handle errors reading the config file
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("FATAL ERROR CONFIG FILE"))
	}
}

//LoadServer server config
func LoadServer() error {
	str := v.GetString("host.address")
	port := v.GetInt64("host.port")
	fmt.Println(str, port)
	return nil
}

