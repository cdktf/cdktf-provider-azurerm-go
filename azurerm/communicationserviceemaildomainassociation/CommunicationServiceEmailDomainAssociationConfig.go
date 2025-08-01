// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package communicationserviceemaildomainassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CommunicationServiceEmailDomainAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/communication_service_email_domain_association#communication_service_id CommunicationServiceEmailDomainAssociation#communication_service_id}.
	CommunicationServiceId *string `field:"required" json:"communicationServiceId" yaml:"communicationServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/communication_service_email_domain_association#email_service_domain_id CommunicationServiceEmailDomainAssociation#email_service_domain_id}.
	EmailServiceDomainId *string `field:"required" json:"emailServiceDomainId" yaml:"emailServiceDomainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/communication_service_email_domain_association#id CommunicationServiceEmailDomainAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/communication_service_email_domain_association#timeouts CommunicationServiceEmailDomainAssociation#timeouts}
	Timeouts *CommunicationServiceEmailDomainAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

