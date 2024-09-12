package mqtt

import (
	"fmt"
	"log"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var initMqttClient sync.Once
var MqttClient mClient

type brokerConfig struct {
	host string
	port string
}

type mClient struct {
	client       mqtt.Client
	topics       map[string]mqtt.MessageHandler
	topicMutex   *sync.RWMutex
	messageMutex *sync.Mutex
	broker       brokerConfig
	clientId     string
}

func NewMqttClient(clientId string, host string, port string, topic string, handler mqtt.MessageHandler) *mClient {
	initMqttClient.Do(func() {
		MqttClient = mClient{
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
			log.Println("Connected to MQTT Broker")

			client.Subscribe(topic, 1, handler)
			log.Println("Subscribed to ", topic)
		}

		var onConnectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
			log.Println("Connection lost to MQTT Broker")
		}

		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%s", MqttClient.broker.host, MqttClient.broker.port))
		opts.OnConnect = onConnectHandler
		opts.OnConnectionLost = onConnectionLostHandler
		opts.SetClientID(MqttClient.clientId)

		MqttClient.client = mqtt.NewClient(opts)
		if token := MqttClient.client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	})
	return &MqttClient
}

func (c *mClient) RegisterTopic(topic string, handler mqtt.MessageHandler) {
	c.topicMutex.Lock()
	defer c.topicMutex.Unlock()
	c.topics[topic] = handler
	token := c.client.Subscribe(topic, 1, handler)
	if token.WaitTimeout(time.Second*5) && token.Error() != nil {
		log.Printf("error in register topic: %s", token.Error())
	}
}

func (c *mClient) DeregisterTopic(topic string) {
	c.topicMutex.Lock()
	defer c.topicMutex.Unlock()
	token := c.client.Unsubscribe(topic)
	delete(c.topics, topic)
	if token.WaitTimeout(time.Second*5) && token.Error() != nil {
		log.Printf("error in deregister topic: %s", token.Error())
	}
}
