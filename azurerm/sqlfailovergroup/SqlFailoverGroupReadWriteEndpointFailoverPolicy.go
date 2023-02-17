package sqlfailovergroup


type SqlFailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_failover_group#mode SqlFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_failover_group#grace_minutes SqlFailoverGroup#grace_minutes}.
	GraceMinutes *float64 `field:"optional" json:"graceMinutes" yaml:"graceMinutes"`
}

