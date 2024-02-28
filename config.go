package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type collectorConfig struct {
	Debug          bool   `mapstructure:"debug"`
	SampleData     bool   `mapstructure:"sample_data"`
	SampleDataPath string `mapstructure:"sample_data_path"`
}

type globalConfig struct {
	Collector collectorConfig `mapstructure:"collector"`
	OneView   oneviewConfig   `mapstructure:"oneview"`
}

type oneviewConfig struct {
	Domain string   `mapstructure:"domain"`
	Hosts  []string `mapstructure:"hosts"`
	Pass   string   `mapstructure:"pass"`
	User   string   `mapstructure:"user"`
}

var config globalConfig

func loadConfig() {

	log.Println("Loading configuration...")

	// load the config
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/hostdb-collector-oneview")
	viper.AddConfigPath(".")

	// load env vars
	viper.SetEnvPrefix("hostdb_collector_oneview")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// read the config file, and handle any errors
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error config file: %s", err))
	}

	// unmarshal into our struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(fmt.Errorf("unable to decode into struct, %v", err))
	}

	// debug
	if config.Collector.Debug {
		log.Println(fmt.Sprintf("%v", config))
	}

}
