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

package badger

import "./parsers"

const (
	// ProviderTravisCI is the constant for Travis CI
	ProviderTravisCI = "travisci"
	// ProviderAppveyor is the constant for AppVeyor
	ProviderAppveyor = "appveyor"
)

// BadgeTemplates is the structure for the template JSON
type BadgeTemplates struct {
	Passing string `json:"Passing"`
	Failing string `json:"Failing"`
	Unknown string `json:"Unknown"`
}

// BadgeTemplateConfig is the structure for the template config JSON
type BadgeTemplateConfig struct {
	Background string         `json:"Background"`
	Badges     BadgeTemplates `json:"Badges"`
}

// OverlayPosition is the structure for the template opverlay position JSON
type OverlayPosition struct {
	Left int `json:"Left"`
	Top  int `json:"Top"`
}

// BadgeOverlay is the structure for specifying an overlay
type BadgeOverlay struct {
	Provider string          `json:"Provider"`
	Position OverlayPosition `json:"Position"`
}

// BadgeConfig is the configuration for a specific badge
type BadgeConfig struct {
	Template BadgeTemplateConfig `json:"Template"`
	Overlays []BadgeOverlay      `json:"Overlays"`
}

// PageConfig is the configuration for a badge's page
type PageConfig struct {
	Template string `json:"Template"`
}

// StatusConfig provides the structure for status configuration
type StatusConfig struct {
	Type     string `json:"Type"`
	Provider string `json:"Provider"`
	URL      string `json:"Url"`
}

// ProjectConfig is the JSON structure for project configurations
type ProjectConfig struct {
	Name     string         `json:"Name"`
	Statuses []StatusConfig `json:"Statuses"`
	Badge    BadgeConfig    `json:"Badge"`
	Page     PageConfig     `json:"Page"`
}

// PageData is teh setup for a project page
type PageData struct {
	ProjectName string
	Overall     parsers.ProviderResult
	Providers   map[string]parsers.ProviderResult
}

// LogOutputConfig sets up the output formats for log files
type LogOutputConfig struct {
	Enabled      bool    `json:"Enabled"`
	Format       string  `json:"Format"`
	RotateSizeMB float64 `json:"RotateSizeMB"`
}

// LogConfig specifies the base setup for loggin
type LogConfig struct {
	Level   string          `json:"Level"`
	Console LogOutputConfig `json:"Console"`
	File    LogOutputConfig `json:"File"`
}

// ServerConfig is the setup for the HTTP server
type ServerConfig struct {
	IP       string `json:"IP"`
	Port     int    `json:"Port"`
	BasePath string `json:"BasePath"`
}

// Config is the general configuration for badger
type Config struct {
	Log          LogConfig    `json:"Log"`
	Server       ServerConfig `json:"Server"`
	ProjectsPath string       `json:"ProjectsPath"`
}
