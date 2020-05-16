package rtmp_test

import (
	"testing"

	"github.com/kris-nova/rtmp"
	"github.com/kris-nova/rtmp/server"
	"github.com/stretchr/testify/assert"
)

func TestNewServerMakesNewServers(t *testing.T) {
	s, err := rtmp.NewServer(":1935")

	assert.IsType(t, new(server.Server), s)
	assert.Nil(t, err)
}
