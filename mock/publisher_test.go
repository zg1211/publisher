package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	mp := NewPublisher()
	err := mp.Publish("test topic", []byte("test message"))
	assert.Nil(t, err)
	expected := []byte("test message")
	actual := <-mp.C("test topic")
	assert.Equal(t, expected, actual)

	err = mp.Publish("test topic", []byte("test message"))
	assert.Nil(t, err)
	expected = []byte("test message")
	actual = <-mp.C("test topic")
	assert.Equal(t, expected, actual)
}
