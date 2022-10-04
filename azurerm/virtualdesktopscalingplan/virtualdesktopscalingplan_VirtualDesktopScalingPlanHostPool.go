package virtualdesktopscalingplan


type VirtualDesktopScalingPlanHostPool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#hostpool_id VirtualDesktopScalingPlan#hostpool_id}.
	HostpoolId *string `field:"required" json:"hostpoolId" yaml:"hostpoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#scaling_plan_enabled VirtualDesktopScalingPlan#scaling_plan_enabled}.
	ScalingPlanEnabled interface{} `field:"required" json:"scalingPlanEnabled" yaml:"scalingPlanEnabled"`
}

