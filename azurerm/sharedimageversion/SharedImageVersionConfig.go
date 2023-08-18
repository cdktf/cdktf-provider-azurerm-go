package sharedimageversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SharedImageVersionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#gallery_name SharedImageVersion#gallery_name}.
	GalleryName *string `field:"required" json:"galleryName" yaml:"galleryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#image_name SharedImageVersion#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#location SharedImageVersion#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#name SharedImageVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#resource_group_name SharedImageVersion#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// target_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#target_region SharedImageVersion#target_region}
	TargetRegion interface{} `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#blob_uri SharedImageVersion#blob_uri}.
	BlobUri *string `field:"optional" json:"blobUri" yaml:"blobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#end_of_life_date SharedImageVersion#end_of_life_date}.
	EndOfLifeDate *string `field:"optional" json:"endOfLifeDate" yaml:"endOfLifeDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#exclude_from_latest SharedImageVersion#exclude_from_latest}.
	ExcludeFromLatest interface{} `field:"optional" json:"excludeFromLatest" yaml:"excludeFromLatest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#id SharedImageVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#managed_image_id SharedImageVersion#managed_image_id}.
	ManagedImageId *string `field:"optional" json:"managedImageId" yaml:"managedImageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#os_disk_snapshot_id SharedImageVersion#os_disk_snapshot_id}.
	OsDiskSnapshotId *string `field:"optional" json:"osDiskSnapshotId" yaml:"osDiskSnapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#replication_mode SharedImageVersion#replication_mode}.
	ReplicationMode *string `field:"optional" json:"replicationMode" yaml:"replicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#storage_account_id SharedImageVersion#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#tags SharedImageVersion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/shared_image_version#timeouts SharedImageVersion#timeouts}
	Timeouts *SharedImageVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

