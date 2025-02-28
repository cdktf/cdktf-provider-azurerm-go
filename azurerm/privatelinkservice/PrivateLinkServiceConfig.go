// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatelinkservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivateLinkServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#load_balancer_frontend_ip_configuration_ids PrivateLinkService#load_balancer_frontend_ip_configuration_ids}.
	LoadBalancerFrontendIpConfigurationIds *[]*string `field:"required" json:"loadBalancerFrontendIpConfigurationIds" yaml:"loadBalancerFrontendIpConfigurationIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#location PrivateLinkService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#name PrivateLinkService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// nat_ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#nat_ip_configuration PrivateLinkService#nat_ip_configuration}
	NatIpConfiguration interface{} `field:"required" json:"natIpConfiguration" yaml:"natIpConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#resource_group_name PrivateLinkService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#auto_approval_subscription_ids PrivateLinkService#auto_approval_subscription_ids}.
	AutoApprovalSubscriptionIds *[]*string `field:"optional" json:"autoApprovalSubscriptionIds" yaml:"autoApprovalSubscriptionIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#enable_proxy_protocol PrivateLinkService#enable_proxy_protocol}.
	EnableProxyProtocol interface{} `field:"optional" json:"enableProxyProtocol" yaml:"enableProxyProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#fqdns PrivateLinkService#fqdns}.
	Fqdns *[]*string `field:"optional" json:"fqdns" yaml:"fqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#id PrivateLinkService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#tags PrivateLinkService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#timeouts PrivateLinkService#timeouts}
	Timeouts *PrivateLinkServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/private_link_service#visibility_subscription_ids PrivateLinkService#visibility_subscription_ids}.
	VisibilitySubscriptionIds *[]*string `field:"optional" json:"visibilitySubscriptionIds" yaml:"visibilitySubscriptionIds"`
}

