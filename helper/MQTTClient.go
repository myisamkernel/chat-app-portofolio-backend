package helper

import (
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

var mqttClient mqtt.Client

func CommonMQTTOptionSetter(clientOption *mqtt.ClientOptions) *mqtt.ClientOptions {
	return clientOption

}

func MQTTConnect() {
	clientOption := mqtt.NewClientOptions()
	clientOption = CommonMQTTOptionSetter(clientOption)
	clientOption.SetUsername("admin@gmail.com")
	clientOption.SetPassword("asdf1234")
	clientOption.SetAutoReconnect(false)
	clientOption.AddBroker("tcp://" + os.Getenv("MQTT_HOST") + ":" + os.Getenv("MQTT_PORT"))
	mqttClient = mqtt.NewClient(clientOption)
	mqttClient.Connect().Wait()
}

func MqttSubscribe(topic uuid.UUID) {
	mqttClient.Subscribe(topic.String(), 2, func(c mqtt.Client, m mqtt.Message) {

	})
}

func ReconnectMqtt() {
	if !mqttClient.IsConnected() {
		mqttClient.Connect().Wait()
	}
}
