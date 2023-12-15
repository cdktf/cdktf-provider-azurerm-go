// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceconfiguration


type MaintenanceConfigurationInstallPatchesWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/maintenance_configuration#classifications_to_include MaintenanceConfiguration#classifications_to_include}.
	ClassificationsToInclude *[]*string `field:"optional" json:"classificationsToInclude" yaml:"classificationsToInclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/maintenance_configuration#kb_numbers_to_exclude MaintenanceConfiguration#kb_numbers_to_exclude}.
	KbNumbersToExclude *[]*string `field:"optional" json:"kbNumbersToExclude" yaml:"kbNumbersToExclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/maintenance_configuration#kb_numbers_to_include MaintenanceConfiguration#kb_numbers_to_include}.
	KbNumbersToInclude *[]*string `field:"optional" json:"kbNumbersToInclude" yaml:"kbNumbersToInclude"`
}

