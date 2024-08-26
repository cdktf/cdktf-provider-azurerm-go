// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchsharedprivatelinkservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SearchSharedPrivateLinkServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#name SearchSharedPrivateLinkService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#search_service_id SearchSharedPrivateLinkService#search_service_id}.
	SearchServiceId *string `field:"required" json:"searchServiceId" yaml:"searchServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#subresource_name SearchSharedPrivateLinkService#subresource_name}.
	SubresourceName *string `field:"required" json:"subresourceName" yaml:"subresourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#target_resource_id SearchSharedPrivateLinkService#target_resource_id}.
	TargetResourceId *string `field:"required" json:"targetResourceId" yaml:"targetResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#id SearchSharedPrivateLinkService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#request_message SearchSharedPrivateLinkService#request_message}.
	RequestMessage *string `field:"optional" json:"requestMessage" yaml:"requestMessage"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/search_shared_private_link_service#timeouts SearchSharedPrivateLinkService#timeouts}
	Timeouts *SearchSharedPrivateLinkServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

