package hpccache


type HpcCacheDirectoryLdapBind struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#dn HpcCache#dn}.
	Dn *string `field:"required" json:"dn" yaml:"dn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#password HpcCache#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
}

