/*
Copyright 2019 hatech Authors
*/

package analytics

import (
	"fmt"

	ga "github.com/jpillora/go-ogle-analytics"
)

const (
	// ClientID contains TrackingID of the application
	clientID = "UA-92076314-21"

	// supported event categories

	// Category category notifies installation of a component of Litmus Infrastructure
	category = "Cnbr-Infra"

	// supported event actions

	// Action is sent when the installation is triggered
	action = "Installation"

	// supported event labels

	// Label denotes event is associated to which Litmus component
	label = "Chaos-Operator"
)

// TriggerAnalytics is responsible for sending out events
func TriggerAnalytics() error {
	client, err := ga.NewClient(clientID)
	if err != nil {
		return fmt.Errorf("new client generation failed, error : %s", err)
	}
	// sets the clientUUID to operator uid
	ClientUUID, err = getUID()
	if err != nil {
		return err
	}
	client.ClientID(ClientUUID)
	if err := client.Send(ga.NewEvent(category, action).Label(label)); err != nil {
		return fmt.Errorf("analytics event sending failed, error: %s", err)
	}
	return nil
}
