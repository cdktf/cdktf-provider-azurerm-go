package mediaservicesaccountfilter


type MediaServicesAccountFilterTrackSelectionCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#operation MediaServicesAccountFilter#operation}.
	Operation *string `field:"required" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#property MediaServicesAccountFilter#property}.
	Property *string `field:"required" json:"property" yaml:"property"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#value MediaServicesAccountFilter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

