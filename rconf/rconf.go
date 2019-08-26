package rconf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"../worker"
	"github.com/sevlyar/go-daemon"
)

func GetDeamonConfig() (cntxt *daemon.Context) {
	contextConfig, _ := os.Open("config.json")
	configDecoder := json.NewDecoder(contextConfig)
	cntxt = &daemon.Context{}
	err := configDecoder.Decode(&cntxt)
	if err != nil {
		log.Fatal("Unable decode config: ", err)
	}
	fmt.Println(ValidateDeamonConfig(cntxt))

	return cntxt
}
func GetWorkerConfig() (Wconfig *worker.WConfig) {
	WorkerConfig, _ := os.Open("config.json")
	configDecoder := json.NewDecoder(WorkerConfig)
	Wconfig = &worker.WConfig{}
	err := configDecoder.Decode(&Wconfig)
	if err != nil {
		log.Fatal("Unable decode worker config: ", err)
	}

	return Wconfig
}

func ValidateDeamonConfig(cntxt *daemon.Context) bool {
	fmt.Println(cntxt.LogFileName)
	fmt.Println("-----------------------")

	return true

}
func ValidateDeamonDBConfig(cntxt *daemon.Context) bool {
	fmt.Println(cntxt.LogFileName)
	fmt.Println("-----------------------")

	return true

}
func ValidateWorkerConfig(Wconfig *worker.WConfig) bool {
	fmt.Println(Wconfig.WorkerSleepTime)
	fmt.Println("-----------------------")

	return true

}
func match(pattern string, text string) bool {
	matched, _ := regexp.Match(pattern, []byte(text))
	if matched {
		fmt.Println("√", pattern, ":", text)
	} else {
		fmt.Println("X", pattern, ":", text)
	}
}
