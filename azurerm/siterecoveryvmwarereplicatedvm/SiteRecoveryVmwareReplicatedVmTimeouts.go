// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicatedvm


type SiteRecoveryVmwareReplicatedVmTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/site_recovery_vmware_replicated_vm#create SiteRecoveryVmwareReplicatedVm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/site_recovery_vmware_replicated_vm#delete SiteRecoveryVmwareReplicatedVm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/site_recovery_vmware_replicated_vm#read SiteRecoveryVmwareReplicatedVm#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/site_recovery_vmware_replicated_vm#update SiteRecoveryVmwareReplicatedVm#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

