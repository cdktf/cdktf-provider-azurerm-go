package mobilenetworkslice


type MobileNetworkSliceSingleNetworkSliceSelectionAssistanceInformation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#slice_service_type MobileNetworkSlice#slice_service_type}.
	SliceServiceType *float64 `field:"required" json:"sliceServiceType" yaml:"sliceServiceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_slice#slice_differentiator MobileNetworkSlice#slice_differentiator}.
	SliceDifferentiator *string `field:"optional" json:"sliceDifferentiator" yaml:"sliceDifferentiator"`
}

