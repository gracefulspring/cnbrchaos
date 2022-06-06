/*
Copyright 2021 hatech Authors
*/

// To create logs for debugging or detailing, please follow this syntax.
// use function log.Info
// in parameters give the name of the log / error (string) ,
// with the variable name for the value(string)
// and then the value to log (any datatype)
// All values should be in key : value pairs only
// For eg. : log.Info("name_of_the_log","variable_name_for_the_value",value, ......)
// For eg. : log.Error(err,"error_statement","variable_name",value)
// For eg. : log.Printf
//("error statement %q other variables %s/%s",targetValue, object.Namespace, object.Name)
// For eg. : log.Errorf
//("unable to reconcile object %s/%s: %v", object.Namespace, object.Name, err)
// This logger uses a structured logging schema in JSON format, which will / can be used further
// to access the values in the logger.

package types

import (
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gracefulspring/cnbrchaos/pkg/apis/cnbrchaos/v1alpha1"
	cnbrchaosv1alpha1 "github.com/gracefulspring/cnbrchaos/pkg/apis/cnbrchaos/v1alpha1"
	"github.com/gracefulspring/cnbrchaos/pkg/controller/utils"
)

var (
	// AppLabelKey contains the application label key
	AppLabelKey string

	// DefaultAnnotationCheck contains the default value (true) of the annotationCheck
	DefaultAnnotationCheck = "false"

	// AppLabelValue contains the application label value
	AppLabelValue string

	// Log with default name ie: controller_chaosengine
	Log = logf.Log.WithName("controller_chaosengine")

	// DefaultChaosRunnerImage contains the default value of runner resource
	DefaultChaosRunnerImage = "cnbrchaos/chaos-runner:latest"

	// ResultCRDName contains name of the chaosresult CRD
	ResultCRDName = "chaosresults.cnbrchaos.io"
)

// ApplicationInfo contains the chaos details for target application
type ApplicationInfo struct {
	Namespace          string
	Label              string
	ExperimentList     []cnbrchaosv1alpha1.ExperimentList
	ServiceAccountName string
	Kind               string
}

//EngineInfo Related information
type EngineInfo struct {
	Instance       *cnbrchaosv1alpha1.ChaosEngine
	AppInfo        *ApplicationInfo
	ConfigMaps     []v1alpha1.ConfigMap
	Secrets        []v1alpha1.Secret
	VolumeOpts     utils.VolumeOpts
	AppExperiments []string
}
