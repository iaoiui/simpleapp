package main

import (
	"fmt"
	"github.com/iaoiui/simpleapp"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	os.Exit(Run())
}

func Run() int {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	csv, err := simpleapp.ReadCSV(filepath.Join(homeDirectory, "test.csv"))
	if err != nil {
		fmt.Errorf("cannot ReadCSV")
		return 1
	}
	fmt.Println(csv.Records()[1:])

	return 0
}
