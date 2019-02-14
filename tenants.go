package octopusdeploy

import (
	"fmt"
	"time"

	"github.com/dghubble/sling"
)

type Tenant struct {
	ID                  string              `json:"Id"`
	Name                string              `json:"Name"`
	TenantTags          []string            `json:"TenantTags"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments"`
	SpaceID             string              `json:"SpaceId"`
	LastModifiedOn      *time.Time          `json:"LastModifiedOn,omitempty"`
	LastModifiedBy      *time.Time          `json:"LastModifiedBy,omitempty"`
	Links               map[string]string   `json:"Links"`
}

type Tenants struct {
	Items []Tenant `json:"Items"`
	PagedResults
}

type TenantService struct {
	sling *sling.Sling
}

func NewTenantService(sling *sling.Sling) *TenantService {
	return &TenantService{
		sling: sling,
	}
}

// Get returns a single machine with a given MachineID
func (s *TenantService) Get(TenantID string) (*Tenant, error) {
	path := fmt.Sprintf("tenants/%s", TenantID)
	resp, err := apiGet(s.sling, new(Tenant), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Tenant), nil
}

// GetAll returns all registered machines
func (s *TenantService) GetAll() (*[]Tenant, error) {
	var p []Tenant
	path := "tenants"
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Tenants), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*Tenants)
		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}
	return &p, nil
}
