// Package user provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/discord-gophers/goapi-gen version v0.3.0 DO NOT EDIT.
package user

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/discord-gophers/goapi-gen/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// PostUserJSONBody defines parameters for PostUser.
type PostUserJSONBody struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// PutUserUUIDJSONBody defines parameters for PutUserUUID.
type PutUserUUIDJSONBody struct {
	Name *string `json:"name,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody PostUserJSONBody

// Bind implements render.Binder.
func (PostUserJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// PutUserUUIDJSONRequestBody defines body for PutUserUUID for application/json ContentType.
type PutUserUUIDJSONRequestBody PutUserUUIDJSONBody

// Bind implements render.Binder.
func (PutUserUUIDJSONRequestBody) Bind(*http.Request) error {
	return nil
}

// Response is a common response struct for all the API calls.
// A Response object may be instantiated via functions for specific operation responses.
// It may also be instantiated directly, for the purpose of responding with a single status code.
type Response struct {
	body        interface{}
	Code        int
	contentType string
}

// Render implements the render.Renderer interface. It sets the Content-Type header
// and status code based on the response definition.
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", resp.contentType)
	render.Status(r, resp.Code)
	return nil
}

// Status is a builder method to override the default status code for a response.
func (resp *Response) Status(code int) *Response {
	resp.Code = code
	return resp
}

// ContentType is a builder method to override the default content type for a response.
func (resp *Response) ContentType(contentType string) *Response {
	resp.contentType = contentType
	return resp
}

// MarshalJSON implements the json.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(resp.body)
}

// MarshalXML implements the xml.Marshaler interface.
// This is used to only marshal the body of the response.
func (resp *Response) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.Encode(resp.body)
}

// GetUserJSON200Response is a constructor method for a GetUser response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserJSON200Response(body []struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	UUID  *string `json:"uuid,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PostUserJSON200Response is a constructor method for a PostUser response.
// A *Response is returned with the configured status code and content type from the spec.
func PostUserJSON200Response(body struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	UUID  *string `json:"uuid,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// GetUserUUIDJSON200Response is a constructor method for a GetUserUUID response.
// A *Response is returned with the configured status code and content type from the spec.
func GetUserUUIDJSON200Response(body struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	UUID  *string `json:"uuid,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// PutUserUUIDJSON200Response is a constructor method for a PutUserUUID response.
// A *Response is returned with the configured status code and content type from the spec.
func PutUserUUIDJSON200Response(body struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
	UUID  *string `json:"uuid,omitempty"`
}) *Response {
	return &Response{
		body:        body,
		Code:        200,
		contentType: "application/json",
	}
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of users
	// (GET /user)
	GetUser(w http.ResponseWriter, r *http.Request) *Response
	// Create a new user
	// (POST /user)
	PostUser(w http.ResponseWriter, r *http.Request) *Response
	// Delete user by UUID
	// (DELETE /user/{uuid})
	DeleteUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response
	// Get user details by UUID
	// (GET /user/{uuid})
	GetUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response
	// Update user details by UUID
	// (PUT /user/{uuid})
	PutUserUUID(w http.ResponseWriter, r *http.Request, uuid string) *Response
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler          ServerInterface
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetUser(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// PostUser operation middleware
func (siw *ServerInterfaceWrapper) PostUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PostUser(w, r)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// DeleteUserUUID operation middleware
func (siw *ServerInterfaceWrapper) DeleteUserUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "uuid" -------------
	var uuid string

	if err := runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "uuid"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.DeleteUserUUID(w, r, uuid)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// GetUserUUID operation middleware
func (siw *ServerInterfaceWrapper) GetUserUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "uuid" -------------
	var uuid string

	if err := runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "uuid"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.GetUserUUID(w, r, uuid)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

// PutUserUUID operation middleware
func (siw *ServerInterfaceWrapper) PutUserUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// ------------- Path parameter "uuid" -------------
	var uuid string

	if err := runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid); err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{err, "uuid"})
		return
	}

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := siw.Handler.PutUserUUID(w, r, uuid)
		if resp != nil {
			if resp.body != nil {
				render.Render(w, r, resp)
			} else {
				w.WriteHeader(resp.Code)
			}
		}
	})

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter %s: %v", err.paramName, err.err)
}

func (err UnescapedCookieParamError) Unwrap() error { return err.err }

type UnmarshalingParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err UnmarshalingParamError) Error() string {
	return fmt.Sprintf("error unmarshaling parameter %s as JSON: %v", err.paramName, err.err)
}

func (err UnmarshalingParamError) Unwrap() error { return err.err }

type RequiredParamError struct {
	err       error
	paramName string
}

// Error implements error.
func (err RequiredParamError) Error() string {
	if err.err == nil {
		return fmt.Sprintf("query parameter %s is required, but not found", err.paramName)
	} else {
		return fmt.Sprintf("query parameter %s is required, but errored: %s", err.paramName, err.err)
	}
}

func (err RequiredParamError) Unwrap() error { return err.err }

type RequiredHeaderError struct {
	paramName string
}

// Error implements error.
func (err RequiredHeaderError) Error() string {
	return fmt.Sprintf("header parameter %s is required, but not found", err.paramName)
}

type InvalidParamFormatError struct {
	err       error
	paramName string
}

// Error implements error.
func (err InvalidParamFormatError) Error() string {
	return fmt.Sprintf("invalid format for parameter %s: %v", err.paramName, err.err)
}

func (err InvalidParamFormatError) Unwrap() error { return err.err }

type TooManyValuesForParamError struct {
	NumValues int
	paramName string
}

// Error implements error.
func (err TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("expected one value for %s, got %d", err.paramName, err.NumValues)
}

// ParameterName is an interface that is implemented by error types that are
// relevant to a specific parameter.
type ParameterError interface {
	error
	// ParamName is the name of the parameter that the error is referring to.
	ParamName() string
}

func (err UnescapedCookieParamError) ParamName() string  { return err.paramName }
func (err UnmarshalingParamError) ParamName() string     { return err.paramName }
func (err RequiredParamError) ParamName() string         { return err.paramName }
func (err RequiredHeaderError) ParamName() string        { return err.paramName }
func (err InvalidParamFormatError) ParamName() string    { return err.paramName }
func (err TooManyValuesForParamError) ParamName() string { return err.paramName }

type ServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

type ServerOption func(*ServerOptions)

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface, opts ...ServerOption) http.Handler {
	options := &ServerOptions{
		BaseURL:    "/",
		BaseRouter: chi.NewRouter(),
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
	}

	for _, f := range opts {
		f(options)
	}

	r := options.BaseRouter
	wrapper := ServerInterfaceWrapper{
		Handler:          si,
		ErrorHandlerFunc: options.ErrorHandlerFunc,
	}

	r.Route(options.BaseURL, func(r chi.Router) {
		r.Get("/user", wrapper.GetUser)
		r.Post("/user", wrapper.PostUser)
		r.Delete("/user/{uuid}", wrapper.DeleteUserUUID)
		r.Get("/user/{uuid}", wrapper.GetUserUUID)
		r.Put("/user/{uuid}", wrapper.PutUserUUID)
	})
	return r
}

func WithRouter(r chi.Router) ServerOption {
	return func(s *ServerOptions) {
		s.BaseRouter = r
	}
}

func WithServerBaseURL(url string) ServerOption {
	return func(s *ServerOptions) {
		s.BaseURL = url
	}
}

func WithErrorHandler(handler func(w http.ResponseWriter, r *http.Request, err error)) ServerOption {
	return func(s *ServerOptions) {
		s.ErrorHandlerFunc = handler
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RUTWsbMRD9K2LOS+y2t721DRRDC7n4VHpQpLGjsPqoZpSyLPrvZbSYuPbWpaQYSk4S",
	"o/l6741mAhN9igEDE/RT7cCFXYR+AnY8IPSwJczqiw56jx4Dq/d3G+jgCTO5GKCHNzfrmzXUDmLCoJOD",
	"Ht41UwdJ84NkhVUhzHLZI8sRE2bNLoaNhR4+IUsR6CAjpRgIW9Db9VoOEwNjaGE6pcGZFrh6JKk+AZkH",
	"9FpujtG3wJQlP7s5DXrthgZoTIKHOLuwl4aD9rj4UIqzCw+1O1ji/SMahmeDzlmPUMXFIpnsEs/sfHbE",
	"Ku6UEEAtBRXvdR5n3Eqr4VePDlKkBZLuIj2z9L0g8Ydox78i6IW8nMMXkzTjMlroOResL9TwetKdS9UG",
	"3WTUjFZRMQaJdmUYxhPZPjYXpVXAH0219t5GfDVJ/SoNWByQ8VzH22aXWtvt5rZ9kqw9sojff53ASSvy",
	"ceCAcsZ0ynR3xNop2G/LKizAnbu8CHduuAFV96NqXdfu4k++OrL/Z74ssnYDqYycHT5dpl4WRDmOOuI/",
	"laUlUa7B/7/YPq95xxzULMn+addsm8tvhqDW+jMAAP//1FwjH78HAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}