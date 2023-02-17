package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRule struct {
	// metric_trigger block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#metric_trigger MonitorAutoscaleSetting#metric_trigger}
	MetricTrigger *MonitorAutoscaleSettingProfileRuleMetricTrigger `field:"required" json:"metricTrigger" yaml:"metricTrigger"`
	// scale_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_autoscale_setting#scale_action MonitorAutoscaleSetting#scale_action}
	ScaleAction *MonitorAutoscaleSettingProfileRuleScaleAction `field:"required" json:"scaleAction" yaml:"scaleAction"`
}

