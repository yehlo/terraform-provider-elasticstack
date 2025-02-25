/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the Model4xxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model4xxResponse{}

// Model4xxResponse struct for Model4xxResponse
type Model4xxResponse struct {
	StatusCode float32 `json:"statusCode"`
	Error      string  `json:"error"`
	Message    string  `json:"message"`
}

// NewModel4xxResponse instantiates a new Model4xxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel4xxResponse(statusCode float32, error_ string, message string) *Model4xxResponse {
	this := Model4xxResponse{}
	this.StatusCode = statusCode
	this.Error = error_
	this.Message = message
	return &this
}

// NewModel4xxResponseWithDefaults instantiates a new Model4xxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel4xxResponseWithDefaults() *Model4xxResponse {
	this := Model4xxResponse{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *Model4xxResponse) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *Model4xxResponse) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *Model4xxResponse) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetError returns the Error field value
func (o *Model4xxResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Model4xxResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Model4xxResponse) SetError(v string) {
	o.Error = v
}

// GetMessage returns the Message field value
func (o *Model4xxResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Model4xxResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Model4xxResponse) SetMessage(v string) {
	o.Message = v
}

func (o Model4xxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model4xxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["error"] = o.Error
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableModel4xxResponse struct {
	value *Model4xxResponse
	isSet bool
}

func (v NullableModel4xxResponse) Get() *Model4xxResponse {
	return v.value
}

func (v *NullableModel4xxResponse) Set(val *Model4xxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModel4xxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModel4xxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel4xxResponse(val *Model4xxResponse) *NullableModel4xxResponse {
	return &NullableModel4xxResponse{value: val, isSet: true}
}

func (v NullableModel4xxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel4xxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
