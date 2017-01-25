package protoparser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	testProtoFile = "_examples/sample.proto"
)


func TestReadRPCs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	rpcs := ReadRPCs(testProtoFile)
	assert.Equal(t, "ProcedureCall_1",rpcs[0].Name)
	assert.False(t, rpcs[0].Input.IsStream)

	assert.Equal(t, "ProcedureCall_2",rpcs[1].Name)
	assert.False(t, rpcs[1].Input.IsStream)
}
