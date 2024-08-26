// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lboutboundrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbOutboundRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#backend_address_pool_id LbOutboundRule#backend_address_pool_id}.
	BackendAddressPoolId *string `field:"required" json:"backendAddressPoolId" yaml:"backendAddressPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#loadbalancer_id LbOutboundRule#loadbalancer_id}.
	LoadbalancerId *string `field:"required" json:"loadbalancerId" yaml:"loadbalancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#name LbOutboundRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#protocol LbOutboundRule#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#allocated_outbound_ports LbOutboundRule#allocated_outbound_ports}.
	AllocatedOutboundPorts *float64 `field:"optional" json:"allocatedOutboundPorts" yaml:"allocatedOutboundPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#enable_tcp_reset LbOutboundRule#enable_tcp_reset}.
	EnableTcpReset interface{} `field:"optional" json:"enableTcpReset" yaml:"enableTcpReset"`
	// frontend_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#frontend_ip_configuration LbOutboundRule#frontend_ip_configuration}
	FrontendIpConfiguration interface{} `field:"optional" json:"frontendIpConfiguration" yaml:"frontendIpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#id LbOutboundRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#idle_timeout_in_minutes LbOutboundRule#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/lb_outbound_rule#timeouts LbOutboundRule#timeouts}
	Timeouts *LbOutboundRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

