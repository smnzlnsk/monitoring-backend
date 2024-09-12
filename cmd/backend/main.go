package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	mymqtt "github.com/smnzlnsk/monitoring-backend/mqtt"
	"log"
	"os"
)

func mqttMessageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Println("Received message: ", string(msg.Payload()))
}

func main() {
	mymqtt.NewMqttClient(
		"backend-mac",
		os.Getenv("MQTT_URL"),
		os.Getenv("MQTT_PORT"),
		"metrics/done",
		mqttMessageHandler)
	for {
	}
}
