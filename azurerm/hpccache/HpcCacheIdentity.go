package hpccache


type HpcCacheIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#identity_ids HpcCache#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#type HpcCache#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

