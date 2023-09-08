// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktophostpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualDesktopHostPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#load_balancer_type VirtualDesktopHostPool#load_balancer_type}.
	LoadBalancerType *string `field:"required" json:"loadBalancerType" yaml:"loadBalancerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#location VirtualDesktopHostPool#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#name VirtualDesktopHostPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#resource_group_name VirtualDesktopHostPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#type VirtualDesktopHostPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#custom_rdp_properties VirtualDesktopHostPool#custom_rdp_properties}.
	CustomRdpProperties *string `field:"optional" json:"customRdpProperties" yaml:"customRdpProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#description VirtualDesktopHostPool#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#friendly_name VirtualDesktopHostPool#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#id VirtualDesktopHostPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#maximum_sessions_allowed VirtualDesktopHostPool#maximum_sessions_allowed}.
	MaximumSessionsAllowed *float64 `field:"optional" json:"maximumSessionsAllowed" yaml:"maximumSessionsAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#personal_desktop_assignment_type VirtualDesktopHostPool#personal_desktop_assignment_type}.
	PersonalDesktopAssignmentType *string `field:"optional" json:"personalDesktopAssignmentType" yaml:"personalDesktopAssignmentType"`
	// Preferred App Group type to display.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#preferred_app_group_type VirtualDesktopHostPool#preferred_app_group_type}
	PreferredAppGroupType *string `field:"optional" json:"preferredAppGroupType" yaml:"preferredAppGroupType"`
	// scheduled_agent_updates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#scheduled_agent_updates VirtualDesktopHostPool#scheduled_agent_updates}
	ScheduledAgentUpdates *VirtualDesktopHostPoolScheduledAgentUpdates `field:"optional" json:"scheduledAgentUpdates" yaml:"scheduledAgentUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#start_vm_on_connect VirtualDesktopHostPool#start_vm_on_connect}.
	StartVmOnConnect interface{} `field:"optional" json:"startVmOnConnect" yaml:"startVmOnConnect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#tags VirtualDesktopHostPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#timeouts VirtualDesktopHostPool#timeouts}
	Timeouts *VirtualDesktopHostPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/virtual_desktop_host_pool#validate_environment VirtualDesktopHostPool#validate_environment}.
	ValidateEnvironment interface{} `field:"optional" json:"validateEnvironment" yaml:"validateEnvironment"`
}

