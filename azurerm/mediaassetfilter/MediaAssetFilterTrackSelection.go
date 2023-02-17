package mediaassetfilter


type MediaAssetFilterTrackSelection struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_asset_filter#condition MediaAssetFilter#condition}
	Condition interface{} `field:"required" json:"condition" yaml:"condition"`
}

