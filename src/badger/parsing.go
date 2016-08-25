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

package badger

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"./parsers"
)

// New creates a new instance of the specified parser
func NewParser(parserType string) (parsers.Parser, error) {
	switch strings.ToLower(parserType) {
	case ProviderTravisCI:
		return &parsers.TravisCIParser{}, nil
	case ProviderAppveyor:
		return &parsers.AppveyorParser{}, nil
	default:
		return nil, errors.New("No parser found for " + parserType)
	}
}

// FetchStatus fetches the current status from a provider and returns
// the parsed result
func FetchStatus(status StatusConfig) (parsers.ProviderResult, error) {
	var result parsers.ProviderResult
	parser, err := NewParser(status.Provider)
	if err != nil {
		return result, err
	}
	result.ProperName = parser.Name()
	result.IsSuccess = false

	// Get the content from the URL
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest("GET", status.Url, nil)
	if err != nil {
		return result, err
	}
	// Set the accept header so that we get JSON results
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return result, errors.New("Unable to fetch status")
	}
	if response.StatusCode != http.StatusOK {
		return result, errors.New("Unable to fetch status: '" + status.Provider + "':" + response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if len(body) == 0 {
		return result, errors.New("Data provided to parse is blank")
	}

	result, err = parser.Parse(body)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FetchAllStatuses fetches all provider statuses by calling FetchStatus for
// each statuses provided. Does not return an error, errors are inserted into
// returned provider results.
func FetchAllStatuses(statuses []StatusConfig) (parsers.ProviderResult, map[string]parsers.ProviderResult) {
	providerStatuses := make(map[string]parsers.ProviderResult)
	overallStatus := parsers.ProviderResult{
		ProperName: "Overall",
		Status:     parsers.ProviderStatusSuccess,
	}
	for _, status := range statuses {
		providerStatus, err := FetchStatus(status)
		if err != nil {
			providerStatus.Status = parsers.ProviderStatusUnknown
			providerStatus.Error = err.Error()
			overallStatus.Status = parsers.ProviderStatusFailed
		} else {
			if providerStatus.Status != parsers.ProviderStatusSuccess {
				overallStatus.Status = parsers.ProviderStatusFailed
			}
		}
		// Always add
		providerStatuses[strings.ToLower(status.Provider)] = providerStatus
	}
	return overallStatus, providerStatuses
}
