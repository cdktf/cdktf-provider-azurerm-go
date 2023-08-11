package monitorautoscalesetting


type MonitorAutoscaleSettingPredictive struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_autoscale_setting#scale_mode MonitorAutoscaleSetting#scale_mode}.
	ScaleMode *string `field:"required" json:"scaleMode" yaml:"scaleMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/monitor_autoscale_setting#look_ahead_time MonitorAutoscaleSetting#look_ahead_time}.
	LookAheadTime *string `field:"optional" json:"lookAheadTime" yaml:"lookAheadTime"`
}

