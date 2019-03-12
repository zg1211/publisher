package mqtt

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	qos              = byte(1)
	retained         = false
	tokenWaitTimeout = 3 * time.Second
)

type publisher struct {
	c     mqtt.Client
	topic string
}

func NewPublisher(client mqtt.Client) (*publisher, error) {
	return &publisher{
		c: client,
	}, nil
}

func (p *publisher) Publish(topic string, msg []byte) error {
	token := p.c.Publish(topic, qos, retained, msg)
	if token.WaitTimeout(tokenWaitTimeout); token.Error() != nil {
		return token.Error()
	}

	return nil
}
