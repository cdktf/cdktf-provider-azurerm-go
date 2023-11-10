// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LabServiceLabConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#location LabServiceLab#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#name LabServiceLab#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#resource_group_name LabServiceLab#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#security LabServiceLab#security}
	Security *LabServiceLabSecurity `field:"required" json:"security" yaml:"security"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#title LabServiceLab#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#virtual_machine LabServiceLab#virtual_machine}
	VirtualMachine *LabServiceLabVirtualMachine `field:"required" json:"virtualMachine" yaml:"virtualMachine"`
	// auto_shutdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#auto_shutdown LabServiceLab#auto_shutdown}
	AutoShutdown *LabServiceLabAutoShutdown `field:"optional" json:"autoShutdown" yaml:"autoShutdown"`
	// connection_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#connection_setting LabServiceLab#connection_setting}
	ConnectionSetting *LabServiceLabConnectionSetting `field:"optional" json:"connectionSetting" yaml:"connectionSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#description LabServiceLab#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#id LabServiceLab#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#lab_plan_id LabServiceLab#lab_plan_id}.
	LabPlanId *string `field:"optional" json:"labPlanId" yaml:"labPlanId"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#network LabServiceLab#network}
	Network *LabServiceLabNetwork `field:"optional" json:"network" yaml:"network"`
	// roster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#roster LabServiceLab#roster}
	Roster *LabServiceLabRoster `field:"optional" json:"roster" yaml:"roster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#tags LabServiceLab#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/lab_service_lab#timeouts LabServiceLab#timeouts}
	Timeouts *LabServiceLabTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

