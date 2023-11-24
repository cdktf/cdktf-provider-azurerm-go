// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubsharedaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubSharedAccessPolicyAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#iothub_name IothubSharedAccessPolicyA#iothub_name}.
	IothubName *string `field:"required" json:"iothubName" yaml:"iothubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#name IothubSharedAccessPolicyA#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#resource_group_name IothubSharedAccessPolicyA#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#device_connect IothubSharedAccessPolicyA#device_connect}.
	DeviceConnect interface{} `field:"optional" json:"deviceConnect" yaml:"deviceConnect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#id IothubSharedAccessPolicyA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#registry_read IothubSharedAccessPolicyA#registry_read}.
	RegistryRead interface{} `field:"optional" json:"registryRead" yaml:"registryRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#registry_write IothubSharedAccessPolicyA#registry_write}.
	RegistryWrite interface{} `field:"optional" json:"registryWrite" yaml:"registryWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#service_connect IothubSharedAccessPolicyA#service_connect}.
	ServiceConnect interface{} `field:"optional" json:"serviceConnect" yaml:"serviceConnect"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iothub_shared_access_policy#timeouts IothubSharedAccessPolicyA#timeouts}
	Timeouts *IothubSharedAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

