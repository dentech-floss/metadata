# metadata

The GCP metadata server provides key pieces of information in regards to the context of where the service is running at, like if we're running on GPC or not and if so then in which project id. This lib wraps the [compute/metadata](https://pkg.go.dev/cloud.google.com/go/compute/metadata) lib to get hold of this information.

In case your service does not make use of [runsd](https://github.com/ahmetb/runsd) then this lib can be used to create [Identity Tokens](https://cloud.google.com/run/docs/securing/service-identity#identity_tokens) which might be required in order to communicate with other Cloud Run service. 

[Access Tokens](https://cloud.google.com/run/docs/securing/service-identity#access_tokens) are also possible to create, which is required when calling GCP API's.

## Install

```
go get github.com/dentech-floss/metadata@v0.1.0
```

## Usage

```go
package example

import (
    "github.com/dentech-floss/metadata/pkg/metadata"
    "github.com/dentech-floss/publisher/pkg/publisher"
    "github.com/dentech-floss/revision/pkg/revision"
    "github.com/dentech-floss/telemetry/pkg/telemetry"
)

func main() {

    metadata := metadata.NewMetadata()

    shutdownTracing := telemetry.SetupTracing(
        ctx,
        &telemetry.TracingConfig{
            ServiceName:           revision.ServiceName,
            ServiceVersion:        revision.ServiceVersion,
            DeploymentEnvironment: metadata.ProjectID,
            OtlpExporterEnabled:   metadata.OnGCP,
        },
    )
    defer shutdownTracing()

    publisher := publisher.NewPublisher(
        &publisher.PublisherConfig{
            OnGCP:     metadata.OnGCP,
            ProjectId: metadata.ProjectID,
        },
    )
}
```