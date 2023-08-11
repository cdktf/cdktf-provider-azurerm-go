package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/site_recovery_replicated_vm#create SiteRecoveryReplicatedVm#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/site_recovery_replicated_vm#delete SiteRecoveryReplicatedVm#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/site_recovery_replicated_vm#read SiteRecoveryReplicatedVm#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/site_recovery_replicated_vm#update SiteRecoveryReplicatedVm#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

