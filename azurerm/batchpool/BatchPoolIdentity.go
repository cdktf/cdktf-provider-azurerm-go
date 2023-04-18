package batchpool


type BatchPoolIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/batch_pool#identity_ids BatchPool#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

