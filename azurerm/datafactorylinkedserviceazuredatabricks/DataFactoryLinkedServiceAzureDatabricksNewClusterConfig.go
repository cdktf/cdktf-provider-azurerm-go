// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuredatabricks


type DataFactoryLinkedServiceAzureDatabricksNewClusterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#cluster_version DataFactoryLinkedServiceAzureDatabricks#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#node_type DataFactoryLinkedServiceAzureDatabricks#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#custom_tags DataFactoryLinkedServiceAzureDatabricks#custom_tags}.
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#driver_node_type DataFactoryLinkedServiceAzureDatabricks#driver_node_type}.
	DriverNodeType *string `field:"optional" json:"driverNodeType" yaml:"driverNodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#init_scripts DataFactoryLinkedServiceAzureDatabricks#init_scripts}.
	InitScripts *[]*string `field:"optional" json:"initScripts" yaml:"initScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#log_destination DataFactoryLinkedServiceAzureDatabricks#log_destination}.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#max_number_of_workers DataFactoryLinkedServiceAzureDatabricks#max_number_of_workers}.
	MaxNumberOfWorkers *float64 `field:"optional" json:"maxNumberOfWorkers" yaml:"maxNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#min_number_of_workers DataFactoryLinkedServiceAzureDatabricks#min_number_of_workers}.
	MinNumberOfWorkers *float64 `field:"optional" json:"minNumberOfWorkers" yaml:"minNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#spark_config DataFactoryLinkedServiceAzureDatabricks#spark_config}.
	SparkConfig *map[string]*string `field:"optional" json:"sparkConfig" yaml:"sparkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/data_factory_linked_service_azure_databricks#spark_environment_variables DataFactoryLinkedServiceAzureDatabricks#spark_environment_variables}.
	SparkEnvironmentVariables *map[string]*string `field:"optional" json:"sparkEnvironmentVariables" yaml:"sparkEnvironmentVariables"`
}

