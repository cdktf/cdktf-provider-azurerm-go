// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubdpscertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubDpsCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#certificate_content IothubDpsCertificate#certificate_content}.
	CertificateContent *string `field:"required" json:"certificateContent" yaml:"certificateContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#iot_dps_name IothubDpsCertificate#iot_dps_name}.
	IotDpsName *string `field:"required" json:"iotDpsName" yaml:"iotDpsName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#name IothubDpsCertificate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#resource_group_name IothubDpsCertificate#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#id IothubDpsCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#is_verified IothubDpsCertificate#is_verified}.
	IsVerified interface{} `field:"optional" json:"isVerified" yaml:"isVerified"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/iothub_dps_certificate#timeouts IothubDpsCertificate#timeouts}
	Timeouts *IothubDpsCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

