// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilter struct {
	// bool_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#bool_equals EventgridEventSubscription#bool_equals}
	BoolEquals interface{} `field:"optional" json:"boolEquals" yaml:"boolEquals"`
	// is_not_null block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#is_not_null EventgridEventSubscription#is_not_null}
	IsNotNull interface{} `field:"optional" json:"isNotNull" yaml:"isNotNull"`
	// is_null_or_undefined block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#is_null_or_undefined EventgridEventSubscription#is_null_or_undefined}
	IsNullOrUndefined interface{} `field:"optional" json:"isNullOrUndefined" yaml:"isNullOrUndefined"`
	// number_greater_than block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_greater_than EventgridEventSubscription#number_greater_than}
	NumberGreaterThan interface{} `field:"optional" json:"numberGreaterThan" yaml:"numberGreaterThan"`
	// number_greater_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_greater_than_or_equals EventgridEventSubscription#number_greater_than_or_equals}
	NumberGreaterThanOrEquals interface{} `field:"optional" json:"numberGreaterThanOrEquals" yaml:"numberGreaterThanOrEquals"`
	// number_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_in EventgridEventSubscription#number_in}
	NumberIn interface{} `field:"optional" json:"numberIn" yaml:"numberIn"`
	// number_in_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_in_range EventgridEventSubscription#number_in_range}
	NumberInRange interface{} `field:"optional" json:"numberInRange" yaml:"numberInRange"`
	// number_less_than block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_less_than EventgridEventSubscription#number_less_than}
	NumberLessThan interface{} `field:"optional" json:"numberLessThan" yaml:"numberLessThan"`
	// number_less_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_less_than_or_equals EventgridEventSubscription#number_less_than_or_equals}
	NumberLessThanOrEquals interface{} `field:"optional" json:"numberLessThanOrEquals" yaml:"numberLessThanOrEquals"`
	// number_not_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_not_in EventgridEventSubscription#number_not_in}
	NumberNotIn interface{} `field:"optional" json:"numberNotIn" yaml:"numberNotIn"`
	// number_not_in_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#number_not_in_range EventgridEventSubscription#number_not_in_range}
	NumberNotInRange interface{} `field:"optional" json:"numberNotInRange" yaml:"numberNotInRange"`
	// string_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_begins_with EventgridEventSubscription#string_begins_with}
	StringBeginsWith interface{} `field:"optional" json:"stringBeginsWith" yaml:"stringBeginsWith"`
	// string_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_contains EventgridEventSubscription#string_contains}
	StringContains interface{} `field:"optional" json:"stringContains" yaml:"stringContains"`
	// string_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_ends_with EventgridEventSubscription#string_ends_with}
	StringEndsWith interface{} `field:"optional" json:"stringEndsWith" yaml:"stringEndsWith"`
	// string_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_in EventgridEventSubscription#string_in}
	StringIn interface{} `field:"optional" json:"stringIn" yaml:"stringIn"`
	// string_not_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_not_begins_with EventgridEventSubscription#string_not_begins_with}
	StringNotBeginsWith interface{} `field:"optional" json:"stringNotBeginsWith" yaml:"stringNotBeginsWith"`
	// string_not_contains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_not_contains EventgridEventSubscription#string_not_contains}
	StringNotContains interface{} `field:"optional" json:"stringNotContains" yaml:"stringNotContains"`
	// string_not_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_not_ends_with EventgridEventSubscription#string_not_ends_with}
	StringNotEndsWith interface{} `field:"optional" json:"stringNotEndsWith" yaml:"stringNotEndsWith"`
	// string_not_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/eventgrid_event_subscription#string_not_in EventgridEventSubscription#string_not_in}
	StringNotIn interface{} `field:"optional" json:"stringNotIn" yaml:"stringNotIn"`
}

