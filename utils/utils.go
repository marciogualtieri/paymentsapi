package utils

import (
	"os"
	"path"
)

/*
GetProjectPath returns the full path of the project.
*/
func GetProjectPath() string {
	goPath := os.Getenv("GOPATH")
	return path.Join(goPath, "src/github.com/marciogualtieri/paymentsapi")
}
