// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationTargetAzureQueryTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/automation_software_update_configuration#tag AutomationSoftwareUpdateConfiguration#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/automation_software_update_configuration#values AutomationSoftwareUpdateConfiguration#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

