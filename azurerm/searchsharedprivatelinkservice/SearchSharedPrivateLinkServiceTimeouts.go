package searchsharedprivatelinkservice


type SearchSharedPrivateLinkServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/search_shared_private_link_service#create SearchSharedPrivateLinkService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/search_shared_private_link_service#delete SearchSharedPrivateLinkService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/search_shared_private_link_service#read SearchSharedPrivateLinkService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/search_shared_private_link_service#update SearchSharedPrivateLinkService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
