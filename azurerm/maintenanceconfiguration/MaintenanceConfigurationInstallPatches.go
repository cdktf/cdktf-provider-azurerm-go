// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceconfiguration


type MaintenanceConfigurationInstallPatches struct {
	// linux block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/maintenance_configuration#linux MaintenanceConfiguration#linux}
	Linux interface{} `field:"optional" json:"linux" yaml:"linux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/maintenance_configuration#reboot MaintenanceConfiguration#reboot}.
	Reboot *string `field:"optional" json:"reboot" yaml:"reboot"`
	// windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/maintenance_configuration#windows MaintenanceConfiguration#windows}
	Windows interface{} `field:"optional" json:"windows" yaml:"windows"`
}

