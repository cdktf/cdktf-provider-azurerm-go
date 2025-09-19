// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package communicationserviceemaildomainassociation


type CommunicationServiceEmailDomainAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/communication_service_email_domain_association#create CommunicationServiceEmailDomainAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/communication_service_email_domain_association#delete CommunicationServiceEmailDomainAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/communication_service_email_domain_association#read CommunicationServiceEmailDomainAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

