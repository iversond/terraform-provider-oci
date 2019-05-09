// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

// +build !record
// +build !replay

package httpreplay

import (
	"net/http"
)

// InstallRecorder does no-op.
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return client, nil
}

// SetScenario lets the recorder know which test is currently executing
func SetScenario(name string) error {
	debugLogf("Not recording. %s", name)
	return nil
}

// SaveScenario saves the recorded service calls for the current scenario
func SaveScenario() error {
	return nil
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return false
}
