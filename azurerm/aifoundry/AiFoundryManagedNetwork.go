// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aifoundry


type AiFoundryManagedNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/ai_foundry#isolation_mode AiFoundry#isolation_mode}.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
}

