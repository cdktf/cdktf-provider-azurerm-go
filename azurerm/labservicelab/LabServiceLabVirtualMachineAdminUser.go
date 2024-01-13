// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab


type LabServiceLabVirtualMachineAdminUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/lab_service_lab#password LabServiceLab#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/lab_service_lab#username LabServiceLab#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

