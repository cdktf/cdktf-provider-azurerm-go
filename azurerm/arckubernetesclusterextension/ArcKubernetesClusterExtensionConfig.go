// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesclusterextension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcKubernetesClusterExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#cluster_id ArcKubernetesClusterExtension#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#extension_type ArcKubernetesClusterExtension#extension_type}.
	ExtensionType *string `field:"required" json:"extensionType" yaml:"extensionType"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#identity ArcKubernetesClusterExtension#identity}
	Identity *ArcKubernetesClusterExtensionIdentity `field:"required" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#name ArcKubernetesClusterExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#configuration_protected_settings ArcKubernetesClusterExtension#configuration_protected_settings}.
	ConfigurationProtectedSettings *map[string]*string `field:"optional" json:"configurationProtectedSettings" yaml:"configurationProtectedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#configuration_settings ArcKubernetesClusterExtension#configuration_settings}.
	ConfigurationSettings *map[string]*string `field:"optional" json:"configurationSettings" yaml:"configurationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#id ArcKubernetesClusterExtension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#release_namespace ArcKubernetesClusterExtension#release_namespace}.
	ReleaseNamespace *string `field:"optional" json:"releaseNamespace" yaml:"releaseNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#release_train ArcKubernetesClusterExtension#release_train}.
	ReleaseTrain *string `field:"optional" json:"releaseTrain" yaml:"releaseTrain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#target_namespace ArcKubernetesClusterExtension#target_namespace}.
	TargetNamespace *string `field:"optional" json:"targetNamespace" yaml:"targetNamespace"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#timeouts ArcKubernetesClusterExtension#timeouts}
	Timeouts *ArcKubernetesClusterExtensionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/arc_kubernetes_cluster_extension#version ArcKubernetesClusterExtension#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

