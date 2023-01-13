package mediaservicesaccountfilter


type MediaServicesAccountFilterTrackSelection struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_services_account_filter#condition MediaServicesAccountFilter#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
}

