/*
 * SMS API
 *
 * With the Nexmo SMS API you can send SMS from your account and lookup messages both messages that you've sent as well as messages sent to your virtual numbers. Numbers are specified in E.164 format. More SMS API documentation is at <https://developer.nexmo.com/messaging/sms/overview>
 *
 * API version: 1.0.5
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sms

import (
	"bytes"
	_context "context"
	"fmt"
	"io/ioutil"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// SendAnSmsOpts Optional parameters for the method 'SendAnSms'
type SendAnSmsOpts struct {
    ApiSecret optional.String
    Sig optional.String
    Text optional.String
    Ttl optional.Int32
    StatusReportReq optional.Bool
    Callback optional.String
    MessageClass optional.Int32
    Type_ optional.String
    Vcard optional.String
    Vcal optional.String
    Body optional.String
    Udh optional.String
    ProtocolId optional.Int32
    Title optional.String
    Url optional.String
    Validity optional.String
    ClientRef optional.String
    AccountRef optional.String
}

/*
SendAnSms Send an SMS
Send an outbound SMS from your Nexmo account
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param format The format of the response
 * @param apiKey Your API key
 * @param from The name or number the message should be sent from. Alphanumeric senderID's are not supported in all countries, see [Global Messaging](https://developer.nexmo.com/messaging/sms/guides/global-messaging#country-specific-features) for more details. If alphanumeric, spaces will be ignored. Numbers are specified in E.164 format.
 * @param to The number that the message should be sent to. Numbers are specified in E.164 format.
 * @param optional nil or *SendAnSmsOpts - Optional Parameters:
 * @param "ApiSecret" (optional.String) -  Your API secret. Required unless `sig` is provided
 * @param "Sig" (optional.String) -  The hash of the request parameters in alphabetical order, a timestamp and the signature secret. See [Signing Requests](/concepts/guides/signing-messages) for more details.
 * @param "Text" (optional.String) -  The body of the message being sent. If your message contains characters that can be encoded according to the GSM Standard and Extended tables then you can set the `type` to `text`. If your message contains characters outside this range, then you will need to set the `type` to `unicode`.
 * @param "Ttl" (optional.Int32) -  **Advanced**: The duration in milliseconds the delivery of an SMS will be attempted.§§ By default Nexmo attempt delivery for 72 hours, however the maximum effective value depends on the operator and is typically 24 - 48 hours. We recommend this value should be kept at its default or at least 30 minutes.
 * @param "StatusReportReq" (optional.Bool) -  **Advanced**: Boolean indicating if you like to receive a [Delivery Receipt](https://developer.nexmo.com/messaging/sms/building-blocks/receive-a-delivery-receipt).
 * @param "Callback" (optional.String) -  **Advanced**: The webhook endpoint the delivery receipt for this sms is sent to. This parameter overrides the webhook endpoint you set in Dashboard.
 * @param "MessageClass" (optional.Int32) -  **Advanced**: The Data Coding Scheme value of the message
 * @param "Type_" (optional.String) -  **Advanced**: The format of the message body
 * @param "Vcard" (optional.String) -  **Advanced**: A business card in [vCard format](https://en.wikipedia.org/wiki/VCard). Depends on `type` parameter having the value `vcard`.
 * @param "Vcal" (optional.String) -  **Advanced**: A calendar event in [vCal format](https://en.wikipedia.org/wiki/VCal). Depends on `type` parameter having the value `vcal`.
 * @param "Body" (optional.String) -  **Advanced**: Hex encoded binary data. Depends on `type` parameter having the value `binary`.
 * @param "Udh" (optional.String) -  **Advanced**: Your custom Hex encoded [User Data Header](https://en.wikipedia.org/wiki/User_Data_Header). Depends on `type` parameter having the value `binary`.
 * @param "ProtocolId" (optional.Int32) -  **Advanced**: The value of the [protocol identifier](https://en.wikipedia.org/wiki/GSM_03.40#Protocol_Identifier) to use. Ensure that the value is aligned with `udh`.
 * @param "Title" (optional.String) -  **Advanced**: The title for a wappush SMS. Depends on `type` parameter having the value `wappush`.
 * @param "Url" (optional.String) -  **Advanced**: The URL of your website. Depends on `type` parameter having the value `wappush`.
 * @param "Validity" (optional.String) -  **Advanced**: The availability for an SMS in milliseconds. Depends on `type` parameter having the value `wappush`.
 * @param "ClientRef" (optional.String) -  **Advanced**: You can optionally include your own reference of up to 40 characters.
 * @param "AccountRef" (optional.String) -  **Advanced**: An optional string used to identify separate accounts using the SMS endpoint for billing purposes. To use this feature, please email [support@nexmo.com](mailto:support@nexmo.com)
@return Sms
*/
func (a *DefaultApiService) SendAnSms(ctx _context.Context, format string, apiKey string, from string, to string, localVarOptionals *SendAnSmsOpts) (Sms, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Sms
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", format)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(apiKey) < 8 {
		return localVarReturnValue, nil, reportError("apiKey must have at least 8 elements")
	}
	if strlen(apiKey) > 8 {
		return localVarReturnValue, nil, reportError("apiKey must have less than 8 elements")
	}
	if strlen(to) < 7 {
		return localVarReturnValue, nil, reportError("to must have at least 7 elements")
	}
	if strlen(to) > 15 {
		return localVarReturnValue, nil, reportError("to must have less than 15 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("api_key", parameterToString(apiKey, ""))
	if localVarOptionals != nil && localVarOptionals.ApiSecret.IsSet() {
		localVarFormParams.Add("api_secret", parameterToString(localVarOptionals.ApiSecret.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sig.IsSet() {
		localVarFormParams.Add("sig", parameterToString(localVarOptionals.Sig.Value(), ""))
	}
	localVarFormParams.Add("from", parameterToString(from, ""))
	localVarFormParams.Add("to", parameterToString(to, ""))
	if localVarOptionals != nil && localVarOptionals.Text.IsSet() {
		localVarFormParams.Add("text", parameterToString(localVarOptionals.Text.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ttl.IsSet() {
		localVarFormParams.Add("ttl", parameterToString(localVarOptionals.Ttl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StatusReportReq.IsSet() {
		localVarFormParams.Add("status-report-req", parameterToString(localVarOptionals.StatusReportReq.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Callback.IsSet() {
		localVarFormParams.Add("callback", parameterToString(localVarOptionals.Callback.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MessageClass.IsSet() {
		localVarFormParams.Add("message-class", parameterToString(localVarOptionals.MessageClass.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		localVarFormParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Vcard.IsSet() {
		localVarFormParams.Add("vcard", parameterToString(localVarOptionals.Vcard.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Vcal.IsSet() {
		localVarFormParams.Add("vcal", parameterToString(localVarOptionals.Vcal.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarFormParams.Add("body", parameterToString(localVarOptionals.Body.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Udh.IsSet() {
		localVarFormParams.Add("udh", parameterToString(localVarOptionals.Udh.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProtocolId.IsSet() {
		localVarFormParams.Add("protocol-id", parameterToString(localVarOptionals.ProtocolId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Title.IsSet() {
		localVarFormParams.Add("title", parameterToString(localVarOptionals.Title.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Url.IsSet() {
		localVarFormParams.Add("url", parameterToString(localVarOptionals.Url.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Validity.IsSet() {
		localVarFormParams.Add("validity", parameterToString(localVarOptionals.Validity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientRef.IsSet() {
		localVarFormParams.Add("client-ref", parameterToString(localVarOptionals.ClientRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountRef.IsSet() {
		localVarFormParams.Add("account-ref", parameterToString(localVarOptionals.AccountRef.Value(), ""))
	}
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

	// hack to reinstate the body in case we need it
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v Sms
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
