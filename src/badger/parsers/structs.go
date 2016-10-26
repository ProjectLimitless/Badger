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
 * Badger. If not, see http://www.apache.org/licenses/LICENSE-2.0.
 */

// Package parsers provides parsers for different CI providers
package parsers

import "time"

const (
	// ProviderStatusSuccess is the constant for success
	ProviderStatusSuccess = "Passing"
	// ProviderStatusFailed is the constant for failure
	ProviderStatusFailed = "Failing"
	// ProviderStatusUnknown is the constant for unknown statuses
	ProviderStatusUnknown = "Unknown"
)

// Parser interface defines the functionality required by a parser
type Parser interface {
	Parse(raw []byte) (ProviderResult, error)
	Name() string
}

// ProviderResult creats a standard result set for multiple CI tools
type ProviderResult struct {
	// The proper name of the CI tool that provided this result
	ProperName string
	Provider   string
	// The current status from the provider
	// as an integer 'ProviderStatusXXX'
	Status string
	// For templating, easy access field
	IsSuccess bool
	// Last commit user's name
	CommitUser string
	// The last commit message if the provider has it
	CommitMessage string
	// The last build time as provided
	BuildDateTime time.Time
	// Any error that occurred
	Error string
}
