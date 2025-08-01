// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package voiceservicescommunicationsgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VoiceServicesCommunicationsGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#codecs VoiceServicesCommunicationsGateway#codecs}.
	Codecs *string `field:"required" json:"codecs" yaml:"codecs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#connectivity VoiceServicesCommunicationsGateway#connectivity}.
	Connectivity *string `field:"required" json:"connectivity" yaml:"connectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#e911_type VoiceServicesCommunicationsGateway#e911_type}.
	E911Type *string `field:"required" json:"e911Type" yaml:"e911Type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#location VoiceServicesCommunicationsGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#name VoiceServicesCommunicationsGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#platforms VoiceServicesCommunicationsGateway#platforms}.
	Platforms *[]*string `field:"required" json:"platforms" yaml:"platforms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#resource_group_name VoiceServicesCommunicationsGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// service_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#service_location VoiceServicesCommunicationsGateway#service_location}
	ServiceLocation interface{} `field:"required" json:"serviceLocation" yaml:"serviceLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#api_bridge VoiceServicesCommunicationsGateway#api_bridge}.
	ApiBridge *string `field:"optional" json:"apiBridge" yaml:"apiBridge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#auto_generated_domain_name_label_scope VoiceServicesCommunicationsGateway#auto_generated_domain_name_label_scope}.
	AutoGeneratedDomainNameLabelScope *string `field:"optional" json:"autoGeneratedDomainNameLabelScope" yaml:"autoGeneratedDomainNameLabelScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#emergency_dial_strings VoiceServicesCommunicationsGateway#emergency_dial_strings}.
	EmergencyDialStrings *[]*string `field:"optional" json:"emergencyDialStrings" yaml:"emergencyDialStrings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#id VoiceServicesCommunicationsGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#microsoft_teams_voicemail_pilot_number VoiceServicesCommunicationsGateway#microsoft_teams_voicemail_pilot_number}.
	MicrosoftTeamsVoicemailPilotNumber *string `field:"optional" json:"microsoftTeamsVoicemailPilotNumber" yaml:"microsoftTeamsVoicemailPilotNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#on_prem_mcp_enabled VoiceServicesCommunicationsGateway#on_prem_mcp_enabled}.
	OnPremMcpEnabled interface{} `field:"optional" json:"onPremMcpEnabled" yaml:"onPremMcpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#tags VoiceServicesCommunicationsGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/voice_services_communications_gateway#timeouts VoiceServicesCommunicationsGateway#timeouts}
	Timeouts *VoiceServicesCommunicationsGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

