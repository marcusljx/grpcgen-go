package protoparser

import (
	"bufio"
	"log"
	"os"

	"regexp"

	"github.com/marcusljx/grpcgen-go/protorep"
)

const (
	rpcRegex = `rpc (\w+)\((\w+)\) returns \((\w+)\)`
)

var(
	defaultProtoParser = &ProtoParser{}
)

type ProtoParser struct {
	rpcs     []*protorep.RPC
}

func (p *ProtoParser) AddRPC(rpc *protorep.RPC) {
	p.rpcs = append(p.rpcs, rpc)
}

func (p *ProtoParser) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(rpcRegex)
	for {
		// Scan each line
		if ok := scanner.Scan(); ok {
			tokens := re.FindStringSubmatch(scanner.Text())
			if len(tokens) > 0 {
				// Formulate RPC object
				rpcObj := &protorep.RPC{
					Name: tokens[1],
					Input: protorep.Communication{ // Todo: stream not supported yet
						IsStream: false,
						Message:  tokens[2],
					},
					Output: protorep.Communication{
						IsStream: false,
						Message:  tokens[3],
					},
				}
				// Add to result
				p.AddRPC(rpcObj)
			}
		} else {
			break
		}
	}
}

func (p *ProtoParser) GetRPCs() ([]*protorep.RPC) {
	return p.rpcs
}

func ReadRPCs(filepath string) (rpcs []*protorep.RPC) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("ReadRPCs(%s)::%v", filepath, err)
	}
	defer f.Close()

	defaultProtoParser.Parse(f)
	return defaultProtoParser.GetRPCs()
}
