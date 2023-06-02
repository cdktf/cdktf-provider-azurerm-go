package virtualdesktopscalingplan


type VirtualDesktopScalingPlanSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#days_of_week VirtualDesktopScalingPlan#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#name VirtualDesktopScalingPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#off_peak_load_balancing_algorithm VirtualDesktopScalingPlan#off_peak_load_balancing_algorithm}.
	OffPeakLoadBalancingAlgorithm *string `field:"required" json:"offPeakLoadBalancingAlgorithm" yaml:"offPeakLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#off_peak_start_time VirtualDesktopScalingPlan#off_peak_start_time}.
	OffPeakStartTime *string `field:"required" json:"offPeakStartTime" yaml:"offPeakStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#peak_load_balancing_algorithm VirtualDesktopScalingPlan#peak_load_balancing_algorithm}.
	PeakLoadBalancingAlgorithm *string `field:"required" json:"peakLoadBalancingAlgorithm" yaml:"peakLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#peak_start_time VirtualDesktopScalingPlan#peak_start_time}.
	PeakStartTime *string `field:"required" json:"peakStartTime" yaml:"peakStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_capacity_threshold_percent VirtualDesktopScalingPlan#ramp_down_capacity_threshold_percent}.
	RampDownCapacityThresholdPercent *float64 `field:"required" json:"rampDownCapacityThresholdPercent" yaml:"rampDownCapacityThresholdPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_force_logoff_users VirtualDesktopScalingPlan#ramp_down_force_logoff_users}.
	RampDownForceLogoffUsers interface{} `field:"required" json:"rampDownForceLogoffUsers" yaml:"rampDownForceLogoffUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_load_balancing_algorithm VirtualDesktopScalingPlan#ramp_down_load_balancing_algorithm}.
	RampDownLoadBalancingAlgorithm *string `field:"required" json:"rampDownLoadBalancingAlgorithm" yaml:"rampDownLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_minimum_hosts_percent VirtualDesktopScalingPlan#ramp_down_minimum_hosts_percent}.
	RampDownMinimumHostsPercent *float64 `field:"required" json:"rampDownMinimumHostsPercent" yaml:"rampDownMinimumHostsPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_notification_message VirtualDesktopScalingPlan#ramp_down_notification_message}.
	RampDownNotificationMessage *string `field:"required" json:"rampDownNotificationMessage" yaml:"rampDownNotificationMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_start_time VirtualDesktopScalingPlan#ramp_down_start_time}.
	RampDownStartTime *string `field:"required" json:"rampDownStartTime" yaml:"rampDownStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_stop_hosts_when VirtualDesktopScalingPlan#ramp_down_stop_hosts_when}.
	RampDownStopHostsWhen *string `field:"required" json:"rampDownStopHostsWhen" yaml:"rampDownStopHostsWhen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_down_wait_time_minutes VirtualDesktopScalingPlan#ramp_down_wait_time_minutes}.
	RampDownWaitTimeMinutes *float64 `field:"required" json:"rampDownWaitTimeMinutes" yaml:"rampDownWaitTimeMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_up_load_balancing_algorithm VirtualDesktopScalingPlan#ramp_up_load_balancing_algorithm}.
	RampUpLoadBalancingAlgorithm *string `field:"required" json:"rampUpLoadBalancingAlgorithm" yaml:"rampUpLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_up_start_time VirtualDesktopScalingPlan#ramp_up_start_time}.
	RampUpStartTime *string `field:"required" json:"rampUpStartTime" yaml:"rampUpStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_up_capacity_threshold_percent VirtualDesktopScalingPlan#ramp_up_capacity_threshold_percent}.
	RampUpCapacityThresholdPercent *float64 `field:"optional" json:"rampUpCapacityThresholdPercent" yaml:"rampUpCapacityThresholdPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/virtual_desktop_scaling_plan#ramp_up_minimum_hosts_percent VirtualDesktopScalingPlan#ramp_up_minimum_hosts_percent}.
	RampUpMinimumHostsPercent *float64 `field:"optional" json:"rampUpMinimumHostsPercent" yaml:"rampUpMinimumHostsPercent"`
}

