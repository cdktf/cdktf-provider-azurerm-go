// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SharedImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#gallery_name SharedImage#gallery_name}.
	GalleryName *string `field:"required" json:"galleryName" yaml:"galleryName"`
	// identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#identifier SharedImage#identifier}
	Identifier *SharedImageIdentifier `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#location SharedImage#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#name SharedImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#os_type SharedImage#os_type}.
	OsType *string `field:"required" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#resource_group_name SharedImage#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#accelerated_network_support_enabled SharedImage#accelerated_network_support_enabled}.
	AcceleratedNetworkSupportEnabled interface{} `field:"optional" json:"acceleratedNetworkSupportEnabled" yaml:"acceleratedNetworkSupportEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#architecture SharedImage#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#confidential_vm_enabled SharedImage#confidential_vm_enabled}.
	ConfidentialVmEnabled interface{} `field:"optional" json:"confidentialVmEnabled" yaml:"confidentialVmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#confidential_vm_supported SharedImage#confidential_vm_supported}.
	ConfidentialVmSupported interface{} `field:"optional" json:"confidentialVmSupported" yaml:"confidentialVmSupported"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#description SharedImage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#disk_controller_type_nvme_enabled SharedImage#disk_controller_type_nvme_enabled}.
	DiskControllerTypeNvmeEnabled interface{} `field:"optional" json:"diskControllerTypeNvmeEnabled" yaml:"diskControllerTypeNvmeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#disk_types_not_allowed SharedImage#disk_types_not_allowed}.
	DiskTypesNotAllowed *[]*string `field:"optional" json:"diskTypesNotAllowed" yaml:"diskTypesNotAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#end_of_life_date SharedImage#end_of_life_date}.
	EndOfLifeDate *string `field:"optional" json:"endOfLifeDate" yaml:"endOfLifeDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#eula SharedImage#eula}.
	Eula *string `field:"optional" json:"eula" yaml:"eula"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#hibernation_enabled SharedImage#hibernation_enabled}.
	HibernationEnabled interface{} `field:"optional" json:"hibernationEnabled" yaml:"hibernationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#hyper_v_generation SharedImage#hyper_v_generation}.
	HyperVGeneration *string `field:"optional" json:"hyperVGeneration" yaml:"hyperVGeneration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#id SharedImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#max_recommended_memory_in_gb SharedImage#max_recommended_memory_in_gb}.
	MaxRecommendedMemoryInGb *float64 `field:"optional" json:"maxRecommendedMemoryInGb" yaml:"maxRecommendedMemoryInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#max_recommended_vcpu_count SharedImage#max_recommended_vcpu_count}.
	MaxRecommendedVcpuCount *float64 `field:"optional" json:"maxRecommendedVcpuCount" yaml:"maxRecommendedVcpuCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#min_recommended_memory_in_gb SharedImage#min_recommended_memory_in_gb}.
	MinRecommendedMemoryInGb *float64 `field:"optional" json:"minRecommendedMemoryInGb" yaml:"minRecommendedMemoryInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#min_recommended_vcpu_count SharedImage#min_recommended_vcpu_count}.
	MinRecommendedVcpuCount *float64 `field:"optional" json:"minRecommendedVcpuCount" yaml:"minRecommendedVcpuCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#privacy_statement_uri SharedImage#privacy_statement_uri}.
	PrivacyStatementUri *string `field:"optional" json:"privacyStatementUri" yaml:"privacyStatementUri"`
	// purchase_plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#purchase_plan SharedImage#purchase_plan}
	PurchasePlan *SharedImagePurchasePlan `field:"optional" json:"purchasePlan" yaml:"purchasePlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#release_note_uri SharedImage#release_note_uri}.
	ReleaseNoteUri *string `field:"optional" json:"releaseNoteUri" yaml:"releaseNoteUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#specialized SharedImage#specialized}.
	Specialized interface{} `field:"optional" json:"specialized" yaml:"specialized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#tags SharedImage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#timeouts SharedImage#timeouts}
	Timeouts *SharedImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#trusted_launch_enabled SharedImage#trusted_launch_enabled}.
	TrustedLaunchEnabled interface{} `field:"optional" json:"trustedLaunchEnabled" yaml:"trustedLaunchEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/shared_image#trusted_launch_supported SharedImage#trusted_launch_supported}.
	TrustedLaunchSupported interface{} `field:"optional" json:"trustedLaunchSupported" yaml:"trustedLaunchSupported"`
}

