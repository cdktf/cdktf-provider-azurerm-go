// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinepacketcapture


type VirtualMachinePacketCaptureStorageLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/virtual_machine_packet_capture#file_path VirtualMachinePacketCapture#file_path}.
	FilePath *string `field:"optional" json:"filePath" yaml:"filePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/virtual_machine_packet_capture#storage_account_id VirtualMachinePacketCapture#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
}

