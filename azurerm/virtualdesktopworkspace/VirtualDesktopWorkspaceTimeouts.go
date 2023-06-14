package virtualdesktopworkspace


type VirtualDesktopWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/virtual_desktop_workspace#create VirtualDesktopWorkspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/virtual_desktop_workspace#delete VirtualDesktopWorkspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/virtual_desktop_workspace#read VirtualDesktopWorkspace#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/virtual_desktop_workspace#update VirtualDesktopWorkspace#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

