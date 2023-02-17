package loadtest


type LoadTestIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/load_test#type LoadTest#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

