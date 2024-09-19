package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/smnzlnsk/monitoring-backend/logging"
	"github.com/smnzlnsk/monitoring-backend/metrics"
	"go.uber.org/zap"
)

func DefaultHandler(_ mqtt.Client, msg mqtt.Message) {
	logging.Logger.Debug("Received message.")
	err := metrics.GetDecoder().DecodeMetrics(msg.Payload())
	if err != nil {
		logging.Logger.Error("error handling metrics", zap.Error(err))
	}
}
