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

type Producer struct {
	c     mqtt.Client
	topic string
}

func NewProducer(client mqtt.Client, topic string) (*Producer, error) {
	return &Producer{
		c:     client,
		topic: topic,
	}, nil
}

func (p *Producer) Publish(msg []byte) error {
	token := p.c.Publish(p.topic, qos, retained, msg)
	if token.WaitTimeout(tokenWaitTimeout); token.Error() != nil {
		return token.Error()
	}

	return nil
}
