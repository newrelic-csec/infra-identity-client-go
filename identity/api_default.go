// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
/*
 * Identity API specifications
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.5.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package identity

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// ConnectPostOpts Optional parameters for the method 'ConnectPost'
type ConnectPostOpts struct {
    ContentEncoding optional.String
}

/*
ConnectPost Connect using a fingerprint
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param connectRequest
 * @param optional nil or *ConnectPostOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
@return ConnectResponse
*/
func (a *DefaultApiService) ConnectPost(ctx _context.Context, userAgent string, xLicenseKey string, connectRequest ConnectRequest, localVarOptionals *ConnectPostOpts) (ConnectResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connect"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	// body params
	localVarPostBody = &connectRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ConnectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ConnectPutOpts Optional parameters for the method 'ConnectPut'
type ConnectPutOpts struct {
    ContentEncoding optional.String
}

/*
ConnectPut Reconnect using a fingerprint and an entityID
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param reconnectRequest
 * @param optional nil or *ConnectPutOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
@return ConnectResponse
*/
func (a *DefaultApiService) ConnectPut(ctx _context.Context, userAgent string, xLicenseKey string, reconnectRequest ReconnectRequest, localVarOptionals *ConnectPutOpts) (ConnectResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/connect"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	// body params
	localVarPostBody = &reconnectRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ConnectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// DisconnectPutOpts Optional parameters for the method 'DisconnectPut'
type DisconnectPutOpts struct {
    ContentEncoding optional.String
}

/*
DisconnectPut disconnects an entity
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param disconnectRequest
 * @param optional nil or *DisconnectPutOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
*/
func (a *DefaultApiService) DisconnectPut(ctx _context.Context, userAgent string, xLicenseKey string, disconnectRequest DisconnectRequest, localVarOptionals *DisconnectPutOpts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/disconnect"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	// body params
	localVarPostBody = &disconnectRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// LookupBatchPostOpts Optional parameters for the method 'LookupBatchPost'
type LookupBatchPostOpts struct {
    ContentEncoding optional.String
}

/*
LookupBatchPost lookup batch for list of entities, given their entityNames
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param lookupRequest
 * @param optional nil or *LookupBatchPostOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
@return []map[string]interface{}
*/
func (a *DefaultApiService) LookupBatchPost(ctx _context.Context, userAgent string, xLicenseKey string, lookupRequest []LookupRequest, localVarOptionals *LookupBatchPostOpts) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lookup/batch"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if len(lookupRequest) < 1 {
		return localVarReturnValue, nil, reportError("lookupRequest must have at least 1 elements")
	}
	if len(lookupRequest) > 1000 {
		return localVarReturnValue, nil, reportError("lookupRequest must have less than 1000 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	// body params
	localVarPostBody = &lookupRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 207 {
			var v []map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// LookupPostOpts Optional parameters for the method 'LookupPost'
type LookupPostOpts struct {
    ContentEncoding optional.String
}

/*
LookupPost lookup for an entity, given the entityName
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param lookupRequest
 * @param optional nil or *LookupPostOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
@return LookupResponse
*/
func (a *DefaultApiService) LookupPost(ctx _context.Context, userAgent string, xLicenseKey string, lookupRequest LookupRequest, localVarOptionals *LookupPostOpts) (LookupResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LookupResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lookup"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	// body params
	localVarPostBody = &lookupRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v LookupResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RegisterBatchPostOpts Optional parameters for the method 'RegisterBatchPost'
type RegisterBatchPostOpts struct {
    ContentEncoding optional.String
    XNRIAgentEntityId optional.Int64
}

/*
RegisterBatchPost Register integration entities in batch of a max size of 1000 and 1MB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param registerRequest
 * @param optional nil or *RegisterBatchPostOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
 * @param "XNRIAgentEntityId" (optional.Int64) - 
@return []RegisterBatchEntityResponse
*/
func (a *DefaultApiService) RegisterBatchPost(ctx _context.Context, userAgent string, xLicenseKey string, registerRequest []RegisterRequest, localVarOptionals *RegisterBatchPostOpts) ([]RegisterBatchEntityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []RegisterBatchEntityResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/register/batch"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if len(registerRequest) < 1 {
		return localVarReturnValue, nil, reportError("registerRequest must have at least 1 elements")
	}
	if len(registerRequest) > 1000 {
		return localVarReturnValue, nil, reportError("registerRequest must have less than 1000 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	if localVarOptionals != nil && localVarOptionals.XNRIAgentEntityId.IsSet() {
		localVarHeaderParams["X-NRI-Agent-Entity-Id"] = parameterToString(localVarOptionals.XNRIAgentEntityId.Value(), "")
	}
	// body params
	localVarPostBody = &registerRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 207 {
			var v []RegisterBatchEntityResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RegisterBatchPutOpts Optional parameters for the method 'RegisterBatchPut'
type RegisterBatchPutOpts struct {
    ContentEncoding optional.String
    XNRIAgentEntityId optional.Int64
}

/*
RegisterBatchPut Re-register integration entities in batch of a max size of 1000 and 1MB
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param reRegisterRequest
 * @param optional nil or *RegisterBatchPutOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
 * @param "XNRIAgentEntityId" (optional.Int64) - 
@return []ReRegisterBatchEntityResponse
*/
func (a *DefaultApiService) RegisterBatchPut(ctx _context.Context, userAgent string, xLicenseKey string, reRegisterRequest []ReRegisterRequest, localVarOptionals *RegisterBatchPutOpts) ([]ReRegisterBatchEntityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ReRegisterBatchEntityResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/register/batch"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if len(reRegisterRequest) < 1 {
		return localVarReturnValue, nil, reportError("reRegisterRequest must have at least 1 elements")
	}
	if len(reRegisterRequest) > 1000 {
		return localVarReturnValue, nil, reportError("reRegisterRequest must have less than 1000 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	if localVarOptionals != nil && localVarOptionals.XNRIAgentEntityId.IsSet() {
		localVarHeaderParams["X-NRI-Agent-Entity-Id"] = parameterToString(localVarOptionals.XNRIAgentEntityId.Value(), "")
	}
	// body params
	localVarPostBody = &reRegisterRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []ReRegisterBatchEntityResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RegisterPostOpts Optional parameters for the method 'RegisterPost'
type RegisterPostOpts struct {
    ContentEncoding optional.String
    XNRIAgentEntityId optional.Int64
}

/*
RegisterPost Register integration entity
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param registerRequest
 * @param optional nil or *RegisterPostOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
 * @param "XNRIAgentEntityId" (optional.Int64) - 
@return RegisterResponse
*/
func (a *DefaultApiService) RegisterPost(ctx _context.Context, userAgent string, xLicenseKey string, registerRequest RegisterRequest, localVarOptionals *RegisterPostOpts) (RegisterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegisterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/register"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	if localVarOptionals != nil && localVarOptionals.XNRIAgentEntityId.IsSet() {
		localVarHeaderParams["X-NRI-Agent-Entity-Id"] = parameterToString(localVarOptionals.XNRIAgentEntityId.Value(), "")
	}
	// body params
	localVarPostBody = &registerRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v RegisterResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RegisterPutOpts Optional parameters for the method 'RegisterPut'
type RegisterPutOpts struct {
    ContentEncoding optional.String
    XNRIAgentEntityId optional.Int64
}

/*
RegisterPut Re-register integration entity
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userAgent
 * @param xLicenseKey
 * @param reRegisterRequest
 * @param optional nil or *RegisterPutOpts - Optional Parameters:
 * @param "ContentEncoding" (optional.String) - 
 * @param "XNRIAgentEntityId" (optional.Int64) - 
@return ReRegisterResponse
*/
func (a *DefaultApiService) RegisterPut(ctx _context.Context, userAgent string, xLicenseKey string, reRegisterRequest ReRegisterRequest, localVarOptionals *RegisterPutOpts) (ReRegisterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReRegisterResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/register"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ContentEncoding.IsSet() {
		localVarHeaderParams["Content-Encoding"] = parameterToString(localVarOptionals.ContentEncoding.Value(), "")
	}
	localVarHeaderParams["User-Agent"] = parameterToString(userAgent, "")
	localVarHeaderParams["X-License-Key"] = parameterToString(xLicenseKey, "")
	if localVarOptionals != nil && localVarOptionals.XNRIAgentEntityId.IsSet() {
		localVarHeaderParams["X-NRI-Agent-Entity-Id"] = parameterToString(localVarOptionals.XNRIAgentEntityId.Value(), "")
	}
	// body params
	localVarPostBody = &reRegisterRequest
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ReRegisterResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
