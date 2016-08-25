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

import (
	"os"
	"testing"

	parsers "."
)

var appVeyorParser parsers.Parser
var travisCIParser parsers.Parser
var appVeyorJson string
var travisCIJson string

func TestMain(m *testing.M) {
	appVeyorParser = &parsers.AppveyorParser{}
	travisCIParser = &parsers.TravisCIParser{}
	appVeyorJson = "{\"project\": {\"projectId\": 220088,\"accountId\": 44354,\"accountName\": \"donovansolms\",\"builds\": [],\"name\": \"ioRPC\",\"slug\": \"iorpc\",\"repositoryType\": \"gitHub\",\"repositoryScm\": \"git\",\"repositoryName\": \"ProjectLimitless/ioRPC\",\"repositoryBranch\": \"master\",\"isPrivate\": false,\"skipBranchesWithoutAppveyorYml\": false,\"enableSecureVariablesInPullRequests\": false,\"enableSecureVariablesInPullRequestsFromSameRepo\": false,\"enableDeploymentInPullRequests\": false,\"rollingBuilds\": false,\"alwaysBuildClosedPullRequests\": false,\"nuGetFeed\": {\"id\": \"iorpc-f1rq241u6kft\",\"name\": \"Project ioRPC\",\"publishingEnabled\": false,\"created\": \"2016-07-29T13:22:10.5478665+00:00\"},\"securityDescriptor\": {\"accessRightDefinitions\": [{\"name\": \"View\",\"description\": \"View\"},{\"name\": \"RunBuild\",\"description\": \"Run build\"},{\"name\": \"Update\",\"description\": \"Update settings\"},{\"name\": \"Delete\",\"description\": \"Delete project\"}],\"roleAces\": [{\"roleId\": 76364,\"name\": \"Administrator\",\"isAdmin\": true,\"accessRights\": [{\"name\": \"View\",\"allowed\": true},{\"name\": \"RunBuild\",\"allowed\": true},{\"name\": \"Update\",\"allowed\": true},{\"name\": \"Delete\",\"allowed\": true}]},{\"roleId\": 76365,\"name\": \"User\",\"isAdmin\": false,\"accessRights\": [{\"name\": \"View\"},{\"name\": \"RunBuild\"},{\"name\": \"Update\"},{\"name\": \"Delete\"}]}]},\"created\": \"2016-07-29T13:22:07.938561+00:00\",\"updated\": \"2016-08-25T09:44:16.0887202+00:00\"},\"build\": {\"buildId\": 4654641,\"jobs\": [{\"jobId\": \"x3k55m2x16hfi7c1\",\"name\": \"\",\"allowFailure\": false,\"messagesCount\": 0,\"compilationMessagesCount\": 17,\"compilationErrorsCount\": 0,\"compilationWarningsCount\": 17,\"testsCount\": 18,\"passedTestsCount\": 18,\"failedTestsCount\": 0,\"artifactsCount\": 1,\"status\": \"success\",\"started\": \"2016-08-25T11:03:34.1692307+00:00\",\"finished\": \"2016-08-25T11:04:23.3755601+00:00\",\"created\": \"2016-08-25T11:03:25.2931592+00:00\",\"updated\": \"2016-08-25T11:04:23.3755601+00:00\"}],\"buildNumber\": 32,\"version\": \"1.0.0.32\",\"message\": \"Clean up comments\",\"branch\": \"master\",\"isTag\": false,\"commitId\": \"48e98e50dbdc0a94a899f8c39baeb1f713183870\",\"authorName\": \"Donovan Solms\",\"authorUsername\": \"donovansolms\",\"committerName\": \"Donovan Solms\",\"committerUsername\": \"donovansolms\",\"committed\": \"2016-08-25T11:03:14+00:00\",\"messages\": [],\"status\": \"success\",\"started\": \"2016-08-25T11:03:34.184853+00:00\",\"finished\": \"2016-08-25T11:04:23.5318057+00:00\",\"created\": \"2016-08-25T11:03:22.808839+00:00\",\"updated\": \"2016-08-25T11:04:23.5318057+00:00\"}}"
	travisCIJson = "[{\"id\":155018968,\"repository_id\":9577945,\"number\":\"32\",\"state\":\"finished\",\"result\":0,\"started_at\":\"2016-08-25T11:05:46Z\",\"finished_at\":\"2016-08-25T11:07:04Z\",\"duration\":78,\"commit\":\"48e98e50dbdc0a94a899f8c39baeb1f713183870\",\"branch\":\"master\",\"message\":\"Clean up comments\",\"event_type\":\"push\"},{\"id\":155014619,\"repository_id\":9577945,\"number\":\"31\",\"state\":\"finished\",\"result\":0,\"started_at\":\"2016-08-25T10:45:32Z\",\"finished_at\":\"2016-08-25T10:46:58Z\",\"duration\":86,\"commit\":\"90efd0e524f0832bfa98cd02ceb63ed86990f147\",\"branch\":\"master\",\"message\":\"Enable XML documentation\",\"event_type\":\"push\"},{\"id\":155010910,\"repository_id\":9577945,\"number\":\"30\",\"state\":\"finished\",\"result\":0,\"started_at\":\"2016-08-25T10:23:58Z\",\"finished_at\":\"2016-08-25T10:25:30Z\",\"duration\":92,\"commit\":\"eb1abd7422457a1db5ea8f5d0bfc37a96cb4fffc\",\"branch\":\"master\",\"message\":\"Updated nuget project icon\",\"event_type\":\"push\"}]"
	os.Exit(m.Run())
}
