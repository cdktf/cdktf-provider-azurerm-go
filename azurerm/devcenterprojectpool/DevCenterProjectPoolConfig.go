// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcenterprojectpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevCenterProjectPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#dev_box_definition_name DevCenterProjectPool#dev_box_definition_name}.
	DevBoxDefinitionName *string `field:"required" json:"devBoxDefinitionName" yaml:"devBoxDefinitionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#dev_center_attached_network_name DevCenterProjectPool#dev_center_attached_network_name}.
	DevCenterAttachedNetworkName *string `field:"required" json:"devCenterAttachedNetworkName" yaml:"devCenterAttachedNetworkName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#dev_center_project_id DevCenterProjectPool#dev_center_project_id}.
	DevCenterProjectId *string `field:"required" json:"devCenterProjectId" yaml:"devCenterProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#local_administrator_enabled DevCenterProjectPool#local_administrator_enabled}.
	LocalAdministratorEnabled interface{} `field:"required" json:"localAdministratorEnabled" yaml:"localAdministratorEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#location DevCenterProjectPool#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#name DevCenterProjectPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#id DevCenterProjectPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#stop_on_disconnect_grace_period_minutes DevCenterProjectPool#stop_on_disconnect_grace_period_minutes}.
	StopOnDisconnectGracePeriodMinutes *float64 `field:"optional" json:"stopOnDisconnectGracePeriodMinutes" yaml:"stopOnDisconnectGracePeriodMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#tags DevCenterProjectPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/dev_center_project_pool#timeouts DevCenterProjectPool#timeouts}
	Timeouts *DevCenterProjectPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

