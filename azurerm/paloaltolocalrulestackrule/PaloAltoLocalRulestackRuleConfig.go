// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoLocalRulestackRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#action PaloAltoLocalRulestackRule#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#applications PaloAltoLocalRulestackRule#applications}.
	Applications *[]*string `field:"required" json:"applications" yaml:"applications"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#destination PaloAltoLocalRulestackRule#destination}
	Destination *PaloAltoLocalRulestackRuleDestination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#name PaloAltoLocalRulestackRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#priority PaloAltoLocalRulestackRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#rulestack_id PaloAltoLocalRulestackRule#rulestack_id}.
	RulestackId *string `field:"required" json:"rulestackId" yaml:"rulestackId"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#source PaloAltoLocalRulestackRule#source}
	Source *PaloAltoLocalRulestackRuleSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#audit_comment PaloAltoLocalRulestackRule#audit_comment}.
	AuditComment *string `field:"optional" json:"auditComment" yaml:"auditComment"`
	// category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#category PaloAltoLocalRulestackRule#category}
	Category *PaloAltoLocalRulestackRuleCategory `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#decryption_rule_type PaloAltoLocalRulestackRule#decryption_rule_type}.
	DecryptionRuleType *string `field:"optional" json:"decryptionRuleType" yaml:"decryptionRuleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#description PaloAltoLocalRulestackRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#enabled PaloAltoLocalRulestackRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#id PaloAltoLocalRulestackRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#inspection_certificate_id PaloAltoLocalRulestackRule#inspection_certificate_id}.
	InspectionCertificateId *string `field:"optional" json:"inspectionCertificateId" yaml:"inspectionCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#logging_enabled PaloAltoLocalRulestackRule#logging_enabled}.
	LoggingEnabled interface{} `field:"optional" json:"loggingEnabled" yaml:"loggingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#negate_destination PaloAltoLocalRulestackRule#negate_destination}.
	NegateDestination interface{} `field:"optional" json:"negateDestination" yaml:"negateDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#negate_source PaloAltoLocalRulestackRule#negate_source}.
	NegateSource interface{} `field:"optional" json:"negateSource" yaml:"negateSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#protocol PaloAltoLocalRulestackRule#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#protocol_ports PaloAltoLocalRulestackRule#protocol_ports}.
	ProtocolPorts *[]*string `field:"optional" json:"protocolPorts" yaml:"protocolPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#tags PaloAltoLocalRulestackRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/palo_alto_local_rulestack_rule#timeouts PaloAltoLocalRulestackRule#timeouts}
	Timeouts *PaloAltoLocalRulestackRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

