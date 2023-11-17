// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabricmanagedcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricManagedClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#client_connection_port ServiceFabricManagedCluster#client_connection_port}.
	ClientConnectionPort *float64 `field:"required" json:"clientConnectionPort" yaml:"clientConnectionPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#http_gateway_port ServiceFabricManagedCluster#http_gateway_port}.
	HttpGatewayPort *float64 `field:"required" json:"httpGatewayPort" yaml:"httpGatewayPort"`
	// lb_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#lb_rule ServiceFabricManagedCluster#lb_rule}
	LbRule interface{} `field:"required" json:"lbRule" yaml:"lbRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#location ServiceFabricManagedCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#name ServiceFabricManagedCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#resource_group_name ServiceFabricManagedCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#authentication ServiceFabricManagedCluster#authentication}
	Authentication *ServiceFabricManagedClusterAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#backup_service_enabled ServiceFabricManagedCluster#backup_service_enabled}.
	BackupServiceEnabled interface{} `field:"optional" json:"backupServiceEnabled" yaml:"backupServiceEnabled"`
	// custom_fabric_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#custom_fabric_setting ServiceFabricManagedCluster#custom_fabric_setting}
	CustomFabricSetting interface{} `field:"optional" json:"customFabricSetting" yaml:"customFabricSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#dns_name ServiceFabricManagedCluster#dns_name}.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#dns_service_enabled ServiceFabricManagedCluster#dns_service_enabled}.
	DnsServiceEnabled interface{} `field:"optional" json:"dnsServiceEnabled" yaml:"dnsServiceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#id ServiceFabricManagedCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// node_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#node_type ServiceFabricManagedCluster#node_type}
	NodeType interface{} `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#password ServiceFabricManagedCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#sku ServiceFabricManagedCluster#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#tags ServiceFabricManagedCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#timeouts ServiceFabricManagedCluster#timeouts}
	Timeouts *ServiceFabricManagedClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#upgrade_wave ServiceFabricManagedCluster#upgrade_wave}.
	UpgradeWave *string `field:"optional" json:"upgradeWave" yaml:"upgradeWave"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/service_fabric_managed_cluster#username ServiceFabricManagedCluster#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

