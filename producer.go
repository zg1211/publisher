package producer

type Producer interface {
	Publish(msg []byte) error
}
