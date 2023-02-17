package mssqlmanagedinstancefailovergroup


type MssqlManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_failover_group#mode MssqlManagedInstanceFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_managed_instance_failover_group#grace_minutes MssqlManagedInstanceFailoverGroup#grace_minutes}.
	GraceMinutes *float64 `field:"optional" json:"graceMinutes" yaml:"graceMinutes"`
}

