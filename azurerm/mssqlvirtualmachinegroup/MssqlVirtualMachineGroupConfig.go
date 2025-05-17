// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachinegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#location MssqlVirtualMachineGroup#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#name MssqlVirtualMachineGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#resource_group_name MssqlVirtualMachineGroup#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#sql_image_offer MssqlVirtualMachineGroup#sql_image_offer}.
	SqlImageOffer *string `field:"required" json:"sqlImageOffer" yaml:"sqlImageOffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#sql_image_sku MssqlVirtualMachineGroup#sql_image_sku}.
	SqlImageSku *string `field:"required" json:"sqlImageSku" yaml:"sqlImageSku"`
	// wsfc_domain_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#wsfc_domain_profile MssqlVirtualMachineGroup#wsfc_domain_profile}
	WsfcDomainProfile *MssqlVirtualMachineGroupWsfcDomainProfile `field:"required" json:"wsfcDomainProfile" yaml:"wsfcDomainProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#id MssqlVirtualMachineGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#tags MssqlVirtualMachineGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/mssql_virtual_machine_group#timeouts MssqlVirtualMachineGroup#timeouts}
	Timeouts *MssqlVirtualMachineGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

