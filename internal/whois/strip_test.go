package whois

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripCommentsIISFormat(t *testing.T) {
	// IIS (.se/.nu) responses start with a "#" disclaimer whose text matches
	// the generic "registered" pattern in expiryRE before the real field does.
	body := strings.Join([]string{
		"# Copyright (c) 1997- The Swedish Internet Foundation.",
		"# All rights reserved.",
		"# Any use of this material to target advertising or",
		"# similar activities is forbidden and will be prosecuted.",
		"# Result of search for registered domain names under",
		"# the .se top level domain.",
		"domain:           example.se",
		"state:            active",
		"expires:          2027-03-22",
		"status:           ok",
	}, "\n")

	result := expiryRE.FindStringSubmatch(stripComments(body))
	require.GreaterOrEqual(t, len(result), 3)
	require.Equal(t, "2027-03-22", strings.TrimSpace(result[2]))
}

func TestStripCommentsKeepsRegularLines(t *testing.T) {
	body := "Domain Name: example.com\nRegistry Expiry Date: 2027-01-01T00:00:00Z\n"
	require.Equal(t, body, stripComments(body))
}
