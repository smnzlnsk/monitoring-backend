package mqtt

import (
	"fmt"
	"github.com/smnzlnsk/monitoring-backend/logging"
	"go.uber.org/zap"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type brokerConfig struct {
	host string
	port string
}

type MqttClient struct {
	client       mqtt.Client
	topics       map[string]mqtt.MessageHandler
	topicMutex   *sync.RWMutex
	messageMutex *sync.Mutex
	broker       brokerConfig
	clientId     string
}

func NewMqttClient(clientId string, host string, port string, topic string, handler mqtt.MessageHandler) *MqttClient {
	c := MqttClient{
		topics:       make(map[string]mqtt.MessageHandler),
		topicMutex:   &sync.RWMutex{},
		messageMutex: &sync.Mutex{},
		broker: brokerConfig{
			host: host,
			port: port,
		},
		clientId: clientId,
	}

	var onConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		logging.Logger.Info("Connected to MQTT Broker")

		client.Subscribe(topic, 1, handler)
		logging.Logger.Info(fmt.Sprintf("Subscribed to %s", topic))
	}

	var onConnectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		logging.Logger.Error("Connection lost to MQTT Broker")
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", c.broker.host, c.broker.port))
	opts.OnConnect = onConnectHandler
	opts.OnConnectionLost = onConnectionLostHandler
	opts.SetClientID(c.clientId)

	c.client = mqtt.NewClient(opts)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		logging.Logger.Panic("could not connect to broker", zap.Error(token.Error()))
	}
	return &c
}

func (c *MqttClient) RegisterTopic(topic string, handler mqtt.MessageHandler) {
	c.topicMutex.Lock()
	defer c.topicMutex.Unlock()
	c.topics[topic] = handler
	token := c.client.Subscribe(topic, 1, handler)
	if token.WaitTimeout(time.Second*5) && token.Error() != nil {
		logging.Logger.Error("error in register topic", zap.Error(token.Error()))
	}
}

func (c *MqttClient) DeregisterTopic(topic string) {
	c.topicMutex.Lock()
	defer c.topicMutex.Unlock()
	token := c.client.Unsubscribe(topic)
	delete(c.topics, topic)
	if token.WaitTimeout(time.Second*5) && token.Error() != nil {
		logging.Logger.Error("error in deregister topic", zap.Error(token.Error()))
	}
}
