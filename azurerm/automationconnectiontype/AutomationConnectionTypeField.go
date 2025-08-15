// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationconnectiontype


type AutomationConnectionTypeField struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/automation_connection_type#name AutomationConnectionType#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/automation_connection_type#type AutomationConnectionType#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/automation_connection_type#is_encrypted AutomationConnectionType#is_encrypted}.
	IsEncrypted interface{} `field:"optional" json:"isEncrypted" yaml:"isEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/automation_connection_type#is_optional AutomationConnectionType#is_optional}.
	IsOptional interface{} `field:"optional" json:"isOptional" yaml:"isOptional"`
}

