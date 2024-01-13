// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningworkspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningWorkspaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#application_insights_id MachineLearningWorkspace#application_insights_id}.
	ApplicationInsightsId *string `field:"required" json:"applicationInsightsId" yaml:"applicationInsightsId"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#identity MachineLearningWorkspace#identity}
	Identity *MachineLearningWorkspaceIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#key_vault_id MachineLearningWorkspace#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#location MachineLearningWorkspace#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#name MachineLearningWorkspace#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#resource_group_name MachineLearningWorkspace#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#storage_account_id MachineLearningWorkspace#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#container_registry_id MachineLearningWorkspace#container_registry_id}.
	ContainerRegistryId *string `field:"optional" json:"containerRegistryId" yaml:"containerRegistryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#description MachineLearningWorkspace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#encryption MachineLearningWorkspace#encryption}
	Encryption *MachineLearningWorkspaceEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#friendly_name MachineLearningWorkspace#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#high_business_impact MachineLearningWorkspace#high_business_impact}.
	HighBusinessImpact interface{} `field:"optional" json:"highBusinessImpact" yaml:"highBusinessImpact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#id MachineLearningWorkspace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#image_build_compute_name MachineLearningWorkspace#image_build_compute_name}.
	ImageBuildComputeName *string `field:"optional" json:"imageBuildComputeName" yaml:"imageBuildComputeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#primary_user_assigned_identity MachineLearningWorkspace#primary_user_assigned_identity}.
	PrimaryUserAssignedIdentity *string `field:"optional" json:"primaryUserAssignedIdentity" yaml:"primaryUserAssignedIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#public_access_behind_virtual_network_enabled MachineLearningWorkspace#public_access_behind_virtual_network_enabled}.
	PublicAccessBehindVirtualNetworkEnabled interface{} `field:"optional" json:"publicAccessBehindVirtualNetworkEnabled" yaml:"publicAccessBehindVirtualNetworkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#public_network_access_enabled MachineLearningWorkspace#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#sku_name MachineLearningWorkspace#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#tags MachineLearningWorkspace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#timeouts MachineLearningWorkspace#timeouts}
	Timeouts *MachineLearningWorkspaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/machine_learning_workspace#v1_legacy_mode_enabled MachineLearningWorkspace#v1_legacy_mode_enabled}.
	V1LegacyModeEnabled interface{} `field:"optional" json:"v1LegacyModeEnabled" yaml:"v1LegacyModeEnabled"`
}

