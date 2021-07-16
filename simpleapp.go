package simpleapp

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

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

// Env read environment variables
func Env(key string) string {
	value := os.Getenv(key)
	return value
}

func ExampleEnv() {
	debug := Env("DEBUG")
	fmt.Println("debug mode is ", debug)
}

func GetS3Object(bucket, item string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	file, err := os.Create(item)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	return nil
}

func ExampleGetS3Object() {

	bucket := Env("BUCKET")
	item := Env("ITEM")

	fmt.Println(bucket, item)

	GetS3Object(bucket, item)
}

func ListBuckets() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	svc := s3.New(sess)
	result, _ := svc.ListBuckets(nil)
	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
