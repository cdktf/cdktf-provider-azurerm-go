// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermcdnfrontdoorendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermCdnFrontdoorEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/data-sources/cdn_frontdoor_endpoint#name DataAzurermCdnFrontdoorEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/data-sources/cdn_frontdoor_endpoint#profile_name DataAzurermCdnFrontdoorEndpoint#profile_name}.
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/data-sources/cdn_frontdoor_endpoint#resource_group_name DataAzurermCdnFrontdoorEndpoint#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/data-sources/cdn_frontdoor_endpoint#id DataAzurermCdnFrontdoorEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/data-sources/cdn_frontdoor_endpoint#timeouts DataAzurermCdnFrontdoorEndpoint#timeouts}
	Timeouts *DataAzurermCdnFrontdoorEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

