/*
Copyright 2021 hatech Authors
*/

package apis

import (
	"github.com/vossss/cnbrchaos/chaos-operator/pkg/apis/cnbrchaos/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
