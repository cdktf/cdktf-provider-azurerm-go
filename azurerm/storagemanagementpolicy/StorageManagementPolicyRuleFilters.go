package storagemanagementpolicy


type StorageManagementPolicyRuleFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#blob_types StorageManagementPolicy#blob_types}.
	BlobTypes *[]*string `field:"required" json:"blobTypes" yaml:"blobTypes"`
	// match_blob_index_tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#match_blob_index_tag StorageManagementPolicy#match_blob_index_tag}
	MatchBlobIndexTag interface{} `field:"optional" json:"matchBlobIndexTag" yaml:"matchBlobIndexTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#prefix_match StorageManagementPolicy#prefix_match}.
	PrefixMatch *[]*string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
}
