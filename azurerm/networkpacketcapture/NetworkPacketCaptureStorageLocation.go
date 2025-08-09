// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpacketcapture


type NetworkPacketCaptureStorageLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/network_packet_capture#file_path NetworkPacketCapture#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/network_packet_capture#storage_account_id NetworkPacketCapture#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
}

