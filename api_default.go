/*
 * Logbook of the World Query API
 *
 * LoTW provides a web service that accepts RESTful queries that report QSOs   satisfying specified criteria:    * accepted by LoTW after a specified date   * confirmed by LoTW after a specified date   * with a specified callsign   * with an operator in a specified DXCC entity   * in a specified mode   * on a specified band   * at a specified date and timeusing a specified station callsign
 *
 * API version: 1.0
 * Contact: k0swe@arrl.net
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lotw

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

// QueryOpts Optional parameters for the method 'Query'
type QueryOpts struct {
	QsoQsl        optional.String
	QsoQslsince   optional.String
	QsoQsorxsince optional.String
	QsoOwncall    optional.String
	QsoCallsign   optional.String
	QsoMode       optional.String
	QsoDxcc       optional.String
	QsoStartdate  optional.String
	QsoStarttime  optional.String
	QsoEnddate    optional.String
	QsoEndtime    optional.String
	QsoMydetail   optional.String
	QsoQsldetail  optional.String
	QsoWithown    optional.String
}

/*
Query Querying LoTW for Acceptance and Confirmation of Submitted QSOs
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param login Note that while the user's primary call sign is usually the username, this is not always the case and should not be assumed.
 * @param password
 * @param qsoQuery If absent, ADIF file will contain no QSO records
 * @param optional nil or *QueryOpts - Optional Parameters:
 * @param "QsoQsl" (optional.String) -  If \"yes\", only QSL records are returned
 * @param "QsoQslsince" (optional.String) -  Returns QSL records received (matched or updated) on or after the specified date. Will also accept date/time in \"YYYY-MM-DD HH:MM:SS\" format. Ignored unless qso_qsl=\"yes\".
 * @param "QsoQsorxsince" (optional.String) -  Returns QSO records received (uploaded) on or after the specified date. Will also accept date/time in \"YYYY-MM-DD HH:MM:SS\" format. Ignored unless qso_qsl=\"no\".
 * @param "QsoOwncall" (optional.String) -  Returns only records whose \"own\" call sign matches.
 * @param "QsoCallsign" (optional.String) -  Returns only records whose \"worked\" call sign matches.
 * @param "QsoMode" (optional.String) -  Returns only records whose mode matches. Mode must be one of the allowed modes.
 * @param "QsoDxcc" (optional.String) -  Returns only records whose DXCC entity matches. (This implies qso_qsl=\"yes\" since the DXCC entity of un-QSL'd stations isn't known to LoTW.) Value must be the ARRL DXCC entity number.
 * @param "QsoStartdate" (optional.String) -  Returns only records with a QSO date on or after the specified value.
 * @param "QsoStarttime" (optional.String) -  Returns only records with a QSO time at or after the specified value on the starting date. This value is ignored if qso_startdate is not provided.
 * @param "QsoEnddate" (optional.String) -  Returns only records with a QSO date on or before the specified value.
 * @param "QsoEndtime" (optional.String) -  Returns only records with a QSO time at or before the specified value on the ending date. This value is ignored if qso_enddate is not provided.
 * @param "QsoMydetail" (optional.String) -  If \"yes\", returns fields that contain the Logging station's location data, if any.
 * @param "QsoQsldetail" (optional.String) -  If \"yes\", returns fields that contain the QSLing station's location data, if any.
 * @param "QsoWithown" (optional.String) -  If \"yes\", each record contains the STATION_CALLSIGN and APP_LoTW_OWNCALL fields to identify the \"own\" call sign used for the QSO.
@return string
*/
func (a *DefaultApiService) Query(ctx _context.Context, login string, password string, qsoQuery int32, localVarOptionals *QueryOpts) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/lotwuser/lotwreport.adi"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("login", parameterToString(login, ""))
	localVarQueryParams.Add("password", parameterToString(password, ""))
	localVarQueryParams.Add("qso_query", parameterToString(qsoQuery, ""))
	if localVarOptionals != nil && localVarOptionals.QsoQsl.IsSet() {
		localVarQueryParams.Add("qso_qsl", parameterToString(localVarOptionals.QsoQsl.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoQslsince.IsSet() {
		localVarQueryParams.Add("qso_qslsince", parameterToString(localVarOptionals.QsoQslsince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoQsorxsince.IsSet() {
		localVarQueryParams.Add("qso_qsorxsince", parameterToString(localVarOptionals.QsoQsorxsince.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoOwncall.IsSet() {
		localVarQueryParams.Add("qso_owncall", parameterToString(localVarOptionals.QsoOwncall.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoCallsign.IsSet() {
		localVarQueryParams.Add("qso_callsign", parameterToString(localVarOptionals.QsoCallsign.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoMode.IsSet() {
		localVarQueryParams.Add("qso_mode", parameterToString(localVarOptionals.QsoMode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoDxcc.IsSet() {
		localVarQueryParams.Add("qso_dxcc", parameterToString(localVarOptionals.QsoDxcc.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoStartdate.IsSet() {
		localVarQueryParams.Add("qso_startdate", parameterToString(localVarOptionals.QsoStartdate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoStarttime.IsSet() {
		localVarQueryParams.Add("qso_starttime", parameterToString(localVarOptionals.QsoStarttime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoEnddate.IsSet() {
		localVarQueryParams.Add("qso_enddate", parameterToString(localVarOptionals.QsoEnddate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoEndtime.IsSet() {
		localVarQueryParams.Add("qso_endtime", parameterToString(localVarOptionals.QsoEndtime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoMydetail.IsSet() {
		localVarQueryParams.Add("qso_mydetail", parameterToString(localVarOptionals.QsoMydetail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoQsldetail.IsSet() {
		localVarQueryParams.Add("qso_qsldetail", parameterToString(localVarOptionals.QsoQsldetail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.QsoWithown.IsSet() {
		localVarQueryParams.Add("qso_withown", parameterToString(localVarOptionals.QsoWithown.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-arrl-adif"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
