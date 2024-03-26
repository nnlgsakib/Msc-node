package main

import (
	_ "embed"

	"github.com/Mind-chain/mind/command/root"
	"github.com/Mind-chain/mind/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
