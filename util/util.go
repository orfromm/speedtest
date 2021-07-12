package util

import (
	"math/rand"
	"net"
)

// Urandom produces a random stream of bytes
func Urandom(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Int31())
	}

	return b
}

// GetNetwork finds the network for the given name (lo, en0, ...)
// based on https://stackoverflow.com/a/31551220
func GetNetwork(name string) *net.IPNet {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil
	}
	addrs, _ := iface.Addrs()
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				return ipnet
			}
		}
	}

	return nil
}
