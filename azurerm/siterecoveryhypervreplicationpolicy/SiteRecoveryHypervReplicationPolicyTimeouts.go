package siterecoveryhypervreplicationpolicy


type SiteRecoveryHypervReplicationPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy#create SiteRecoveryHypervReplicationPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy#delete SiteRecoveryHypervReplicationPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy#read SiteRecoveryHypervReplicationPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy#update SiteRecoveryHypervReplicationPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
