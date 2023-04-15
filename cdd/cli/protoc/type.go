package protoc

type Protoc interface {
	Exec(input string, printExecStatement bool) error
	AddProtoPath(...string)
	AddProtocGenOut(...ProtocGenOut)
}

type protoc struct {
	protoPaths    []string
	protocGenOuts []ProtocGenOut
}

type ProtobufVersion int

const ProtobufVersion1 ProtobufVersion = 1
const ProtobufVersion2 ProtobufVersion = 2

type ProtocGenOut struct {
	Name       string
	Opts       map[string]string
	OutputPath string
	Version    ProtobufVersion
}
