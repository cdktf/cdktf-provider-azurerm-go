// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabricmanagedcluster


type ServiceFabricManagedClusterNodeType struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#application_port_range ServiceFabricManagedCluster#application_port_range}.
	ApplicationPortRange *string `field:"required" json:"applicationPortRange" yaml:"applicationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#data_disk_size_gb ServiceFabricManagedCluster#data_disk_size_gb}.
	DataDiskSizeGb *float64 `field:"required" json:"dataDiskSizeGb" yaml:"dataDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#ephemeral_port_range ServiceFabricManagedCluster#ephemeral_port_range}.
	EphemeralPortRange *string `field:"required" json:"ephemeralPortRange" yaml:"ephemeralPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#name ServiceFabricManagedCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_image_offer ServiceFabricManagedCluster#vm_image_offer}.
	VmImageOffer *string `field:"required" json:"vmImageOffer" yaml:"vmImageOffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_image_publisher ServiceFabricManagedCluster#vm_image_publisher}.
	VmImagePublisher *string `field:"required" json:"vmImagePublisher" yaml:"vmImagePublisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_image_sku ServiceFabricManagedCluster#vm_image_sku}.
	VmImageSku *string `field:"required" json:"vmImageSku" yaml:"vmImageSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_image_version ServiceFabricManagedCluster#vm_image_version}.
	VmImageVersion *string `field:"required" json:"vmImageVersion" yaml:"vmImageVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_instance_count ServiceFabricManagedCluster#vm_instance_count}.
	VmInstanceCount *float64 `field:"required" json:"vmInstanceCount" yaml:"vmInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_size ServiceFabricManagedCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#capacities ServiceFabricManagedCluster#capacities}.
	Capacities *map[string]*string `field:"optional" json:"capacities" yaml:"capacities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#data_disk_type ServiceFabricManagedCluster#data_disk_type}.
	DataDiskType *string `field:"optional" json:"dataDiskType" yaml:"dataDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#multiple_placement_groups_enabled ServiceFabricManagedCluster#multiple_placement_groups_enabled}.
	MultiplePlacementGroupsEnabled interface{} `field:"optional" json:"multiplePlacementGroupsEnabled" yaml:"multiplePlacementGroupsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#placement_properties ServiceFabricManagedCluster#placement_properties}.
	PlacementProperties *map[string]*string `field:"optional" json:"placementProperties" yaml:"placementProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#primary ServiceFabricManagedCluster#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#stateless ServiceFabricManagedCluster#stateless}.
	Stateless interface{} `field:"optional" json:"stateless" yaml:"stateless"`
	// vm_secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/service_fabric_managed_cluster#vm_secrets ServiceFabricManagedCluster#vm_secrets}
	VmSecrets interface{} `field:"optional" json:"vmSecrets" yaml:"vmSecrets"`
}

