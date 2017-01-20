package serverwriter

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/marcusljx/grpcgen-go/functions"
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
	outputRootPath      string
	ServerPackageString string
	Package             string
	PackagePath         string
	template            *template.Template
}

func NewServerWriter(outputPath, packagePath, templatesPath string) *ServerWriter {
	// Create Server Template
	tmpl := template.
		New("ServerWriter").
		Funcs(funcMap)

	// Read all serverwriter templates
	files, err := ioutil.ReadDir(templatesPath)
	functions.CheckFatal(err)

	for _, fInfo := range files {
		_, err := tmpl.ParseFiles(filepath.Join(templatesPath, fInfo.Name()))
		functions.CheckFatal(err)
	}

	return &ServerWriter{
		outputRootPath:      outputPath,
		ServerPackageString: "server",
		Package:             filepath.Base(packagePath),
		PackagePath:         packagePath,
		template:            tmpl,
	}
}

func (s *ServerWriter) Create() {
	// Setup File
	err := os.Mkdir(s.outputRootPath, os.ModePerm)
	functions.CheckFatal(err)

	s.CreateStartServerMainFile()
	s.CreateServerLogicStructFile()
	s.CreateServerLogicFuncFile()
}

func (s *ServerWriter) CreateServerLogicFuncFile() {
	//TODO
}

func (s *ServerWriter) CreateStartServerMainFile() {
	f, err := os.Create(filepath.Join(filepath.Dir(s.outputRootPath), "start_server.go"))
	functions.CheckFatal(err)
	defer f.Close()

	err = s.template.ExecuteTemplate(f, startServerMainFileTemplatePath, s)
	functions.CheckFatal(err)
}

func (s *ServerWriter) CreateServerLogicStructFile() {
	f, err := s.newFile("server_struct.go")
	functions.CheckFatal(err)
	defer f.Close()

	err = s.template.ExecuteTemplate(f, serverLogicStructTemplatePath, s)
	functions.CheckFatal(err)
}

//------------------------------------------- LOCAL FUNCTIONS
func (s *ServerWriter) newFile(filename string) (*os.File, error) {
	return os.Create(filepath.Join(s.outputRootPath, filename))
}
