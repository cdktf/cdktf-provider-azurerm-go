// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#admin_username OrchestratedVirtualMachineScaleSet#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#admin_password OrchestratedVirtualMachineScaleSet#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// admin_ssh_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#admin_ssh_key OrchestratedVirtualMachineScaleSet#admin_ssh_key}
	AdminSshKey interface{} `field:"optional" json:"adminSshKey" yaml:"adminSshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#computer_name_prefix OrchestratedVirtualMachineScaleSet#computer_name_prefix}.
	ComputerNamePrefix *string `field:"optional" json:"computerNamePrefix" yaml:"computerNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#disable_password_authentication OrchestratedVirtualMachineScaleSet#disable_password_authentication}.
	DisablePasswordAuthentication interface{} `field:"optional" json:"disablePasswordAuthentication" yaml:"disablePasswordAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#patch_assessment_mode OrchestratedVirtualMachineScaleSet#patch_assessment_mode}.
	PatchAssessmentMode *string `field:"optional" json:"patchAssessmentMode" yaml:"patchAssessmentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#patch_mode OrchestratedVirtualMachineScaleSet#patch_mode}.
	PatchMode *string `field:"optional" json:"patchMode" yaml:"patchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#provision_vm_agent OrchestratedVirtualMachineScaleSet#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/orchestrated_virtual_machine_scale_set#secret OrchestratedVirtualMachineScaleSet#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

