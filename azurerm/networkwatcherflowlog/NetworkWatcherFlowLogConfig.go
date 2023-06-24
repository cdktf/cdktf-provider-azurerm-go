package networkwatcherflowlog

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkWatcherFlowLogConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#enabled NetworkWatcherFlowLog#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#name NetworkWatcherFlowLog#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#network_security_group_id NetworkWatcherFlowLog#network_security_group_id}.
	NetworkSecurityGroupId *string `field:"required" json:"networkSecurityGroupId" yaml:"networkSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#network_watcher_name NetworkWatcherFlowLog#network_watcher_name}.
	NetworkWatcherName *string `field:"required" json:"networkWatcherName" yaml:"networkWatcherName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#resource_group_name NetworkWatcherFlowLog#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#retention_policy NetworkWatcherFlowLog#retention_policy}
	RetentionPolicy *NetworkWatcherFlowLogRetentionPolicy `field:"required" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#storage_account_id NetworkWatcherFlowLog#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#id NetworkWatcherFlowLog#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#location NetworkWatcherFlowLog#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#tags NetworkWatcherFlowLog#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#timeouts NetworkWatcherFlowLog#timeouts}
	Timeouts *NetworkWatcherFlowLogTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#traffic_analytics NetworkWatcherFlowLog#traffic_analytics}
	TrafficAnalytics *NetworkWatcherFlowLogTrafficAnalytics `field:"optional" json:"trafficAnalytics" yaml:"trafficAnalytics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/network_watcher_flow_log#version NetworkWatcherFlowLog#version}.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

