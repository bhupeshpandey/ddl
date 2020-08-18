package config

import (
	model "com.webfrm.app/models"
	"fmt"
	"github.com/spf13/viper"

)

func ReadConfig() *model.Configurations {
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./../")

	viper.SetConfigType("yml")

	var configuration model.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))

	return &configuration
}
