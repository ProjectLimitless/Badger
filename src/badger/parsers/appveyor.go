/**
 * This file is part of Badger.
 * Copyright © 2016 Donovan Solms.
 * Project Limitless
 * https://www.projectlimitless.io
 *
 * Badger and Project Limitless is free software: you can redistribute it and/or modify
 * it under the terms of the Apache License Version 2.0.
 *
 * You should have received a copy of the Apache License Version 2.0 with
 * Badger. If not, see <http://www.apache.org/licenses/LICENSE-2.0>.
 */

package parsers

import (
	"encoding/json"
	"strings"
	"time"
)

type AppveyorParser struct {
}

type AppveyorData struct {
	Build struct {
		BuildID           int       `json:"buildId"`
		BuildNumber       int       `json:"buildNumber"`
		Version           string    `json:"version"`
		Message           string    `json:"message"`
		Branch            string    `json:"branch"`
		IsTag             bool      `json:"isTag"`
		CommitID          string    `json:"commitId"`
		CommitterName     string    `json:"committerName"`
		CommitterUsername string    `json:"committerUsername"`
		Committed         time.Time `json:"committed"`
		Status            string    `json:"status"`
		Started           time.Time `json:"started"`
		Finished          time.Time `json:"finished"`
		Created           time.Time `json:"created"`
		Updated           time.Time `json:"updated"`
	} `json:"build"`
}

// Parse parses the json bytes into a provider result
func (parser *AppveyorParser) Parse(raw []byte) (ProviderResult, error) {
	var result ProviderResult
	result.ProperName = parser.Name()
	result.Provider = "AppVeyor"
	var data AppveyorData
	err := json.Unmarshal(raw, &data)
	if err != nil {
		return result, err
	}

	switch strings.ToLower(data.Build.Status) {
	case "success":
		result.Status = ProviderStatusSuccess
		result.IsSuccess = true
	case "failed":
		result.Status = ProviderStatusFailed
	default:
		result.Status = ProviderStatusUnknown
	}
	result.BuildDateTime = data.Build.Finished
	result.CommitMessage = data.Build.Message
	result.CommitUser = data.Build.CommitterName

	return result, nil
}

// Name returns the Proper name of the provider for the parser
func (parser *AppveyorParser) Name() string {
	return "AppVeyor"
}
