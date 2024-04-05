// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceconfiguration


type MaintenanceConfigurationInstallPatchesLinux struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/maintenance_configuration#classifications_to_include MaintenanceConfiguration#classifications_to_include}.
	ClassificationsToInclude *[]*string `field:"optional" json:"classificationsToInclude" yaml:"classificationsToInclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/maintenance_configuration#package_names_mask_to_exclude MaintenanceConfiguration#package_names_mask_to_exclude}.
	PackageNamesMaskToExclude *[]*string `field:"optional" json:"packageNamesMaskToExclude" yaml:"packageNamesMaskToExclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/maintenance_configuration#package_names_mask_to_include MaintenanceConfiguration#package_names_mask_to_include}.
	PackageNamesMaskToInclude *[]*string `field:"optional" json:"packageNamesMaskToInclude" yaml:"packageNamesMaskToInclude"`
}

