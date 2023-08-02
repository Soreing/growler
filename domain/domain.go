// Package domain containing business logic, contracts and models
package domain

import (
	"github.com/Soreing/growler/domain/usecases"
	"github.com/google/wire"
)

// DependencySet dependencies provided by the domain
var DependencySet = wire.NewSet(
	usecases.NewUseCases,
)
