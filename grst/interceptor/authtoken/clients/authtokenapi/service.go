package authtokenapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
)

const (
	timeout = 10
)

type client struct {
	baseURL           string
	validateTokenPath string
	httpClient        *http.Client
}

func NewClient(baseURL string, validateTokenPath string) AuthTokenClient {
	return &client{
		baseURL:           baseURL,
		validateTokenPath: validateTokenPath,
		httpClient:        &http.Client{Timeout: 10 * time.Second},
	}
}

type ValidateTokenReq struct {
	AccessToken string `json:"access_token"`
}

func (s *client) ValidateToken(reqBody io.Reader) (*SuccessResp, *ErrorResp) {
	req, err := http.NewRequest("POST", s.baseURL+"/"+s.validateTokenPath, reqBody)
	if err != nil {
		return nil, &ErrorResp{HTTPStatus: 500, GRPCStatus: codes.Unknown, Code: 98102, Message: err.Error()}
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, &ErrorResp{HTTPStatus: 500, GRPCStatus: codes.Unknown, Code: 98103, Message: err.Error()}
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &ErrorResp{HTTPStatus: 500, GRPCStatus: codes.Unknown, Code: 98104, Message: err.Error()}
	}

	var result SuccessResp
	if resp.StatusCode >= 400 {
		var errResult ErrorResp
		err = json.Unmarshal(respBody, &errResult)
		if err != nil {
			return nil, &ErrorResp{HTTPStatus: 500, GRPCStatus: codes.Unknown, Code: 98105, Message: err.Error()}
		}
		return nil, &errResult
	} else {
		err = json.Unmarshal(respBody, &result)
		if err != nil {
			return nil, &ErrorResp{HTTPStatus: 500, GRPCStatus: codes.Unknown, Code: 98106, Message: err.Error()}
		}
	}
	return &result, nil
}
