// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclustertrustedaccessrolebinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterTrustedAccessRoleBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#kubernetes_cluster_id KubernetesClusterTrustedAccessRoleBinding#kubernetes_cluster_id}.
	KubernetesClusterId *string `field:"required" json:"kubernetesClusterId" yaml:"kubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#name KubernetesClusterTrustedAccessRoleBinding#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#roles KubernetesClusterTrustedAccessRoleBinding#roles}.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#source_resource_id KubernetesClusterTrustedAccessRoleBinding#source_resource_id}.
	SourceResourceId *string `field:"required" json:"sourceResourceId" yaml:"sourceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#id KubernetesClusterTrustedAccessRoleBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kubernetes_cluster_trusted_access_role_binding#timeouts KubernetesClusterTrustedAccessRoleBinding#timeouts}
	Timeouts *KubernetesClusterTrustedAccessRoleBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

