package octopusdeploy

import (
	"fmt"

	"net/url"

	"github.com/dghubble/sling"
)

type LifecycleService struct {
	sling *sling.Sling
}

func NewLifecycleService(sling *sling.Sling) *LifecycleService {
	return &LifecycleService{
		sling: sling,
	}
}

type Lifecycles struct {
	Items []Lifecycle `json:"Items"`
	PagedResults
}

type Lifecycle struct {
	ID                      string          `json:"Id,omitempty"`
	Name                    string          `json:"Name"`
	Description             string          `json:"Description,omitempty"`
	ReleaseRetentionPolicy  RetentionPeriod `json:"ReleaseRetentionPolicy,omitempty"`
	TentacleRetentionPolicy RetentionPeriod `json:"TentacleRetentionPolicy,omitempty"`
	Phases                  []Phase         `json:"Phases"`
}

type RetentionUnit string

const (
	RetentionUnit_Days  = RetentionUnit("Days")
	RetentionUnit_Items = RetentionUnit("Items")
)

type RetentionPeriod struct {
	Unit              RetentionUnit `json:"Unit"`
	QuantityToKeep    int32         `json:"QuantityToKeep"`
	ShouldKeepForever bool          `json:"ShouldKeepForever"`
}

type Phase struct {
	ID                                 string           `json:"Id,omitempty"`
	Name                               string           `json:"Name"`
	MinimumEnvironmentsBeforePromotion int32            `json:"MinimumEnvironmentsBeforePromotion"`
	IsOptionalPhase                    bool             `json:"IsOptionalPhase"`
	ReleaseRetentionPolicy             *RetentionPeriod `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy            *RetentionPeriod `json:"TentacleRetentionPolicy"`
	AutomaticDeploymentTargets         []string         `json:"AutomaticDeploymentTargets"`
	OptionalDeploymentTargets          []string         `json:"OptionalDeploymentTargets"`
}

func NewLifecycle(name string) *Lifecycle {
	return &Lifecycle{
		Name:   name,
		Phases: []Phase{},
		TentacleRetentionPolicy: RetentionPeriod{
			Unit: RetentionUnit_Days,
		},
		ReleaseRetentionPolicy: RetentionPeriod{
			Unit: RetentionUnit_Days,
		},
	}
}

// Get returns a single lifecycle by its lifecycleid in Octopus Deploy
func (s *LifecycleService) Get(LifecycleID string) (*Lifecycle, error) {
	path := fmt.Sprintf("lifecycles/%s", LifecycleID)
	resp, err := apiGet(s.sling, new(Lifecycle), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

// GetAll returns all lifecycles in Octopus Deploy
func (s *LifecycleService) GetAll() (*[]Lifecycle, error) {
	return s.get("")
}

func (s *LifecycleService) get(query string) (*[]Lifecycle, error) {
	var p []Lifecycle

	path := "lifecycles?take=2147483647"
	if query != "" {
		path = fmt.Sprintf("%s&%s", path, query)
	}

	loadNextPage := true

	for loadNextPage { // Older Octopus Servers do not accept the take parameter, so the only choice is to page through them
		resp, err := apiGet(s.sling, new(Lifecycles), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Lifecycles)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName gets an existing lifecycle by its lifecycle name in Octopus Deploy
func (s *LifecycleService) GetByName(lifecycleName string) (*Lifecycle, error) {
	var foundLifecycle Lifecycle
	lifecycles, err := s.get(fmt.Sprintf("partialName=%s", url.PathEscape(lifecycleName)))

	if err != nil {
		return nil, err
	}

	for _, lifecycle := range *lifecycles {
		if lifecycle.Name == lifecycleName {
			return &lifecycle, nil
		}
	}

	return &foundLifecycle, fmt.Errorf("no lifecycle found with lifecycle name %s", lifecycleName)
}

// Add adds an new lifecycle in Octopus Deploy
func (s *LifecycleService) Add(lifecycle *Lifecycle) (*Lifecycle, error) {
	resp, err := apiAdd(s.sling, lifecycle, new(Lifecycle), "lifecycles")

	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}

// Delete deletes an existing lifecycle in Octopus Deploy
func (s *LifecycleService) Delete(lifecycleid string) error {
	path := fmt.Sprintf("lifecycles/%s", lifecycleid)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an existing lifecycle in Octopus Deploy
func (s *LifecycleService) Update(lifecycle *Lifecycle) (*Lifecycle, error) {
	path := fmt.Sprintf("lifecycles/%s", lifecycle.ID)
	resp, err := apiUpdate(s.sling, lifecycle, new(Lifecycle), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Lifecycle), nil
}
