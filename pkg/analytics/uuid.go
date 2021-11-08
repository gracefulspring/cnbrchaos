/*
Copyright 2021 hatech Authors
*/

package analytics

import (
	"fmt"
	"os"

	clientset "github.com/vossss/cnbrchaos/chaos-operator/pkg/kubernetes"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClientUUID contains clientUUID for analytics
var ClientUUID string

// it derives the UID of the chaos-operator deployment
// and used it for the analytics
func getUID() (string, error) {
	// creates kubernetes client
	clients, err := clientset.CreateClientSet()
	if err != nil {
		return "", err
	}
	// deriving operator pod name & namespace
	podName := os.Getenv("POD_NAME")
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podName == "" || podNamespace == "" {
		return podName, fmt.Errorf("POD_NAME or POD_NAMESPACE ENV not set")
	}
	// get operator pod details
	pod, err := clients.CoreV1().Pods(podNamespace).Get(podName, v1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("unable to get %s pod in %s namespace", podName, podNamespace)
	}
	return getOperatorUID(pod, clients)
}

// it returns the deployment name, derived from the owner references
func getDeploymentName(pod *core_v1.Pod, clients *kubernetes.Clientset) (string, error) {
	for _, own := range pod.OwnerReferences {
		if own.Kind == "ReplicaSet" {
			rs, err := clients.AppsV1().ReplicaSets(pod.Namespace).Get(own.Name, v1.GetOptions{})
			if err != nil {
				return "", err
			}
			for _, own := range rs.OwnerReferences {
				if own.Kind == "Deployment" {
					return own.Name, nil
				}
			}
		}
	}
	return "", fmt.Errorf("no deployment found for %v pod", pod.Name)
}

// it returns the uid of the chaos-operator deployment
func getOperatorUID(pod *core_v1.Pod, clients *kubernetes.Clientset) (string, error) {
	// derive the deployment name belongs to operator pod
	deployName, err := getDeploymentName(pod, clients)
	if err != nil {
		return "", err
	}

	deploy, err := clients.AppsV1().Deployments(pod.Namespace).Get(deployName, v1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("unable to get %s deployment in %s namespace", deployName, pod.Namespace)
	}
	if string(deploy.UID) == "" {
		return "", fmt.Errorf("unable to find the deployment uid")
	}
	return string(deploy.UID), nil
}