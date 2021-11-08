/*
Copyright 2021 hatech Authors
*/

package controller

import (
	"github.com/vossss/cnbrchaos/pkg/controller/chaosengine"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, chaosengine.Add)
}
