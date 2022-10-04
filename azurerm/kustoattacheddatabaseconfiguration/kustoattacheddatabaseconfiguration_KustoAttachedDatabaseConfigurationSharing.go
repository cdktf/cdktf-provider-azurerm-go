package kustoattacheddatabaseconfiguration


type KustoAttachedDatabaseConfigurationSharing struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#external_tables_to_exclude KustoAttachedDatabaseConfiguration#external_tables_to_exclude}.
	ExternalTablesToExclude *[]*string `field:"optional" json:"externalTablesToExclude" yaml:"externalTablesToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#external_tables_to_include KustoAttachedDatabaseConfiguration#external_tables_to_include}.
	ExternalTablesToInclude *[]*string `field:"optional" json:"externalTablesToInclude" yaml:"externalTablesToInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#materialized_views_to_exclude KustoAttachedDatabaseConfiguration#materialized_views_to_exclude}.
	MaterializedViewsToExclude *[]*string `field:"optional" json:"materializedViewsToExclude" yaml:"materializedViewsToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#materialized_views_to_include KustoAttachedDatabaseConfiguration#materialized_views_to_include}.
	MaterializedViewsToInclude *[]*string `field:"optional" json:"materializedViewsToInclude" yaml:"materializedViewsToInclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#tables_to_exclude KustoAttachedDatabaseConfiguration#tables_to_exclude}.
	TablesToExclude *[]*string `field:"optional" json:"tablesToExclude" yaml:"tablesToExclude"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_attached_database_configuration#tables_to_include KustoAttachedDatabaseConfiguration#tables_to_include}.
	TablesToInclude *[]*string `field:"optional" json:"tablesToInclude" yaml:"tablesToInclude"`
}

