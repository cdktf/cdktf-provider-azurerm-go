package sqlfailovergroup


type SqlFailoverGroupReadonlyEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_failover_group#mode SqlFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

