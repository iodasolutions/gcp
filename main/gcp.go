package main

import (
	"github.com/iodasolutions/gcp"
	"github.com/iodasolutions/xbee-common/provider"
)

func main() {
	var p gcp.Provider
	var a gcp.Admin
	provider.Execute(p, a)
}
