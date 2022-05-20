package metadata

import (
	"fmt"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

type Metadata struct {
	OnGCP     bool
	ProjectID string
}

func NewMetadata() *Metadata {

	// if this process is running on GCE then we're on GCP :)
	onGCP := metadata.OnGCE()

	projectID := ""
	if onGCP {
		var err error
		projectID, err = metadata.ProjectID()
		if err != nil {
			panic(err)
		}
	}

	return &Metadata{
		OnGCP:     onGCP,
		ProjectID: projectID,
	}
}

// Returns the current instance's numeric project ID.
func (m *Metadata) NumericProjectID() (string, error) {
	if m.OnGCP {
		return metadata.NumericProjectID()
	}
	return "", nil
}

// Returns the current VM's numeric instance ID.
func (m *Metadata) InstanceID() (string, error) {
	if m.OnGCP {
		return metadata.InstanceID()
	}
	return "", nil
}

// Returns the Cloud Run geographical region.
func (m *Metadata) Region() (string, error) {
	if m.OnGCP {
		return metadata.Get("instance/region")
	}
	return "", nil
}

// Returns an OIDC token to call another services that can validate an identity token.
// On Cloud Run, the provided "audience" shall be the URL of the service you want to invoke.
// https://cloud.google.com/run/docs/securing/service-identity#identity_tokens
func (m *Metadata) IdentityToken(audience string) (string, error) {
	if m.OnGCP {
		return metadata.Get(fmt.Sprintf("instance/service-accounts/default/identity?audience=%s", audience))
	}
	return "", nil
}

// Returns an access token required to call GCP API's with.
// The provided "scopes" shall be a list of the OAuth scopes requested.
// https://cloud.google.com/run/docs/securing/service-identity#access_tokens
func (m *Metadata) AccessToken(scopes []string) (string, error) {
	if m.OnGCP {
		return metadata.Get(fmt.Sprintf("instance/service-accounts/default/token?scopes=%s", strings.Join(scopes[:], ",")))
	}
	return "", nil
}
