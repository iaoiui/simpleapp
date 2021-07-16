package main

import (
	"errors"
	"fmt"
	"github.com/iaoiui/simpleapp"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	os.Exit(Run())
}

// LoadDotEnv Load .env file
func loadDotEnv() {
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

func isDebug() (bool, error) {
	var debug bool = false
	if simpleapp.Env("DEBUG") != "" {
		var err error
		debug, err = strconv.ParseBool(simpleapp.Env("DEBUG"))
		if err != nil {
			return debug, errors.New("DEBUG env is not bool")
		}
	}

	//fmt.Println("debug mode is ", debug)
	return debug, nil
}

func exampleCheckDebugMode() {
	if _, err := isDebug(); err != nil {
		fmt.Errorf("cannot check debug mode")
	}
	// Output: debug mode is  true
}

func Run() int {

	loadDotEnv()
	runWebServer()
	return 0
}

func handler(w http.ResponseWriter, r *http.Request) {
	text := "Hello, World\n"
	fmt.Fprintf(w, text)

	debug, err := isDebug()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	if debug {
		fmt.Fprintf(w, "debug mode\n")
	}
}
func runWebServer() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":8080", nil)
}
