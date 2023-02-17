package batchpool


type BatchPoolNodePlacement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#policy BatchPool#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

