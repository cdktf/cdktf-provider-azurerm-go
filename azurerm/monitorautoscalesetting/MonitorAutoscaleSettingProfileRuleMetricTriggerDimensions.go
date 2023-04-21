package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRuleMetricTriggerDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/monitor_autoscale_setting#name MonitorAutoscaleSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/monitor_autoscale_setting#operator MonitorAutoscaleSetting#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/monitor_autoscale_setting#values MonitorAutoscaleSetting#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

