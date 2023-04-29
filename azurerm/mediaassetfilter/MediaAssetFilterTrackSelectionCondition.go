package mediaassetfilter


type MediaAssetFilterTrackSelectionCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/media_asset_filter#operation MediaAssetFilter#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/media_asset_filter#property MediaAssetFilter#property}.
	Property *string `field:"optional" json:"property" yaml:"property"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/media_asset_filter#value MediaAssetFilter#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

