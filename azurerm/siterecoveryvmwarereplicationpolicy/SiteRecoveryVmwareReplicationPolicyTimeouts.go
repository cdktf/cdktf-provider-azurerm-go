// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicationpolicy


type SiteRecoveryVmwareReplicationPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_vmware_replication_policy#create SiteRecoveryVmwareReplicationPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_vmware_replication_policy#delete SiteRecoveryVmwareReplicationPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_vmware_replication_policy#read SiteRecoveryVmwareReplicationPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/site_recovery_vmware_replication_policy#update SiteRecoveryVmwareReplicationPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

