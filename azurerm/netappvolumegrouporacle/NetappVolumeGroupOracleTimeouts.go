// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegrouporacle


type NetappVolumeGroupOracleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/netapp_volume_group_oracle#create NetappVolumeGroupOracle#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/netapp_volume_group_oracle#delete NetappVolumeGroupOracle#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/netapp_volume_group_oracle#read NetappVolumeGroupOracle#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/netapp_volume_group_oracle#update NetappVolumeGroupOracle#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

