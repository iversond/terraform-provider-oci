// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

// +build replay

package httpreplay

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var recorder *Recorder

// SetScenario creates a new recorder for this scenario
func SetScenario(name string) error {
	var err error
	if recorder, err = NewRecorderAsMode(name, ModeReplaying); err == nil {
		// cleanup existing files in /tmp folder
		RemoveContents("/tmp")
		recorder.SetMatcher(matcher)
		recorder.SetTransformer(transformer)
	}
	return err
}

// SaveScenario does nothing when recording
func SaveScenario() error {
	recorder = nil
	return nil
}

// InstallRecorder puts the recording transport into the http client, then returns a type that is compatible with the SDK's HTTPRequestDispatcher
func InstallRecorder(client *http.Client) (HTTPRecordingClient, error) {
	return InstallRecorderForRecodReplay(client, recorder)
}

// ShouldRetryImmediately returns true if replaying
func ShouldRetryImmediately() bool {
	return true
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		if strings.Contains(name, ".yaml") {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
