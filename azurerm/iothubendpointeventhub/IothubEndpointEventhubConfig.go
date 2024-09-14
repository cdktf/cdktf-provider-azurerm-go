// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubendpointeventhub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubEndpointEventhubConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#iothub_id IothubEndpointEventhub#iothub_id}.
	IothubId *string `field:"required" json:"iothubId" yaml:"iothubId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#name IothubEndpointEventhub#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#resource_group_name IothubEndpointEventhub#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#authentication_type IothubEndpointEventhub#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#connection_string IothubEndpointEventhub#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#endpoint_uri IothubEndpointEventhub#endpoint_uri}.
	EndpointUri *string `field:"optional" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#entity_path IothubEndpointEventhub#entity_path}.
	EntityPath *string `field:"optional" json:"entityPath" yaml:"entityPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#id IothubEndpointEventhub#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#identity_id IothubEndpointEventhub#identity_id}.
	IdentityId *string `field:"optional" json:"identityId" yaml:"identityId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/iothub_endpoint_eventhub#timeouts IothubEndpointEventhub#timeouts}
	Timeouts *IothubEndpointEventhubTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

