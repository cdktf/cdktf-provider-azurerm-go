package monitormetricalert


type MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/monitor_metric_alert#component_id MonitorMetricAlert#component_id}.
	ComponentId *string `field:"required" json:"componentId" yaml:"componentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/monitor_metric_alert#failed_location_count MonitorMetricAlert#failed_location_count}.
	FailedLocationCount *float64 `field:"required" json:"failedLocationCount" yaml:"failedLocationCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.68.0/docs/resources/monitor_metric_alert#web_test_id MonitorMetricAlert#web_test_id}.
	WebTestId *string `field:"required" json:"webTestId" yaml:"webTestId"`
}

