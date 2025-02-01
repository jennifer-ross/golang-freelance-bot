package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang-freelance-bot/logger"
	"time"

	"github.com/valyala/fasthttp"
)

// BaseApi represents the base API client structure.
// It contains the context, base URI, and an HTTP client.
type BaseApi struct {
	ctx     context.Context  // Context for storing request-scoped values (e.g., auth tokens)
	baseUri string           // Base URI for API requests
	client  *fasthttp.Client // Fast HTTP client for making requests
}

// BaseApiResponse represents the structure of the API response.
// It contains the headers and body of the response.
type BaseApiResponse struct {
	Headers map[string]string // Response headers
	Body    []byte            // Response body
}

// Constants for context keys used for authentication.
const AuthCtxKey = "fasthttpsauth"       // Key for storing authentication status in context
const AuthKeyCtxKey = "fasthttpsauthkey" // Key for storing authentication token in context

// ParamsCtxKey is a context key used for storing query parameters in the request.
const ParamsCtxKey = "fasthttpparams" // Key for storing query parameters in context

// New creates and returns a new BaseApi instance.
// It initializes the HTTP client with default timeouts and settings.
func New(uri string) *BaseApi {
	// You may read the timeouts from some config
	readTimeout, _ := time.ParseDuration("500ms")      // Timeout for reading responses
	writeTimeout, _ := time.ParseDuration("500ms")     // Timeout for writing requests
	maxIdleConnDuration, _ := time.ParseDuration("1h") // Maximum idle connection duration

	return &BaseApi{
		ctx:     context.Background(), // Initialize context
		baseUri: uri,                  // Set base URI
		client: &fasthttp.Client{ // Configure HTTP client
			ReadTimeout:                   readTimeout,
			WriteTimeout:                  writeTimeout,
			MaxIdleConnDuration:           maxIdleConnDuration,
			NoDefaultUserAgentHeader:      true, // Disable default User-Agent header
			DisableHeaderNamesNormalizing: true, // Disable header names normalization
			DisablePathNormalizing:        true, // Disable path normalization
			// Increase DNS cache time to an hour instead of the default minute
			Dial: (&fasthttp.TCPDialer{
				Concurrency:      4096,      // Maximum number of concurrent connections
				DNSCacheDuration: time.Hour, // DNS cache duration
			}).Dial,
		},
	}
}

// Request sends an HTTP request to the specified path with the given method.
// It returns an error (if any) and a BaseApiResponse containing the response body and headers.
func (api *BaseApi) Request(path string, method string) (error, *BaseApiResponse) {
	// Validate the HTTP method
	if !isValidMethod(method) {
		var errorStr = fmt.Sprintf("Not a valid HTTP method: %s", method)
		logger.Get("app").Error().Msg(errorStr)
		return errors.New(errorStr), nil
	}

	// Acquire an HTTP request object
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // Release the request object after use

	// Set request URI and method
	req.SetRequestURI(fmt.Sprintf("%s/%s", api.baseUri, path))
	req.Header.SetMethod(method)

	// Check if authentication is required and set the Authorization header
	if isAuth, ok := api.ctx.Value(AuthCtxKey).(bool); ok && isAuth {
		if authKey, ok := api.ctx.Value(AuthKeyCtxKey).(string); ok && authKey != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authKey))
		}
	}

	// Set query parameters if the method for GET and body parameters for other methods provided in the context.
	api.setParamsFromCtx(req, method)

	// Acquire an HTTP response object
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // Release the response object after use

	// Send the request and handle errors
	err := api.client.Do(req, resp)
	if err != nil {
		logger.Get("app").Error().Err(err).Msg("Request failed")
		return err, nil
	}

	// Extract the response body
	body := make([]byte, len(resp.Body()))
	copy(body, resp.Body())

	// Extract the response headers
	headers := make(map[string]string)
	resp.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})

	// Clear the context for the next request
	api.ctx = context.Background()

	// Return the response
	return nil, &BaseApiResponse{
		Body:    body,
		Headers: headers,
	}
}

// setParamsFromCtx sets the request parameters from the context.
// If the HTTP method is GET, it adds query parameters to the request URI.
// For other methods, it marshals the parameters into JSON and sets them as the request body.
func (api *BaseApi) setParamsFromCtx(req *fasthttp.Request, method string) {
	if params, ok := api.ctx.Value(ParamsCtxKey).(map[string]string); ok && params != nil {
		// TODO add logic to choose different body format from ctx. For now, it uses JSON.
		if method == fasthttp.MethodGet {
			for paramKey, paramVal := range params {
				req.URI().QueryArgs().AddBytesKV([]byte(paramKey), []byte(paramVal))
			}
		} else {
			jsonData, err := json.Marshal(params)

			if err != nil {
				req.SetBody(jsonData)
			}
		}
	}
}

// Get sends a GET request to the specified path.
// It returns the API response.
func (api *BaseApi) Get(path string) *BaseApiResponse {
	var _, resp = api.Request(path, fasthttp.MethodGet)
	return resp
}

// Post sends a POST request to the specified path.
// It returns the API response.
func (api *BaseApi) Post(path string) *BaseApiResponse {
	var _, resp = api.Request(path, fasthttp.MethodPost)
	return resp
}

// AuthBasic sets the authentication token for the API client.
// It enables basic authentication and stores the token in the context.
func (api *BaseApi) AuthBasic(token string) *BaseApi {
	api.ctx = context.WithValue(api.ctx, AuthCtxKey, true)     // Enable authentication
	api.ctx = context.WithValue(api.ctx, AuthKeyCtxKey, token) // Store the token
	return api
}

// Params sets the provided parameters in the context for the API client.
func (api *BaseApi) Params(params map[string]string) *BaseApi {
	api.ctx = context.WithValue(api.ctx, ParamsCtxKey, params) // Store the parameters
	return api
}

// isValidMethod checks if the provided HTTP method is valid.
// Valid methods are GET, POST, PUT, and DELETE.
func isValidMethod(method string) bool {
	return method == fasthttp.MethodGet || method == fasthttp.MethodPost || method == fasthttp.MethodPut || method == fasthttp.MethodDelete
}
