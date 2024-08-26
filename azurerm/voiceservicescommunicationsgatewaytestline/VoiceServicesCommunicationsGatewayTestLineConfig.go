// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package voiceservicescommunicationsgatewaytestline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VoiceServicesCommunicationsGatewayTestLineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#location VoiceServicesCommunicationsGatewayTestLine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#name VoiceServicesCommunicationsGatewayTestLine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#phone_number VoiceServicesCommunicationsGatewayTestLine#phone_number}.
	PhoneNumber *string `field:"required" json:"phoneNumber" yaml:"phoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#purpose VoiceServicesCommunicationsGatewayTestLine#purpose}.
	Purpose *string `field:"required" json:"purpose" yaml:"purpose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#voice_services_communications_gateway_id VoiceServicesCommunicationsGatewayTestLine#voice_services_communications_gateway_id}.
	VoiceServicesCommunicationsGatewayId *string `field:"required" json:"voiceServicesCommunicationsGatewayId" yaml:"voiceServicesCommunicationsGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#id VoiceServicesCommunicationsGatewayTestLine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#tags VoiceServicesCommunicationsGatewayTestLine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/voice_services_communications_gateway_test_line#timeouts VoiceServicesCommunicationsGatewayTestLine#timeouts}
	Timeouts *VoiceServicesCommunicationsGatewayTestLineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

