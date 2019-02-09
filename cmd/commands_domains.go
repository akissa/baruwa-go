// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cmd cmdline client for the Baruwa REST API
package cmd

import (
	"fmt"
	"log"

	"github.com/baruwa-enterprise/baruwa-go/api"
	prettyjson "github.com/hokaccha/go-prettyjson"
	cli "github.com/jawher/mow.cli"
)

func domainShow(cmd *cli.Cmd) {
}

func domainCreate(cmd *cli.Cmd) {
}

func domainUpdate(cmd *cli.Cmd) {
}

func domainDelete(cmd *cli.Cmd) {
}

func domainsList(cmd *cli.Cmd) {
	cmd.Action = func() {
		var b []byte
		var err error
		var c *api.Client
		var d *api.DomainList

		if c, err = GetClient(); err != nil {
			log.Fatal(err)
		}

		if d, err = c.GetDomains(nil); err != nil {
			log.Fatal(err)
		}

		if len(d.Items) == 0 {
			fmt.Println()
			return
		}

		if b, err = prettyjson.Marshal(d); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", b)
	}
}
