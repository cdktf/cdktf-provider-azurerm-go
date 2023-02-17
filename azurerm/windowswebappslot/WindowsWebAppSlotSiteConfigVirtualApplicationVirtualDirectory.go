package windowswebappslot


type WindowsWebAppSlotSiteConfigVirtualApplicationVirtualDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#physical_path WindowsWebAppSlot#physical_path}.
	PhysicalPath *string `field:"optional" json:"physicalPath" yaml:"physicalPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_web_app_slot#virtual_path WindowsWebAppSlot#virtual_path}.
	VirtualPath *string `field:"optional" json:"virtualPath" yaml:"virtualPath"`
}

