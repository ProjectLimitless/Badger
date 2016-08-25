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

package parsers_test

import "testing"

func TestAppveyorName(t *testing.T) {
	expected := "AppVeyor"
	v := appVeyorParser.Name()
	if v != expected {
		t.Errorf("AppVeyor parser should set name to '%s' and not '%s'", expected, v)
	}
}

func TestAppveyorParse(t *testing.T) {
	parseResult, err := appVeyorParser.Parse([]byte(appVeyorJson))
	if err != nil {
		t.Errorf("Unable to parse AppVeyor JSON: %s", err.Error())
	}

	t.Run("IsSuccess", func(t *testing.T) {
		if parseResult.IsSuccess != true {
			t.Errorf("IsSuccess should be true")
		}
	})

	t.Run("Status", func(t *testing.T) {
		if parseResult.Status != "Passing" {
			t.Errorf("Status should be '%s' and not '%s'", "success", parseResult.Status)
		}
	})

	t.Run("CommitUser", func(t *testing.T) {
		if parseResult.CommitUser != "Donovan Solms" {
			t.Errorf("CommitUser should be '%s' and not '%s'", "Donovan Solms", parseResult.CommitUser)
		}
	})

	t.Run("CommitMessage", func(t *testing.T) {
		if parseResult.CommitMessage != "Clean up comments" {
			t.Errorf("CommitMessage should be '%s' and not '%s'", "Clean up comments", parseResult.CommitMessage)
		}
	})
}

func TestAppveyorParseInvalidJSON(t *testing.T) {
	_, err := appVeyorParser.Parse([]byte("{name:}"))
	if err == nil {
		t.Error("Parsing should have returned an error for invalid JSON")
	}
}
