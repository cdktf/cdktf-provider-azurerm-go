package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRuleScaleAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/monitor_autoscale_setting#cooldown MonitorAutoscaleSetting#cooldown}.
	Cooldown *string `field:"required" json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/monitor_autoscale_setting#direction MonitorAutoscaleSetting#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/monitor_autoscale_setting#type MonitorAutoscaleSetting#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/monitor_autoscale_setting#value MonitorAutoscaleSetting#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

