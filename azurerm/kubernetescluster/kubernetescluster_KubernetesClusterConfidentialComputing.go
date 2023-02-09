package kubernetescluster


type KubernetesClusterConfidentialComputing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#sgx_quote_helper_enabled KubernetesCluster#sgx_quote_helper_enabled}.
	SgxQuoteHelperEnabled interface{} `field:"required" json:"sgxQuoteHelperEnabled" yaml:"sgxQuoteHelperEnabled"`
}

