// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/eventgridsystemtopiceventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription}.
type EventgridSystemTopicEventSubscription interface {
	cdktf.TerraformResource
	AdvancedFilter() EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference
	AdvancedFilteringOnArraysEnabled() interface{}
	SetAdvancedFilteringOnArraysEnabled(val interface{})
	AdvancedFilteringOnArraysEnabledInput() interface{}
	AdvancedFilterInput() *EventgridSystemTopicEventSubscriptionAdvancedFilter
	AzureFunctionEndpoint() EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference
	AzureFunctionEndpointInput() *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint
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
	DeadLetterIdentity() EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference
	DeadLetterIdentityInput() *EventgridSystemTopicEventSubscriptionDeadLetterIdentity
	DeliveryIdentity() EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference
	DeliveryIdentityInput() *EventgridSystemTopicEventSubscriptionDeliveryIdentity
	DeliveryProperty() EventgridSystemTopicEventSubscriptionDeliveryPropertyList
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RetryPolicy() EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference
	RetryPolicyInput() *EventgridSystemTopicEventSubscriptionRetryPolicy
	ServiceBusQueueEndpointId() *string
	SetServiceBusQueueEndpointId(val *string)
	ServiceBusQueueEndpointIdInput() *string
	ServiceBusTopicEndpointId() *string
	SetServiceBusTopicEndpointId(val *string)
	ServiceBusTopicEndpointIdInput() *string
	StorageBlobDeadLetterDestination() EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	StorageBlobDeadLetterDestinationInput() *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination
	StorageQueueEndpoint() EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference
	StorageQueueEndpointInput() *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint
	SubjectFilter() EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference
	SubjectFilterInput() *EventgridSystemTopicEventSubscriptionSubjectFilter
	SystemTopic() *string
	SetSystemTopic(val *string)
	SystemTopicInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EventgridSystemTopicEventSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebhookEndpoint() EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference
	WebhookEndpointInput() *EventgridSystemTopicEventSubscriptionWebhookEndpoint
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedFilter(value *EventgridSystemTopicEventSubscriptionAdvancedFilter)
	PutAzureFunctionEndpoint(value *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint)
	PutDeadLetterIdentity(value *EventgridSystemTopicEventSubscriptionDeadLetterIdentity)
	PutDeliveryIdentity(value *EventgridSystemTopicEventSubscriptionDeliveryIdentity)
	PutDeliveryProperty(value interface{})
	PutRetryPolicy(value *EventgridSystemTopicEventSubscriptionRetryPolicy)
	PutStorageBlobDeadLetterDestination(value *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination)
	PutStorageQueueEndpoint(value *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint)
	PutSubjectFilter(value *EventgridSystemTopicEventSubscriptionSubjectFilter)
	PutTimeouts(value *EventgridSystemTopicEventSubscriptionTimeouts)
	PutWebhookEndpoint(value *EventgridSystemTopicEventSubscriptionWebhookEndpoint)
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
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EventgridSystemTopicEventSubscription
type jsiiProxy_EventgridSystemTopicEventSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AdvancedFilter() EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference
	_jsii_.Get(
		j,
		"advancedFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AdvancedFilteringOnArraysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AdvancedFilteringOnArraysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AdvancedFilterInput() *EventgridSystemTopicEventSubscriptionAdvancedFilter {
	var returns *EventgridSystemTopicEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"advancedFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AzureFunctionEndpoint() EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference
	_jsii_.Get(
		j,
		"azureFunctionEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) AzureFunctionEndpointInput() *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint
	_jsii_.Get(
		j,
		"azureFunctionEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeadLetterIdentity() EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference {
	var returns EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference
	_jsii_.Get(
		j,
		"deadLetterIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeadLetterIdentityInput() *EventgridSystemTopicEventSubscriptionDeadLetterIdentity {
	var returns *EventgridSystemTopicEventSubscriptionDeadLetterIdentity
	_jsii_.Get(
		j,
		"deadLetterIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeliveryIdentity() EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference {
	var returns EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference
	_jsii_.Get(
		j,
		"deliveryIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeliveryIdentityInput() *EventgridSystemTopicEventSubscriptionDeliveryIdentity {
	var returns *EventgridSystemTopicEventSubscriptionDeliveryIdentity
	_jsii_.Get(
		j,
		"deliveryIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeliveryProperty() EventgridSystemTopicEventSubscriptionDeliveryPropertyList {
	var returns EventgridSystemTopicEventSubscriptionDeliveryPropertyList
	_jsii_.Get(
		j,
		"deliveryProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DeliveryPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) EventDeliverySchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) EventDeliverySchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) EventhubEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) EventhubEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ExpirationTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ExpirationTimeUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) HybridConnectionEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) HybridConnectionEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) IncludedEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) IncludedEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) RetryPolicy() EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference {
	var returns EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) RetryPolicyInput() *EventgridSystemTopicEventSubscriptionRetryPolicy {
	var returns *EventgridSystemTopicEventSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ServiceBusQueueEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ServiceBusQueueEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ServiceBusTopicEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) ServiceBusTopicEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) StorageBlobDeadLetterDestination() EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference {
	var returns EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) StorageBlobDeadLetterDestinationInput() *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination {
	var returns *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) StorageQueueEndpoint() EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference {
	var returns EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference
	_jsii_.Get(
		j,
		"storageQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) StorageQueueEndpointInput() *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint
	_jsii_.Get(
		j,
		"storageQueueEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SubjectFilter() EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference {
	var returns EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SubjectFilterInput() *EventgridSystemTopicEventSubscriptionSubjectFilter {
	var returns *EventgridSystemTopicEventSubscriptionSubjectFilter
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SystemTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SystemTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Timeouts() EventgridSystemTopicEventSubscriptionTimeoutsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) WebhookEndpoint() EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference {
	var returns EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference
	_jsii_.Get(
		j,
		"webhookEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) WebhookEndpointInput() *EventgridSystemTopicEventSubscriptionWebhookEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"webhookEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription} Resource.
func NewEventgridSystemTopicEventSubscription(scope constructs.Construct, id *string, config *EventgridSystemTopicEventSubscriptionConfig) EventgridSystemTopicEventSubscription {
	_init_.Initialize()

	if err := validateNewEventgridSystemTopicEventSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridSystemTopicEventSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription} Resource.
func NewEventgridSystemTopicEventSubscription_Override(e EventgridSystemTopicEventSubscription, scope constructs.Construct, id *string, config *EventgridSystemTopicEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetAdvancedFilteringOnArraysEnabled(val interface{}) {
	if err := j.validateSetAdvancedFilteringOnArraysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedFilteringOnArraysEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetEventDeliverySchema(val *string) {
	if err := j.validateSetEventDeliverySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventDeliverySchema",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetEventhubEndpointId(val *string) {
	if err := j.validateSetEventhubEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetExpirationTimeUtc(val *string) {
	if err := j.validateSetExpirationTimeUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTimeUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetHybridConnectionEndpointId(val *string) {
	if err := j.validateSetHybridConnectionEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridConnectionEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetIncludedEventTypes(val *[]*string) {
	if err := j.validateSetIncludedEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedEventTypes",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetServiceBusQueueEndpointId(val *string) {
	if err := j.validateSetServiceBusQueueEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusQueueEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetServiceBusTopicEndpointId(val *string) {
	if err := j.validateSetServiceBusTopicEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusTopicEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription)SetSystemTopic(val *string) {
	if err := j.validateSetSystemTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemTopic",
		val,
	)
}

// Generates CDKTF code for importing a EventgridSystemTopicEventSubscription resource upon running "cdktf plan <stack-name>".
func EventgridSystemTopicEventSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEventgridSystemTopicEventSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func EventgridSystemTopicEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridSystemTopicEventSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridSystemTopicEventSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridSystemTopicEventSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridSystemTopicEventSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridSystemTopicEventSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventgridSystemTopicEventSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutAdvancedFilter(value *EventgridSystemTopicEventSubscriptionAdvancedFilter) {
	if err := e.validatePutAdvancedFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAdvancedFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutAzureFunctionEndpoint(value *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint) {
	if err := e.validatePutAzureFunctionEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAzureFunctionEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeadLetterIdentity(value *EventgridSystemTopicEventSubscriptionDeadLetterIdentity) {
	if err := e.validatePutDeadLetterIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeliveryIdentity(value *EventgridSystemTopicEventSubscriptionDeliveryIdentity) {
	if err := e.validatePutDeliveryIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeliveryProperty(value interface{}) {
	if err := e.validatePutDeliveryPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryProperty",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutRetryPolicy(value *EventgridSystemTopicEventSubscriptionRetryPolicy) {
	if err := e.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutStorageBlobDeadLetterDestination(value *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination) {
	if err := e.validatePutStorageBlobDeadLetterDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageBlobDeadLetterDestination",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutStorageQueueEndpoint(value *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint) {
	if err := e.validatePutStorageQueueEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageQueueEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutSubjectFilter(value *EventgridSystemTopicEventSubscriptionSubjectFilter) {
	if err := e.validatePutSubjectFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSubjectFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutTimeouts(value *EventgridSystemTopicEventSubscriptionTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutWebhookEndpoint(value *EventgridSystemTopicEventSubscriptionWebhookEndpoint) {
	if err := e.validatePutWebhookEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putWebhookEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetAdvancedFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetAdvancedFilteringOnArraysEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilteringOnArraysEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetAzureFunctionEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureFunctionEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetDeadLetterIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeadLetterIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetDeliveryIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetDeliveryProperty() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryProperty",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetEventDeliverySchema() {
	_jsii_.InvokeVoid(
		e,
		"resetEventDeliverySchema",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetEventhubEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetEventhubEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetExpirationTimeUtc() {
	_jsii_.InvokeVoid(
		e,
		"resetExpirationTimeUtc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetHybridConnectionEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetHybridConnectionEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetIncludedEventTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetIncludedEventTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetServiceBusQueueEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusQueueEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetServiceBusTopicEndpointId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusTopicEndpointId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetStorageBlobDeadLetterDestination() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageBlobDeadLetterDestination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetStorageQueueEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageQueueEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ResetWebhookEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetWebhookEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

