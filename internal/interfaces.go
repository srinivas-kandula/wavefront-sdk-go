// Package internal offers helper interfaces that are internal to the Wavefront Go SDK.
// Interfaces within this package are not guaranteed to be backwards compatible between releases.
package internal

import "net/http"

// Reporter is an interface for reporting data to a Wavefront service.
type Reporter interface {
	Report(format string, pointLines string) (*http.Response, error)
	Server() string
}

type Flusher interface {
	Flush() error
	GetFailureCount() int64
	Start()
}

type ConnectionHandler interface {
	Connect() error
	Connected() bool
	Close()
	SendData(lines string) error

	Flusher
}
