package monitormetricalert


type MonitorMetricAlertCriteriaDimension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_metric_alert#name MonitorMetricAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_metric_alert#operator MonitorMetricAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_metric_alert#values MonitorMetricAlert#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

