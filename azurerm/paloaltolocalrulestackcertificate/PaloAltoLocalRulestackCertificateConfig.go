// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoLocalRulestackCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#name PaloAltoLocalRulestackCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#rulestack_id PaloAltoLocalRulestackCertificate#rulestack_id}.
	RulestackId *string `field:"required" json:"rulestackId" yaml:"rulestackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#audit_comment PaloAltoLocalRulestackCertificate#audit_comment}.
	AuditComment *string `field:"optional" json:"auditComment" yaml:"auditComment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#description PaloAltoLocalRulestackCertificate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#id PaloAltoLocalRulestackCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#key_vault_certificate_id PaloAltoLocalRulestackCertificate#key_vault_certificate_id}.
	KeyVaultCertificateId *string `field:"optional" json:"keyVaultCertificateId" yaml:"keyVaultCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#self_signed PaloAltoLocalRulestackCertificate#self_signed}.
	SelfSigned interface{} `field:"optional" json:"selfSigned" yaml:"selfSigned"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/palo_alto_local_rulestack_certificate#timeouts PaloAltoLocalRulestackCertificate#timeouts}
	Timeouts *PaloAltoLocalRulestackCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

