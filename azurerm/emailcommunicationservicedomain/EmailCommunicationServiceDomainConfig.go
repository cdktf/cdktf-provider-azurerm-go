// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailcommunicationservicedomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailCommunicationServiceDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#domain_management EmailCommunicationServiceDomain#domain_management}.
	DomainManagement *string `field:"required" json:"domainManagement" yaml:"domainManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#email_service_id EmailCommunicationServiceDomain#email_service_id}.
	EmailServiceId *string `field:"required" json:"emailServiceId" yaml:"emailServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#name EmailCommunicationServiceDomain#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#id EmailCommunicationServiceDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#tags EmailCommunicationServiceDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#timeouts EmailCommunicationServiceDomain#timeouts}
	Timeouts *EmailCommunicationServiceDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/email_communication_service_domain#user_engagement_tracking_enabled EmailCommunicationServiceDomain#user_engagement_tracking_enabled}.
	UserEngagementTrackingEnabled interface{} `field:"optional" json:"userEngagementTrackingEnabled" yaml:"userEngagementTrackingEnabled"`
}

