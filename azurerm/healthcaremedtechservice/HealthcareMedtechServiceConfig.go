// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaremedtechservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareMedtechServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#device_mapping_json HealthcareMedtechService#device_mapping_json}.
	DeviceMappingJson *string `field:"required" json:"deviceMappingJson" yaml:"deviceMappingJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#eventhub_consumer_group_name HealthcareMedtechService#eventhub_consumer_group_name}.
	EventhubConsumerGroupName *string `field:"required" json:"eventhubConsumerGroupName" yaml:"eventhubConsumerGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#eventhub_name HealthcareMedtechService#eventhub_name}.
	EventhubName *string `field:"required" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#eventhub_namespace_name HealthcareMedtechService#eventhub_namespace_name}.
	EventhubNamespaceName *string `field:"required" json:"eventhubNamespaceName" yaml:"eventhubNamespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#location HealthcareMedtechService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#name HealthcareMedtechService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#workspace_id HealthcareMedtechService#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#id HealthcareMedtechService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#identity HealthcareMedtechService#identity}
	Identity *HealthcareMedtechServiceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#tags HealthcareMedtechService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_medtech_service#timeouts HealthcareMedtechService#timeouts}
	Timeouts *HealthcareMedtechServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

