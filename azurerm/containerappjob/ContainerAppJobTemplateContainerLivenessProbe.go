// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappjob


type ContainerAppJobTemplateContainerLivenessProbe struct {
	// The port number on which to connect. Possible values are between `1` and `65535`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#port ContainerAppJob#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Type of probe. Possible values are `TCP`, `HTTP`, and `HTTPS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#transport ContainerAppJob#transport}
	Transport *string `field:"required" json:"transport" yaml:"transport"`
	// The number of consecutive failures required to consider this probe as failed.
	//
	// Possible values are between `1` and `10`. Defaults to `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#failure_count_threshold ContainerAppJob#failure_count_threshold}
	FailureCountThreshold *float64 `field:"optional" json:"failureCountThreshold" yaml:"failureCountThreshold"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#header ContainerAppJob#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The probe hostname.
	//
	// Defaults to the pod IP address. Setting a value for `Host` in `headers` can be used to override this for `http` and `https` type probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#host ContainerAppJob#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The time in seconds to wait after the container has started before the probe is started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#initial_delay ContainerAppJob#initial_delay}
	InitialDelay *float64 `field:"optional" json:"initialDelay" yaml:"initialDelay"`
	// How often, in seconds, the probe should run. Possible values are between `1` and `240`. Defaults to `10`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#interval_seconds ContainerAppJob#interval_seconds}
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The URI to use with the `host` for http type probes.
	//
	// Not valid for `TCP` type probes. Defaults to `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#path ContainerAppJob#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Time in seconds after which the probe times out. Possible values are between `1` an `240`. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/container_app_job#timeout ContainerAppJob#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

