package protorep

type Communication struct {
	IsStream bool
	Message  interface{}
}

type RPC struct {
	Name   string
	Input  Communication
	Output Communication
}

type Service struct {
	RPCs []RPC
}

type ProtoRep struct {
	Syntax  string
	Package string
	Options map[string]interface{}
	Service Service
}
