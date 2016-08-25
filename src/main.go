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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sync"

	"./badger"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration JSON file")
	flag.Parse()
	fmt.Println("Using config file:", *configPath)

	// Check if the config file exists
	fileBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println("Unable to read config file:", err.Error())
		panic(err)
	}

	// Parse the config
	var config badger.Config
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		fmt.Println("Unable to parse config file:", err.Error())
		panic(err)
	}

	fmt.Println("Setting up Badger...")

	badgerBadger, err := badger.New(config)
	if err != nil {
		fmt.Println("Unable to create Badger instance: %s", err.Error())
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		badgerBadger.Start()
	}()

	fmt.Sprintln("Badger is running on %s:%d", config.Server.IP, config.Server.Port)
	waitGroup.Wait()

	fmt.Println("Shutdown Badger")
}
