//go:build tools
// +build tools

package main

import (
	_ "github.com/ethereum/go-ethereum/cmd/abigen"
	_ "github.com/onsi/ginkgo/v2/ginkgo"
)
