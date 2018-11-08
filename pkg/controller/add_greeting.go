package controller

import (
	"github.com/darthlukan/hello-operator/pkg/controller/greeting"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, greeting.Add)
}
