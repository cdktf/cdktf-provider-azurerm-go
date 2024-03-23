// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapphybridconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebAppHybridConnectionConfig struct {
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
	// The hostname of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#hostname WebAppHybridConnection#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The port to use for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#port WebAppHybridConnection#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The ID of the Relay Hybrid Connection to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#relay_id WebAppHybridConnection#relay_id}
	RelayId *string `field:"required" json:"relayId" yaml:"relayId"`
	// The ID of the Web App for this Hybrid Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#web_app_id WebAppHybridConnection#web_app_id}
	WebAppId *string `field:"required" json:"webAppId" yaml:"webAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#id WebAppHybridConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the Relay key with `Send` permission to use. Defaults to `RootManageSharedAccessKey`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#send_key_name WebAppHybridConnection#send_key_name}
	SendKeyName *string `field:"optional" json:"sendKeyName" yaml:"sendKeyName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/web_app_hybrid_connection#timeouts WebAppHybridConnection#timeouts}
	Timeouts *WebAppHybridConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

