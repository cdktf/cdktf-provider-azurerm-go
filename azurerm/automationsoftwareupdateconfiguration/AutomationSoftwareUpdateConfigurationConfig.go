// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationSoftwareUpdateConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#automation_account_id AutomationSoftwareUpdateConfiguration#automation_account_id}.
	AutomationAccountId *string `field:"required" json:"automationAccountId" yaml:"automationAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#name AutomationSoftwareUpdateConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#schedule AutomationSoftwareUpdateConfiguration#schedule}
	Schedule *AutomationSoftwareUpdateConfigurationSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#duration AutomationSoftwareUpdateConfiguration#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#id AutomationSoftwareUpdateConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// linux block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#linux AutomationSoftwareUpdateConfiguration#linux}
	Linux *AutomationSoftwareUpdateConfigurationLinux `field:"optional" json:"linux" yaml:"linux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#non_azure_computer_names AutomationSoftwareUpdateConfiguration#non_azure_computer_names}.
	NonAzureComputerNames *[]*string `field:"optional" json:"nonAzureComputerNames" yaml:"nonAzureComputerNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#operating_system AutomationSoftwareUpdateConfiguration#operating_system}.
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// post_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#post_task AutomationSoftwareUpdateConfiguration#post_task}
	PostTask *AutomationSoftwareUpdateConfigurationPostTask `field:"optional" json:"postTask" yaml:"postTask"`
	// pre_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#pre_task AutomationSoftwareUpdateConfiguration#pre_task}
	PreTask *AutomationSoftwareUpdateConfigurationPreTask `field:"optional" json:"preTask" yaml:"preTask"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#target AutomationSoftwareUpdateConfiguration#target}
	Target *AutomationSoftwareUpdateConfigurationTarget `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#timeouts AutomationSoftwareUpdateConfiguration#timeouts}
	Timeouts *AutomationSoftwareUpdateConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#virtual_machine_ids AutomationSoftwareUpdateConfiguration#virtual_machine_ids}.
	VirtualMachineIds *[]*string `field:"optional" json:"virtualMachineIds" yaml:"virtualMachineIds"`
	// windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/automation_software_update_configuration#windows AutomationSoftwareUpdateConfiguration#windows}
	Windows *AutomationSoftwareUpdateConfigurationWindows `field:"optional" json:"windows" yaml:"windows"`
}

