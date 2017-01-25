package serverwriter

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"log"

	"fmt"

	"github.com/marcusljx/grpcgen-go/functions"
	"github.com/marcusljx/grpcgen-go/protoparser"
	"github.com/marcusljx/grpcgen-go/protorep"
)

const (
	startServerMainFileTemplatePath = "start_server_main.tmpl"
	serverLogicStructTemplatePath   = "server_logic_struct.tmpl"
	serverLogicFuncTemplatePath     = "server_logic_func.tmpl"
)

var (
	funcMap = template.FuncMap{
		"func_ToTitleCase": strings.Title,
		"func_PathDir":     filepath.Dir,
	}
)

type ServerWriter struct {
	ServiceRootFullPath string
	ServiceName         string

	ServerLogicName       string
	ServerPackageFullPath string
	ServerImportPath      string

	ProtoPackageFullPath   string
	ProtoPackageImportPath string
	ProtoFileFullPath      string

	Functions []*protorep.RPC
	template  *template.Template
}

func createTemplateObject(templatesDirPath string) *template.Template {
	// Create Server Template
	tmpl := template.
		New("ServerWriter").
		Funcs(funcMap)
	// Read all serverwriter templates
	files, err := ioutil.ReadDir(templatesDirPath)
	functions.CheckFatal(err)
	for _, fInfo := range files {
		_, err := tmpl.ParseFiles(filepath.Join(templatesDirPath, fInfo.Name()))
		functions.CheckFatal(err)
	}
	return tmpl
}

func NewServerWriter(gopathOutputPath, serverTemplatesFullPath string) *ServerWriter {
	serviceName := filepath.Base(gopathOutputPath)
	serviceRootFullPath := functions.QualifyFromGopathSrc(gopathOutputPath)

	protoPackageImportPath := filepath.Join(gopathOutputPath, serviceName)
	protoPackageFullPath := functions.QualifyFromGopathSrc(protoPackageImportPath)
	protoFileFullPath := filepath.Join(protoPackageFullPath, fmt.Sprintf("%s.%s", serviceName, "proto"))

	serverImportPath := filepath.Join(gopathOutputPath, "server")
	serverPackageFullPath := functions.QualifyFromGopathSrc(serverImportPath)

	return &ServerWriter{
		ServiceRootFullPath:    serviceRootFullPath,
		ServiceName:            serviceName,
		ServerPackageFullPath:  serverPackageFullPath,
		ServerImportPath:       serverImportPath,
		ProtoPackageFullPath:   protoPackageFullPath,
		ProtoPackageImportPath: protoPackageImportPath,
		ProtoFileFullPath:      protoFileFullPath,
		template:               createTemplateObject(serverTemplatesFullPath),
		Functions:              protoparser.ReadRPCs(protoFileFullPath),
	}
}

func (s *ServerWriter) Create() {
	// Setup File
	err := os.Mkdir(s.ServerPackageFullPath, os.ModePerm)
	if err != nil {
		if os.IsExist(err) {
			log.Printf("WARNING: server package already exists. It will be overwritten.")
		}
	}

	s.CreateStartServerMainFile()
	s.CreateServerLogicStructFile()
	s.CreateServerLogicFuncFile()
}

func (s *ServerWriter) CreateServerLogicFuncFile() {
	f, err := s.newServerPackageFile("server_logic.go")
	functions.CheckFatal(err)
	defer f.Close()

	err = s.template.ExecuteTemplate(f, serverLogicFuncTemplatePath, s)

}

func (s *ServerWriter) CreateStartServerMainFile() {
	f, err := os.Create(filepath.Join(s.ServiceRootFullPath, "start_server.go"))
	functions.CheckFatal(err)
	defer f.Close()

	err = s.template.ExecuteTemplate(f, startServerMainFileTemplatePath, s)
	functions.CheckFatal(err)
}

func (s *ServerWriter) CreateServerLogicStructFile() {
	f, err := s.newServerPackageFile("server_struct.go")
	functions.CheckFatal(err)
	defer f.Close()

	err = s.template.ExecuteTemplate(f, serverLogicStructTemplatePath, s)
	functions.CheckFatal(err)
}

//------------------------------------------- LOCAL FUNCTIONS
func (s *ServerWriter) newServerPackageFile(filename string) (*os.File, error) {
	return os.Create(filepath.Join(s.ServerPackageFullPath, filename))
}
