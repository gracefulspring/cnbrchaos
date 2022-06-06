/*
Copyright 2021 hatech Authors
*/

package resource

import (
	"fmt"
	"os"
	"strings"

	chaosTypes "github.com/gracefulspring/cnbrchaos/pkg/controller/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// Annotations on app to enable chaos on it
const (
	ChaosAnnotationValue      = "true"
	DefaultChaosAnnotationKey = "cnbrchaos.io/chaos"
)

var (
	// ChaosAnnotationKey is global variable used as the Key for annotation check.
	ChaosAnnotationKey = GetAnnotationKey()
)

// GetAnnotationKey returns the annotation to be used while validating applications.
func GetAnnotationKey() string {

	annotationKey := os.Getenv("CUSTOM_ANNOTATION")
	if len(annotationKey) != 0 {
		return annotationKey
	}
	return DefaultChaosAnnotationKey

}

// CheckChaosAnnotation will check for the annotation of required resources
func CheckChaosAnnotation(engine *chaosTypes.EngineInfo, clientset kubernetes.Interface, dynamicClientSet dynamic.Interface) (*chaosTypes.EngineInfo, error) {

	switch strings.ToLower(engine.AppInfo.Kind) {
	case "deployment", "deployments":
		engine, err := CheckDeploymentAnnotation(clientset, engine)
		if err != nil {
			return engine, fmt.Errorf("resource type 'deployment', err: %+v", err)
		}
	case "statefulset", "statefulsets":
		engine, err := CheckStatefulSetAnnotation(clientset, engine)
		if err != nil {
			return engine, fmt.Errorf("resource type 'statefulset', err: %+v", err)
		}
	case "daemonset", "daemonsets":
		engine, err := CheckDaemonSetAnnotation(clientset, engine)
		if err != nil {
			return engine, fmt.Errorf("resource type 'daemonset', err: %+v", err)
		}
	case "deploymentconfig", "deploymentconfigs":
		engine, err := CheckDeploymentConfigAnnotation(dynamicClientSet, engine)
		if err != nil {
			return engine, fmt.Errorf("resource type 'deploymentconfig', err: %+v", err)
		}
	case "rollout", "rollouts":
		engine, err := CheckRolloutAnnotation(dynamicClientSet, engine)
		if err != nil {
			return engine, fmt.Errorf("resource type 'rollout', err: %+v", err)
		}
	default:
		return engine, fmt.Errorf("resource type '%s' not supported for induce chaos", engine.AppInfo.Kind)
	}
	return engine, nil
}

// IsChaosEnabled check for the given annotation value
func IsChaosEnabled(annotationValue string) bool {
	return annotationValue == ChaosAnnotationValue
}
