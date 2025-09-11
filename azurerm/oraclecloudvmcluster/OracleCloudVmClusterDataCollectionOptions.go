// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oraclecloudvmcluster


type OracleCloudVmClusterDataCollectionOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/oracle_cloud_vm_cluster#diagnostics_events_enabled OracleCloudVmCluster#diagnostics_events_enabled}.
	DiagnosticsEventsEnabled interface{} `field:"optional" json:"diagnosticsEventsEnabled" yaml:"diagnosticsEventsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/oracle_cloud_vm_cluster#health_monitoring_enabled OracleCloudVmCluster#health_monitoring_enabled}.
	HealthMonitoringEnabled interface{} `field:"optional" json:"healthMonitoringEnabled" yaml:"healthMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/oracle_cloud_vm_cluster#incident_logs_enabled OracleCloudVmCluster#incident_logs_enabled}.
	IncidentLogsEnabled interface{} `field:"optional" json:"incidentLogsEnabled" yaml:"incidentLogsEnabled"`
}

