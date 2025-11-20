// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinegalleryapplicationassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineGalleryApplicationAssignmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#gallery_application_version_id VirtualMachineGalleryApplicationAssignment#gallery_application_version_id}.
	GalleryApplicationVersionId *string `field:"required" json:"galleryApplicationVersionId" yaml:"galleryApplicationVersionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#virtual_machine_id VirtualMachineGalleryApplicationAssignment#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#configuration_blob_uri VirtualMachineGalleryApplicationAssignment#configuration_blob_uri}.
	ConfigurationBlobUri *string `field:"optional" json:"configurationBlobUri" yaml:"configurationBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#id VirtualMachineGalleryApplicationAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#order VirtualMachineGalleryApplicationAssignment#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#tag VirtualMachineGalleryApplicationAssignment#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/virtual_machine_gallery_application_assignment#timeouts VirtualMachineGalleryApplicationAssignment#timeouts}
	Timeouts *VirtualMachineGalleryApplicationAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

