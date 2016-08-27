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

// Package badger provides the core functionality to run a Badger server
package badger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"image/draw"
	_ "image/jpeg"
	png "image/png"

	"./parsers"
	"github.com/donovansolms/lumberjack"
	"github.com/gorilla/mux"
	logging "github.com/op/go-logging"
	"github.com/tylerb/graceful"
)

type Badger struct {
	log         *logging.Logger
	router      *mux.Router
	bindAddress string
	Projects    map[string]ProjectConfig
	PagesPath   string
	BadgesPath  string
	cacheSince  string
	cacheUntil  string
}

// New creates a new instance of Badger
func New(config Config) (Badger, error) {

	// set up logging
	log := NewLog(config.Log)

	log.Debug("Setting up Badger...")
	var projectsPath string
	if config.ProjectsPath == "" {
		projectsPath = "projects"
	} else {
		var err error
		projectsPath, err = filepath.Abs(config.ProjectsPath)
		if err != nil {
			return Badger{}, errors.New("Unable to get project path: " + err.Error())
		}
	}

	if config.Server.IP == "" {
		return Badger{}, errors.New("You must specify a bind address")
	}
	if config.Server.Port == 0 {
		return Badger{}, errors.New("You must specify a bind port")
	}

	// Get all the project configs
	files, err := ioutil.ReadDir(projectsPath)
	if err != nil {
		return Badger{}, errors.New("Unable to open project path '" + projectsPath + "'")
	}
	badger := Badger{
		log:         log,
		bindAddress: fmt.Sprintf("%s:%d", config.Server.IP, config.Server.Port),
		Projects:    make(map[string]ProjectConfig),
		PagesPath:   "pages",
		BadgesPath:  "badges",
	}

	basePath := config.Server.BasePath

	router := mux.NewRouter()
	router.HandleFunc(basePath+"/", badger.RootHandler)
	router.HandleFunc(basePath+"/{project}", badger.ProjectPageHandler)
	router.HandleFunc(basePath+"/{project}/badge", badger.ProjectBadgeHandler)
	// serve CSS files directly
	cssServer := http.StripPrefix(basePath+"/css/", http.FileServer(http.Dir("./pages/css/")))
	router.PathPrefix(basePath + "/css/").Handler(cssServer)
	// serve JS files directly
	jsServer := http.StripPrefix(basePath+"/js/", http.FileServer(http.Dir("./pages/js/")))
	router.PathPrefix(basePath + "/js/").Handler(jsServer)
	// serve image files directly
	imgServer := http.StripPrefix(basePath+"/i/", http.FileServer(http.Dir("./pages/i/")))
	router.PathPrefix(basePath + "/i/").Handler(imgServer)

	badger.router = router

	// Parse all the project files
	log.Debug("Loading project files...")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".bbproj" {
			continue
		}

		fileBytes, err := ioutil.ReadFile(filepath.Join(projectsPath, file.Name()))
		if err != nil {
			log.Warning("Unable to read project file '%s': %s", filepath.Join(projectsPath, file.Name()), err.Error())
			continue
		}
		var projectConfig ProjectConfig
		err = json.Unmarshal(fileBytes, &projectConfig)
		if err != nil {
			log.Warning("Unable to parse project file '%s': %s", file.Name(), err.Error())
			continue
		}

		badger.Projects[strings.ToLower(projectConfig.Name)] = projectConfig
		log.Debug("Project '%s' loaded", projectConfig.Name)
	}
	// Proper english for config vs configs count
	if len(badger.Projects) == 0 {
		log.Error("No project configs loaded")
		return Badger{}, errors.New("No project configs loaded")
	} else if len(badger.Projects) == 1 {
		log.Info("Loaded %d project config", len(badger.Projects))
	} else {
		log.Info("Loaded %d project configs", len(badger.Projects))
	}

	badger.cacheSince = time.Now().Format(http.TimeFormat)
	badger.cacheUntil = time.Now().Add(time.Second * 60).Format(http.TimeFormat)

	return badger, nil
}

// Start starts the HTPP server
func (badger *Badger) Start() {
	// https://github.com/golang/go/issues/4674 so I use graceful
	// instead of http.ListenAndServe(badger.bindAddress, badger.router)
	graceful.Run(badger.bindAddress, 10*time.Second, badger.router)
}

// RootHandler handles calls to the root path and renders
// all the projects loaded on the root.html template
func (badger *Badger) RootHandler(w http.ResponseWriter, r *http.Request) {
	badger.log.Debug("Request received")
}

// ProjectPageHandler handles calls to /{project}
func (badger *Badger) ProjectPageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	project = strings.ToLower(project)
	badger.log.Debug("Request received for project page '%s'", project)

	if projectConfig, ok := badger.Projects[project]; ok {

		pagePath := filepath.Join(badger.PagesPath, project+".html")
		badger.log.Debug("Loading project page at %s", pagePath)

		page, err := template.ParseFiles(pagePath)
		if err != nil {
			badger.log.Warning("Project page not found: %s. Using default.", err.Error())
			// load the default page
			page, err = template.ParseFiles(filepath.Join(badger.PagesPath, "default.html"))
			if err != nil {
				badger.log.Error("Default page does not exist at '%s': %s", "default.html", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Default page could not be loaded '%s': %s", "default.html", err.Error())))
				return
			}
		}

		overallStatus, providerStatuses := FetchAllStatuses(projectConfig.Statuses)

		pageData := PageData{
			ProjectName: projectConfig.Name,
			Overall:     overallStatus,
			Providers:   providerStatuses,
		}

		err = page.Execute(w, pageData)
		if err != nil {
			badger.log.Error("Unable to execute template for '%'", project, err.Error())
		}
	} else {
		badger.log.Error("Project config not found for project '%s'", project)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Project config not found for project '%s'", project)))
		return
	}
}

// ProjectBadgeHandler handles calls to /{project}/badge
func (badger *Badger) ProjectBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	project = strings.ToLower(project)
	badger.log.Debug("Request received for project badge '%s'", project)

	if projectConfig, ok := badger.Projects[project]; ok {

		// Check if this project has an overlay section
		if len(projectConfig.Badge.Overlays) == 0 {
			badger.log.Error("No overlays found for project '%s'", project)
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(fmt.Sprintf("No overlays found for project '%s'", project)))
			return
		}

		backgroundPath := filepath.Join(badger.BadgesPath, projectConfig.Badge.Template.Background)
		badger.log.Debug("Building project badge from %s", backgroundPath)

		reader, err := os.Open(backgroundPath)
		if err != nil {
			badger.log.Warning("Project badge not found: %s. Using default.", err.Error())
			// load the default badge background
			reader, err = os.Open(filepath.Join(badger.BadgesPath, "default.png"))
			if err != nil {
				badger.log.Error("Default badge does not exist at '%s': %s", "default.png", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Default badge could not be loaded '%s': %s", "default.png", err.Error())))
				return
			}
		}
		backgroundImageRaw, imageType, err := image.Decode(reader)
		if err != nil {
			badger.log.Error("Error loading background image: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error loading background image: %s", err.Error())))
			return
		}
		badger.log.Debug("Background image loaded as type '%s'", imageType)
		bounds := backgroundImageRaw.Bounds()
		backgroundImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(backgroundImage, backgroundImage.Bounds(), backgroundImageRaw, bounds.Min, draw.Src)

		overallStatus, providerStatuses := FetchAllStatuses(projectConfig.Statuses)
		_ = overallStatus

		// map statuses to a map based on proper name
		providerStatusMap := make(map[string]parsers.ProviderResult)
		for _, status := range providerStatuses {
			badger.log.Debug("Mapping provider '%s'", status.ProperName)
			providerStatusMap[status.Provider] = status
		}
		badger.log.Debug("Overlays available: %d", len(projectConfig.Badge.Overlays))
		for _, overlay := range projectConfig.Badge.Overlays {
			if status, ok := providerStatusMap[overlay.Provider]; ok {
				badger.log.Debug("Overlaying provider '%s' status: %s", overlay.Provider, status.Status)
				var imageReader io.Reader
				switch status.Status {
				case parsers.ProviderStatusSuccess:
					imageReader, err = os.Open(filepath.Join(badger.BadgesPath, projectConfig.Badge.Template.Badges.Passing))
				case parsers.ProviderStatusFailed:
					imageReader, err = os.Open(filepath.Join(badger.BadgesPath, projectConfig.Badge.Template.Badges.Failing))
				case parsers.ProviderStatusUnknown:
					imageReader, err = os.Open(filepath.Join(badger.BadgesPath, projectConfig.Badge.Template.Badges.Unknown))
				}
				if err != nil {
					badger.log.Error("Unable to load status badge: %s", err.Error())
					continue
				}

				overlayImage, _, err := image.Decode(imageReader)
				if err != nil {
					badger.log.Error("Status badge error: %s", err.Error())
					continue
				}
				topLeft := image.Point{
					X: overlay.Position.Left,
					Y: overlay.Position.Top,
				}
				badgePlacement := overlayImage.Bounds().Add(topLeft)
				draw.Draw(backgroundImage, badgePlacement, overlayImage, image.ZP, draw.Over)
			} else {
				badger.log.Warning("Overlay provider '%s' not available in listed providers", overlay.Provider)
			}
		}
		img := image.Image(backgroundImage)
		w.Header().Set("Cache-Control", "no-cache, private")
		w.Header().Set("Last-Modified", badger.cacheSince)
		w.Header().Set("Expires", badger.cacheUntil)
		writeImage(badger.log, w, img)
		badger.log.Info("Badge rendered")

	} else {
		badger.log.Error("Project config not found for project '%s'", project)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Project config not found for project '%s'", project)))
		return
	}
}

// writeImage encodes an image 'img' in png format and writes it into ResponseWriter.
func writeImage(log *logging.Logger, w http.ResponseWriter, img image.Image) {
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		log.Error("Unable to encode image: %s", err.Error())
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		log.Error("Unable to write image to the HTTP output: %s", err.Error())
	}
}

// NewLog creates a new instance of the logger
func NewLog(logConfig LogConfig) *logging.Logger {

	// Collection of logging backends
	var backendFormatters []logging.Backend
	backendFormatters = make([]logging.Backend, 0)

	// If Console logging is enabled
	if logConfig.Console.Enabled == true {
		var format = logging.MustStringFormatter(
			logConfig.Console.Format,
		)
		backend := logging.NewLogBackend(os.Stderr, "", 0)
		backendFormatter := logging.NewBackendFormatter(backend, format)
		backendLeveled := logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "")
		backendFormatters = append(backendFormatters, backendLeveled)
	}
	if logConfig.File.Enabled == true {
		var format = logging.MustStringFormatter(
			logConfig.File.Format,
		)
		logfile := &lumberjack.Logger{
			Filename:   "logs/badger.log",
			MaxSize:    int(logConfig.File.RotateSizeMB),
			MaxBackups: 3,  // Number of files to keep
			MaxAge:     30, // Days to keep
		}
		backend := logging.NewLogBackend(logfile, "", 0)
		backendFormatter := logging.NewBackendFormatter(backend, format)
		backendLeveled := logging.AddModuleLevel(backendFormatter)
		backendLeveled.SetLevel(logging.DEBUG, "")
		backendFormatters = append(backendFormatters, backendLeveled)
	}

	logging.SetBackend(backendFormatters...)

	switch logConfig.Level {
	case "DEBUG":
		logging.SetLevel(logging.DEBUG, "")
	case "INFO":
		logging.SetLevel(logging.INFO, "")
	case "WARNING":
		logging.SetLevel(logging.WARNING, "")
	case "ERROR":
		logging.SetLevel(logging.ERROR, "")
	case "CRITICAL":
		logging.SetLevel(logging.CRITICAL, "")
	}

	log := logging.MustGetLogger("Badger")

	return log
}
