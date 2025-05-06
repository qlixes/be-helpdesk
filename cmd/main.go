package main

import (
	"fmt"

	"github.com/qlixes/helpdesk/internal/provider"
)

func main() {

	fmt.Println(provider.Cfg.App.Host)
}
