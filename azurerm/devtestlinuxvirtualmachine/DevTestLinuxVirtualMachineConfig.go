// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestlinuxvirtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevTestLinuxVirtualMachineConfig struct {
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
	// gallery_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#gallery_image_reference DevTestLinuxVirtualMachine#gallery_image_reference}
	GalleryImageReference *DevTestLinuxVirtualMachineGalleryImageReference `field:"required" json:"galleryImageReference" yaml:"galleryImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#lab_name DevTestLinuxVirtualMachine#lab_name}.
	LabName *string `field:"required" json:"labName" yaml:"labName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#lab_subnet_name DevTestLinuxVirtualMachine#lab_subnet_name}.
	LabSubnetName *string `field:"required" json:"labSubnetName" yaml:"labSubnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#lab_virtual_network_id DevTestLinuxVirtualMachine#lab_virtual_network_id}.
	LabVirtualNetworkId *string `field:"required" json:"labVirtualNetworkId" yaml:"labVirtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#location DevTestLinuxVirtualMachine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#name DevTestLinuxVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#resource_group_name DevTestLinuxVirtualMachine#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#size DevTestLinuxVirtualMachine#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#storage_type DevTestLinuxVirtualMachine#storage_type}.
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#username DevTestLinuxVirtualMachine#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#allow_claim DevTestLinuxVirtualMachine#allow_claim}.
	AllowClaim interface{} `field:"optional" json:"allowClaim" yaml:"allowClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#disallow_public_ip_address DevTestLinuxVirtualMachine#disallow_public_ip_address}.
	DisallowPublicIpAddress interface{} `field:"optional" json:"disallowPublicIpAddress" yaml:"disallowPublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#id DevTestLinuxVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inbound_nat_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#inbound_nat_rule DevTestLinuxVirtualMachine#inbound_nat_rule}
	InboundNatRule interface{} `field:"optional" json:"inboundNatRule" yaml:"inboundNatRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#notes DevTestLinuxVirtualMachine#notes}.
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#password DevTestLinuxVirtualMachine#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#ssh_key DevTestLinuxVirtualMachine#ssh_key}.
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#tags DevTestLinuxVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/dev_test_linux_virtual_machine#timeouts DevTestLinuxVirtualMachine#timeouts}
	Timeouts *DevTestLinuxVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

