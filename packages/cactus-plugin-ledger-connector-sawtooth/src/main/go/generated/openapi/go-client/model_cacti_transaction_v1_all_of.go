/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
)

// checks if the CactiTransactionV1AllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CactiTransactionV1AllOf{}

// CactiTransactionV1AllOf struct for CactiTransactionV1AllOf
type CactiTransactionV1AllOf struct {
	// Decoded payload of sawtooth transaction.
	PayloadDecoded interface{} `json:"payload_decoded"`
}

// NewCactiTransactionV1AllOf instantiates a new CactiTransactionV1AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCactiTransactionV1AllOf(payloadDecoded interface{}) *CactiTransactionV1AllOf {
	this := CactiTransactionV1AllOf{}
	this.PayloadDecoded = payloadDecoded
	return &this
}

// NewCactiTransactionV1AllOfWithDefaults instantiates a new CactiTransactionV1AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCactiTransactionV1AllOfWithDefaults() *CactiTransactionV1AllOf {
	this := CactiTransactionV1AllOf{}
	return &this
}

// GetPayloadDecoded returns the PayloadDecoded field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CactiTransactionV1AllOf) GetPayloadDecoded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PayloadDecoded
}

// GetPayloadDecodedOk returns a tuple with the PayloadDecoded field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CactiTransactionV1AllOf) GetPayloadDecodedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PayloadDecoded) {
		return nil, false
	}
	return &o.PayloadDecoded, true
}

// SetPayloadDecoded sets field value
func (o *CactiTransactionV1AllOf) SetPayloadDecoded(v interface{}) {
	o.PayloadDecoded = v
}

func (o CactiTransactionV1AllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CactiTransactionV1AllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PayloadDecoded != nil {
		toSerialize["payload_decoded"] = o.PayloadDecoded
	}
	return toSerialize, nil
}

type NullableCactiTransactionV1AllOf struct {
	value *CactiTransactionV1AllOf
	isSet bool
}

func (v NullableCactiTransactionV1AllOf) Get() *CactiTransactionV1AllOf {
	return v.value
}

func (v *NullableCactiTransactionV1AllOf) Set(val *CactiTransactionV1AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCactiTransactionV1AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCactiTransactionV1AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCactiTransactionV1AllOf(val *CactiTransactionV1AllOf) *NullableCactiTransactionV1AllOf {
	return &NullableCactiTransactionV1AllOf{value: val, isSet: true}
}

func (v NullableCactiTransactionV1AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCactiTransactionV1AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


