// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main cmdline client for the Baruwa REST API
package main

import (
	"os"

	"github.com/baruwa-enterprise/baruwa-go/cmd"
)

func main() {
	cli := cmd.NewCLI()
	cli.RegisterCommands()
	cli.Run(os.Args)
}
