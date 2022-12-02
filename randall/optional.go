package randall

import "encoding/json"

type OptionalUInt struct {
	value    uint
	hasValue bool
}

func NewOptionalUInt(value uint) OptionalUInt {
	return OptionalUInt{
		value:    value,
		hasValue: true,
	}
}

func (o *OptionalUInt) Value() uint {
	return o.value
}

func (o *OptionalUInt) SetValue(value uint) {
	o.value = value
	o.hasValue = true
}

func (o *OptionalUInt) SetNil() {
	o.value = 0
	o.hasValue = false
}

func (o *OptionalUInt) HasValue() bool {
	return o.hasValue
}

// MarshalJSON implements the encoding json interface.
func (o OptionalUInt) MarshalJSON() ([]byte, error) {
	if o.hasValue {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding json interface.
func (o *OptionalUInt) UnmarshalJSON(data []byte) error {
	var value *uint

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		o.SetValue(*value)
	} else {
		o.SetNil()
	}

	return nil
}

type OptionalFloat32 struct {
	value    float32
	hasValue bool
}

func NewOptionalFloat32(value float32) OptionalFloat32 {
	return OptionalFloat32{
		value:    value,
		hasValue: true,
	}
}

func (o *OptionalFloat32) Value() float32 {
	return o.value
}

func (o *OptionalFloat32) SetValue(value float32) {
	o.value = value
	o.hasValue = true
}

func (o *OptionalFloat32) SetNil() {
	o.value = 0
	o.hasValue = false
}

func (o *OptionalFloat32) HasValue() bool {
	return o.hasValue
}

// MarshalJSON implements the encoding json interface.
func (o OptionalFloat32) MarshalJSON() ([]byte, error) {
	if o.hasValue {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding json interface.
func (o *OptionalFloat32) UnmarshalJSON(data []byte) error {
	var value *float32

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		o.SetValue(*value)
	} else {
		o.SetNil()
	}

	return nil
}

type OptionalBool struct {
	value    bool
	hasValue bool
}

func NewOptionalBool(value bool) OptionalBool {
	return OptionalBool{
		value:    value,
		hasValue: true,
	}
}

func (o *OptionalBool) Value() bool {
	return o.value
}

func (o *OptionalBool) SetValue(value bool) {
	o.value = value
	o.hasValue = true
}

func (o *OptionalBool) SetNil() {
	o.value = false
	o.hasValue = false
}

func (o *OptionalBool) HasValue() bool {
	return o.hasValue
}

// MarshalJSON implements the encoding json interface.
func (o OptionalBool) MarshalJSON() ([]byte, error) {
	if o.hasValue {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding json interface.
func (o *OptionalBool) UnmarshalJSON(data []byte) error {
	var value *bool

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		o.SetValue(*value)
	} else {
		o.SetNil()
	}

	return nil
}

type OptionalString struct {
	value    string
	hasValue bool
}

func NewOptionalString(value string) OptionalString {
	return OptionalString{
		value:    value,
		hasValue: true,
	}
}

func (o *OptionalString) Value() string {
	return o.value
}

func (o *OptionalString) SetValue(value string) {
	o.value = value
	o.hasValue = true
}

func (o *OptionalString) SetNil() {
	o.value = ""
	o.hasValue = false
}

func (o *OptionalString) HasValue() bool {
	return o.hasValue
}

// MarshalJSON implements the encoding json interface.
func (o OptionalString) MarshalJSON() ([]byte, error) {
	if o.hasValue {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// UnmarshalJSON implements the encoding json interface.
func (o *OptionalString) UnmarshalJSON(data []byte) error {
	var value *string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if value != nil {
		o.SetValue(*value)
	} else {
		o.SetNil()
	}

	return nil
}

func FOptionalUInt(v uint) *uint {
	return &v
}

func FOptionalFloat32(v float32) *float32 {
	return &v
}

func FOptionalBool(v bool) *bool {
	return &v
}

func FOptionalString(v string) *string {
	return &v
}

func FOptional[T any](v T) *T {
	return &v
}

// func getUrlValues(providers ...QueryStringProvider) (url.Values, error) {
// 	if len(providers) == 0 {
// 		return url.Values{}, errors.New("getUrlValues() called with empty params slice")
// 	}

// 	values := make(url.Values)

// 	for _, provider := range providers {
// 		newValues, err := provider.AddQuery(values)

// 		if err != nil {
// 			return url.Values{}, err
// 		}

// 		values = newValues
// 	}

// 	return values, nil
// }
