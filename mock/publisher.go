package mock

// concurrent is not supported
type publisher struct {
	cs map[string]chan []byte
}

func NewPublisher() *publisher {
	return &publisher{
		cs: make(map[string]chan []byte),
	}
}

func (p *publisher) Publish(topic string, msg []byte) error {
	if c, ok := p.cs[topic]; ok {
		c <- msg

		return nil
	}

	c := make(chan []byte, 100)
	p.cs[topic] = c
	c <- msg

	return nil
}

func (p *publisher) C(topic string) <-chan []byte {
	return p.cs[topic]
}
