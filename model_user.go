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

// User User.  Represents any possible value for a user (subject or principal). Can be a: - Specific user object e.g.: 'user:will', 'folder:marketing', 'org:contoso', ...) - Specific userset (e.g. 'group:engineering#member') - Public-typed wildcard (e.g. 'user:*')  See https://openfga.dev/docs/concepts#what-is-a-user
type User struct {
	Object   *FgaObject     `json:"object,omitempty"yaml:"object,omitempty"`
	Userset  *UsersetUser   `json:"userset,omitempty"yaml:"userset,omitempty"`
	Wildcard *TypedWildcard `json:"wildcard,omitempty"yaml:"wildcard,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *User) GetObject() FgaObject {
	if o == nil || o.Object == nil {
		var ret FgaObject
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetObjectOk() (*FgaObject, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *User) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given FgaObject and assigns it to the Object field.
func (o *User) SetObject(v FgaObject) {
	o.Object = &v
}

// GetUserset returns the Userset field value if set, zero value otherwise.
func (o *User) GetUserset() UsersetUser {
	if o == nil || o.Userset == nil {
		var ret UsersetUser
		return ret
	}
	return *o.Userset
}

// GetUsersetOk returns a tuple with the Userset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsersetOk() (*UsersetUser, bool) {
	if o == nil || o.Userset == nil {
		return nil, false
	}
	return o.Userset, true
}

// HasUserset returns a boolean if a field has been set.
func (o *User) HasUserset() bool {
	if o != nil && o.Userset != nil {
		return true
	}

	return false
}

// SetUserset gets a reference to the given UsersetUser and assigns it to the Userset field.
func (o *User) SetUserset(v UsersetUser) {
	o.Userset = &v
}

// GetWildcard returns the Wildcard field value if set, zero value otherwise.
func (o *User) GetWildcard() TypedWildcard {
	if o == nil || o.Wildcard == nil {
		var ret TypedWildcard
		return ret
	}
	return *o.Wildcard
}

// GetWildcardOk returns a tuple with the Wildcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetWildcardOk() (*TypedWildcard, bool) {
	if o == nil || o.Wildcard == nil {
		return nil, false
	}
	return o.Wildcard, true
}

// HasWildcard returns a boolean if a field has been set.
func (o *User) HasWildcard() bool {
	if o != nil && o.Wildcard != nil {
		return true
	}

	return false
}

// SetWildcard gets a reference to the given TypedWildcard and assigns it to the Wildcard field.
func (o *User) SetWildcard(v TypedWildcard) {
	o.Wildcard = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Userset != nil {
		toSerialize["userset"] = o.Userset
	}
	if o.Wildcard != nil {
		toSerialize["wildcard"] = o.Wildcard
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
