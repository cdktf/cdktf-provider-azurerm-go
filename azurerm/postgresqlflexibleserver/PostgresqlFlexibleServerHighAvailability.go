package postgresqlflexibleserver


type PostgresqlFlexibleServerHighAvailability struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#mode PostgresqlFlexibleServer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#standby_availability_zone PostgresqlFlexibleServer#standby_availability_zone}.
	StandbyAvailabilityZone *string `field:"optional" json:"standbyAvailabilityZone" yaml:"standbyAvailabilityZone"`
}

