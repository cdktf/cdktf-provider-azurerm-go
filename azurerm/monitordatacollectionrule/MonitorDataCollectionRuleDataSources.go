// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule


type MonitorDataCollectionRuleDataSources struct {
	// data_import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#data_import MonitorDataCollectionRule#data_import}
	DataImport *MonitorDataCollectionRuleDataSourcesDataImport `field:"optional" json:"dataImport" yaml:"dataImport"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#extension MonitorDataCollectionRule#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// iis_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#iis_log MonitorDataCollectionRule#iis_log}
	IisLog interface{} `field:"optional" json:"iisLog" yaml:"iisLog"`
	// log_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#log_file MonitorDataCollectionRule#log_file}
	LogFile interface{} `field:"optional" json:"logFile" yaml:"logFile"`
	// performance_counter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#performance_counter MonitorDataCollectionRule#performance_counter}
	PerformanceCounter interface{} `field:"optional" json:"performanceCounter" yaml:"performanceCounter"`
	// platform_telemetry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#platform_telemetry MonitorDataCollectionRule#platform_telemetry}
	PlatformTelemetry interface{} `field:"optional" json:"platformTelemetry" yaml:"platformTelemetry"`
	// prometheus_forwarder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#prometheus_forwarder MonitorDataCollectionRule#prometheus_forwarder}
	PrometheusForwarder interface{} `field:"optional" json:"prometheusForwarder" yaml:"prometheusForwarder"`
	// syslog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#syslog MonitorDataCollectionRule#syslog}
	Syslog interface{} `field:"optional" json:"syslog" yaml:"syslog"`
	// windows_event_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#windows_event_log MonitorDataCollectionRule#windows_event_log}
	WindowsEventLog interface{} `field:"optional" json:"windowsEventLog" yaml:"windowsEventLog"`
	// windows_firewall_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/monitor_data_collection_rule#windows_firewall_log MonitorDataCollectionRule#windows_firewall_log}
	WindowsFirewallLog interface{} `field:"optional" json:"windowsFirewallLog" yaml:"windowsFirewallLog"`
}

