// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicationrecoveryplan


type SiteRecoveryReplicationRecoveryPlanBootRecoveryGroup struct {
	// post_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/site_recovery_replication_recovery_plan#post_action SiteRecoveryReplicationRecoveryPlan#post_action}
	PostAction interface{} `field:"optional" json:"postAction" yaml:"postAction"`
	// pre_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/site_recovery_replication_recovery_plan#pre_action SiteRecoveryReplicationRecoveryPlan#pre_action}
	PreAction interface{} `field:"optional" json:"preAction" yaml:"preAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/site_recovery_replication_recovery_plan#replicated_protected_items SiteRecoveryReplicationRecoveryPlan#replicated_protected_items}.
	ReplicatedProtectedItems *[]*string `field:"optional" json:"replicatedProtectedItems" yaml:"replicatedProtectedItems"`
}

