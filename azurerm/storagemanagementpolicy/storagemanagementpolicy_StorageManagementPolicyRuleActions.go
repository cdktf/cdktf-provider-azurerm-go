package storagemanagementpolicy


type StorageManagementPolicyRuleActions struct {
	// base_blob block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#base_blob StorageManagementPolicy#base_blob}
	BaseBlob *StorageManagementPolicyRuleActionsBaseBlob `field:"optional" json:"baseBlob" yaml:"baseBlob"`
	// snapshot block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#snapshot StorageManagementPolicy#snapshot}
	Snapshot *StorageManagementPolicyRuleActionsSnapshot `field:"optional" json:"snapshot" yaml:"snapshot"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#version StorageManagementPolicy#version}
	Version *StorageManagementPolicyRuleActionsVersion `field:"optional" json:"version" yaml:"version"`
}
