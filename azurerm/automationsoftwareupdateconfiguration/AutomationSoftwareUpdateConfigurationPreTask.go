// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationPreTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_software_update_configuration#parameters AutomationSoftwareUpdateConfiguration#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_software_update_configuration#source AutomationSoftwareUpdateConfiguration#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

