package octopusdeploy

import (
	"errors"
	"strings"
	"time"

	"github.com/dghubble/sling"
)

type DeploymentItem struct {
	ID                       string            `json:"Id"`
	RealeaseID               string            `json:"ReleaseId"`
	EnvironmentID            string            `json:"EnvironmentId"`
	TenantID                 string            `json:"TenantId"`
	ForcePackageDownload     bool              `json:"ForcePackageDownload"`
	ForcePackageRedeployment bool              `json:"ForcePackageRedeployment"`
	SkipActions              []string          `json:"SkipActions"`
	SpecificMachineIds       []string          `json:"SpecificMachineIds"`
	ExcludedMachineIds       []string          `json:"ExcludedMachineIds"`
	DeploymentProcessID      string            `json:"DeploymentProcessId"`
	ManifestVariableSetID    string            `json:"ManifestVariableSetId"`
	TaskID                   string            `json:"TaskId"`
	ProjectID                string            `json:"ProjectId"`
	ChannelID                string            `json:"ChannelId"`
	UseGuidedFailure         bool              `json:"UseGuidedFailure"`
	Comments                 string            `json:"Comments"`
	FormValues               map[string]string `json:"FormValues"`
	QueueTime                time.Time         `json:"QueueTime"`
	QueueTimeExpiry          time.Time         `json:"QueueTimeExpiry"`
	Name                     string            `json:"Name"`
	LastModifiedOn           time.Time         `json:"LastModifiedOn,omitempty"`
	LastModifiedBy           time.Time         `json:"LastModifiedBy,omitempty"`
	Links                    map[string]string `json:"Links,omitempty"`
}

type Deployments struct {
	ID             string            `json:"Id"`
	ItemType       string            `json:"ItemType"`
	TotalResults   int64             `json:"TotalResults"`
	ItemsPerPage   int64             `json:"ItemsPerPage"`
	NumberOfPages  int64             `json:"NumberOfPages"`
	LastPageNumber int64             `json:"LastPageNumber"`
	Items          []DeploymentItem  `json:"Items"`
	TenantID       string            `json:"TenantId"`
	LastModifiedOn string            `json:"LastModifiedOn,omitempty"`
	LastModifiedBy string            `json:"LastModifiedBy,omitempty"`
	Links          map[string]string `json:"Links,omitempty"`
}

type DeploymentService struct {
	sling *sling.Sling
}

func NewDeploymentService(sling *sling.Sling) *DeploymentService {
	return &DeploymentService{
		sling: sling,
	}
}

func (s *DeploymentService) All() (*Deployments, error) {
	path := "deployments"
	resp, err := apiGet(s.sling, new(Deployments), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Deployments), nil
}

func (s *DeploymentService) FindLatestForTenant(tenantID string, terminateAt ...int) (*DeploymentItem, error) {
	path := "deployments"

	for {
		resp, err := apiGet(s.sling, new(Deployments), path)
		if err != nil {
			return nil, err
		}

		deployments := resp.(*Deployments)

		for _, deployment := range deployments.Items {
			if deployment.TenantID == tenantID {
				return &deployment, nil
			}
		}

		var ok bool
		path, ok = deployments.Links["Page.Next"]
		if !ok {
			break
		}
		path = strings.Replace(path, "/api/", "", 1)
	}

	return nil, errors.New("not found")
}
