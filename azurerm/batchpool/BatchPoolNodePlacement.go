package batchpool


type BatchPoolNodePlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/batch_pool#policy BatchPool#policy}.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

