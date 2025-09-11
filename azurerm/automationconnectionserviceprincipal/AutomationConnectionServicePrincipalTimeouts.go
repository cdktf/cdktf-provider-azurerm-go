// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationconnectionserviceprincipal


type AutomationConnectionServicePrincipalTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_connection_service_principal#create AutomationConnectionServicePrincipal#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_connection_service_principal#delete AutomationConnectionServicePrincipal#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_connection_service_principal#read AutomationConnectionServicePrincipal#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/automation_connection_service_principal#update AutomationConnectionServicePrincipal#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

