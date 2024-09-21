// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadtest


type LoadTestEncryption struct {
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/load_test#identity LoadTest#identity}
	Identity *LoadTestEncryptionIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/load_test#key_url LoadTest#key_url}.
	KeyUrl *string `field:"required" json:"keyUrl" yaml:"keyUrl"`
}

