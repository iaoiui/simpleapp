package simpleapp

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func Run() {
	fmt.Println("Hello SimpleApp")
}

type CSV struct {
	records [][]string
}

func (c *CSV) Records() [][]string {
	return c.records
}

// ReadCSV read specified csv file
func ReadCSV(filepath string) (CSV, error) {
	f, err := os.Open(filepath)
	if err != nil {
		//fmt.Errorf("cannot open file")
		return CSV{}, errors.New("cannot open file")
	}

	reader := csv.NewReader(f)
	var records [][]string
	records, err = reader.ReadAll()
	if err != nil {
		return CSV{}, errors.New("cannot read csv file")
	}
	//fmt.Println(records[1:])
	return CSV{records}, nil
}

func ExampleReadCSV() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	csv, err := ReadCSV(filepath.Join(homeDirectory, "test.csv"))
	if err != nil {
		fmt.Errorf("cannot ReadCSV")
	}
	fmt.Println(csv.Records()[1:])
	// Output: [[Alice 28] [Bob 29]]
}

func WriteCSV(content string, fileName string) {

}

// Env read environment variables
func Env(key string) string {
	value := os.Getenv(key)
	return value
}

func ExampleEnv() {
	debug := Env("DEBUG")
	fmt.Println("debug mode is ", debug)
}
