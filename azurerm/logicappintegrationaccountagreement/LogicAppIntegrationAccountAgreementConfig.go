// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountagreement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppIntegrationAccountAgreementConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#agreement_type LogicAppIntegrationAccountAgreement#agreement_type}.
	AgreementType *string `field:"required" json:"agreementType" yaml:"agreementType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#content LogicAppIntegrationAccountAgreement#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// guest_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#guest_identity LogicAppIntegrationAccountAgreement#guest_identity}
	GuestIdentity *LogicAppIntegrationAccountAgreementGuestIdentity `field:"required" json:"guestIdentity" yaml:"guestIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#guest_partner_name LogicAppIntegrationAccountAgreement#guest_partner_name}.
	GuestPartnerName *string `field:"required" json:"guestPartnerName" yaml:"guestPartnerName"`
	// host_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#host_identity LogicAppIntegrationAccountAgreement#host_identity}
	HostIdentity *LogicAppIntegrationAccountAgreementHostIdentity `field:"required" json:"hostIdentity" yaml:"hostIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#host_partner_name LogicAppIntegrationAccountAgreement#host_partner_name}.
	HostPartnerName *string `field:"required" json:"hostPartnerName" yaml:"hostPartnerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#integration_account_name LogicAppIntegrationAccountAgreement#integration_account_name}.
	IntegrationAccountName *string `field:"required" json:"integrationAccountName" yaml:"integrationAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#name LogicAppIntegrationAccountAgreement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#resource_group_name LogicAppIntegrationAccountAgreement#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#id LogicAppIntegrationAccountAgreement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#metadata LogicAppIntegrationAccountAgreement#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/logic_app_integration_account_agreement#timeouts LogicAppIntegrationAccountAgreement#timeouts}
	Timeouts *LogicAppIntegrationAccountAgreementTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

