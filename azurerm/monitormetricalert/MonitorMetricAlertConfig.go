package monitormetricalert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorMetricAlertConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#name MonitorMetricAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#resource_group_name MonitorMetricAlert#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#scopes MonitorMetricAlert#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#action MonitorMetricAlert#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// application_insights_web_test_location_availability_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#application_insights_web_test_location_availability_criteria MonitorMetricAlert#application_insights_web_test_location_availability_criteria}
	ApplicationInsightsWebTestLocationAvailabilityCriteria *MonitorMetricAlertApplicationInsightsWebTestLocationAvailabilityCriteria `field:"optional" json:"applicationInsightsWebTestLocationAvailabilityCriteria" yaml:"applicationInsightsWebTestLocationAvailabilityCriteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#auto_mitigate MonitorMetricAlert#auto_mitigate}.
	AutoMitigate interface{} `field:"optional" json:"autoMitigate" yaml:"autoMitigate"`
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#criteria MonitorMetricAlert#criteria}
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#description MonitorMetricAlert#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamic_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#dynamic_criteria MonitorMetricAlert#dynamic_criteria}
	DynamicCriteria *MonitorMetricAlertDynamicCriteria `field:"optional" json:"dynamicCriteria" yaml:"dynamicCriteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#enabled MonitorMetricAlert#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#frequency MonitorMetricAlert#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#id MonitorMetricAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#severity MonitorMetricAlert#severity}.
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#tags MonitorMetricAlert#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The location of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#target_resource_location MonitorMetricAlert#target_resource_location}
	TargetResourceLocation *string `field:"optional" json:"targetResourceLocation" yaml:"targetResourceLocation"`
	// The resource type (e.g. Microsoft.Compute/virtualMachines) of the target pluginsdk. Required when using subscription, resource group scope or multiple scopes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#target_resource_type MonitorMetricAlert#target_resource_type}
	TargetResourceType *string `field:"optional" json:"targetResourceType" yaml:"targetResourceType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#timeouts MonitorMetricAlert#timeouts}
	Timeouts *MonitorMetricAlertTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_metric_alert#window_size MonitorMetricAlert#window_size}.
	WindowSize *string `field:"optional" json:"windowSize" yaml:"windowSize"`
}

