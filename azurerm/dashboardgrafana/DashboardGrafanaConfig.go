package dashboardgrafana

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardGrafanaConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#location DashboardGrafana#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#name DashboardGrafana#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#resource_group_name DashboardGrafana#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#api_key_enabled DashboardGrafana#api_key_enabled}.
	ApiKeyEnabled interface{} `field:"optional" json:"apiKeyEnabled" yaml:"apiKeyEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#auto_generated_domain_name_label_scope DashboardGrafana#auto_generated_domain_name_label_scope}.
	AutoGeneratedDomainNameLabelScope *string `field:"optional" json:"autoGeneratedDomainNameLabelScope" yaml:"autoGeneratedDomainNameLabelScope"`
	// azure_monitor_workspace_integrations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#azure_monitor_workspace_integrations DashboardGrafana#azure_monitor_workspace_integrations}
	AzureMonitorWorkspaceIntegrations interface{} `field:"optional" json:"azureMonitorWorkspaceIntegrations" yaml:"azureMonitorWorkspaceIntegrations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#deterministic_outbound_ip_enabled DashboardGrafana#deterministic_outbound_ip_enabled}.
	DeterministicOutboundIpEnabled interface{} `field:"optional" json:"deterministicOutboundIpEnabled" yaml:"deterministicOutboundIpEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#id DashboardGrafana#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#identity DashboardGrafana#identity}
	Identity *DashboardGrafanaIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#public_network_access_enabled DashboardGrafana#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#sku DashboardGrafana#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#tags DashboardGrafana#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#timeouts DashboardGrafana#timeouts}
	Timeouts *DashboardGrafanaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dashboard_grafana#zone_redundancy_enabled DashboardGrafana#zone_redundancy_enabled}.
	ZoneRedundancyEnabled interface{} `field:"optional" json:"zoneRedundancyEnabled" yaml:"zoneRedundancyEnabled"`
}
