package descriptor

import (
	"fmt"
	"log"
	"strings"

	cddext "github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/cdd/api"
	annotations "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MethodDescriptorExt struct {
	*descriptorpb.MethodDescriptorProto
	Repository   *DescriptorRepository
	ServiceExt   *ServiceDescriptorExt
	RequestType  *MessageDescriptorExt
	ResponseType *MessageDescriptorExt
	HttpRule     *annotations.HttpRule
	HttpMethod   string
	PathTemplate string
	Auth         *cddext.Auth
}

func (MethodDescriptorExt) New(svcext *ServiceDescriptorExt, method *descriptorpb.MethodDescriptorProto) *MethodDescriptorExt {
	mthext := &MethodDescriptorExt{
		MethodDescriptorProto: method,
		Repository:            svcext.FileExt.Repository,
		ServiceExt:            svcext,
		RequestType:           nil,
		ResponseType:          nil,
		HttpRule:              nil,
		HttpMethod:            "",
		PathTemplate:          "",
	}
	if msgext, ok := mthext.Repository.Messages[*mthext.InputType]; ok {
		mthext.RequestType = msgext
	}
	if msgext, ok := mthext.Repository.Messages[*mthext.OutputType]; ok {
		mthext.ResponseType = msgext
	}
	mthext.HttpRule, mthext.HttpMethod, mthext.PathTemplate = parseExtHTTPRule(method)
	mthext.Auth = parseExtAuth(method)
	if mthext.Auth == nil {
		mthext.Auth = &cddext.Auth{Needauth: false, Roles: []string{"*"}}
	} else if len(mthext.Auth.Roles) == 0 {
		mthext.Auth.Roles = []string{"*"}
	}
	return mthext
}

func (mthext *MethodDescriptorExt) GetIdentifier() string {
	return mthext.ServiceExt.GetIdentifier() + "." + mthext.GetName()
}

/*Parse annotation http rule*/
func parseExtHTTPRule(method *descriptorpb.MethodDescriptorProto) (httpRule *annotations.HttpRule, httpMethod string, pathTemplate string) {
	if method.Options == nil {
		return nil, "", ""
	} else if !proto.HasExtension(method.Options, annotations.E_Http) {
		return nil, "", ""
	}

	ext := proto.GetExtension(method.Options, annotations.E_Http)
	opts, ok := ext.(*annotations.HttpRule)
	if !ok {
		log.Println(fmt.Errorf("extension is %T; want an HttpRule", ext))
		return nil, "", ""
	}
	httpRule = opts
	httpMethod, pathTemplate = "", ""
	switch {
	case httpRule.GetGet() != "":
		httpMethod = "GET"
		pathTemplate = httpRule.GetGet()
	case httpRule.GetPut() != "":
		httpMethod = "PUT"
		pathTemplate = httpRule.GetPut()
	case httpRule.GetPost() != "":
		httpMethod = "POST"
		pathTemplate = httpRule.GetPost()
	case httpRule.GetPost() != "":
		httpMethod = "DELETE"
		pathTemplate = httpRule.GetDelete()
	case httpRule.GetPatch() != "":
		httpMethod = "PATCH"
		pathTemplate = httpRule.GetPatch()
	case httpRule.GetCustom() != nil:
		httpMethod = httpRule.GetCustom().Kind
		pathTemplate = httpRule.GetCustom().Path
	default:
	}
	pathTemplate = parsePathTemplateToGinFormat(pathTemplate)
	return
}

func parsePathTemplateToGinFormat(pathTemplate string) string {
	paths := strings.Split(pathTemplate, "/")

	for i, p := range paths {
		if len(p) <= 0 {
			continue
		} else if string(p[0]) == "{" && string(p[len(p)-1]) == "}" {
			paths[i] = strings.Replace(paths[i], "{", "", -1)
			paths[i] = strings.Replace(paths[i], "}", "", -1)
			paths[i] = strings.ToLower(paths[i])
			paths[i] = ":" + paths[i]
		}
	}

	return strings.Join(paths, "/")
}

func parseExtAuth(method *descriptorpb.MethodDescriptorProto) *cddext.Auth {
	if method.Options == nil {
		return nil
	} else if !proto.HasExtension(method.Options, cddext.E_Auth) {
		return nil
	}

	ext := proto.GetExtension(method.Options, cddext.E_Auth)
	opts, ok := ext.(*cddext.Auth)
	if !ok {
		log.Println(fmt.Errorf("[parseExtAuth] extension is %T; want an Auth", ext))
		return nil
	}
	return opts
}
