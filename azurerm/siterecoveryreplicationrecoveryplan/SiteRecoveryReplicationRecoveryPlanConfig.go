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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#name SiteRecoveryReplicationRecoveryPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#recovery_vault_id SiteRecoveryReplicationRecoveryPlan#recovery_vault_id}.
	RecoveryVaultId *string `field:"required" json:"recoveryVaultId" yaml:"recoveryVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#source_recovery_fabric_id SiteRecoveryReplicationRecoveryPlan#source_recovery_fabric_id}.
	SourceRecoveryFabricId *string `field:"required" json:"sourceRecoveryFabricId" yaml:"sourceRecoveryFabricId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#target_recovery_fabric_id SiteRecoveryReplicationRecoveryPlan#target_recovery_fabric_id}.
	TargetRecoveryFabricId *string `field:"required" json:"targetRecoveryFabricId" yaml:"targetRecoveryFabricId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#id SiteRecoveryReplicationRecoveryPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// recovery_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#recovery_group SiteRecoveryReplicationRecoveryPlan#recovery_group}
	RecoveryGroup interface{} `field:"optional" json:"recoveryGroup" yaml:"recoveryGroup"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/site_recovery_replication_recovery_plan#timeouts SiteRecoveryReplicationRecoveryPlan#timeouts}
	Timeouts *SiteRecoveryReplicationRecoveryPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

