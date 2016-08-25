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
	"errors"
	"time"
)

type TravisCIParser struct {
}

type TravisCIBuild struct {
	ID           int       `json:"id"`
	RepositoryID int       `json:"repository_id"`
	Number       string    `json:"number"`
	State        string    `json:"state"`
	Result       int       `json:"result"`
	StartedAt    time.Time `json:"started_at"`
	FinishedAt   time.Time `json:"finished_at"`
	Duration     int       `json:"duration"`
	Commit       string    `json:"commit"`
	Branch       string    `json:"branch"`
	Message      string    `json:"message"`
	EventType    string    `json:"event_type"`
}

// Parse parses the json bytes into a provider result
func (parser *TravisCIParser) Parse(raw []byte) (ProviderResult, error) {
	var result ProviderResult
	result.ProperName = parser.Name()
	result.Provider = "TravisCI"
	var builds []TravisCIBuild
	err := json.Unmarshal(raw, &builds)
	if err != nil {
		return result, err
	}

	if len(builds) > 0 {
		build := builds[0]
		switch build.Result {
		case 0:
			result.Status = ProviderStatusSuccess
			result.IsSuccess = true
		case 1:
			result.Status = ProviderStatusFailed
		default:
			result.Status = ProviderStatusUnknown
		}
		result.BuildDateTime = build.FinishedAt
		result.CommitMessage = build.Message
		result.CommitUser = "Unknown"
		return result, nil

	} else {
		return result, errors.New("No builds found for Travis CI")
	}
	return result, nil
}

// Name returns the Proper name of the provider for the parser
func (parser *TravisCIParser) Name() string {
	return "Travis CI"
}
