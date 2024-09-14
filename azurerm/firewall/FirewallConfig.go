// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewall

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirewallConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#location Firewall#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#name Firewall#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#resource_group_name Firewall#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#sku_name Firewall#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#sku_tier Firewall#sku_tier}.
	SkuTier *string `field:"required" json:"skuTier" yaml:"skuTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#dns_proxy_enabled Firewall#dns_proxy_enabled}.
	DnsProxyEnabled interface{} `field:"optional" json:"dnsProxyEnabled" yaml:"dnsProxyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#dns_servers Firewall#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#firewall_policy_id Firewall#firewall_policy_id}.
	FirewallPolicyId *string `field:"optional" json:"firewallPolicyId" yaml:"firewallPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#id Firewall#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#ip_configuration Firewall#ip_configuration}
	IpConfiguration interface{} `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// management_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#management_ip_configuration Firewall#management_ip_configuration}
	ManagementIpConfiguration *FirewallManagementIpConfiguration `field:"optional" json:"managementIpConfiguration" yaml:"managementIpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#private_ip_ranges Firewall#private_ip_ranges}.
	PrivateIpRanges *[]*string `field:"optional" json:"privateIpRanges" yaml:"privateIpRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#tags Firewall#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#threat_intel_mode Firewall#threat_intel_mode}.
	ThreatIntelMode *string `field:"optional" json:"threatIntelMode" yaml:"threatIntelMode"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#timeouts Firewall#timeouts}
	Timeouts *FirewallTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_hub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#virtual_hub Firewall#virtual_hub}
	VirtualHub *FirewallVirtualHub `field:"optional" json:"virtualHub" yaml:"virtualHub"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/firewall#zones Firewall#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

