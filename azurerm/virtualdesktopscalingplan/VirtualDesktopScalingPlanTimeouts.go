package virtualdesktopscalingplan


type VirtualDesktopScalingPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/virtual_desktop_scaling_plan#create VirtualDesktopScalingPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/virtual_desktop_scaling_plan#delete VirtualDesktopScalingPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/virtual_desktop_scaling_plan#read VirtualDesktopScalingPlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/virtual_desktop_scaling_plan#update VirtualDesktopScalingPlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

