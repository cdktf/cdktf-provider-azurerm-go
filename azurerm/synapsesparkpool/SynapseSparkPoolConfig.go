// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesparkpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseSparkPoolConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#name SynapseSparkPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#node_size SynapseSparkPool#node_size}.
	NodeSize *string `field:"required" json:"nodeSize" yaml:"nodeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#node_size_family SynapseSparkPool#node_size_family}.
	NodeSizeFamily *string `field:"required" json:"nodeSizeFamily" yaml:"nodeSizeFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#synapse_workspace_id SynapseSparkPool#synapse_workspace_id}.
	SynapseWorkspaceId *string `field:"required" json:"synapseWorkspaceId" yaml:"synapseWorkspaceId"`
	// auto_pause block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#auto_pause SynapseSparkPool#auto_pause}
	AutoPause *SynapseSparkPoolAutoPause `field:"optional" json:"autoPause" yaml:"autoPause"`
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#auto_scale SynapseSparkPool#auto_scale}
	AutoScale *SynapseSparkPoolAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#cache_size SynapseSparkPool#cache_size}.
	CacheSize *float64 `field:"optional" json:"cacheSize" yaml:"cacheSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#compute_isolation_enabled SynapseSparkPool#compute_isolation_enabled}.
	ComputeIsolationEnabled interface{} `field:"optional" json:"computeIsolationEnabled" yaml:"computeIsolationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#dynamic_executor_allocation_enabled SynapseSparkPool#dynamic_executor_allocation_enabled}.
	DynamicExecutorAllocationEnabled interface{} `field:"optional" json:"dynamicExecutorAllocationEnabled" yaml:"dynamicExecutorAllocationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#id SynapseSparkPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// library_requirement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#library_requirement SynapseSparkPool#library_requirement}
	LibraryRequirement *SynapseSparkPoolLibraryRequirement `field:"optional" json:"libraryRequirement" yaml:"libraryRequirement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#max_executors SynapseSparkPool#max_executors}.
	MaxExecutors *float64 `field:"optional" json:"maxExecutors" yaml:"maxExecutors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#min_executors SynapseSparkPool#min_executors}.
	MinExecutors *float64 `field:"optional" json:"minExecutors" yaml:"minExecutors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#node_count SynapseSparkPool#node_count}.
	NodeCount *float64 `field:"optional" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#session_level_packages_enabled SynapseSparkPool#session_level_packages_enabled}.
	SessionLevelPackagesEnabled interface{} `field:"optional" json:"sessionLevelPackagesEnabled" yaml:"sessionLevelPackagesEnabled"`
	// spark_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#spark_config SynapseSparkPool#spark_config}
	SparkConfig *SynapseSparkPoolSparkConfig `field:"optional" json:"sparkConfig" yaml:"sparkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#spark_events_folder SynapseSparkPool#spark_events_folder}.
	SparkEventsFolder *string `field:"optional" json:"sparkEventsFolder" yaml:"sparkEventsFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#spark_log_folder SynapseSparkPool#spark_log_folder}.
	SparkLogFolder *string `field:"optional" json:"sparkLogFolder" yaml:"sparkLogFolder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#spark_version SynapseSparkPool#spark_version}.
	SparkVersion *string `field:"optional" json:"sparkVersion" yaml:"sparkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#tags SynapseSparkPool#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/synapse_spark_pool#timeouts SynapseSparkPool#timeouts}
	Timeouts *SynapseSparkPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

