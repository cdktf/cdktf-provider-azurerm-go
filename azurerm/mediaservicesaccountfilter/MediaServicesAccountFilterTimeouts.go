package mediaservicesaccountfilter


type MediaServicesAccountFilterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#create MediaServicesAccountFilter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#delete MediaServicesAccountFilter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#read MediaServicesAccountFilter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#update MediaServicesAccountFilter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
