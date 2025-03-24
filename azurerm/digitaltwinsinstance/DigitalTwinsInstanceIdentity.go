// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package digitaltwinsinstance


type DigitalTwinsInstanceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_instance#type DigitalTwinsInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/digital_twins_instance#identity_ids DigitalTwinsInstance#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

