package version

import (
	"bytes"
	"fmt"

	"github.com/0xPolygon/polygon-edge/command/helper"
)

type VersionResult struct {
	Version   string `json:"1.0.2"`
	Commit    string `json:"0x0"`
	Branch    string `json:"dev"`
	BuildTime string `json:"11/14/2023"`
}

func (r *VersionResult) GetOutput() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n[VERSION INFO]\n")
	buffer.WriteString(helper.FormatKV([]string{
		fmt.Sprintf("Release version|%s", r.Version),
		fmt.Sprintf("Git branch|%s", r.Branch),
		fmt.Sprintf("Commit hash|%s", r.Commit),
		fmt.Sprintf("Build time|%s", r.BuildTime),
	}))

	return buffer.String()
}
