package siterecoveryreplicationrecoveryplan


type SiteRecoveryReplicationRecoveryPlanRecoveryGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replication_recovery_plan#type SiteRecoveryReplicationRecoveryPlan#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// post_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replication_recovery_plan#post_action SiteRecoveryReplicationRecoveryPlan#post_action}
	PostAction interface{} `field:"optional" json:"postAction" yaml:"postAction"`
	// pre_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replication_recovery_plan#pre_action SiteRecoveryReplicationRecoveryPlan#pre_action}
	PreAction interface{} `field:"optional" json:"preAction" yaml:"preAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replication_recovery_plan#replicated_protected_items SiteRecoveryReplicationRecoveryPlan#replicated_protected_items}.
	ReplicatedProtectedItems *[]*string `field:"optional" json:"replicatedProtectedItems" yaml:"replicatedProtectedItems"`
}
