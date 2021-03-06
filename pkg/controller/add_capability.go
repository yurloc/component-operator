package controller

import (
	"github.com/snowdrop/component-operator/pkg/controller/capability"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, capability.New)
}
