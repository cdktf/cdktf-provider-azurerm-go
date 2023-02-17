package batchpool


type BatchPoolStartTaskUserIdentityAutoUser struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#elevation_level BatchPool#elevation_level}.
	ElevationLevel *string `field:"optional" json:"elevationLevel" yaml:"elevationLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#scope BatchPool#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

