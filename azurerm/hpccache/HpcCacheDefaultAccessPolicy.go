package hpccache


type HpcCacheDefaultAccessPolicy struct {
	// access_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#access_rule HpcCache#access_rule}
	AccessRule interface{} `field:"required" json:"accessRule" yaml:"accessRule"`
}

