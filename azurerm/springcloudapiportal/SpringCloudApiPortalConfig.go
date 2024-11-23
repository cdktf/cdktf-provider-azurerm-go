// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudapiportal

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudApiPortalConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#name SpringCloudApiPortal#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#spring_cloud_service_id SpringCloudApiPortal#spring_cloud_service_id}.
	SpringCloudServiceId *string `field:"required" json:"springCloudServiceId" yaml:"springCloudServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#api_try_out_enabled SpringCloudApiPortal#api_try_out_enabled}.
	ApiTryOutEnabled interface{} `field:"optional" json:"apiTryOutEnabled" yaml:"apiTryOutEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#gateway_ids SpringCloudApiPortal#gateway_ids}.
	GatewayIds *[]*string `field:"optional" json:"gatewayIds" yaml:"gatewayIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#https_only_enabled SpringCloudApiPortal#https_only_enabled}.
	HttpsOnlyEnabled interface{} `field:"optional" json:"httpsOnlyEnabled" yaml:"httpsOnlyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#id SpringCloudApiPortal#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#instance_count SpringCloudApiPortal#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#public_network_access_enabled SpringCloudApiPortal#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// sso block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#sso SpringCloudApiPortal#sso}
	Sso *SpringCloudApiPortalSso `field:"optional" json:"sso" yaml:"sso"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/spring_cloud_api_portal#timeouts SpringCloudApiPortal#timeouts}
	Timeouts *SpringCloudApiPortalTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

