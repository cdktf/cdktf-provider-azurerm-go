package synapsesqlpool


type SynapseSqlPoolRestore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/synapse_sql_pool#point_in_time SynapseSqlPool#point_in_time}.
	PointInTime *string `field:"required" json:"pointInTime" yaml:"pointInTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/synapse_sql_pool#source_database_id SynapseSqlPool#source_database_id}.
	SourceDatabaseId *string `field:"required" json:"sourceDatabaseId" yaml:"sourceDatabaseId"`
}

