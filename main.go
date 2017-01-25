package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/marcusljx/grpcgen-go/functions"
	"github.com/marcusljx/grpcgen-go/serverwriter"
)

var (
	gopathOutputPath = flag.String("gopath-output", "", "[REQUIRED] Import path for package containing .proto file. For some service 'AAA', this value is usually 'github.com/.../.../AAA/AAA'")

	grpcgenROOT = functions.QualifyFromGopathSrc(filepath.Join("github.com", "marcusljx", "grpcgen-go"))
)

func main() {
	flag.Parse()
	serverTemplatesFullPath := filepath.Join(grpcgenROOT, "serverwriter", "templates")

	log.Printf("gopathOutputPath = %s\n", *gopathOutputPath)
	log.Printf("serverTemplatesFullPath = %s\n", serverTemplatesFullPath)
	s := serverwriter.NewServerWriter(*gopathOutputPath, serverTemplatesFullPath)
	s.Create()
}
