// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterUpgradePolicy struct {
	// delta_health_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#delta_health_policy ServiceFabricCluster#delta_health_policy}
	DeltaHealthPolicy *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy `field:"optional" json:"deltaHealthPolicy" yaml:"deltaHealthPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#force_restart_enabled ServiceFabricCluster#force_restart_enabled}.
	ForceRestartEnabled interface{} `field:"optional" json:"forceRestartEnabled" yaml:"forceRestartEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#health_check_retry_timeout ServiceFabricCluster#health_check_retry_timeout}.
	HealthCheckRetryTimeout *string `field:"optional" json:"healthCheckRetryTimeout" yaml:"healthCheckRetryTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#health_check_stable_duration ServiceFabricCluster#health_check_stable_duration}.
	HealthCheckStableDuration *string `field:"optional" json:"healthCheckStableDuration" yaml:"healthCheckStableDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#health_check_wait_duration ServiceFabricCluster#health_check_wait_duration}.
	HealthCheckWaitDuration *string `field:"optional" json:"healthCheckWaitDuration" yaml:"healthCheckWaitDuration"`
	// health_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#health_policy ServiceFabricCluster#health_policy}
	HealthPolicy *ServiceFabricClusterUpgradePolicyHealthPolicy `field:"optional" json:"healthPolicy" yaml:"healthPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#upgrade_domain_timeout ServiceFabricCluster#upgrade_domain_timeout}.
	UpgradeDomainTimeout *string `field:"optional" json:"upgradeDomainTimeout" yaml:"upgradeDomainTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#upgrade_replica_set_check_timeout ServiceFabricCluster#upgrade_replica_set_check_timeout}.
	UpgradeReplicaSetCheckTimeout *string `field:"optional" json:"upgradeReplicaSetCheckTimeout" yaml:"upgradeReplicaSetCheckTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/service_fabric_cluster#upgrade_timeout ServiceFabricCluster#upgrade_timeout}.
	UpgradeTimeout *string `field:"optional" json:"upgradeTimeout" yaml:"upgradeTimeout"`
}

