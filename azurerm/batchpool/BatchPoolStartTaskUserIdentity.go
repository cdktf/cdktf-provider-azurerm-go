package batchpool


type BatchPoolStartTaskUserIdentity struct {
	// auto_user block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_user BatchPool#auto_user}
	AutoUser *BatchPoolStartTaskUserIdentityAutoUser `field:"optional" json:"autoUser" yaml:"autoUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#user_name BatchPool#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

