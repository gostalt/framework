package service

import (
	"github.com/gostalt/framework/maker"

	"github.com/sarulabs/di"
)

// FileGeneratorServiceProvider is responsible for loading in a
// number of "generators", such as entity and repository creators.
//
// TODO: Maybe rename this? Not super happy with the name.
type FileGeneratorServiceProvider struct {
	BaseProvider
}

// Register loads the various Makers into the container.
func (p FileGeneratorServiceProvider) Register(b *di.Builder) {
	b.Add(
		di.Def{
			Name: "EntityMaker",
			Build: func(c di.Container) (interface{}, error) {
				// TODO: Don't hardcode the path in here.
				// Probably have to move the config to the container
				// and retrieve it from there, with a fallback.
				return maker.EntityMaker{Path: "app/entity/"}, nil
			},
		},
		di.Def{
			Name: "RepositoryMaker",
			Build: func(c di.Container) (interface{}, error) {
				return maker.RepositoryMaker{Path: "app/repository/"}, nil
			},
		},
		di.Def{
			Name: "HandlerMaker",
			Build: func(c di.Container) (interface{}, error) {
				return maker.HandlerMaker{Path: "app/http/handler/"}, nil
			},
		},
	)
}
