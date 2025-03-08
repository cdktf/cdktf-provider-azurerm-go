// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaredicomservice


type HealthcareDicomServiceStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/healthcare_dicom_service#file_system_name HealthcareDicomService#file_system_name}.
	FileSystemName *string `field:"required" json:"fileSystemName" yaml:"fileSystemName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/healthcare_dicom_service#storage_account_id HealthcareDicomService#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
}

