package main

import (
	"errors"
	"fmt"
	"github.com/iaoiui/simpleapp"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	os.Exit(Run())
}

// LoadDotEnv Load .env file
func LoadDotEnv() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Error getting current working directory")
	}
	fmt.Println(filepath.Join(cwd, ".env"))
	err = godotenv.Load(filepath.Join(cwd, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func CheckDebugMode() error {
	var debug bool = false
	if simpleapp.Env("DEBUG") != "" {
		var err error
		debug, err = strconv.ParseBool(simpleapp.Env("DEBUG"))
		if err != nil {
			return errors.New("DEBUG env is not bool")
		}
	}

	fmt.Println("debug mode is ", debug)
	return nil
}

func ExampleCheckDebugMode() {
	if err := CheckDebugMode(); err != nil {
		fmt.Errorf("cannot check debug mode")
	}
	// Output: debug mode is  true
}

func Run() int {

	LoadDotEnv()

	return 0
}
