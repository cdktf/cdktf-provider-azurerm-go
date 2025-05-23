// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nginxdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NginxDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#location NginxDeployment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#name NginxDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#resource_group_name NginxDeployment#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#sku NginxDeployment#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#automatic_upgrade_channel NginxDeployment#automatic_upgrade_channel}.
	AutomaticUpgradeChannel *string `field:"optional" json:"automaticUpgradeChannel" yaml:"automaticUpgradeChannel"`
	// auto_scale_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#auto_scale_profile NginxDeployment#auto_scale_profile}
	AutoScaleProfile interface{} `field:"optional" json:"autoScaleProfile" yaml:"autoScaleProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#capacity NginxDeployment#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#diagnose_support_enabled NginxDeployment#diagnose_support_enabled}.
	DiagnoseSupportEnabled interface{} `field:"optional" json:"diagnoseSupportEnabled" yaml:"diagnoseSupportEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#email NginxDeployment#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// frontend_private block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#frontend_private NginxDeployment#frontend_private}
	FrontendPrivate interface{} `field:"optional" json:"frontendPrivate" yaml:"frontendPrivate"`
	// frontend_public block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#frontend_public NginxDeployment#frontend_public}
	FrontendPublic *NginxDeploymentFrontendPublic `field:"optional" json:"frontendPublic" yaml:"frontendPublic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#id NginxDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#identity NginxDeployment#identity}
	Identity *NginxDeploymentIdentity `field:"optional" json:"identity" yaml:"identity"`
	// logging_storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#logging_storage_account NginxDeployment#logging_storage_account}
	LoggingStorageAccount interface{} `field:"optional" json:"loggingStorageAccount" yaml:"loggingStorageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#managed_resource_group NginxDeployment#managed_resource_group}.
	ManagedResourceGroup *string `field:"optional" json:"managedResourceGroup" yaml:"managedResourceGroup"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#network_interface NginxDeployment#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#tags NginxDeployment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#timeouts NginxDeployment#timeouts}
	Timeouts *NginxDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// web_application_firewall block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/nginx_deployment#web_application_firewall NginxDeployment#web_application_firewall}
	WebApplicationFirewall *NginxDeploymentWebApplicationFirewall `field:"optional" json:"webApplicationFirewall" yaml:"webApplicationFirewall"`
}

