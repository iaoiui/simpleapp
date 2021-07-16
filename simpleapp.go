package simpleapp

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
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

func WriteCSV(content string, fileName string) {

}
