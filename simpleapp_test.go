package simpleapp

import (
	"log"
	"os/user"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestReadCSV(t *testing.T) {
	type args struct {
		filepath string
	}
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDirectory := user.HomeDir
	tests := []struct {
		name string
		args args
		want CSV
	}{
		{"read csv",
			args{filepath.Join(homeDirectory, "test.csv")},
			CSV{[][]string{{"name", "age"}, {"Alice", "28"}, {"Bob", "29"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ReadCSV(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
