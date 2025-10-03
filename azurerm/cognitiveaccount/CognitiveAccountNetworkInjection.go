// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitiveaccount


type CognitiveAccountNetworkInjection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/cognitive_account#scenario CognitiveAccount#scenario}.
	Scenario *string `field:"required" json:"scenario" yaml:"scenario"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/cognitive_account#subnet_id CognitiveAccount#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

