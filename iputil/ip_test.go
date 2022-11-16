package iputil

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocal(t *testing.T) {
	locals := []string{
		"10.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"172.17.0.0/12",
		"172.18.0.0/12",
		"172.19.0.0/12",
		"172.20.0.0/12",
		"172.21.0.0/12",
		"172.22.0.0/12",
		"172.23.0.0/12",
		"172.24.0.0/12",
		"172.25.0.0/12",
		"172.26.0.0/12",
		"172.27.0.0/12",
		"172.28.0.0/12",
		"172.29.0.0/12",
		"172.30.0.0/12",
		"172.31.0.0/12",
		"192.168.0.0/16",
	}
	for _, local := range locals {
		assert.True(t, true, IsLocal(net.ParseIP(local).To4()))
	}
}

var ip2longMapping = map[string]uint{
	"0.0.0.0":         0,
	"118.195.209.151": 1992544663,
	"127.0.0.1":       2130706433,
	"182.150.20.6":    3063288838,
	"192.168.1.1":     3232235777,
	"255.255.255.255": 4294967295,
}

func TestParse(t *testing.T) {
	for ipStr, long := range ip2longMapping {
		ip, err := Parse(long)
		assert.Nil(t, err)
		assert.Equal(t, ip.String(), ipStr)
	}
}

func TestToLong(t *testing.T) {
	for ipStr, long := range ip2longMapping {
		l, err := ToLong(net.ParseIP(ipStr))
		assert.Nil(t, err)
		assert.Equal(t, long, l)
	}
}
