package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/iaoiui/simpleapp"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	var err error
	debug, err = strconv.ParseBool(simpleapp.Env("DEBUG", "false"))
	if err != nil {
		return debug, errors.New("DEBUG env is not bool")
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
	//
	loadDotEnv()

	flag.Parse()

	var c myCollector
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(*addr, nil))

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
	var port int = 8080
	var err error
	port, err = strconv.Atoi(simpleapp.Env("PORT", string(port)))
	if err != nil {
		fmt.Errorf("cannot losf PORT environment")
	}
	//if simpleapp.Env("PORT", string(port)) == "" {
	//	port, err = strconv.Atoi(simpleapp.Env("PORT", string(port)))
	//	if err != nil {
	//		fmt.Errorf("cannot losf PORT environment")
	//	}
	//}

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// metrics
// Metricsの定義
const (
	namespace = "SampleMetric"
)

type myCollector struct{} // 今回働いてくれるインスタンス

// metricsの記述子 「metricsの中に埋め込む情報の1つ（名前、#HELP に乗せる情報）であり、後にグラフで表示させるための数値以外のもの」

var (
	exampleCount = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "example_count",
		Help:      "example counter help",
	})
	exampleGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "example_gauge",
		Help:      "example gauge help",
	})
)

// Describe と Collect

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- exampleCount.Desc()
	ch <- exampleGauge.Desc()
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	exampleValue := float64(12345)

	ch <- prometheus.MustNewConstMetric(
		exampleCount.Desc(),
		prometheus.CounterValue,
		float64(exampleValue),
	)
	ch <- prometheus.MustNewConstMetric(
		exampleGauge.Desc(),
		prometheus.GaugeValue,
		float64(exampleValue),
	)
}

var addr = flag.String("listen-address", "127.0.0.1:5000", "The address to listen on for HTTP requests.")
