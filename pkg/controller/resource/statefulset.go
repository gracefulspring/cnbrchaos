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

// CheckStatefulSetAnnotation will check the annotation of StatefulSet
func CheckStatefulSetAnnotation(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, error) {
	targetAppList, err := getStatefulSetLists(clientset, engine)
	if err != nil {
		return engine, err
	}
	engine, chaosEnabledStatefulSet, err := checkForChaosEnabledStatefulSet(targetAppList, engine)
	if err != nil {
		return engine, err
	}
	if chaosEnabledStatefulSet == 0 {
		return engine, errors.New("no chaos-candidate found")
	}
	return engine, nil
}

// getStatefulSetLists will list the statefulset which having the chaos label
func getStatefulSetLists(clientset kubernetes.Interface, engine *chaosTypes.EngineInfo) (*appsV1.StatefulSetList, error) {
	targetAppList, err := clientset.AppsV1().StatefulSets(engine.AppInfo.Namespace).List(context.Background(), metaV1.ListOptions{
		LabelSelector: engine.Instance.Spec.Appinfo.Applabel,
		FieldSelector: ""})
	if err != nil {
		return nil, fmt.Errorf("error while listing statefulsets with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	if len(targetAppList.Items) == 0 {
		return nil, fmt.Errorf("no statefulset apps with matching labels %s", engine.Instance.Spec.Appinfo.Applabel)
	}
	return targetAppList, err
}

// checkForChaosEnabledStatefulSet will check and count the total chaos enabled application
func checkForChaosEnabledStatefulSet(targetAppList *appsV1.StatefulSetList, engine *chaosTypes.EngineInfo) (*chaosTypes.EngineInfo, int, error) {
	chaosEnabledStatefulSet := 0
	for _, statefulSet := range targetAppList.Items {
		annotationValue := statefulSet.ObjectMeta.GetAnnotations()[ChaosAnnotationKey]
		if IsChaosEnabled(annotationValue) {
			chaosTypes.Log.Info("chaos candidate of", "kind:", engine.AppInfo.Kind, "appName: ", statefulSet.ObjectMeta.Name, "appUUID: ", statefulSet.ObjectMeta.UID)
			chaosEnabledStatefulSet++
		}
	}
	return engine, chaosEnabledStatefulSet, nil
}
