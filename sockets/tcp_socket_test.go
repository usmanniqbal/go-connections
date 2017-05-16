package sockets

import (
	"testing"
	"crypto/tls"

	"github.com/docker/go-connections/tlsconfig"
)

const (
	key = "../tlsconfig/fixtures/key.pem"
	cert = "../tlsconfig/fixtures/cert.pem"
)

func TestNewTCPSocket(t *testing.T) {
	addr := "localhost:1234"
	versions := []uint16{
		tls.VersionTLS11,
		tls.VersionTLS12,
	}

	for _, v := range versions {
		tlsConfig, err := tlsconfig.Server(tlsconfig.Options{
			MinVersion: v,
			CertFile:   cert,
			KeyFile:    key,
		})

		if err != nil {
			t.Fatal("Unable to configure server TLS", err)
		}

		if tlsConfig.MinVersion != v {
			t.Fatal("Unexpected minimum TLS version: ", tlsConfig.MinVersion)
		}

		listener, err := NewTCPSocket(addr, tlsConfig)
		listener.Close()

		if err != nil {
			t.Fatal(err)
		}

	}
}
