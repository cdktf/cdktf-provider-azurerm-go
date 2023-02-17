package kubernetescluster


type KubernetesClusterKubeletIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#client_id KubernetesCluster#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#object_id KubernetesCluster#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#user_assigned_identity_id KubernetesCluster#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

