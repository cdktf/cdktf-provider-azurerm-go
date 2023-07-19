package kubernetescluster


type KubernetesClusterConfidentialComputing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#sgx_quote_helper_enabled KubernetesCluster#sgx_quote_helper_enabled}.
	SgxQuoteHelperEnabled interface{} `field:"required" json:"sgxQuoteHelperEnabled" yaml:"sgxQuoteHelperEnabled"`
}

