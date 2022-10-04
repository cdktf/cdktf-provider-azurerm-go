package kubernetescluster


type KubernetesClusterLinuxProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#admin_username KubernetesCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// ssh_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#ssh_key KubernetesCluster#ssh_key}
	SshKey *KubernetesClusterLinuxProfileSshKey `field:"required" json:"sshKey" yaml:"sshKey"`
}

