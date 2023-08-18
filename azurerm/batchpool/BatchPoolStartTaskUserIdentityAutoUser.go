package batchpool


type BatchPoolStartTaskUserIdentityAutoUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/batch_pool#elevation_level BatchPool#elevation_level}.
	ElevationLevel *string `field:"optional" json:"elevationLevel" yaml:"elevationLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/batch_pool#scope BatchPool#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

