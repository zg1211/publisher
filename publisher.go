package publisher

type Publisher interface {
	Publish(topic string, msg []byte) error
}
