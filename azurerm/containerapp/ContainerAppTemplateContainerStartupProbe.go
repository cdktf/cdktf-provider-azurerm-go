// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppTemplateContainerStartupProbe struct {
	// The port number on which to connect. Possible values are between `1` and `65535`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#port ContainerApp#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Type of probe. Possible values are `TCP`, `HTTP`, and `HTTPS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#transport ContainerApp#transport}
	Transport *string `field:"required" json:"transport" yaml:"transport"`
	// The number of consecutive failures required to consider this probe as failed.
	//
	// Possible values are between `1` and `30`. Defaults to `3`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#failure_count_threshold ContainerApp#failure_count_threshold}
	FailureCountThreshold *float64 `field:"optional" json:"failureCountThreshold" yaml:"failureCountThreshold"`
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#header ContainerApp#header}
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// The probe hostname.
	//
	// Defaults to the pod IP address. Setting a value for `Host` in `headers` can be used to override this for `http` and `https` type probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#host ContainerApp#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The number of seconds elapsed after the container has started before the probe is initiated.
	//
	// Possible values are between `0` and `60`. Defaults to `0` seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#initial_delay ContainerApp#initial_delay}
	InitialDelay *float64 `field:"optional" json:"initialDelay" yaml:"initialDelay"`
	// How often, in seconds, the probe should run. Possible values are between `1` and `240`. Defaults to `10`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#interval_seconds ContainerApp#interval_seconds}
	IntervalSeconds *float64 `field:"optional" json:"intervalSeconds" yaml:"intervalSeconds"`
	// The URI to use with the `host` for http type probes.
	//
	// Not valid for `TCP` type probes. Defaults to `/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#path ContainerApp#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Time in seconds after which the probe times out. Possible values are between `1` an `240`. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/container_app#timeout ContainerApp#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

