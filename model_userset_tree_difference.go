/**
 * Go SDK for OpenFGA
 *
 * API version: 1.x
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// UsersetTreeDifference struct for UsersetTreeDifference
type UsersetTreeDifference struct {
	Base     Node `json:"base"yaml:"base"`
	Subtract Node `json:"subtract"yaml:"subtract"`
}

// NewUsersetTreeDifference instantiates a new UsersetTreeDifference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersetTreeDifference(base Node, subtract Node) *UsersetTreeDifference {
	this := UsersetTreeDifference{}
	this.Base = base
	this.Subtract = subtract
	return &this
}

// NewUsersetTreeDifferenceWithDefaults instantiates a new UsersetTreeDifference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersetTreeDifferenceWithDefaults() *UsersetTreeDifference {
	this := UsersetTreeDifference{}
	return &this
}

// GetBase returns the Base field value
func (o *UsersetTreeDifference) GetBase() Node {
	if o == nil {
		var ret Node
		return ret
	}

	return o.Base
}

// GetBaseOk returns a tuple with the Base field value
// and a boolean to check if the value has been set.
func (o *UsersetTreeDifference) GetBaseOk() (*Node, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Base, true
}

// SetBase sets field value
func (o *UsersetTreeDifference) SetBase(v Node) {
	o.Base = v
}

// GetSubtract returns the Subtract field value
func (o *UsersetTreeDifference) GetSubtract() Node {
	if o == nil {
		var ret Node
		return ret
	}

	return o.Subtract
}

// GetSubtractOk returns a tuple with the Subtract field value
// and a boolean to check if the value has been set.
func (o *UsersetTreeDifference) GetSubtractOk() (*Node, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtract, true
}

// SetSubtract sets field value
func (o *UsersetTreeDifference) SetSubtract(v Node) {
	o.Subtract = v
}

func (o UsersetTreeDifference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base"] = o.Base
	toSerialize["subtract"] = o.Subtract
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableUsersetTreeDifference struct {
	value *UsersetTreeDifference
	isSet bool
}

func (v NullableUsersetTreeDifference) Get() *UsersetTreeDifference {
	return v.value
}

func (v *NullableUsersetTreeDifference) Set(val *UsersetTreeDifference) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersetTreeDifference) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersetTreeDifference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersetTreeDifference(val *UsersetTreeDifference) *NullableUsersetTreeDifference {
	return &NullableUsersetTreeDifference{value: val, isSet: true}
}

func (v NullableUsersetTreeDifference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersetTreeDifference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
