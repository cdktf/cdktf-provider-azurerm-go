// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractivitylogalert


type MonitorActivityLogAlertCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#category MonitorActivityLogAlert#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#caller MonitorActivityLogAlert#caller}.
	Caller *string `field:"optional" json:"caller" yaml:"caller"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#level MonitorActivityLogAlert#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#levels MonitorActivityLogAlert#levels}.
	Levels *[]*string `field:"optional" json:"levels" yaml:"levels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#operation_name MonitorActivityLogAlert#operation_name}.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#recommendation_category MonitorActivityLogAlert#recommendation_category}.
	RecommendationCategory *string `field:"optional" json:"recommendationCategory" yaml:"recommendationCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#recommendation_impact MonitorActivityLogAlert#recommendation_impact}.
	RecommendationImpact *string `field:"optional" json:"recommendationImpact" yaml:"recommendationImpact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#recommendation_type MonitorActivityLogAlert#recommendation_type}.
	RecommendationType *string `field:"optional" json:"recommendationType" yaml:"recommendationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_group MonitorActivityLogAlert#resource_group}.
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_groups MonitorActivityLogAlert#resource_groups}.
	ResourceGroups *[]*string `field:"optional" json:"resourceGroups" yaml:"resourceGroups"`
	// resource_health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_health MonitorActivityLogAlert#resource_health}
	ResourceHealth *MonitorActivityLogAlertCriteriaResourceHealth `field:"optional" json:"resourceHealth" yaml:"resourceHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_id MonitorActivityLogAlert#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_ids MonitorActivityLogAlert#resource_ids}.
	ResourceIds *[]*string `field:"optional" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_provider MonitorActivityLogAlert#resource_provider}.
	ResourceProvider *string `field:"optional" json:"resourceProvider" yaml:"resourceProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_providers MonitorActivityLogAlert#resource_providers}.
	ResourceProviders *[]*string `field:"optional" json:"resourceProviders" yaml:"resourceProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_type MonitorActivityLogAlert#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#resource_types MonitorActivityLogAlert#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// service_health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#service_health MonitorActivityLogAlert#service_health}
	ServiceHealth *MonitorActivityLogAlertCriteriaServiceHealth `field:"optional" json:"serviceHealth" yaml:"serviceHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#status MonitorActivityLogAlert#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#statuses MonitorActivityLogAlert#statuses}.
	Statuses *[]*string `field:"optional" json:"statuses" yaml:"statuses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#sub_status MonitorActivityLogAlert#sub_status}.
	SubStatus *string `field:"optional" json:"subStatus" yaml:"subStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/monitor_activity_log_alert#sub_statuses MonitorActivityLogAlert#sub_statuses}.
	SubStatuses *[]*string `field:"optional" json:"subStatuses" yaml:"subStatuses"`
}

