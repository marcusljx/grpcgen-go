package main

import (
	"flag"

	"log"

	"path/filepath"

	"github.com/marcusljx/grpcgen-go/serverwriter"
)

const (
	_templates_server = "serverwriter/templates"
)

var (
	serverPackagePath = flag.String("server-path", "./server", "Path to generate server package.")
	packagePath       = flag.String("go-package", "", "[REQUIRED] Import path for package containing .proto file. For some service 'AAA', this value is usually 'github.com/.../.../AAA/AAA'")
)

func main() {
	flag.Parse()
	log.Printf("packagePath = %s\n", *packagePath)

	serverTemplatesPath, _ := filepath.Abs(_templates_server)
	s := serverwriter.NewServerWriter(*serverPackagePath, *packagePath, serverTemplatesPath)
	s.Create()
}
