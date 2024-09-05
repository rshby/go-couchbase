package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// LoadConfig is method to load env
func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("failed load env : %s", err.Error())
	}
}

func CouchbaseHost() string {
	return viper.GetString("couchbase.host")
}

func CouchbaseUser() string {
	return viper.GetString("couchbase.user")
}

func CouchbasePassword() string {
	return viper.GetString("couchbase.password")
}

func CouchbaseBucket() string {
	return viper.GetString("couchbase.bucket")
}
