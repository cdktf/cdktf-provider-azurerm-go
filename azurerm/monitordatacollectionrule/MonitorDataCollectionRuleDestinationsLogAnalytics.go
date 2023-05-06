package monitordatacollectionrule


type MonitorDataCollectionRuleDestinationsLogAnalytics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_data_collection_rule#name MonitorDataCollectionRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/monitor_data_collection_rule#workspace_resource_id MonitorDataCollectionRule#workspace_resource_id}.
	WorkspaceResourceId *string `field:"required" json:"workspaceResourceId" yaml:"workspaceResourceId"`
}

