package searchservice


type SearchServiceIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/search_service#type SearchService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

