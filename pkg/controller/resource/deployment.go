/*
Copyright 2021 hatech Authors
*/

package resource

import (
	"errors"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	chaosTypes "github.com/vossss/cnbrchaos/pkg/controller/types"
)

// CheckDeploymentAnnotation will check the annotation of deployment
func CheckDeploymentAnnotation(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, error) {
	targetAppList, err := getDeploymentLists(clientset, engine)
	if err != nil {
		return engine, err
	}
	engine, chaosEnabledDeployment, err := checkForChaosEnabledDeployment(targetAppList, engine)
	if err != nil {
		return engine, err
	}
	if chaosEnabledDeployment == 0 {
		return engine, errors.New("no chaos-candidate found")
	}
	return engine, nil
}

// getDeploymentLists will list the deployments which having the chaos label
func getDeploymentLists(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*v1.DeploymentList, error) {
	targetAppList, err := clientset.AppsV1().Deployments(engine.AppInfo.Namespace).List(metaV1.ListOptions{
		LabelSelector: engine.Instance.Spec.Appinfo.Applabel,
		FieldSelector: ""})
	if err != nil {
		return nil, fmt.Errorf("error while listing deployments with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	if len(targetAppList.Items) == 0 {
		return nil, fmt.Errorf("no deployments apps with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	return targetAppList, err
}

// checkForChaosEnabledDeployment will check and count the total chaos enabled application
func checkForChaosEnabledDeployment(targetAppList *v1.DeploymentList, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, int, error) {
	chaosEnabledDeployment := 0
	for _, deployment := range targetAppList.Items {
		annotationValue := deployment.ObjectMeta.GetAnnotations()[ChaosAnnotationKey]
		if IsChaosEnabled(annotationValue) {
			chaosTypes.Log.Info("chaos candidate of", "kind:", engine.AppInfo.Kind, "appName: ", deployment.ObjectMeta.Name, "appUUID: ", deployment.ObjectMeta.UID)
			chaosEnabledDeployment++
		}
	}
	return engine, chaosEnabledDeployment, nil
}
