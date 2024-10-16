// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build windows

package runtime

import (
	"crypto/sha256"
	"fmt"

	"github.com/elastic/elastic-agent/internal/pkg/agent/application/info"
	"github.com/elastic/elastic-agent/internal/pkg/agent/application/paths"
)

func getShipperAddr(componentID string) string {
	// when installed the address is fixed to a location
	if info.RunningInstalled() {
		return fmt.Sprintf(paths.ShipperSocketPipePattern, componentID)
	}

	// not install, adjust the path based on data path
	data := paths.Data()
	// entire string cannot be longer than 256 characters, this forces the
	// length to always be 87 characters (but unique per data path)
	return fmt.Sprintf(`\\.\pipe\elastic-agent-%x-%s-pipe`, sha256.Sum256([]byte(data)), componentID)
}
