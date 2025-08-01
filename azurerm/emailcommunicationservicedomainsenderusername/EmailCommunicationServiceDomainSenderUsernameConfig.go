// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailcommunicationservicedomainsenderusername

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailCommunicationServiceDomainSenderUsernameConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/email_communication_service_domain_sender_username#email_service_domain_id EmailCommunicationServiceDomainSenderUsername#email_service_domain_id}.
	EmailServiceDomainId *string `field:"required" json:"emailServiceDomainId" yaml:"emailServiceDomainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/email_communication_service_domain_sender_username#name EmailCommunicationServiceDomainSenderUsername#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/email_communication_service_domain_sender_username#display_name EmailCommunicationServiceDomainSenderUsername#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/email_communication_service_domain_sender_username#id EmailCommunicationServiceDomainSenderUsername#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/email_communication_service_domain_sender_username#timeouts EmailCommunicationServiceDomainSenderUsername#timeouts}
	Timeouts *EmailCommunicationServiceDomainSenderUsernameTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

