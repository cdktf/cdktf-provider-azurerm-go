// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridpartnerconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridPartnerConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#resource_group_name EventgridPartnerConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#default_maximum_expiration_time_in_days EventgridPartnerConfiguration#default_maximum_expiration_time_in_days}.
	DefaultMaximumExpirationTimeInDays *float64 `field:"optional" json:"defaultMaximumExpirationTimeInDays" yaml:"defaultMaximumExpirationTimeInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#id EventgridPartnerConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// partner_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#partner_authorization EventgridPartnerConfiguration#partner_authorization}
	PartnerAuthorization interface{} `field:"optional" json:"partnerAuthorization" yaml:"partnerAuthorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#tags EventgridPartnerConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/eventgrid_partner_configuration#timeouts EventgridPartnerConfiguration#timeouts}
	Timeouts *EventgridPartnerConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

