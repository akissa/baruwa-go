// BaruwaAPI Golang bindings for Baruwa REST API
// Copyright (C) 2019 Andrew Colin Kissa <andrew@topdog.za.net>

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package api

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// AliasAddress hosts alias addresses
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#alias-addresses
type AliasAddress struct {
	ID      int    `json:"id,omitempty" url:"id,omitempty"`
	Address string `json:"address" url:"address"`
	Enabled bool   `json:"enabled" url:"enabled"`
}

// GetAliasAddress returns an alias address
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-an-existing-alias-address
func (c *Client) GetAliasAddress(aliasID int) (alias *AliasAddress, err error) {
	if aliasID <= 0 {
		err = fmt.Errorf(aliasIDError)
		return
	}

	alias = &AliasAddress{}

	err = c.get(fmt.Sprintf("aliasaddresses/%d", aliasID), nil, alias)

	return
}

// CreateAliasAddress creates an alias address
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-an-alias-address
func (c *Client) CreateAliasAddress(userID int, alias *AliasAddress) (err error) {
	var v url.Values

	if userID <= 0 {
		err = fmt.Errorf(userIDError)
		return
	}

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	v, _ = query.Values(alias)

	err = c.post(fmt.Sprintf("aliasaddresses/%d", userID), v, alias)

	return
}

// UpdateAliasAddress updates an alias address
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-an-alias-address
func (c *Client) UpdateAliasAddress(alias *AliasAddress) (err error) {
	var v url.Values

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if alias.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	v, _ = query.Values(alias)

	err = c.put(fmt.Sprintf("aliasaddresses/%d", alias.ID), v, nil)

	return
}

// DeleteAliasAddress deletes an alias address
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-an-alias-address
func (c *Client) DeleteAliasAddress(alias *AliasAddress) (err error) {
	var v url.Values

	if alias == nil {
		err = fmt.Errorf(aliasParamError)
		return
	}

	if alias.ID <= 0 {
		err = fmt.Errorf(aliasSIDError)
		return
	}

	v, _ = query.Values(alias)

	err = c.delete(fmt.Sprintf("aliasaddresses/%d", alias.ID), v)

	return
}
