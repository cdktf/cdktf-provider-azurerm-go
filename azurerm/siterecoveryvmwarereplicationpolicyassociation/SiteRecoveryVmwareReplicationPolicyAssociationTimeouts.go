// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicationpolicyassociation


type SiteRecoveryVmwareReplicationPolicyAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_vmware_replication_policy_association#create SiteRecoveryVmwareReplicationPolicyAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_vmware_replication_policy_association#delete SiteRecoveryVmwareReplicationPolicyAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/site_recovery_vmware_replication_policy_association#read SiteRecoveryVmwareReplicationPolicyAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

