// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapikey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationInsightsApiKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#application_insights_id ApplicationInsightsApiKey#application_insights_id}.
	ApplicationInsightsId *string `field:"required" json:"applicationInsightsId" yaml:"applicationInsightsId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#name ApplicationInsightsApiKey#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#id ApplicationInsightsApiKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#read_permissions ApplicationInsightsApiKey#read_permissions}.
	ReadPermissions *[]*string `field:"optional" json:"readPermissions" yaml:"readPermissions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#timeouts ApplicationInsightsApiKey#timeouts}
	Timeouts *ApplicationInsightsApiKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/application_insights_api_key#write_permissions ApplicationInsightsApiKey#write_permissions}.
	WritePermissions *[]*string `field:"optional" json:"writePermissions" yaml:"writePermissions"`
}

