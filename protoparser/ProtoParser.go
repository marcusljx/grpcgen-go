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

func ReadRPCs(filepath string) (rpcs []*protorep.RPC) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("ReadRPCs(%s)::%v", filepath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	if err != nil {
		log.Fatalf("ReadRPCs(%s)::%v", filepath, err)
	}

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
				rpcs = append(rpcs, rpcObj)
			}
		} else {
			break
		}
	}

	return
}
