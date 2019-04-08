package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type ProjectVariable struct {
	ProjectID   string                              `json:"ProjectId"`
	ProjectName string                              `json:"ProjectName"`
	Templates   []ActionTemplateParameter           `json:"Templates"`
	Variables   map[string]map[string]PropertyValue `json:"Variables"`
	Links       map[string]string                   `json:"Links"`
}

type LibraryVariable struct {
	LibraryVariableSetID   string                    `json:"LibraryVariableSetId"`
	LibraryVariableSetName string                    `json:"LibraryVariableSetName"`
	Templates              []ActionTemplateParameter `json:"Templates"`
	Variables              map[string]PropertyValue  `json:"Variables"`
	Links                  map[string]string         `json:"Links"`
}

type TenantVariables struct {
	ID               string                     `json:"Id"`
	TenantID         string                     `json:"TenantId"`
	TenantName       string                     `json:"TenantName"`
	ProjectVariables map[string]ProjectVariable `json:"ProjectVariables"`
	LibraryVariables map[string]LibraryVariable `json:"LibraryVariables"`
	SpaceID          string                     `json:"SpaceId"`
	LastModifiedOn   string                     `json:"LastModifiedOn,omitempty"`
	Links            map[string]string          `json:"Links,omitempty"`
}

type TenantVariableService struct {
	sling *sling.Sling
}

func NewTenantVariableService(sling *sling.Sling) *TenantVariableService {
	return &TenantVariableService{
		sling: sling,
	}
}

func (s *TenantVariableService) GetTenant(tenantID string) (*TenantVariables, error) {
	path := fmt.Sprintf("tenants/%s/variables", tenantID)
	resp, err := apiGet(s.sling, new(TenantVariables), path)

	if err != nil {
		return nil, err
	}

	return resp.(*TenantVariables), nil
}
