// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Package cmd cmdline client for the Baruwa REST API
package cmd

import (
	cli "github.com/jawher/mow.cli"
)

// RegisterCommands registers all CLI commands
func (c *CLI) RegisterCommands() {
	// user
	c.Command("user", "manage user accounts", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of a user account", userShow)
		cmd.Command("create", "create a new user account", userCreate)
		cmd.Command("update", "update a user account", userUpdate)
		cmd.Command("delete", "delete a user account", userDelete)
		cmd.Command("alias", "manage user alias addresses", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of an alias address", aliasShow)
			cmd.Command("create", "create a new alias address", aliasCreate)
			cmd.Command("update", "update an alias address", aliasUpdate)
			cmd.Command("delete", "delete an alias address", aliasDelete)
		})
	})
	// users
	c.Command("users", "list user accounts", usersList)
	// domain
	c.Command("domain", "manage domains", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of a domain", domainShow)
		cmd.Command("create", "create a new domain", domainCreate)
		cmd.Command("update", "update a domain", domainUpdate)
		cmd.Command("delete", "delete a domain", domainDelete)
		cmd.Command("alias", "manage alias domains", func(cmd *cli.Cmd) {
			cmd.Command("show", "show detailed information of a domain alias", domainAliasShow)
			cmd.Command("create", "create a new domain alias", domainAliasCreate)
			cmd.Command("update", "update a domain alias", domainAliasUpdate)
			cmd.Command("delete", "delete a domain alias", domainAliasDelete)
		})
	})
	// domains
	c.Command("domains", "list domains", domainsList)
	// organization
	c.Command("organization", "manage organizations", func(cmd *cli.Cmd) {
		cmd.Command("show", "show detailed information of an organization", organizationShow)
		cmd.Command("create", "create a new organization", organizationCreate)
		cmd.Command("update", "update a organization", organizationUpdate)
		cmd.Command("delete", "delete a organization", organizationDelete)
	})
	// organizations
	c.Command("organizations", "list organizations", organizationsList)
	// systemstatus
	c.Command("systemstatus", "show system status", systemStatus)
}
