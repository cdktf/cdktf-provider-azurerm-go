// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceplan


type LabServicePlanDefaultAutoShutdown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_plan#disconnect_delay LabServicePlan#disconnect_delay}.
	DisconnectDelay *string `field:"optional" json:"disconnectDelay" yaml:"disconnectDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_plan#idle_delay LabServicePlan#idle_delay}.
	IdleDelay *string `field:"optional" json:"idleDelay" yaml:"idleDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_plan#no_connect_delay LabServicePlan#no_connect_delay}.
	NoConnectDelay *string `field:"optional" json:"noConnectDelay" yaml:"noConnectDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_plan#shutdown_on_idle LabServicePlan#shutdown_on_idle}.
	ShutdownOnIdle *string `field:"optional" json:"shutdownOnIdle" yaml:"shutdownOnIdle"`
}

