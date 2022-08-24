/*
notifier

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.2
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NotifyBody struct for NotifyBody
type NotifyBody struct {
	Message string `json:"message"`
	UserId string `json:"user_id"`
	Subject string `json:"subject"`
}

// NewNotifyBody instantiates a new NotifyBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyBody(message string, userId string, subject string) *NotifyBody {
	this := NotifyBody{}
	this.Message = message
	this.UserId = userId
	this.Subject = subject
	return &this
}

// NewNotifyBodyWithDefaults instantiates a new NotifyBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyBodyWithDefaults() *NotifyBody {
	this := NotifyBody{}
	return &this
}

// GetMessage returns the Message field value
func (o *NotifyBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotifyBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotifyBody) SetMessage(v string) {
	o.Message = v
}

// GetUserId returns the UserId field value
func (o *NotifyBody) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *NotifyBody) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *NotifyBody) SetUserId(v string) {
	o.UserId = v
}

// GetSubject returns the Subject field value
func (o *NotifyBody) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *NotifyBody) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *NotifyBody) SetSubject(v string) {
	o.Subject = v
}

func (o NotifyBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableNotifyBody struct {
	value *NotifyBody
	isSet bool
}

func (v NullableNotifyBody) Get() *NotifyBody {
	return v.value
}

func (v *NullableNotifyBody) Set(val *NotifyBody) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyBody) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyBody(val *NotifyBody) *NullableNotifyBody {
	return &NullableNotifyBody{value: val, isSet: true}
}

func (v NullableNotifyBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


