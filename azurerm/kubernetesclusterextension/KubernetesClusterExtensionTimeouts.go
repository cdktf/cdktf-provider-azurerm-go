package kubernetesclusterextension


type KubernetesClusterExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/kubernetes_cluster_extension#create KubernetesClusterExtension#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/kubernetes_cluster_extension#delete KubernetesClusterExtension#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/kubernetes_cluster_extension#read KubernetesClusterExtension#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/kubernetes_cluster_extension#update KubernetesClusterExtension#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

