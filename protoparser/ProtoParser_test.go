package protoparser

import "testing"

const (
	fpath = "/home/marcuslow/GoPath/src/myGRPCserver/myGRPC/myGRPC.proto"
)

func TestReadFile(t *testing.T) {
	rpcs := ReadRPCs(fpath)
}
