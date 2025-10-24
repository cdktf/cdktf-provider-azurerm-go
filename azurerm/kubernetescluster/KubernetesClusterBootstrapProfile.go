// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterBootstrapProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/kubernetes_cluster#artifact_source KubernetesCluster#artifact_source}.
	ArtifactSource *string `field:"optional" json:"artifactSource" yaml:"artifactSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/kubernetes_cluster#container_registry_id KubernetesCluster#container_registry_id}.
	ContainerRegistryId *string `field:"optional" json:"containerRegistryId" yaml:"containerRegistryId"`
}

