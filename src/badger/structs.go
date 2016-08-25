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
	ProviderTravisCI = "travisci"
	ProviderAppveyor = "appveyor"
)

type BadgeTemplates struct {
	Passing string `json:"Passing"`
	Failing string `json:"Failing"`
	Unknown string `json:"Unknown"`
}

type BadgeTemplateConfig struct {
	Background string         `json:"Background"`
	Badges     BadgeTemplates `json:"Badges"`
}

type OverlayPosition struct {
	Left int `json:"Left"`
	Top  int `json:"Top"`
}

type BadgeOverlay struct {
	Provider string          `json:"Provider"`
	Position OverlayPosition `json:"Position"`
}

type BadgeConfig struct {
	Template BadgeTemplateConfig `json:"Template"`
	Overlays []BadgeOverlay      `json:"Overlays"`
}

type PageConfig struct {
	Template string `json:"Template"`
}

type StatusConfig struct {
	Type     string `json:"Type"`
	Provider string `json:"Provider"`
	Url      string `json:"Url"`
}

type ProjectConfig struct {
	Name     string         `json:"Name"`
	Statuses []StatusConfig `json:"Statuses"`
	Badge    BadgeConfig    `json:"Badge"`
	Page     PageConfig     `json:"Page"`
}

type PageData struct {
	ProjectName string
	Overall     parsers.ProviderResult
	Providers   map[string]parsers.ProviderResult
}

type LogOutputConfig struct {
	Enabled      bool    `json:"Enabled"`
	Format       string  `json:"Format"`
	RotateSizeMB float64 `json:"RotateSizeMB"`
}

type LogConfig struct {
	Level   string          `json:"Level"`
	Console LogOutputConfig `json:"Console"`
	File    LogOutputConfig `json:"File"`
}

type ServerConfig struct {
	IP       string `json:"IP"`
	Port     int    `json:"Port"`
	BasePath string `json:"BasePath"`
}

type Config struct {
	Log    LogConfig    `json:"Log"`
	Server ServerConfig `json:"Server"`
}
