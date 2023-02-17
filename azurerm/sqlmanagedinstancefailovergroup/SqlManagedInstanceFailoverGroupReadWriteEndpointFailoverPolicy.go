package sqlmanagedinstancefailovergroup


type SqlManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_managed_instance_failover_group#mode SqlManagedInstanceFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sql_managed_instance_failover_group#grace_minutes SqlManagedInstanceFailoverGroup#grace_minutes}.
	GraceMinutes *float64 `field:"optional" json:"graceMinutes" yaml:"graceMinutes"`
}

