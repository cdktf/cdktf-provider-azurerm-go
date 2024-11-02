// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementidentityprovidertwitter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementIdentityProviderTwitterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#api_key ApiManagementIdentityProviderTwitter#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#api_management_name ApiManagementIdentityProviderTwitter#api_management_name}.
	ApiManagementName *string `field:"required" json:"apiManagementName" yaml:"apiManagementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#api_secret_key ApiManagementIdentityProviderTwitter#api_secret_key}.
	ApiSecretKey *string `field:"required" json:"apiSecretKey" yaml:"apiSecretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#resource_group_name ApiManagementIdentityProviderTwitter#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#id ApiManagementIdentityProviderTwitter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/api_management_identity_provider_twitter#timeouts ApiManagementIdentityProviderTwitter#timeouts}
	Timeouts *ApiManagementIdentityProviderTwitterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

