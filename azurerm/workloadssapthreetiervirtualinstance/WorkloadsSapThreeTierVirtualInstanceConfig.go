// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadsSapThreeTierVirtualInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#app_location WorkloadsSapThreeTierVirtualInstance#app_location}.
	AppLocation *string `field:"required" json:"appLocation" yaml:"appLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#environment WorkloadsSapThreeTierVirtualInstance#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#location WorkloadsSapThreeTierVirtualInstance#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#name WorkloadsSapThreeTierVirtualInstance#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#resource_group_name WorkloadsSapThreeTierVirtualInstance#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#sap_fqdn WorkloadsSapThreeTierVirtualInstance#sap_fqdn}.
	SapFqdn *string `field:"required" json:"sapFqdn" yaml:"sapFqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#sap_product WorkloadsSapThreeTierVirtualInstance#sap_product}.
	SapProduct *string `field:"required" json:"sapProduct" yaml:"sapProduct"`
	// three_tier_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#three_tier_configuration WorkloadsSapThreeTierVirtualInstance#three_tier_configuration}
	ThreeTierConfiguration *WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration `field:"required" json:"threeTierConfiguration" yaml:"threeTierConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#id WorkloadsSapThreeTierVirtualInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#identity WorkloadsSapThreeTierVirtualInstance#identity}
	Identity *WorkloadsSapThreeTierVirtualInstanceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#managed_resource_group_name WorkloadsSapThreeTierVirtualInstance#managed_resource_group_name}.
	ManagedResourceGroupName *string `field:"optional" json:"managedResourceGroupName" yaml:"managedResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#managed_resources_network_access_type WorkloadsSapThreeTierVirtualInstance#managed_resources_network_access_type}.
	ManagedResourcesNetworkAccessType *string `field:"optional" json:"managedResourcesNetworkAccessType" yaml:"managedResourcesNetworkAccessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#tags WorkloadsSapThreeTierVirtualInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/workloads_sap_three_tier_virtual_instance#timeouts WorkloadsSapThreeTierVirtualInstance#timeouts}
	Timeouts *WorkloadsSapThreeTierVirtualInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

