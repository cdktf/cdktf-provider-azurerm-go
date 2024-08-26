// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackoutbounduntrustcertificateassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoLocalRulestackOutboundUntrustCertificateAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/palo_alto_local_rulestack_outbound_untrust_certificate_association#certificate_id PaloAltoLocalRulestackOutboundUntrustCertificateAssociation#certificate_id}.
	CertificateId *string `field:"required" json:"certificateId" yaml:"certificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/palo_alto_local_rulestack_outbound_untrust_certificate_association#id PaloAltoLocalRulestackOutboundUntrustCertificateAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/palo_alto_local_rulestack_outbound_untrust_certificate_association#timeouts PaloAltoLocalRulestackOutboundUntrustCertificateAssociation#timeouts}
	Timeouts *PaloAltoLocalRulestackOutboundUntrustCertificateAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

