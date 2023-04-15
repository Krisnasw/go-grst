package context

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
)

type ContextKey string

func (key ContextKey) String() string {
	return string(key)
}

const (
	CONTEXT_CLIENT_APPNAME         ContextKey = "appname"
	CONTEXT_CLIENT_APPVERSION      ContextKey = "appversion"
	CONTEXT_CLIENT_APPVERSIONCODE  ContextKey = "appversioncode"
	CONTEXT_CLIENT_MANUFACTURER    ContextKey = "manufacturer"
	CONTEXT_CLIENT_MODEL           ContextKey = "model"
	CONTEXT_CLIENT_PLATFORM        ContextKey = "platform"
	CONTEXT_CLIENT_PLATFORMVERSION ContextKey = "platformversion"
	CONTEXT_CLIENT_SDKVERSION      ContextKey = "sdkversion"
)

type ClientContext struct {
	AppName         string
	AppVersion      string
	AppVersionCode  string
	Manufacturer    string
	Model           string
	Platform        string
	PlatformVersion string
	SdkVersion      string
	UserAgent       string
}

func GetClientContext(ctx context.Context) (ClientContext, error) {
	result := ClientContext{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return result, fmt.Errorf("context is invalid: %s", fmt.Sprint(ctx))
	}

	result.AppName = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_APPNAME)
	result.AppVersion = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_APPVERSION)
	result.AppVersionCode = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_APPVERSIONCODE)
	result.Manufacturer = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_MANUFACTURER)
	result.Model = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_MODEL)
	result.Platform = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_PLATFORM)
	result.PlatformVersion = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_PLATFORMVERSION)
	result.SdkVersion = parseMetadataToString(md, "grst."+CONTEXT_CLIENT_SDKVERSION)
	result.UserAgent = parseMetadataToString(md, "user-agent") //user-agent is always forwarded, therefore it doesn't need prefix grst
	return result, nil
}

func parseMetadataToString(md metadata.MD, key ContextKey) string {
	result := md.Get(key.String())
	if len(result) > 0 {
		return result[0]
	}
	return ""

}
