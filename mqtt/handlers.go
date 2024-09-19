package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/smnzlnsk/monitoring-backend/metrics"
	"log"
)

func DefaultHandler(_ mqtt.Client, msg mqtt.Message) {
	log.Println("Received message.")
	err := metrics.GetDecoder().DecodeMetrics(msg.Payload())
	if err != nil {
		log.Printf("Error handling metrics: %v", err)
	}
}
