package monitorautoscalesetting


type MonitorAutoscaleSettingProfileRuleMetricTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#metric_name MonitorAutoscaleSetting#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#metric_resource_id MonitorAutoscaleSetting#metric_resource_id}.
	MetricResourceId *string `field:"required" json:"metricResourceId" yaml:"metricResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#operator MonitorAutoscaleSetting#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#statistic MonitorAutoscaleSetting#statistic}.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#threshold MonitorAutoscaleSetting#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#time_aggregation MonitorAutoscaleSetting#time_aggregation}.
	TimeAggregation *string `field:"required" json:"timeAggregation" yaml:"timeAggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#time_grain MonitorAutoscaleSetting#time_grain}.
	TimeGrain *string `field:"required" json:"timeGrain" yaml:"timeGrain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#time_window MonitorAutoscaleSetting#time_window}.
	TimeWindow *string `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#dimensions MonitorAutoscaleSetting#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#divide_by_instance_count MonitorAutoscaleSetting#divide_by_instance_count}.
	DivideByInstanceCount interface{} `field:"optional" json:"divideByInstanceCount" yaml:"divideByInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/monitor_autoscale_setting#metric_namespace MonitorAutoscaleSetting#metric_namespace}.
	MetricNamespace *string `field:"optional" json:"metricNamespace" yaml:"metricNamespace"`
}

