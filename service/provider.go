package service

import "github.com/sarulabs/di/v2"

// Provider is a piece of the app that provides a service for the
// developer, such as a database or logger. Providers can Register
// items into the DI Container, and can perform actions.
type Provider interface {
	Register(*di.Builder)
	Boot(di.Container)
}

// BaseProvider provides a blank canvas to ensure that Providers
// satisfy the Provider interface.
type BaseProvider struct{}

// Register is called on all Providers in sequence when the app
// is first loaded.
func (p BaseProvider) Register(b *di.Builder) {}

// Boot is called when all the Providers have been Registered.
func (p BaseProvider) Boot(c di.Container) {}
