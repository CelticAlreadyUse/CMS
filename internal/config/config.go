package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func InitLoadWithViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func PORT_HTTP() string {
	return viper.GetString("server.port")
}
func DBHOST() string {
	return viper.GetString("mysql.host")
}
func DBPort() string {
	return viper.GetString("mysql.port")
}
func DBUSER() string {
	return viper.GetString("mysql.username")
}
func DBPass() string {
	return viper.GetString("mysql.password")
}
func DBName() string {
	return viper.GetString("mysql.name")
}
func JWTSigningKey() string {
	return viper.GetString("jwt.signing_key")
}
func JWTExp() time.Duration {
	return viper.GetDuration("jwt.exp")
}
