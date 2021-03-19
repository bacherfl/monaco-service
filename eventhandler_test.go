package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	keptn "github.com/keptn/go-utils/pkg/lib/keptn"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
)

/**
 * loads a cloud event from the passed test json file and initializes a keptn object with it
 */
func initializeTestObjects(eventFileName string) (*keptnv2.Keptn, *cloudevents.Event, error) {
	// load sample event
	eventFile, err := ioutil.ReadFile(eventFileName)
	if err != nil {
		return nil, nil, fmt.Errorf("Cant load %s: %s", eventFileName, err.Error())
	}

	incomingEvent := &cloudevents.Event{}
	err = json.Unmarshal(eventFile, incomingEvent)
	if err != nil {
		return nil, nil, fmt.Errorf("Error parsing: %s", err.Error())
	}

	// Add a Fake EventSender to KeptnOptions
	var keptnOptions = keptn.KeptnOpts{
		// 	EventSender: &fake.EventSender{},
	}
	keptnOptions.UseLocalFileSystem = true
	myKeptn, err := keptnv2.NewKeptn(incomingEvent, keptnOptions)

	return myKeptn, incomingEvent, err
}

// Tests HandleMonacoTriggeredEvent
// TODO: Add your test-code
func TestHandleMonacoTriggeredEvent(t *testing.T) {
	myKeptn, incomingEvent, err := initializeTestObjects("test-events/monaco.triggered.json")
	if err != nil {
		t.Error(err)
		return
	}

	specificEvent := &MonacoStartedEventData{}
	err = incomingEvent.DataAs(specificEvent)
	if err != nil {
		t.Errorf("Error getting keptn event data")
	}

	err = HandleMonacoTriggeredEvent(myKeptn, *incomingEvent, specificEvent)
	if err != nil {
		t.Errorf("Error: " + err.Error())
	}
}
