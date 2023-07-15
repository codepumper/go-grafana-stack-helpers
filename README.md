# Grafana Stack Go Utils

Grafana Stack Go Utils is a collection of reusable components and utilities in Go for working with the Grafana stack. This repository provides high-level implementations and examples that demonstrate how to easily utilize the Grafana, Loki, Promtail, and other components within a Go application.

## Features

- Docker-compose and configuration files for setting up the Grafana stack
- Tracing client for distributed tracing using Grafana Tempo or other tracing backends
- Metrics client for collecting and reporting application-specific metrics to Prometheus or other metrics systems
- Logging package for easy integration with Prometheus, Opentelemetry and Grafana Loki

## Tracing

## Metrics

## Logging

At this time, OpenTelemetry for Go does not provide a built-in logging implementation. To address this, we can utilize Promtail, which is part of the Grafana Loki project, for log forwarding to OpenTelemetry.

The logging package uses Promtail for log management. **Promtail** collects logs and pushes them to the **OpenTelemetry collector**, which further processes and forwards the logs to the **Loki exporter**. The exporter translates and transmits the logs to **Loki for storage and analysis**. This setup enables centralized log management and integration with other observability data through the OpenTelemetry ecosystem.

### Usage

To use the Logging Package in your Go project, simply import it:

```go
import "github.com/codepumper/go-grafana-stack-helpers/logging"
```

Here is a quick example that demonstrates how to use the Logging Package:

```go
package main

import (
	"github.com/codepumper/go-grafana-stack-helpers/logging"
)

func main() {
    appName := "my-app"
	labels := "{application=\"" + appName + "\"}"

	err := logging.InitLogger("json", labels, nil)
	if err != nil {
		fmt.Println("Failed to initialize Logger")
	}

	logging.Infof("This is an info message")

    logging.Shutdown()
}

```


---

License
The Packages in this repository are open source and available under the Apache License 2.0. See the LICENSE file for more information.



