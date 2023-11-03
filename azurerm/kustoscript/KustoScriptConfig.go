// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoscript

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KustoScriptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#database_id KustoScript#database_id}.
	DatabaseId *string `field:"required" json:"databaseId" yaml:"databaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#name KustoScript#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#continue_on_errors_enabled KustoScript#continue_on_errors_enabled}.
	ContinueOnErrorsEnabled interface{} `field:"optional" json:"continueOnErrorsEnabled" yaml:"continueOnErrorsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#force_an_update_when_value_changed KustoScript#force_an_update_when_value_changed}.
	ForceAnUpdateWhenValueChanged *string `field:"optional" json:"forceAnUpdateWhenValueChanged" yaml:"forceAnUpdateWhenValueChanged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#id KustoScript#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#sas_token KustoScript#sas_token}.
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#script_content KustoScript#script_content}.
	ScriptContent *string `field:"optional" json:"scriptContent" yaml:"scriptContent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#timeouts KustoScript#timeouts}
	Timeouts *KustoScriptTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kusto_script#url KustoScript#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

