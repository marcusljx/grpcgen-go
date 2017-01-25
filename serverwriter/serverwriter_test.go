package serverwriter

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	templates_location, _ = filepath.Abs("templates")
)

func TestMain(m *testing.M) {
	// SetUp

	// Run Tests
	exitCode := m.Run()

	// TearDown

	// Finish
	os.Exit(exitCode)
}

func TestNewServerWriter(t *testing.T) {
	//s := NewServerWriter("github.com/testing/testPackage", templates_location)
	//assert.IsType(t, &ServerWriter{}, s, "Did not create type ServerWriter")
	//
	//assert.Equal(t, "github.com/testing/testPackage", s.ServiceRootFullPath)
	//assert.Equal(t, "testPackage", s.ServerLogicName)
	//assert.Equal(t, "server", s.ServerPackageString)
	//assert.Equal(t, "github.com/testing/testPackage", s.PackagePath)
	//assert.IsType(t, &template.Template{}, s.template)
}
