package monitorautoscalesetting


type MonitorAutoscaleSettingProfileCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#default MonitorAutoscaleSetting#default}.
	Default *float64 `field:"required" json:"default" yaml:"default"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#maximum MonitorAutoscaleSetting#maximum}.
	Maximum *float64 `field:"required" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#minimum MonitorAutoscaleSetting#minimum}.
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
}

