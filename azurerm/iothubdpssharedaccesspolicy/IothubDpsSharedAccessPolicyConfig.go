// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubdpssharedaccesspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubDpsSharedAccessPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#iothub_dps_name IothubDpsSharedAccessPolicy#iothub_dps_name}.
	IothubDpsName *string `field:"required" json:"iothubDpsName" yaml:"iothubDpsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#name IothubDpsSharedAccessPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#resource_group_name IothubDpsSharedAccessPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#enrollment_read IothubDpsSharedAccessPolicy#enrollment_read}.
	EnrollmentRead interface{} `field:"optional" json:"enrollmentRead" yaml:"enrollmentRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#enrollment_write IothubDpsSharedAccessPolicy#enrollment_write}.
	EnrollmentWrite interface{} `field:"optional" json:"enrollmentWrite" yaml:"enrollmentWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#id IothubDpsSharedAccessPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#registration_read IothubDpsSharedAccessPolicy#registration_read}.
	RegistrationRead interface{} `field:"optional" json:"registrationRead" yaml:"registrationRead"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#registration_write IothubDpsSharedAccessPolicy#registration_write}.
	RegistrationWrite interface{} `field:"optional" json:"registrationWrite" yaml:"registrationWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#service_config IothubDpsSharedAccessPolicy#service_config}.
	ServiceConfig interface{} `field:"optional" json:"serviceConfig" yaml:"serviceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/iothub_dps_shared_access_policy#timeouts IothubDpsSharedAccessPolicy#timeouts}
	Timeouts *IothubDpsSharedAccessPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

