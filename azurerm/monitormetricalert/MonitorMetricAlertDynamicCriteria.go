package monitormetricalert


type MonitorMetricAlertDynamicCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#aggregation MonitorMetricAlert#aggregation}.
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#alert_sensitivity MonitorMetricAlert#alert_sensitivity}.
	AlertSensitivity *string `field:"required" json:"alertSensitivity" yaml:"alertSensitivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#metric_name MonitorMetricAlert#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#metric_namespace MonitorMetricAlert#metric_namespace}.
	MetricNamespace *string `field:"required" json:"metricNamespace" yaml:"metricNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#operator MonitorMetricAlert#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#dimension MonitorMetricAlert#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#evaluation_failure_count MonitorMetricAlert#evaluation_failure_count}.
	EvaluationFailureCount *float64 `field:"optional" json:"evaluationFailureCount" yaml:"evaluationFailureCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#evaluation_total_count MonitorMetricAlert#evaluation_total_count}.
	EvaluationTotalCount *float64 `field:"optional" json:"evaluationTotalCount" yaml:"evaluationTotalCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#ignore_data_before MonitorMetricAlert#ignore_data_before}.
	IgnoreDataBefore *string `field:"optional" json:"ignoreDataBefore" yaml:"ignoreDataBefore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/monitor_metric_alert#skip_metric_validation MonitorMetricAlert#skip_metric_validation}.
	SkipMetricValidation interface{} `field:"optional" json:"skipMetricValidation" yaml:"skipMetricValidation"`
}

