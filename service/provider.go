package service

import "github.com/sarulabs/di/v2"

// Providers are a feature of Gostalt that bootstraps the entire application. Each
// component, such as routing and database interaction, is governed by its own
// service provider.
type Provider interface {
	Register(*di.Builder) error
	Boot(di.Container) error
}

// BaseProvider enables developers to use embedding to ensure that Providers that
// they have created themselves satisfy the Provider interface.
type BaseProvider struct{}

// Register is called on all Providers in sequence when the app is first loaded.
func (p BaseProvider) Register(b *di.Builder) error {
	return nil
}

// Boot is called when all the Providers have been Registered.
func (p BaseProvider) Boot(c di.Container) error {
	return nil
}
