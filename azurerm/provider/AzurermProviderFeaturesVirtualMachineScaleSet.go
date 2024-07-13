// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesVirtualMachineScaleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#force_delete AzurermProvider#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#reimage_on_manual_upgrade AzurermProvider#reimage_on_manual_upgrade}.
	ReimageOnManualUpgrade interface{} `field:"optional" json:"reimageOnManualUpgrade" yaml:"reimageOnManualUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#roll_instances_when_required AzurermProvider#roll_instances_when_required}.
	RollInstancesWhenRequired interface{} `field:"optional" json:"rollInstancesWhenRequired" yaml:"rollInstancesWhenRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs#scale_to_zero_before_deletion AzurermProvider#scale_to_zero_before_deletion}.
	ScaleToZeroBeforeDeletion interface{} `field:"optional" json:"scaleToZeroBeforeDeletion" yaml:"scaleToZeroBeforeDeletion"`
}

