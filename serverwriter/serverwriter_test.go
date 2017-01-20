package serverwriter

import (
	"os"
	"testing"

	"github.com/alecthomas/template"
	"github.com/stretchr/testify/assert"
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
	s := NewServerWriter("output", "github.com/testing/testPackage", "./serverwriter/templates")
	assert.IsType(t, &ServerWriter{}, s, "Did not create type ServerWriter")

	assert.Equal(t, "output", s.outputRootPath)
	assert.Equal(t, "testPackage", s.Package)
	assert.Equal(t, "server", s.ServerPackageString)
	assert.Equal(t, "github.com/testing/testPackage", s.PackagePath)
	assert.IsType(t, &template.Template{}, s.template)
}
