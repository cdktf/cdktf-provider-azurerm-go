// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorylinkedserviceazuredatabricks


type DataFactoryLinkedServiceAzureDatabricksInstancePool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/data_factory_linked_service_azure_databricks#cluster_version DataFactoryLinkedServiceAzureDatabricks#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/data_factory_linked_service_azure_databricks#instance_pool_id DataFactoryLinkedServiceAzureDatabricks#instance_pool_id}.
	InstancePoolId *string `field:"required" json:"instancePoolId" yaml:"instancePoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/data_factory_linked_service_azure_databricks#max_number_of_workers DataFactoryLinkedServiceAzureDatabricks#max_number_of_workers}.
	MaxNumberOfWorkers *float64 `field:"optional" json:"maxNumberOfWorkers" yaml:"maxNumberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/data_factory_linked_service_azure_databricks#min_number_of_workers DataFactoryLinkedServiceAzureDatabricks#min_number_of_workers}.
	MinNumberOfWorkers *float64 `field:"optional" json:"minNumberOfWorkers" yaml:"minNumberOfWorkers"`
}

