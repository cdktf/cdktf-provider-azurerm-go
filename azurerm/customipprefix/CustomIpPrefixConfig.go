// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customipprefix

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomIpPrefixConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#cidr CustomIpPrefix#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#location CustomIpPrefix#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#name CustomIpPrefix#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#resource_group_name CustomIpPrefix#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#commissioning_enabled CustomIpPrefix#commissioning_enabled}.
	CommissioningEnabled interface{} `field:"optional" json:"commissioningEnabled" yaml:"commissioningEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#id CustomIpPrefix#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#internet_advertising_disabled CustomIpPrefix#internet_advertising_disabled}.
	InternetAdvertisingDisabled interface{} `field:"optional" json:"internetAdvertisingDisabled" yaml:"internetAdvertisingDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#parent_custom_ip_prefix_id CustomIpPrefix#parent_custom_ip_prefix_id}.
	ParentCustomIpPrefixId *string `field:"optional" json:"parentCustomIpPrefixId" yaml:"parentCustomIpPrefixId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#roa_validity_end_date CustomIpPrefix#roa_validity_end_date}.
	RoaValidityEndDate *string `field:"optional" json:"roaValidityEndDate" yaml:"roaValidityEndDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#tags CustomIpPrefix#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#timeouts CustomIpPrefix#timeouts}
	Timeouts *CustomIpPrefixTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#wan_validation_signed_message CustomIpPrefix#wan_validation_signed_message}.
	WanValidationSignedMessage *string `field:"optional" json:"wanValidationSignedMessage" yaml:"wanValidationSignedMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/custom_ip_prefix#zones CustomIpPrefix#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

