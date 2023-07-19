package mssqlmanagedinstancefailovergroup


type MssqlManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/mssql_managed_instance_failover_group#mode MssqlManagedInstanceFailoverGroup#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/mssql_managed_instance_failover_group#grace_minutes MssqlManagedInstanceFailoverGroup#grace_minutes}.
	GraceMinutes *float64 `field:"optional" json:"graceMinutes" yaml:"graceMinutes"`
}

