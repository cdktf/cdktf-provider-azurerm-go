package loadtest


type LoadTestIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/load_test#type LoadTest#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

