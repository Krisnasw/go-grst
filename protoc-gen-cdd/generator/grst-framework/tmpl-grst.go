package grstframework

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/compiler/protogen"
)

func increment(i int) int {
	i++
	return i
}

var (
	tmplGrpcRestHeader = template.Must(template.New("grst-header").Funcs(template.FuncMap{
		"ToTitle":     strings.Title,
		"ToCamelCase": strcase.ToCamel,
		"Increment":   increment,
	}).Parse(`
	// Code generated by protoc-gen-cdd. DO NOT EDIT.
	// source: {{.GetName}}
	{{- $pkgName := .GetPackage}}
	package {{$pkgName}}

	import (
		"net/http"
		"strings"

		"github.com/krisnasw/go-grst/grst"
		"google.golang.org/grpc"
		grst_errors "github.com/krisnasw/go-grst/grst/errors"
				
		"github.com/mcuadros/go-defaults"
		"google.golang.org/grpc/codes"
		"google.golang.org/grpc/credentials"
		"gopkg.in/validator.v2"
	)

	type fullMethods struct {
	{{ range $svc := .ServiceExt}}
	{{- range $mth := $svc.MethodExt}}
		{{ ToCamelCase $svc.GetName }}_{{$mth.GetName}} string
	{{- end}}
	{{ end}}
	}

	var FullMethods = fullMethods{
	{{- range $svc := .ServiceExt}}
	{{- range $mth := $svc.MethodExt}}
		{{ToCamelCase $svc.GetName}}_{{$mth.GetName}}: "/{{$pkgName}}.{{$svc.GetName}}/{{$mth.GetName}}",
	{{- end}}
	{{- end}}
	}

	var NeedAuthFullMethods = []string{
	{{- range $svc := .ServiceExt}}
	{{- range $mth := $svc.MethodExt}}
		{{if $mth.Auth.Needauth}} "/{{$pkgName}}.{{$svc.GetName}}/{{$mth.GetName}}", {{end}}
	{{- end}}
	{{- end}}
	}

	type AuthConfig struct {
		NeedAuth bool
		Roles    []string
	}
	var AuthConfigFullMethods = map[string]AuthConfig{
		{{- range $svc := .ServiceExt}}
		{{- range $mth := $svc.MethodExt}}
		"/{{$pkgName}}.{{$svc.GetName}}/{{$mth.GetName}}": AuthConfig{  NeedAuth: {{$mth.Auth.Needauth}}, Roles: []string{ {{- range $role := $mth.Auth.Roles}}  "{{$role}}",  {{- end}} }},
		{{- end}}
		{{- end}}
	}

	var NeedApiKeyFullMethods = []string{
		{{- range $svc := .ServiceExt}}
	{{- range $mth := $svc.MethodExt}}
		{{if $mth.Auth.Needapikey}} "/{{$pkgName}}.{{$svc.GetName}}/{{$mth.GetName}}", {{end}}
	{{- end}}
	{{- end}}
	}

	func ValidateRequest(req interface{}) error {
		defaults.SetDefaults(req)
		if errs := validator.Validate(req); errs != nil {
			validateError := []*grst_errors.ErrorDetail{}
			for field, err := range errs.(validator.ErrorMap) {
				errMessage := strings.Replace(err.Error(), "{field}", field, -1)
				validateError = append(validateError, &grst_errors.ErrorDetail{Code: 999, Field: field, Message: errMessage})
			}
			return grst_errors.New(http.StatusBadRequest, codes.InvalidArgument, 999, "Validation Error", validateError...)
		}
		
		return nil
	}	
	`))

	tmplGrstBody = template.Must(template.New("grst-body").Funcs(template.FuncMap{
		"ToTitle":     strings.Title,
		"ToCamelCase": strcase.ToCamel,
		"Increment":   increment,
	}).Parse(`
	{{$svcName := ToCamelCase .Service.GetName}}
	
	/*==================== {{$svcName}} Section ====================*/

	func Register{{$svcName}}GrstServer(grpcRestServer *grst.Server, hndl {{$svcName}}Server) {
	{{ range $m := .Service.MethodExt}}
		forward_{{$svcName}}_{{$m.GetName}}_0 = {{if $m.GetServerStreaming}}grpcRestServer.GetForwardResponseStream(){{else}}grpcRestServer.GetForwardResponseMessage(){{end}}
		{{ if $m.HttpRule }}
			{{ range $idx, $binding := $m.HttpRule.AdditionalBindings }}
			forward_{{$svcName}}_{{$m.GetName}}_{{Increment $idx}} = {{if $m.GetServerStreaming}}grpcRestServer.GetForwardResponseStream(){{else}}grpcRestServer.GetForwardResponseMessage(){{end}}
			{{ end }}
		{{ end}}
	{{ end}}
		Register{{$svcName}}Server(grpcRestServer.GetGrpcServer(), hndl)
		grpcRestServer.RegisterRestHandler(Register{{$svcName}}Handler)
	}

	func New{{$svcName}}GrstClient(serverHost string, creds *credentials.TransportCredentials, dialOpts ...grpc.DialOption) ({{$svcName}}Client, error) {
		opts := []grpc.DialOption{}
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)))
		opts = append(opts, grpc.WithMaxMsgSize(1024*1024*20))
		if creds == nil {
			opts = append(opts, grpc.WithInsecure())
		} else {
			opts = append(opts, grpc.WithTransportCredentials(*creds))
		}
		opts = append(opts, dialOpts...)
		grpcConn, err := grpc.Dial(serverHost, opts...)
		return New{{$svcName}}Client(grpcConn), err
	}
	`))
)

func applyTemplateGrpcRest(f *descriptor.FileDescriptorExt) (*generator.GeneratorResponseFile, error) {
	w := bytes.NewBuffer(nil)
	var tmplData = struct {
		*descriptor.FileDescriptorExt
	}{
		f,
	}

	if err := tmplGrpcRestHeader.Execute(w, tmplData); err != nil {
		return nil, err
	}

	for _, svc := range f.ServiceExt {
		var data = struct {
			File                          descriptor.FileDescriptorExt
			Service                       descriptor.ServiceDescriptorExt
			ServiceNameFirstLetterCapital string
		}{
			File:                          *f,
			Service:                       *svc,
			ServiceNameFirstLetterCapital: strings.Title(svc.GetName()),
		}

		if err := tmplGrstBody.Execute(w, data); err != nil {
			return nil, err
		}

	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generator.GeneratorResponseFile{
		Filename:     strings.ReplaceAll(f.GetName(), ".proto", "") + ".pb.cdd.go",
		Content:      string(formatted),
		GoImportPath: protogen.GoImportPath(""),
	}, nil
}
