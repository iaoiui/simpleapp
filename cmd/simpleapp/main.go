package main

import (
	"fmt"
	"github.com/iaoiui/simpleapp"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Exit(Run())
}

// LoadDotEnv Load .env file
func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Run() int {
	debug := false
	var err error
	LoadDotEnv()

	if simpleapp.Env("DEBUG") != "" {
		debug, err = strconv.ParseBool(simpleapp.Env("DEBUG"))
		if err != nil {
			fmt.Errorf("DEBUG env is not bool")
			return 1
		}
	}

	fmt.Println("debug mode is ", debug)
	//if (DEBUG)

	return 0
}
