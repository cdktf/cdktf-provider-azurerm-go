// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicationrecoveryplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryReplicationRecoveryPlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// boot_recovery_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#boot_recovery_group SiteRecoveryReplicationRecoveryPlan#boot_recovery_group}
	BootRecoveryGroup interface{} `field:"required" json:"bootRecoveryGroup" yaml:"bootRecoveryGroup"`
	// failover_recovery_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#failover_recovery_group SiteRecoveryReplicationRecoveryPlan#failover_recovery_group}
	FailoverRecoveryGroup *SiteRecoveryReplicationRecoveryPlanFailoverRecoveryGroup `field:"required" json:"failoverRecoveryGroup" yaml:"failoverRecoveryGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#name SiteRecoveryReplicationRecoveryPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#recovery_vault_id SiteRecoveryReplicationRecoveryPlan#recovery_vault_id}.
	RecoveryVaultId *string `field:"required" json:"recoveryVaultId" yaml:"recoveryVaultId"`
	// shutdown_recovery_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#shutdown_recovery_group SiteRecoveryReplicationRecoveryPlan#shutdown_recovery_group}
	ShutdownRecoveryGroup *SiteRecoveryReplicationRecoveryPlanShutdownRecoveryGroup `field:"required" json:"shutdownRecoveryGroup" yaml:"shutdownRecoveryGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#source_recovery_fabric_id SiteRecoveryReplicationRecoveryPlan#source_recovery_fabric_id}.
	SourceRecoveryFabricId *string `field:"required" json:"sourceRecoveryFabricId" yaml:"sourceRecoveryFabricId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#target_recovery_fabric_id SiteRecoveryReplicationRecoveryPlan#target_recovery_fabric_id}.
	TargetRecoveryFabricId *string `field:"required" json:"targetRecoveryFabricId" yaml:"targetRecoveryFabricId"`
	// azure_to_azure_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#azure_to_azure_settings SiteRecoveryReplicationRecoveryPlan#azure_to_azure_settings}
	AzureToAzureSettings *SiteRecoveryReplicationRecoveryPlanAzureToAzureSettings `field:"optional" json:"azureToAzureSettings" yaml:"azureToAzureSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#id SiteRecoveryReplicationRecoveryPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/site_recovery_replication_recovery_plan#timeouts SiteRecoveryReplicationRecoveryPlan#timeouts}
	Timeouts *SiteRecoveryReplicationRecoveryPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

