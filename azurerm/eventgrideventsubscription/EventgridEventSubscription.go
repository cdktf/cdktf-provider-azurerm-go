// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/eventgrideventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription}.
type EventgridEventSubscription interface {
	cdktf.TerraformResource
	AdvancedFilter() EventgridEventSubscriptionAdvancedFilterOutputReference
	AdvancedFilteringOnArraysEnabled() interface{}
	SetAdvancedFilteringOnArraysEnabled(val interface{})
	AdvancedFilteringOnArraysEnabledInput() interface{}
	AdvancedFilterInput() *EventgridEventSubscriptionAdvancedFilter
	AzureFunctionEndpoint() EventgridEventSubscriptionAzureFunctionEndpointOutputReference
	AzureFunctionEndpointInput() *EventgridEventSubscriptionAzureFunctionEndpoint
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeadLetterIdentity() EventgridEventSubscriptionDeadLetterIdentityOutputReference
	DeadLetterIdentityInput() *EventgridEventSubscriptionDeadLetterIdentity
	DeliveryIdentity() EventgridEventSubscriptionDeliveryIdentityOutputReference
	DeliveryIdentityInput() *EventgridEventSubscriptionDeliveryIdentity
	DeliveryProperty() EventgridEventSubscriptionDeliveryPropertyList
	DeliveryPropertyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventDeliverySchema() *string
	SetEventDeliverySchema(val *string)
	EventDeliverySchemaInput() *string
	EventhubEndpointId() *string
	SetEventhubEndpointId(val *string)
	EventhubEndpointIdInput() *string
	ExpirationTimeUtc() *string
	SetExpirationTimeUtc(val *string)
	ExpirationTimeUtcInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HybridConnectionEndpointId() *string
	SetHybridConnectionEndpointId(val *string)
	HybridConnectionEndpointIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludedEventTypes() *[]*string
	SetIncludedEventTypes(val *[]*string)
	IncludedEventTypesInput() *[]*string
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RetryPolicy() EventgridEventSubscriptionRetryPolicyOutputReference
	RetryPolicyInput() *EventgridEventSubscriptionRetryPolicy
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	ServiceBusQueueEndpointId() *string
	SetServiceBusQueueEndpointId(val *string)
	ServiceBusQueueEndpointIdInput() *string
	ServiceBusTopicEndpointId() *string
	SetServiceBusTopicEndpointId(val *string)
	ServiceBusTopicEndpointIdInput() *string
	StorageBlobDeadLetterDestination() EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	StorageBlobDeadLetterDestinationInput() *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	StorageQueueEndpoint() EventgridEventSubscriptionStorageQueueEndpointOutputReference
	StorageQueueEndpointInput() *EventgridEventSubscriptionStorageQueueEndpoint
	SubjectFilter() EventgridEventSubscriptionSubjectFilterOutputReference
	SubjectFilterInput() *EventgridEventSubscriptionSubjectFilter
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EventgridEventSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebhookEndpoint() EventgridEventSubscriptionWebhookEndpointOutputReference
	WebhookEndpointInput() *EventgridEventSubscriptionWebhookEndpoint
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedFilter(value *EventgridEventSubscriptionAdvancedFilter)
	PutAzureFunctionEndpoint(value *EventgridEventSubscriptionAzureFunctionEndpoint)
	PutDeadLetterIdentity(value *EventgridEventSubscriptionDeadLetterIdentity)
	PutDeliveryIdentity(value *EventgridEventSubscriptionDeliveryIdentity)
	PutDeliveryProperty(value interface{})
	PutRetryPolicy(value *EventgridEventSubscriptionRetryPolicy)
	PutStorageBlobDeadLetterDestination(value *EventgridEventSubscriptionStorageBlobDeadLetterDestination)
	PutStorageQueueEndpoint(value *EventgridEventSubscriptionStorageQueueEndpoint)
	PutSubjectFilter(value *EventgridEventSubscriptionSubjectFilter)
	PutTimeouts(value *EventgridEventSubscriptionTimeouts)
	PutWebhookEndpoint(value *EventgridEventSubscriptionWebhookEndpoint)
	ResetAdvancedFilter()
	ResetAdvancedFilteringOnArraysEnabled()
	ResetAzureFunctionEndpoint()
	ResetDeadLetterIdentity()
	ResetDeliveryIdentity()
	ResetDeliveryProperty()
	ResetEventDeliverySchema()
	ResetEventhubEndpointId()
	ResetExpirationTimeUtc()
	ResetHybridConnectionEndpointId()
	ResetId()
	ResetIncludedEventTypes()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetryPolicy()
	ResetServiceBusQueueEndpointId()
	ResetServiceBusTopicEndpointId()
	ResetStorageBlobDeadLetterDestination()
	ResetStorageQueueEndpoint()
	ResetSubjectFilter()
	ResetTimeouts()
	ResetWebhookEndpoint()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EventgridEventSubscription
type jsiiProxy_EventgridEventSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilter() EventgridEventSubscriptionAdvancedFilterOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterOutputReference
	_jsii_.Get(
		j,
		"advancedFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilteringOnArraysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilteringOnArraysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilterInput() *EventgridEventSubscriptionAdvancedFilter {
	var returns *EventgridEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"advancedFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AzureFunctionEndpoint() EventgridEventSubscriptionAzureFunctionEndpointOutputReference {
	var returns EventgridEventSubscriptionAzureFunctionEndpointOutputReference
	_jsii_.Get(
		j,
		"azureFunctionEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AzureFunctionEndpointInput() *EventgridEventSubscriptionAzureFunctionEndpoint {
	var returns *EventgridEventSubscriptionAzureFunctionEndpoint
	_jsii_.Get(
		j,
		"azureFunctionEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeadLetterIdentity() EventgridEventSubscriptionDeadLetterIdentityOutputReference {
	var returns EventgridEventSubscriptionDeadLetterIdentityOutputReference
	_jsii_.Get(
		j,
		"deadLetterIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeadLetterIdentityInput() *EventgridEventSubscriptionDeadLetterIdentity {
	var returns *EventgridEventSubscriptionDeadLetterIdentity
	_jsii_.Get(
		j,
		"deadLetterIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryIdentity() EventgridEventSubscriptionDeliveryIdentityOutputReference {
	var returns EventgridEventSubscriptionDeliveryIdentityOutputReference
	_jsii_.Get(
		j,
		"deliveryIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryIdentityInput() *EventgridEventSubscriptionDeliveryIdentity {
	var returns *EventgridEventSubscriptionDeliveryIdentity
	_jsii_.Get(
		j,
		"deliveryIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryProperty() EventgridEventSubscriptionDeliveryPropertyList {
	var returns EventgridEventSubscriptionDeliveryPropertyList
	_jsii_.Get(
		j,
		"deliveryProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventDeliverySchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventDeliverySchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventhubEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventhubEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ExpirationTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ExpirationTimeUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) HybridConnectionEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) HybridConnectionEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IncludedEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IncludedEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RetryPolicy() EventgridEventSubscriptionRetryPolicyOutputReference {
	var returns EventgridEventSubscriptionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RetryPolicyInput() *EventgridEventSubscriptionRetryPolicy {
	var returns *EventgridEventSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusQueueEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusQueueEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusTopicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusTopicEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageBlobDeadLetterDestination() EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference {
	var returns EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageBlobDeadLetterDestinationInput() *EventgridEventSubscriptionStorageBlobDeadLetterDestination {
	var returns *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageQueueEndpoint() EventgridEventSubscriptionStorageQueueEndpointOutputReference {
	var returns EventgridEventSubscriptionStorageQueueEndpointOutputReference
	_jsii_.Get(
		j,
		"storageQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageQueueEndpointInput() *EventgridEventSubscriptionStorageQueueEndpoint {
	var returns *EventgridEventSubscriptionStorageQueueEndpoint
	_jsii_.Get(
		j,
		"storageQueueEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) SubjectFilter() EventgridEventSubscriptionSubjectFilterOutputReference {
	var returns EventgridEventSubscriptionSubjectFilterOutputReference
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) SubjectFilterInput() *EventgridEventSubscriptionSubjectFilter {
	var returns *EventgridEventSubscriptionSubjectFilter
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Timeouts() EventgridEventSubscriptionTimeoutsOutputReference {
	var returns EventgridEventSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) WebhookEndpoint() EventgridEventSubscriptionWebhookEndpointOutputReference {
	var returns EventgridEventSubscriptionWebhookEndpointOutputReference
	_jsii_.Get(
		j,
		"webhookEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) WebhookEndpointInput() *EventgridEventSubscriptionWebhookEndpoint {
	var returns *EventgridEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"webhookEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription(scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) EventgridEventSubscription {
	_init_.Initialize()

	if err := validateNewEventgridEventSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridEventSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription_Override(e EventgridEventSubscription, scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetAdvancedFilteringOnArraysEnabled(val interface{}) {
	if err := j.validateSetAdvancedFilteringOnArraysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedFilteringOnArraysEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetEventDeliverySchema(val *string) {
	if err := j.validateSetEventDeliverySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventDeliverySchema",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetEventhubEndpointId(val *string) {
	if err := j.validateSetEventhubEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetExpirationTimeUtc(val *string) {
	if err := j.validateSetExpirationTimeUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTimeUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetHybridConnectionEndpointId(val *string) {
	if err := j.validateSetHybridConnectionEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridConnectionEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetIncludedEventTypes(val *[]*string) {
	if err := j.validateSetIncludedEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedEventTypes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetServiceBusQueueEndpointId(val *string) {
	if err := j.validateSetServiceBusQueueEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusQueueEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetServiceBusTopicEndpointId(val *string) {
	if err := j.validateSetServiceBusTopicEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusTopicEndpointId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EventgridEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridEventSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridEventSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventgridEventSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAdvancedFilter(value *EventgridEventSubscriptionAdvancedFilter) {
	if err := e.validatePutAdvancedFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAdvancedFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAzureFunctionEndpoint(value *EventgridEventSubscriptionAzureFunctionEndpoint) {
	if err := e.validatePutAzureFunctionEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAzureFunctionEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeadLetterIdentity(value *EventgridEventSubscriptionDeadLetterIdentity) {
	if err := e.validatePutDeadLetterIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryIdentity(value *EventgridEventSubscriptionDeliveryIdentity) {
	if err := e.validatePutDeliveryIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryProperty(value interface{}) {
	if err := e.validatePutDeliveryPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryProperty",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutRetryPolicy(value *EventgridEventSubscriptionRetryPolicy) {
	if err := e.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageBlobDeadLetterDestination(value *EventgridEventSubscriptionStorageBlobDeadLetterDestination) {
	if err := e.validatePutStorageBlobDeadLetterDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageBlobDeadLetterDestination",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageQueueEndpoint(value *EventgridEventSubscriptionStorageQueueEndpoint) {
	if err := e.validatePutStorageQueueEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageQueueEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutSubjectFilter(value *EventgridEventSubscriptionSubjectFilter) {
	if err := e.validatePutSubjectFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSubjectFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutTimeouts(value *EventgridEventSubscriptionTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutWebhookEndpoint(value *EventgridEventSubscriptionWebhookEndpoint) {
	if err := e.validatePutWebhookEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putWebhookEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAdvancedFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAdvancedFilteringOnArraysEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilteringOnArraysEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAzureFunctionEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureFunctionEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeadLetterIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeadLetterIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeliveryIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeliveryProperty() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryProperty",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetEventDeliverySchema() {
	_jsii_.InvokeVoid(
		e,
		"resetEventDeliverySchema",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetEventhubEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetEventhubEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetExpirationTimeUtc() {
	_jsii_.InvokeVoid(
		e,
		"resetExpirationTimeUtc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetHybridConnectionEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetHybridConnectionEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetIncludedEventTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetIncludedEventTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetServiceBusQueueEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusQueueEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetServiceBusTopicEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusTopicEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetStorageBlobDeadLetterDestination() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageBlobDeadLetterDestination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetStorageQueueEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageQueueEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetWebhookEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetWebhookEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

