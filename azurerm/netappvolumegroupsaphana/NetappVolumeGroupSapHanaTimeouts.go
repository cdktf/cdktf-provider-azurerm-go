// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumegroupsaphana


type NetappVolumeGroupSapHanaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/netapp_volume_group_sap_hana#create NetappVolumeGroupSapHana#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/netapp_volume_group_sap_hana#delete NetappVolumeGroupSapHana#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/netapp_volume_group_sap_hana#read NetappVolumeGroupSapHana#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/netapp_volume_group_sap_hana#update NetappVolumeGroupSapHana#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

