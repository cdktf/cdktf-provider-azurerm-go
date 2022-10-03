package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/eventgrideventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription azurerm_eventgrid_event_subscription}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_EventgridEventSubscription) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription(scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) EventgridEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription_Override(e EventgridEventSubscription, scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetAdvancedFilteringOnArraysEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"advancedFilteringOnArraysEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetEventDeliverySchema(val *string) {
	_jsii_.Set(
		j,
		"eventDeliverySchema",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetEventhubEndpointId(val *string) {
	_jsii_.Set(
		j,
		"eventhubEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetExpirationTimeUtc(val *string) {
	_jsii_.Set(
		j,
		"expirationTimeUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetHybridConnectionEndpointId(val *string) {
	_jsii_.Set(
		j,
		"hybridConnectionEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetIncludedEventTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"includedEventTypes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetLabels(val *[]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetServiceBusQueueEndpointId(val *string) {
	_jsii_.Set(
		j,
		"serviceBusQueueEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription) SetServiceBusTopicEndpointId(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAdvancedFilter(value *EventgridEventSubscriptionAdvancedFilter) {
	_jsii_.InvokeVoid(
		e,
		"putAdvancedFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAzureFunctionEndpoint(value *EventgridEventSubscriptionAzureFunctionEndpoint) {
	_jsii_.InvokeVoid(
		e,
		"putAzureFunctionEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeadLetterIdentity(value *EventgridEventSubscriptionDeadLetterIdentity) {
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryIdentity(value *EventgridEventSubscriptionDeliveryIdentity) {
	_jsii_.InvokeVoid(
		e,
		"putDeliveryIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryProperty(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putDeliveryProperty",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutRetryPolicy(value *EventgridEventSubscriptionRetryPolicy) {
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageBlobDeadLetterDestination(value *EventgridEventSubscriptionStorageBlobDeadLetterDestination) {
	_jsii_.InvokeVoid(
		e,
		"putStorageBlobDeadLetterDestination",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageQueueEndpoint(value *EventgridEventSubscriptionStorageQueueEndpoint) {
	_jsii_.InvokeVoid(
		e,
		"putStorageQueueEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutSubjectFilter(value *EventgridEventSubscriptionSubjectFilter) {
	_jsii_.InvokeVoid(
		e,
		"putSubjectFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutTimeouts(value *EventgridEventSubscriptionTimeouts) {
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutWebhookEndpoint(value *EventgridEventSubscriptionWebhookEndpoint) {
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

type EventgridEventSubscriptionAdvancedFilter struct {
	// bool_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#bool_equals EventgridEventSubscription#bool_equals}
	BoolEquals interface{} `field:"optional" json:"boolEquals" yaml:"boolEquals"`
	// is_not_null block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#is_not_null EventgridEventSubscription#is_not_null}
	IsNotNull interface{} `field:"optional" json:"isNotNull" yaml:"isNotNull"`
	// is_null_or_undefined block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#is_null_or_undefined EventgridEventSubscription#is_null_or_undefined}
	IsNullOrUndefined interface{} `field:"optional" json:"isNullOrUndefined" yaml:"isNullOrUndefined"`
	// number_greater_than block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_greater_than EventgridEventSubscription#number_greater_than}
	NumberGreaterThan interface{} `field:"optional" json:"numberGreaterThan" yaml:"numberGreaterThan"`
	// number_greater_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_greater_than_or_equals EventgridEventSubscription#number_greater_than_or_equals}
	NumberGreaterThanOrEquals interface{} `field:"optional" json:"numberGreaterThanOrEquals" yaml:"numberGreaterThanOrEquals"`
	// number_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_in EventgridEventSubscription#number_in}
	NumberIn interface{} `field:"optional" json:"numberIn" yaml:"numberIn"`
	// number_in_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_in_range EventgridEventSubscription#number_in_range}
	NumberInRange interface{} `field:"optional" json:"numberInRange" yaml:"numberInRange"`
	// number_less_than block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_less_than EventgridEventSubscription#number_less_than}
	NumberLessThan interface{} `field:"optional" json:"numberLessThan" yaml:"numberLessThan"`
	// number_less_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_less_than_or_equals EventgridEventSubscription#number_less_than_or_equals}
	NumberLessThanOrEquals interface{} `field:"optional" json:"numberLessThanOrEquals" yaml:"numberLessThanOrEquals"`
	// number_not_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_not_in EventgridEventSubscription#number_not_in}
	NumberNotIn interface{} `field:"optional" json:"numberNotIn" yaml:"numberNotIn"`
	// number_not_in_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#number_not_in_range EventgridEventSubscription#number_not_in_range}
	NumberNotInRange interface{} `field:"optional" json:"numberNotInRange" yaml:"numberNotInRange"`
	// string_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_begins_with EventgridEventSubscription#string_begins_with}
	StringBeginsWith interface{} `field:"optional" json:"stringBeginsWith" yaml:"stringBeginsWith"`
	// string_contains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_contains EventgridEventSubscription#string_contains}
	StringContains interface{} `field:"optional" json:"stringContains" yaml:"stringContains"`
	// string_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_ends_with EventgridEventSubscription#string_ends_with}
	StringEndsWith interface{} `field:"optional" json:"stringEndsWith" yaml:"stringEndsWith"`
	// string_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_in EventgridEventSubscription#string_in}
	StringIn interface{} `field:"optional" json:"stringIn" yaml:"stringIn"`
	// string_not_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_not_begins_with EventgridEventSubscription#string_not_begins_with}
	StringNotBeginsWith interface{} `field:"optional" json:"stringNotBeginsWith" yaml:"stringNotBeginsWith"`
	// string_not_contains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_not_contains EventgridEventSubscription#string_not_contains}
	StringNotContains interface{} `field:"optional" json:"stringNotContains" yaml:"stringNotContains"`
	// string_not_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_not_ends_with EventgridEventSubscription#string_not_ends_with}
	StringNotEndsWith interface{} `field:"optional" json:"stringNotEndsWith" yaml:"stringNotEndsWith"`
	// string_not_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#string_not_in EventgridEventSubscription#string_not_in}
	StringNotIn interface{} `field:"optional" json:"stringNotIn" yaml:"stringNotIn"`
}

type EventgridEventSubscriptionAdvancedFilterBoolEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionAdvancedFilterBoolEqualsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterBoolEqualsList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterBoolEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterBoolEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterBoolEqualsList_Override(e EventgridEventSubscriptionAdvancedFilterBoolEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() interface{}
	SetValue(val interface{})
	ValueInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetValue(val interface{}) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterIsNotNull struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

type EventgridEventSubscriptionAdvancedFilterIsNotNullList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterIsNotNullList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterIsNotNullList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterIsNotNullList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNotNullList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterIsNotNullList_Override(e EventgridEventSubscriptionAdvancedFilterIsNotNullList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNotNullList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNotNullOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterIsNullOrUndefined struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

type EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList_Override(e EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanList_Override(e EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList_Override(e EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterNumberInList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberInList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberInList_Override(e EventgridEventSubscriptionAdvancedFilterNumberInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberInOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberInOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*float64
	SetValues(val *[]*float64)
	ValuesInput() *[]*float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberInOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberInOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberInRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterNumberInRangeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberInRangeList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberInRangeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberInRangeList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberInRangeList_Override(e EventgridEventSubscriptionAdvancedFilterNumberInRangeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*float64
	SetValues(val *[]*float64)
	ValuesInput() *[]*float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThanList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberLessThanList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberLessThanList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanList_Override(e EventgridEventSubscriptionAdvancedFilterNumberLessThanList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList_Override(e EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberNotIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterNumberNotInList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberNotInList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberNotInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberNotInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberNotInList_Override(e EventgridEventSubscriptionAdvancedFilterNumberNotInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*float64
	SetValues(val *[]*float64)
	ValuesInput() *[]*float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberNotInRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberNotInRangeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberNotInRangeList_Override(e EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*float64
	SetValues(val *[]*float64)
	ValuesInput() *[]*float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterOutputReference interface {
	cdktf.ComplexObject
	BoolEquals() EventgridEventSubscriptionAdvancedFilterBoolEqualsList
	BoolEqualsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionAdvancedFilter
	SetInternalValue(val *EventgridEventSubscriptionAdvancedFilter)
	IsNotNull() EventgridEventSubscriptionAdvancedFilterIsNotNullList
	IsNotNullInput() interface{}
	IsNullOrUndefined() EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	IsNullOrUndefinedInput() interface{}
	NumberGreaterThan() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList
	NumberGreaterThanInput() interface{}
	NumberGreaterThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	NumberGreaterThanOrEqualsInput() interface{}
	NumberIn() EventgridEventSubscriptionAdvancedFilterNumberInList
	NumberInInput() interface{}
	NumberInRange() EventgridEventSubscriptionAdvancedFilterNumberInRangeList
	NumberInRangeInput() interface{}
	NumberLessThan() EventgridEventSubscriptionAdvancedFilterNumberLessThanList
	NumberLessThanInput() interface{}
	NumberLessThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	NumberLessThanOrEqualsInput() interface{}
	NumberNotIn() EventgridEventSubscriptionAdvancedFilterNumberNotInList
	NumberNotInInput() interface{}
	NumberNotInRange() EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList
	NumberNotInRangeInput() interface{}
	StringBeginsWith() EventgridEventSubscriptionAdvancedFilterStringBeginsWithList
	StringBeginsWithInput() interface{}
	StringContains() EventgridEventSubscriptionAdvancedFilterStringContainsList
	StringContainsInput() interface{}
	StringEndsWith() EventgridEventSubscriptionAdvancedFilterStringEndsWithList
	StringEndsWithInput() interface{}
	StringIn() EventgridEventSubscriptionAdvancedFilterStringInList
	StringInInput() interface{}
	StringNotBeginsWith() EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList
	StringNotBeginsWithInput() interface{}
	StringNotContains() EventgridEventSubscriptionAdvancedFilterStringNotContainsList
	StringNotContainsInput() interface{}
	StringNotEndsWith() EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList
	StringNotEndsWithInput() interface{}
	StringNotIn() EventgridEventSubscriptionAdvancedFilterStringNotInList
	StringNotInInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutBoolEquals(value interface{})
	PutIsNotNull(value interface{})
	PutIsNullOrUndefined(value interface{})
	PutNumberGreaterThan(value interface{})
	PutNumberGreaterThanOrEquals(value interface{})
	PutNumberIn(value interface{})
	PutNumberInRange(value interface{})
	PutNumberLessThan(value interface{})
	PutNumberLessThanOrEquals(value interface{})
	PutNumberNotIn(value interface{})
	PutNumberNotInRange(value interface{})
	PutStringBeginsWith(value interface{})
	PutStringContains(value interface{})
	PutStringEndsWith(value interface{})
	PutStringIn(value interface{})
	PutStringNotBeginsWith(value interface{})
	PutStringNotContains(value interface{})
	PutStringNotEndsWith(value interface{})
	PutStringNotIn(value interface{})
	ResetBoolEquals()
	ResetIsNotNull()
	ResetIsNullOrUndefined()
	ResetNumberGreaterThan()
	ResetNumberGreaterThanOrEquals()
	ResetNumberIn()
	ResetNumberInRange()
	ResetNumberLessThan()
	ResetNumberLessThanOrEquals()
	ResetNumberNotIn()
	ResetNumberNotInRange()
	ResetStringBeginsWith()
	ResetStringContains()
	ResetStringEndsWith()
	ResetStringIn()
	ResetStringNotBeginsWith()
	ResetStringNotContains()
	ResetStringNotEndsWith()
	ResetStringNotIn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) BoolEquals() EventgridEventSubscriptionAdvancedFilterBoolEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterBoolEqualsList
	_jsii_.Get(
		j,
		"boolEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) BoolEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InternalValue() *EventgridEventSubscriptionAdvancedFilter {
	var returns *EventgridEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNotNull() EventgridEventSubscriptionAdvancedFilterIsNotNullList {
	var returns EventgridEventSubscriptionAdvancedFilterIsNotNullList
	_jsii_.Get(
		j,
		"isNotNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNotNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNotNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefined() EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	var returns EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	_jsii_.Get(
		j,
		"isNullOrUndefined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefinedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullOrUndefinedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThan() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList
	_jsii_.Get(
		j,
		"numberGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	_jsii_.Get(
		j,
		"numberGreaterThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberIn() EventgridEventSubscriptionAdvancedFilterNumberInList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInList
	_jsii_.Get(
		j,
		"numberIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInRange() EventgridEventSubscriptionAdvancedFilterNumberInRangeList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInRangeList
	_jsii_.Get(
		j,
		"numberInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThan() EventgridEventSubscriptionAdvancedFilterNumberLessThanList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanList
	_jsii_.Get(
		j,
		"numberLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	_jsii_.Get(
		j,
		"numberLessThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotIn() EventgridEventSubscriptionAdvancedFilterNumberNotInList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInList
	_jsii_.Get(
		j,
		"numberNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInRange() EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList
	_jsii_.Get(
		j,
		"numberNotInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringBeginsWith() EventgridEventSubscriptionAdvancedFilterStringBeginsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringBeginsWithList
	_jsii_.Get(
		j,
		"stringBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringContains() EventgridEventSubscriptionAdvancedFilterStringContainsList {
	var returns EventgridEventSubscriptionAdvancedFilterStringContainsList
	_jsii_.Get(
		j,
		"stringContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringEndsWith() EventgridEventSubscriptionAdvancedFilterStringEndsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringEndsWithList
	_jsii_.Get(
		j,
		"stringEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringIn() EventgridEventSubscriptionAdvancedFilterStringInList {
	var returns EventgridEventSubscriptionAdvancedFilterStringInList
	_jsii_.Get(
		j,
		"stringIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWith() EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList
	_jsii_.Get(
		j,
		"stringNotBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotContains() EventgridEventSubscriptionAdvancedFilterStringNotContainsList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotContainsList
	_jsii_.Get(
		j,
		"stringNotContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWith() EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList
	_jsii_.Get(
		j,
		"stringNotEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotIn() EventgridEventSubscriptionAdvancedFilterStringNotInList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotInList
	_jsii_.Get(
		j,
		"stringNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionAdvancedFilterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) SetInternalValue(val *EventgridEventSubscriptionAdvancedFilter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutBoolEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putBoolEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutIsNotNull(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putIsNotNull",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutIsNullOrUndefined(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putIsNullOrUndefined",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThan(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThanOrEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberInRange(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThan(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThanOrEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberNotIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberNotInRange(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberNotInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringBeginsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringContains(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringEndsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotBeginsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotContains(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotEndsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetBoolEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetBoolEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetIsNotNull() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNotNull",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetIsNullOrUndefined() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNullOrUndefined",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringBeginsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringBeginsWithList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringBeginsWithList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringBeginsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringBeginsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringBeginsWithList_Override(e EventgridEventSubscriptionAdvancedFilterStringBeginsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringContainsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringContainsList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringContainsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringContainsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringContainsList_Override(e EventgridEventSubscriptionAdvancedFilterStringContainsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringContainsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringContainsOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringContainsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringEndsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringEndsWithList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringEndsWithList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringEndsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringEndsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringEndsWithList_Override(e EventgridEventSubscriptionAdvancedFilterStringEndsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringInList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringInList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringInList_Override(e EventgridEventSubscriptionAdvancedFilterStringInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringInOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringInOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringInOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringInOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotBeginsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList_Override(e EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringNotContainsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotContainsList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotContainsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotContainsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotContainsList_Override(e EventgridEventSubscriptionAdvancedFilterStringNotContainsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotEndsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotEndsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotEndsWithList_Override(e EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridEventSubscriptionAdvancedFilterStringNotInList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotInList
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotInList_Override(e EventgridEventSubscriptionAdvancedFilterStringNotInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) Get(index *float64) EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterStringNotInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterStringNotInOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterStringNotInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionAzureFunctionEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#function_id EventgridEventSubscription#function_id}.
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#max_events_per_batch EventgridEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#preferred_batch_size_in_kilobytes EventgridEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

type EventgridEventSubscriptionAzureFunctionEndpointOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	FunctionId() *string
	SetFunctionId(val *string)
	FunctionIdInput() *string
	InternalValue() *EventgridEventSubscriptionAzureFunctionEndpoint
	SetInternalValue(val *EventgridEventSubscriptionAzureFunctionEndpoint)
	MaxEventsPerBatch() *float64
	SetMaxEventsPerBatch(val *float64)
	MaxEventsPerBatchInput() *float64
	PreferredBatchSizeInKilobytes() *float64
	SetPreferredBatchSizeInKilobytes(val *float64)
	PreferredBatchSizeInKilobytesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxEventsPerBatch()
	ResetPreferredBatchSizeInKilobytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAzureFunctionEndpointOutputReference
type jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) FunctionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) InternalValue() *EventgridEventSubscriptionAzureFunctionEndpoint {
	var returns *EventgridEventSubscriptionAzureFunctionEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) MaxEventsPerBatch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) MaxEventsPerBatchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) PreferredBatchSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) PreferredBatchSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAzureFunctionEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionAzureFunctionEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAzureFunctionEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAzureFunctionEndpointOutputReference_Override(e EventgridEventSubscriptionAzureFunctionEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAzureFunctionEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetFunctionId(val *string) {
	_jsii_.Set(
		j,
		"functionId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetInternalValue(val *EventgridEventSubscriptionAzureFunctionEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetMaxEventsPerBatch(val *float64) {
	_jsii_.Set(
		j,
		"maxEventsPerBatch",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetPreferredBatchSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"preferredBatchSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ResetMaxEventsPerBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxEventsPerBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ResetPreferredBatchSizeInKilobytes() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredBatchSizeInKilobytes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAzureFunctionEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#name EventgridEventSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#scope EventgridEventSubscription#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// advanced_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#advanced_filter EventgridEventSubscription#advanced_filter}
	AdvancedFilter *EventgridEventSubscriptionAdvancedFilter `field:"optional" json:"advancedFilter" yaml:"advancedFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#advanced_filtering_on_arrays_enabled EventgridEventSubscription#advanced_filtering_on_arrays_enabled}.
	AdvancedFilteringOnArraysEnabled interface{} `field:"optional" json:"advancedFilteringOnArraysEnabled" yaml:"advancedFilteringOnArraysEnabled"`
	// azure_function_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#azure_function_endpoint EventgridEventSubscription#azure_function_endpoint}
	AzureFunctionEndpoint *EventgridEventSubscriptionAzureFunctionEndpoint `field:"optional" json:"azureFunctionEndpoint" yaml:"azureFunctionEndpoint"`
	// dead_letter_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#dead_letter_identity EventgridEventSubscription#dead_letter_identity}
	DeadLetterIdentity *EventgridEventSubscriptionDeadLetterIdentity `field:"optional" json:"deadLetterIdentity" yaml:"deadLetterIdentity"`
	// delivery_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#delivery_identity EventgridEventSubscription#delivery_identity}
	DeliveryIdentity *EventgridEventSubscriptionDeliveryIdentity `field:"optional" json:"deliveryIdentity" yaml:"deliveryIdentity"`
	// delivery_property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#delivery_property EventgridEventSubscription#delivery_property}
	DeliveryProperty interface{} `field:"optional" json:"deliveryProperty" yaml:"deliveryProperty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#event_delivery_schema EventgridEventSubscription#event_delivery_schema}.
	EventDeliverySchema *string `field:"optional" json:"eventDeliverySchema" yaml:"eventDeliverySchema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#eventhub_endpoint_id EventgridEventSubscription#eventhub_endpoint_id}.
	EventhubEndpointId *string `field:"optional" json:"eventhubEndpointId" yaml:"eventhubEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#expiration_time_utc EventgridEventSubscription#expiration_time_utc}.
	ExpirationTimeUtc *string `field:"optional" json:"expirationTimeUtc" yaml:"expirationTimeUtc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#hybrid_connection_endpoint_id EventgridEventSubscription#hybrid_connection_endpoint_id}.
	HybridConnectionEndpointId *string `field:"optional" json:"hybridConnectionEndpointId" yaml:"hybridConnectionEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#id EventgridEventSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#included_event_types EventgridEventSubscription#included_event_types}.
	IncludedEventTypes *[]*string `field:"optional" json:"includedEventTypes" yaml:"includedEventTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#labels EventgridEventSubscription#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#retry_policy EventgridEventSubscription#retry_policy}
	RetryPolicy *EventgridEventSubscriptionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#service_bus_queue_endpoint_id EventgridEventSubscription#service_bus_queue_endpoint_id}.
	ServiceBusQueueEndpointId *string `field:"optional" json:"serviceBusQueueEndpointId" yaml:"serviceBusQueueEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#service_bus_topic_endpoint_id EventgridEventSubscription#service_bus_topic_endpoint_id}.
	ServiceBusTopicEndpointId *string `field:"optional" json:"serviceBusTopicEndpointId" yaml:"serviceBusTopicEndpointId"`
	// storage_blob_dead_letter_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#storage_blob_dead_letter_destination EventgridEventSubscription#storage_blob_dead_letter_destination}
	StorageBlobDeadLetterDestination *EventgridEventSubscriptionStorageBlobDeadLetterDestination `field:"optional" json:"storageBlobDeadLetterDestination" yaml:"storageBlobDeadLetterDestination"`
	// storage_queue_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#storage_queue_endpoint EventgridEventSubscription#storage_queue_endpoint}
	StorageQueueEndpoint *EventgridEventSubscriptionStorageQueueEndpoint `field:"optional" json:"storageQueueEndpoint" yaml:"storageQueueEndpoint"`
	// subject_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#subject_filter EventgridEventSubscription#subject_filter}
	SubjectFilter *EventgridEventSubscriptionSubjectFilter `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#timeouts EventgridEventSubscription#timeouts}
	Timeouts *EventgridEventSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// webhook_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#webhook_endpoint EventgridEventSubscription#webhook_endpoint}
	WebhookEndpoint *EventgridEventSubscriptionWebhookEndpoint `field:"optional" json:"webhookEndpoint" yaml:"webhookEndpoint"`
}

type EventgridEventSubscriptionDeadLetterIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#type EventgridEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#user_assigned_identity EventgridEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

type EventgridEventSubscriptionDeadLetterIdentityOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionDeadLetterIdentity
	SetInternalValue(val *EventgridEventSubscriptionDeadLetterIdentity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserAssignedIdentity() *string
	SetUserAssignedIdentity(val *string)
	UserAssignedIdentityInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetUserAssignedIdentity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionDeadLetterIdentityOutputReference
type jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) InternalValue() *EventgridEventSubscriptionDeadLetterIdentity {
	var returns *EventgridEventSubscriptionDeadLetterIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) UserAssignedIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) UserAssignedIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionDeadLetterIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionDeadLetterIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeadLetterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionDeadLetterIdentityOutputReference_Override(e EventgridEventSubscriptionDeadLetterIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeadLetterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetInternalValue(val *EventgridEventSubscriptionDeadLetterIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) SetUserAssignedIdentity(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentity",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) ResetUserAssignedIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetUserAssignedIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeadLetterIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionDeliveryIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#type EventgridEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#user_assigned_identity EventgridEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

type EventgridEventSubscriptionDeliveryIdentityOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionDeliveryIdentity
	SetInternalValue(val *EventgridEventSubscriptionDeliveryIdentity)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserAssignedIdentity() *string
	SetUserAssignedIdentity(val *string)
	UserAssignedIdentityInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetUserAssignedIdentity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionDeliveryIdentityOutputReference
type jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) InternalValue() *EventgridEventSubscriptionDeliveryIdentity {
	var returns *EventgridEventSubscriptionDeliveryIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) UserAssignedIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) UserAssignedIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionDeliveryIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionDeliveryIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionDeliveryIdentityOutputReference_Override(e EventgridEventSubscriptionDeliveryIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetInternalValue(val *EventgridEventSubscriptionDeliveryIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) SetUserAssignedIdentity(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentity",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) ResetUserAssignedIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetUserAssignedIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionDeliveryProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#header_name EventgridEventSubscription#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#type EventgridEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#secret EventgridEventSubscription#secret}.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#source_field EventgridEventSubscription#source_field}.
	SourceField *string `field:"optional" json:"sourceField" yaml:"sourceField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type EventgridEventSubscriptionDeliveryPropertyList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EventgridEventSubscriptionDeliveryPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionDeliveryPropertyList
type jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionDeliveryPropertyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridEventSubscriptionDeliveryPropertyList {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionDeliveryPropertyList_Override(e EventgridEventSubscriptionDeliveryPropertyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) Get(index *float64) EventgridEventSubscriptionDeliveryPropertyOutputReference {
	var returns EventgridEventSubscriptionDeliveryPropertyOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionDeliveryPropertyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HeaderName() *string
	SetHeaderName(val *string)
	HeaderNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Secret() interface{}
	SetSecret(val interface{})
	SecretInput() interface{}
	SourceField() *string
	SetSourceField(val *string)
	SourceFieldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecret()
	ResetSourceField()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionDeliveryPropertyOutputReference
type jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) HeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) Secret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SourceField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SourceFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionDeliveryPropertyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridEventSubscriptionDeliveryPropertyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionDeliveryPropertyOutputReference_Override(e EventgridEventSubscriptionDeliveryPropertyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionDeliveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetHeaderName(val *string) {
	_jsii_.Set(
		j,
		"headerName",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetSecret(val interface{}) {
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetSourceField(val *string) {
	_jsii_.Set(
		j,
		"sourceField",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ResetSourceField() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceField",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		e,
		"resetValue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionDeliveryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#event_time_to_live EventgridEventSubscription#event_time_to_live}.
	EventTimeToLive *float64 `field:"required" json:"eventTimeToLive" yaml:"eventTimeToLive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#max_delivery_attempts EventgridEventSubscription#max_delivery_attempts}.
	MaxDeliveryAttempts *float64 `field:"required" json:"maxDeliveryAttempts" yaml:"maxDeliveryAttempts"`
}

type EventgridEventSubscriptionRetryPolicyOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventTimeToLive() *float64
	SetEventTimeToLive(val *float64)
	EventTimeToLiveInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionRetryPolicy
	SetInternalValue(val *EventgridEventSubscriptionRetryPolicy)
	MaxDeliveryAttempts() *float64
	SetMaxDeliveryAttempts(val *float64)
	MaxDeliveryAttemptsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionRetryPolicyOutputReference
type jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) EventTimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventTimeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) EventTimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventTimeToLiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) InternalValue() *EventgridEventSubscriptionRetryPolicy {
	var returns *EventgridEventSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) MaxDeliveryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) MaxDeliveryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionRetryPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionRetryPolicyOutputReference_Override(e EventgridEventSubscriptionRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetEventTimeToLive(val *float64) {
	_jsii_.Set(
		j,
		"eventTimeToLive",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetInternalValue(val *EventgridEventSubscriptionRetryPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetMaxDeliveryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxDeliveryAttempts",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionStorageBlobDeadLetterDestination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#storage_account_id EventgridEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#storage_blob_container_name EventgridEventSubscription#storage_blob_container_name}.
	StorageBlobContainerName *string `field:"required" json:"storageBlobContainerName" yaml:"storageBlobContainerName"`
}

type EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	SetInternalValue(val *EventgridEventSubscriptionStorageBlobDeadLetterDestination)
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	StorageBlobContainerName() *string
	SetStorageBlobContainerName(val *string)
	StorageBlobContainerNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
type jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InternalValue() *EventgridEventSubscriptionStorageBlobDeadLetterDestination {
	var returns *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageBlobContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageBlobContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference_Override(e EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetInternalValue(val *EventgridEventSubscriptionStorageBlobDeadLetterDestination) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetStorageBlobContainerName(val *string) {
	_jsii_.Set(
		j,
		"storageBlobContainerName",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionStorageQueueEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#queue_name EventgridEventSubscription#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#storage_account_id EventgridEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#queue_message_time_to_live_in_seconds EventgridEventSubscription#queue_message_time_to_live_in_seconds}.
	QueueMessageTimeToLiveInSeconds *float64 `field:"optional" json:"queueMessageTimeToLiveInSeconds" yaml:"queueMessageTimeToLiveInSeconds"`
}

type EventgridEventSubscriptionStorageQueueEndpointOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionStorageQueueEndpoint
	SetInternalValue(val *EventgridEventSubscriptionStorageQueueEndpoint)
	QueueMessageTimeToLiveInSeconds() *float64
	SetQueueMessageTimeToLiveInSeconds(val *float64)
	QueueMessageTimeToLiveInSecondsInput() *float64
	QueueName() *string
	SetQueueName(val *string)
	QueueNameInput() *string
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetQueueMessageTimeToLiveInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionStorageQueueEndpointOutputReference
type jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) InternalValue() *EventgridEventSubscriptionStorageQueueEndpoint {
	var returns *EventgridEventSubscriptionStorageQueueEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) QueueMessageTimeToLiveInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueMessageTimeToLiveInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) QueueMessageTimeToLiveInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueMessageTimeToLiveInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) QueueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionStorageQueueEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionStorageQueueEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionStorageQueueEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionStorageQueueEndpointOutputReference_Override(e EventgridEventSubscriptionStorageQueueEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionStorageQueueEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetInternalValue(val *EventgridEventSubscriptionStorageQueueEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetQueueMessageTimeToLiveInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"queueMessageTimeToLiveInSeconds",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetQueueName(val *string) {
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) ResetQueueMessageTimeToLiveInSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetQueueMessageTimeToLiveInSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionStorageQueueEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionSubjectFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#case_sensitive EventgridEventSubscription#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#subject_begins_with EventgridEventSubscription#subject_begins_with}.
	SubjectBeginsWith *string `field:"optional" json:"subjectBeginsWith" yaml:"subjectBeginsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#subject_ends_with EventgridEventSubscription#subject_ends_with}.
	SubjectEndsWith *string `field:"optional" json:"subjectEndsWith" yaml:"subjectEndsWith"`
}

type EventgridEventSubscriptionSubjectFilterOutputReference interface {
	cdktf.ComplexObject
	CaseSensitive() interface{}
	SetCaseSensitive(val interface{})
	CaseSensitiveInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionSubjectFilter
	SetInternalValue(val *EventgridEventSubscriptionSubjectFilter)
	SubjectBeginsWith() *string
	SetSubjectBeginsWith(val *string)
	SubjectBeginsWithInput() *string
	SubjectEndsWith() *string
	SetSubjectEndsWith(val *string)
	SubjectEndsWithInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCaseSensitive()
	ResetSubjectBeginsWith()
	ResetSubjectEndsWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionSubjectFilterOutputReference
type jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) CaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) CaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) InternalValue() *EventgridEventSubscriptionSubjectFilter {
	var returns *EventgridEventSubscriptionSubjectFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SubjectBeginsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SubjectBeginsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SubjectEndsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SubjectEndsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionSubjectFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionSubjectFilterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionSubjectFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionSubjectFilterOutputReference_Override(e EventgridEventSubscriptionSubjectFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionSubjectFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetCaseSensitive(val interface{}) {
	_jsii_.Set(
		j,
		"caseSensitive",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetInternalValue(val *EventgridEventSubscriptionSubjectFilter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetSubjectBeginsWith(val *string) {
	_jsii_.Set(
		j,
		"subjectBeginsWith",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetSubjectEndsWith(val *string) {
	_jsii_.Set(
		j,
		"subjectEndsWith",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ResetCaseSensitive() {
	_jsii_.InvokeVoid(
		e,
		"resetCaseSensitive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ResetSubjectBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ResetSubjectEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionSubjectFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#create EventgridEventSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#delete EventgridEventSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#read EventgridEventSubscription#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#update EventgridEventSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type EventgridEventSubscriptionTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionTimeoutsOutputReference
type jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionTimeoutsOutputReference_Override(e EventgridEventSubscriptionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		e,
		"resetCreate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		e,
		"resetDelete",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		e,
		"resetRead",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridEventSubscriptionWebhookEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#url EventgridEventSubscription#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#active_directory_app_id_or_uri EventgridEventSubscription#active_directory_app_id_or_uri}.
	ActiveDirectoryAppIdOrUri *string `field:"optional" json:"activeDirectoryAppIdOrUri" yaml:"activeDirectoryAppIdOrUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#active_directory_tenant_id EventgridEventSubscription#active_directory_tenant_id}.
	ActiveDirectoryTenantId *string `field:"optional" json:"activeDirectoryTenantId" yaml:"activeDirectoryTenantId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#max_events_per_batch EventgridEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#preferred_batch_size_in_kilobytes EventgridEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

type EventgridEventSubscriptionWebhookEndpointOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryAppIdOrUri() *string
	SetActiveDirectoryAppIdOrUri(val *string)
	ActiveDirectoryAppIdOrUriInput() *string
	ActiveDirectoryTenantId() *string
	SetActiveDirectoryTenantId(val *string)
	ActiveDirectoryTenantIdInput() *string
	BaseUrl() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *EventgridEventSubscriptionWebhookEndpoint
	SetInternalValue(val *EventgridEventSubscriptionWebhookEndpoint)
	MaxEventsPerBatch() *float64
	SetMaxEventsPerBatch(val *float64)
	MaxEventsPerBatchInput() *float64
	PreferredBatchSizeInKilobytes() *float64
	SetPreferredBatchSizeInKilobytes(val *float64)
	PreferredBatchSizeInKilobytesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetActiveDirectoryAppIdOrUri()
	ResetActiveDirectoryTenantId()
	ResetMaxEventsPerBatch()
	ResetPreferredBatchSizeInKilobytes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionWebhookEndpointOutputReference
type jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InternalValue() *EventgridEventSubscriptionWebhookEndpoint {
	var returns *EventgridEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionWebhookEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionWebhookEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionWebhookEndpointOutputReference_Override(e EventgridEventSubscriptionWebhookEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetActiveDirectoryAppIdOrUri(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryAppIdOrUri",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetActiveDirectoryTenantId(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryTenantId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetInternalValue(val *EventgridEventSubscriptionWebhookEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetMaxEventsPerBatch(val *float64) {
	_jsii_.Set(
		j,
		"maxEventsPerBatch",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetPreferredBatchSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"preferredBatchSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryAppIdOrUri() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryAppIdOrUri",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetMaxEventsPerBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxEventsPerBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ResetPreferredBatchSizeInKilobytes() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredBatchSizeInKilobytes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionWebhookEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

