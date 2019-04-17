// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type AdvancedSyntaxOption struct {
	value bool
}

func AdvancedSyntax(v bool) *AdvancedSyntaxOption {
	return &AdvancedSyntaxOption{v}
}

func (o *AdvancedSyntaxOption) Get() bool {
	if o == nil {
		return false
	}
	return o.value
}

func (o AdvancedSyntaxOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AdvancedSyntaxOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *AdvancedSyntaxOption) Equal(o2 *AdvancedSyntaxOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func AdvancedSyntaxEqual(o1, o2 *AdvancedSyntaxOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}