package octopusdeploy

import (
	"encoding/json"
	"errors"
)

type ActionTemplateParameter struct {
	DefaultValue    PropertyValue     `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string `json:"DisplaySettings,omitempty"`
	HelpText        string            `json:"HelpText,omitempty"`
	ID              string            `json:"Id,omitempty"`
	Label           string            `json:"Label,omitempty"`
	LastModifiedBy  string            `json:"LastModifiedBy,omitempty"`
	LastModifiedOn  string            `json:"LastModifiedOn,omitempty"`
	Links           Links             `json:"Links,omitempty"`
	Name            string            `json:"Name,omitempty"`
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
