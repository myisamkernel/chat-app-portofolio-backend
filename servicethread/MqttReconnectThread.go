package servicethread

import (
	"chat-app-portfolio/helper"
	"time"
)

func InitMqttReconnectThread() {

	ticker := time.NewTicker(time.Duration(time.Millisecond))

	for {
		select {
		case <-ticker.C:
			helper.ReconnectMqtt()
		}
	}
}
