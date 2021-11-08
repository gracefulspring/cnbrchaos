/*
Copyright 2021 hatech Authors
*/

package utils

import (
	corev1 "k8s.io/api/core/v1"
)

// RemoveString removes a particular string from a slice of strings
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// SetEnv sets the env inside envDetails struct
func (envDetails *ENVDetails) SetEnv(key, value string) *ENVDetails {
	if value != "" {
		envDetails.ENV = append(envDetails.ENV, corev1.EnvVar{
			Name:  key,
			Value: value,
		})
	}
	return envDetails
}
