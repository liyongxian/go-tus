package tus

import (
	"net/http"
)

// Config provides a way to configure the Client depending on your needs.
type Config struct {
	// ChunkSize divide the file into chunks.
	ChunkSize int64
	// Resume enables resumable upload.
	Resume bool
	// OverridePatchMethod allow to by pass proxies sendind a POST request instead of PATCH.
	OverridePatchMethod bool
	// Store map an upload's fingerprint with the corresponding upload URL.
	// If Resume is true the Store is required.
	Store Store
	// Set custom header values used in all requests.
	Header http.Header
	// Set custom Transport settings.
	// Use this if you ahe behind a proxy.
	Transport *http.Transport
}

// DefaultConfig return the default Client configuration.
func DefaultConfig() *Config {
	return &Config{
		ChunkSize:           2 * 1024 * 1024,
		Resume:              false,
		OverridePatchMethod: false,
		Store:               nil,
		Header:              make(http.Header),
		Transport:           nil,
	}
}

// Validate validates the custom configuration.
func (c *Config) Validate() error {
	if c.ChunkSize < 1 {
		return ErrChuckSize
	}

	if c.Resume && c.Store == nil {
		return ErrNilStore
	}

	return nil
}

type Meta struct {
	Dir string
	//Watch string
	User string
	Project string
	Author string
	Token string
}
type Conf struct {
	UploadUrl string
	Meta Meta
}
