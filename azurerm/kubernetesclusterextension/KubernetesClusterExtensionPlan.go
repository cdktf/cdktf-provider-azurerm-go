package kubernetesclusterextension


type KubernetesClusterExtensionPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_extension#name KubernetesClusterExtension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_extension#product KubernetesClusterExtension#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_extension#publisher KubernetesClusterExtension#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_extension#promotion_code KubernetesClusterExtension#promotion_code}.
	PromotionCode *string `field:"optional" json:"promotionCode" yaml:"promotionCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_extension#version KubernetesClusterExtension#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

