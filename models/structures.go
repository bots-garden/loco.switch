package models

import (
	"time"
)

// This structure is used to read a JSON payload from an HTTP request
type Function struct {
	Name     string `json:"name"`
	Revision string `json:"revision"`
	HttpPort int    `json:"httpPort"`
	Status   int    `json:"status"`
	Https    bool   `json:"https"`
	Domain   string `json:"domain"`
	Resource string `json:"resource"` // "/something" or ""
}

// The Resource field is used when registering a function with it's first Endpoint
// See: handlers.RegisterFunction
// And when assing an Endpoint to a function
// See: handlers.AddEndpointToFunction

type Endpoint struct {
	HttpPort int    `json:"httpPort"`
	Https    bool   `json:"https"`
	Domain   string `json:"domain"`
	Resource string `json:"resource"` // "/something" or ""
}

// TODO: add an error counter
type FunctionRecord struct {
	Name      string     `json:"name"`
	Revision  string     `json:"revision"`
	Status    int        `json:"status"`
	Timestamp time.Time  `json:"timestamp"` // record the time the event was requested
	Endpoints []Endpoint `json:"endpoints"` // for scale
}

type FunctionError struct {
	Error    string `json:"error"`
	Name     string `json:"name"`
	Revision string `json:"revision"`
}
