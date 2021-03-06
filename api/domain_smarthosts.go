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

// DomainSmartHost holds domain smarthosts
type DomainSmartHost struct {
	ID          int    `json:"id,omitempty" url:"id,omitempty"`
	Address     string `json:"address" url:"address"`
	Username    string `json:"username" url:"username"`
	Password    string `json:"password,omitempty" url:"password,omitempty"`
	Port        int    `json:"port" url:"port"`
	RequireTLS  bool   `json:"require_tls" url:"require_tls"`
	Enabled     bool   `json:"enabled" url:"enabled"`
	Description string `json:"description" url:"description"`
}

// DomainSmartHostList holds domain smarthosts
type DomainSmartHostList struct {
	Items []DomainSmartHost `json:"items"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

// GetDomainSmartHosts returns a DomainSmartHostList object
// This contains a paginated list of domain smarthosts and links
// to the neighbouring pages.
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#listing-domain-smarthosts
func (c *Client) GetDomainSmartHosts(domainID int, opts *ListOptions) (l *DomainSmartHostList, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	l = &DomainSmartHostList{}

	err = c.get(fmt.Sprintf("domains/smarthosts/%d", domainID), opts, l)

	return
}

// GetDomainSmartHost returns a domain smarthost
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#retrieve-a-domain-smarthost
func (c *Client) GetDomainSmartHost(domainID, serverID int) (server *DomainSmartHost, err error) {
	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if serverID <= 0 {
		err = fmt.Errorf(serverIDError)
		return
	}

	server = &DomainSmartHost{}

	err = c.get(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, serverID), nil, server)

	return
}

// CreateDomainSmartHost creates a domain smarthost
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#create-a-domain-smarthost
func (c *Client) CreateDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	v, _ = query.Values(server)

	err = c.post(fmt.Sprintf("domains/smarthosts/%d", domainID), v, server)

	return
}

// UpdateDomainSmartHost updates a domain smarthost
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#update-a-domain-smarthost
func (c *Client) UpdateDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.put(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, server.ID), v, server)

	return
}

// DeleteDomainSmartHost deletes a domain smarthost
//
// Baruwa API Docs: https://www.baruwa.com/docs/api/#delete-a-domain-smarthost
func (c *Client) DeleteDomainSmartHost(domainID int, server *DomainSmartHost) (err error) {
	var v url.Values

	if domainID <= 0 {
		err = fmt.Errorf(domainIDError)
		return
	}

	if server == nil {
		err = fmt.Errorf(serverParamError)
		return
	}

	if server.ID <= 0 {
		err = fmt.Errorf(serverSIDError)
		return
	}

	v, _ = query.Values(server)

	err = c.delete(fmt.Sprintf("domains/smarthosts/%d/%d", domainID, server.ID), v)

	return
}
