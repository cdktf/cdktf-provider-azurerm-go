// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emailcommunicationservicedomain


type EmailCommunicationServiceDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/email_communication_service_domain#create EmailCommunicationServiceDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/email_communication_service_domain#delete EmailCommunicationServiceDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/email_communication_service_domain#read EmailCommunicationServiceDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/email_communication_service_domain#update EmailCommunicationServiceDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

