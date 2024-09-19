package main

import (
	"github.com/smnzlnsk/monitoring-backend/logging"
	"github.com/smnzlnsk/monitoring-backend/metrics"
	"github.com/smnzlnsk/monitoring-backend/mqtt"
	"log"
	"os"
	"sync"
)

var once sync.Once
var initError error

type backend struct {
	messagingClient *mqtt.MqttClient
}

func main() {
	once.Do(func() {
		err := metrics.InitMarshaler("proto")
		if err != nil {
			initError = err
		}
		err = metrics.InitOakestraDecoder()
		if err != nil {
			initError = err
		}
		err = logging.NewLogger()
		return
	})
	if initError != nil {
		log.Fatal(initError)
	}

	_ = backend{
		messagingClient: mqtt.NewMqttClient(
			"backend-mac",
			os.Getenv("MQTT_URL"),
			os.Getenv("MQTT_PORT"),
			"metrics/done",
			mqtt.DefaultHandler),
	}
	// loop infinitely [for now]
	for {
		continue
	}
}
