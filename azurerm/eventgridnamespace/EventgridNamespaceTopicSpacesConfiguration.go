// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridnamespace


type EventgridNamespaceTopicSpacesConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#alternative_authentication_name_source EventgridNamespace#alternative_authentication_name_source}.
	AlternativeAuthenticationNameSource *[]*string `field:"optional" json:"alternativeAuthenticationNameSource" yaml:"alternativeAuthenticationNameSource"`
	// dynamic_routing_enrichment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#dynamic_routing_enrichment EventgridNamespace#dynamic_routing_enrichment}
	DynamicRoutingEnrichment interface{} `field:"optional" json:"dynamicRoutingEnrichment" yaml:"dynamicRoutingEnrichment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#maximum_client_sessions_per_authentication_name EventgridNamespace#maximum_client_sessions_per_authentication_name}.
	MaximumClientSessionsPerAuthenticationName *float64 `field:"optional" json:"maximumClientSessionsPerAuthenticationName" yaml:"maximumClientSessionsPerAuthenticationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#maximum_session_expiry_in_hours EventgridNamespace#maximum_session_expiry_in_hours}.
	MaximumSessionExpiryInHours *float64 `field:"optional" json:"maximumSessionExpiryInHours" yaml:"maximumSessionExpiryInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#route_topic_id EventgridNamespace#route_topic_id}.
	RouteTopicId *string `field:"optional" json:"routeTopicId" yaml:"routeTopicId"`
	// static_routing_enrichment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/eventgrid_namespace#static_routing_enrichment EventgridNamespace#static_routing_enrichment}
	StaticRoutingEnrichment interface{} `field:"optional" json:"staticRoutingEnrichment" yaml:"staticRoutingEnrichment"`
}

