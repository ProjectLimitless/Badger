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

func TestTravisCIName(t *testing.T) {
	expected := "Travis CI"
	v := travisCIParser.Name()
	if v != expected {
		t.Errorf("TravisCI parser should set name to '%s' and not '%s'", expected, v)
	}
}

func TestTravisCIParse(t *testing.T) {
	parseResult, err := travisCIParser.Parse([]byte(travisCIJson))
	if err != nil {
		t.Errorf("Unable to parse TravisCI JSON: %s", err.Error())
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
		// Travis CI doesn't return the commit user
		if parseResult.CommitUser != "Unknown" {
			t.Errorf("CommitUser should be '%s' and not '%s'", "Unknown", parseResult.CommitUser)
		}
	})

	t.Run("CommitMessage", func(t *testing.T) {
		if parseResult.CommitMessage != "Clean up comments" {
			t.Errorf("CommitMessage should be '%s' and not '%s'", "Clean up comments", parseResult.CommitMessage)
		}
	})
}

func TestTravisCIParseInvalidJSON(t *testing.T) {
	_, err := travisCIParser.Parse([]byte("{name:}"))
	if err == nil {
		t.Error("Parsing should have returned an error for invalid JSON")
	}
}
