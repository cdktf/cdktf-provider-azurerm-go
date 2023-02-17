package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRecurrence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#days MonitorAutoscaleSetting#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#hours MonitorAutoscaleSetting#hours}.
	Hours *[]*float64 `field:"required" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#minutes MonitorAutoscaleSetting#minutes}.
	Minutes *[]*float64 `field:"required" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#timezone MonitorAutoscaleSetting#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

