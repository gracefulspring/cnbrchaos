/*
Copyright 2021 hatech Authors
*/

package resource

import (
	"context"
	"errors"
	"fmt"

	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	chaosTypes "github.com/gracefulspring/cnbrchaos/pkg/controller/types"
)

// CheckDaemonSetAnnotation will check the annotation of DaemonSet
func CheckDaemonSetAnnotation(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, error) {
	targetAppList, err := getDaemonSetLists(clientset, engine)
	if err != nil {
		return engine, err
	}
	engine, chaosEnabledDaemonSet, err := checkForChaosEnabledDaemonSet(targetAppList, engine)
	if err != nil {
		return engine, err
	}
	if chaosEnabledDaemonSet == 0 {
		return engine, errors.New("no chaos-candidate found")
	}
	return engine, nil
}

// getDaemonSetLists will list the daemonSets which having the chaos label
func getDaemonSetLists(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*appsV1.DaemonSetList, error) {
	targetAppList, err := clientset.AppsV1().DaemonSets(engine.AppInfo.Namespace).List(context.Background(), metaV1.ListOptions{
		LabelSelector: engine.Instance.Spec.Appinfo.Applabel,
		FieldSelector: ""})
	if err != nil {
		return nil, fmt.Errorf("error while listing daemonSets with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	if len(targetAppList.Items) == 0 {
		return nil, fmt.Errorf("no daemonSets apps with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	return targetAppList, err
}

// checkForChaosEnabledDaemonSet will check and count the total chaos enabled application
func checkForChaosEnabledDaemonSet(targetAppList *appsV1.DaemonSetList, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, int, error) {
	chaosEnabledDaemonSet := 0
	for _, daemonSet := range targetAppList.Items {
		annotationValue := daemonSet.ObjectMeta.GetAnnotations()[ChaosAnnotationKey]
		if IsChaosEnabled(annotationValue) {
			chaosTypes.Log.Info("chaos candidate of", "kind:", engine.AppInfo.Kind, "appName: ", daemonSet.ObjectMeta.Name, "appUUID: ", daemonSet.ObjectMeta.UID)
			chaosEnabledDaemonSet++
		}
	}
	return engine, chaosEnabledDaemonSet, nil
}
