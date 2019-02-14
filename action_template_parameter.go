package octopusdeploy

import (
	"encoding/json"
	"errors"
)

type ActionTemplateParameter struct {

	// default value
	DefaultValue PropertyValue `json:"DefaultValue,omitempty"`

	// display settings
	DisplaySettings map[string]string `json:"DisplaySettings,omitempty"`

	// help text
	HelpText string `json:"HelpText,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// label
	Label string `json:"Label,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn string `json:"LastModifiedOn,omitempty"` // datetime

	// links
	Links Links `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

func (propertyValue *PropertyValue) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &propertyValue.SensitiveValue); err == nil {
		propertyValue.IsSensitive = true
		return nil
	}
	if err := json.Unmarshal(b, &propertyValue.Value); err == nil {
		propertyValue.IsSensitive = false
		return nil
	}
	return errors.New("failed to unmarshal PropertyValue")
}

type PropertyValue struct {
	IsSensitive    bool           `json:"IsSensitive,omitempty"`
	Value          string         `json:"Value,omitempty"`
	SensitiveValue SensitiveValue `json:"SensitiveValue,omitempty"`
}

type SensitiveValue struct {
	HasValue bool   `json:"HasValue,omitempty"`
	NewValue string `json:"NewValue,omitempty"`
}
