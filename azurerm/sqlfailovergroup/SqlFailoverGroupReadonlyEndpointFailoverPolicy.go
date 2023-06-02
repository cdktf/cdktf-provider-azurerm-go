package sqlfailovergroup


type SqlFailoverGroupReadonlyEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/sql_failover_group#mode SqlFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

