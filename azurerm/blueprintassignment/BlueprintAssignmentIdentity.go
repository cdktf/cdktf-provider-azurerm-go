// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blueprintassignment


type BlueprintAssignmentIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/blueprint_assignment#type BlueprintAssignment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/blueprint_assignment#identity_ids BlueprintAssignment#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

