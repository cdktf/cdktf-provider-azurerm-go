// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpointcustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnEndpointCustomDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#cdn_endpoint_id CdnEndpointCustomDomain#cdn_endpoint_id}.
	CdnEndpointId *string `field:"required" json:"cdnEndpointId" yaml:"cdnEndpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#host_name CdnEndpointCustomDomain#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#name CdnEndpointCustomDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cdn_managed_https block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#cdn_managed_https CdnEndpointCustomDomain#cdn_managed_https}
	CdnManagedHttps *CdnEndpointCustomDomainCdnManagedHttps `field:"optional" json:"cdnManagedHttps" yaml:"cdnManagedHttps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#id CdnEndpointCustomDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#timeouts CdnEndpointCustomDomain#timeouts}
	Timeouts *CdnEndpointCustomDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_managed_https block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/cdn_endpoint_custom_domain#user_managed_https CdnEndpointCustomDomain#user_managed_https}
	UserManagedHttps *CdnEndpointCustomDomainUserManagedHttps `field:"optional" json:"userManagedHttps" yaml:"userManagedHttps"`
}

