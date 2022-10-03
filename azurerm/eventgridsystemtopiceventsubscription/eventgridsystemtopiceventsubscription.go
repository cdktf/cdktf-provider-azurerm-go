package eventgridsystemtopiceventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/eventgridsystemtopiceventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription} Resource.
func NewEventgridSystemTopicEventSubscription(scope constructs.Construct, id *string, config *EventgridSystemTopicEventSubscriptionConfig) EventgridSystemTopicEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription azurerm_eventgrid_system_topic_event_subscription} Resource.
func NewEventgridSystemTopicEventSubscription_Override(e EventgridSystemTopicEventSubscription, scope constructs.Construct, id *string, config *EventgridSystemTopicEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetAdvancedFilteringOnArraysEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"advancedFilteringOnArraysEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetEventDeliverySchema(val *string) {
	_jsii_.Set(
		j,
		"eventDeliverySchema",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetEventhubEndpointId(val *string) {
	_jsii_.Set(
		j,
		"eventhubEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetExpirationTimeUtc(val *string) {
	_jsii_.Set(
		j,
		"expirationTimeUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetHybridConnectionEndpointId(val *string) {
	_jsii_.Set(
		j,
		"hybridConnectionEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetIncludedEventTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"includedEventTypes",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetLabels(val *[]*string) {
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetServiceBusQueueEndpointId(val *string) {
	_jsii_.Set(
		j,
		"serviceBusQueueEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetServiceBusTopicEndpointId(val *string) {
	_jsii_.Set(
		j,
		"serviceBusTopicEndpointId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscription) SetSystemTopic(val *string) {
	_jsii_.Set(
		j,
		"systemTopic",
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
func EventgridSystemTopicEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscription",
		"isConstruct",
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutAdvancedFilter(value *EventgridSystemTopicEventSubscriptionAdvancedFilter) {
	_jsii_.InvokeVoid(
		e,
		"putAdvancedFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutAzureFunctionEndpoint(value *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint) {
	_jsii_.InvokeVoid(
		e,
		"putAzureFunctionEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeadLetterIdentity(value *EventgridSystemTopicEventSubscriptionDeadLetterIdentity) {
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeliveryIdentity(value *EventgridSystemTopicEventSubscriptionDeliveryIdentity) {
	_jsii_.InvokeVoid(
		e,
		"putDeliveryIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutDeliveryProperty(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putDeliveryProperty",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutRetryPolicy(value *EventgridSystemTopicEventSubscriptionRetryPolicy) {
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutStorageBlobDeadLetterDestination(value *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination) {
	_jsii_.InvokeVoid(
		e,
		"putStorageBlobDeadLetterDestination",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutStorageQueueEndpoint(value *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint) {
	_jsii_.InvokeVoid(
		e,
		"putStorageQueueEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutSubjectFilter(value *EventgridSystemTopicEventSubscriptionSubjectFilter) {
	_jsii_.InvokeVoid(
		e,
		"putSubjectFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutTimeouts(value *EventgridSystemTopicEventSubscriptionTimeouts) {
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscription) PutWebhookEndpoint(value *EventgridSystemTopicEventSubscriptionWebhookEndpoint) {
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

type EventgridSystemTopicEventSubscriptionAdvancedFilter struct {
	// bool_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#bool_equals EventgridSystemTopicEventSubscription#bool_equals}
	BoolEquals interface{} `field:"optional" json:"boolEquals" yaml:"boolEquals"`
	// is_not_null block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#is_not_null EventgridSystemTopicEventSubscription#is_not_null}
	IsNotNull interface{} `field:"optional" json:"isNotNull" yaml:"isNotNull"`
	// is_null_or_undefined block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#is_null_or_undefined EventgridSystemTopicEventSubscription#is_null_or_undefined}
	IsNullOrUndefined interface{} `field:"optional" json:"isNullOrUndefined" yaml:"isNullOrUndefined"`
	// number_greater_than block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_greater_than EventgridSystemTopicEventSubscription#number_greater_than}
	NumberGreaterThan interface{} `field:"optional" json:"numberGreaterThan" yaml:"numberGreaterThan"`
	// number_greater_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_greater_than_or_equals EventgridSystemTopicEventSubscription#number_greater_than_or_equals}
	NumberGreaterThanOrEquals interface{} `field:"optional" json:"numberGreaterThanOrEquals" yaml:"numberGreaterThanOrEquals"`
	// number_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_in EventgridSystemTopicEventSubscription#number_in}
	NumberIn interface{} `field:"optional" json:"numberIn" yaml:"numberIn"`
	// number_in_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_in_range EventgridSystemTopicEventSubscription#number_in_range}
	NumberInRange interface{} `field:"optional" json:"numberInRange" yaml:"numberInRange"`
	// number_less_than block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_less_than EventgridSystemTopicEventSubscription#number_less_than}
	NumberLessThan interface{} `field:"optional" json:"numberLessThan" yaml:"numberLessThan"`
	// number_less_than_or_equals block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_less_than_or_equals EventgridSystemTopicEventSubscription#number_less_than_or_equals}
	NumberLessThanOrEquals interface{} `field:"optional" json:"numberLessThanOrEquals" yaml:"numberLessThanOrEquals"`
	// number_not_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_not_in EventgridSystemTopicEventSubscription#number_not_in}
	NumberNotIn interface{} `field:"optional" json:"numberNotIn" yaml:"numberNotIn"`
	// number_not_in_range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#number_not_in_range EventgridSystemTopicEventSubscription#number_not_in_range}
	NumberNotInRange interface{} `field:"optional" json:"numberNotInRange" yaml:"numberNotInRange"`
	// string_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_begins_with EventgridSystemTopicEventSubscription#string_begins_with}
	StringBeginsWith interface{} `field:"optional" json:"stringBeginsWith" yaml:"stringBeginsWith"`
	// string_contains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_contains EventgridSystemTopicEventSubscription#string_contains}
	StringContains interface{} `field:"optional" json:"stringContains" yaml:"stringContains"`
	// string_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_ends_with EventgridSystemTopicEventSubscription#string_ends_with}
	StringEndsWith interface{} `field:"optional" json:"stringEndsWith" yaml:"stringEndsWith"`
	// string_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_in EventgridSystemTopicEventSubscription#string_in}
	StringIn interface{} `field:"optional" json:"stringIn" yaml:"stringIn"`
	// string_not_begins_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_not_begins_with EventgridSystemTopicEventSubscription#string_not_begins_with}
	StringNotBeginsWith interface{} `field:"optional" json:"stringNotBeginsWith" yaml:"stringNotBeginsWith"`
	// string_not_contains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_not_contains EventgridSystemTopicEventSubscription#string_not_contains}
	StringNotContains interface{} `field:"optional" json:"stringNotContains" yaml:"stringNotContains"`
	// string_not_ends_with block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_not_ends_with EventgridSystemTopicEventSubscription#string_not_ends_with}
	StringNotEndsWith interface{} `field:"optional" json:"stringNotEndsWith" yaml:"stringNotEndsWith"`
	// string_not_in block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#string_not_in EventgridSystemTopicEventSubscription#string_not_in}
	StringNotIn interface{} `field:"optional" json:"stringNotIn" yaml:"stringNotIn"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) SetValue(val interface{}) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNull struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefined struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRange struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*float64 `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Values() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ValuesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) SetValues(val *[]*float64) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference interface {
	cdktf.ComplexObject
	BoolEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
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
	InternalValue() *EventgridSystemTopicEventSubscriptionAdvancedFilter
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionAdvancedFilter)
	IsNotNull() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList
	IsNotNullInput() interface{}
	IsNullOrUndefined() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	IsNullOrUndefinedInput() interface{}
	NumberGreaterThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList
	NumberGreaterThanInput() interface{}
	NumberGreaterThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	NumberGreaterThanOrEqualsInput() interface{}
	NumberIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList
	NumberInInput() interface{}
	NumberInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList
	NumberInRangeInput() interface{}
	NumberLessThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList
	NumberLessThanInput() interface{}
	NumberLessThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	NumberLessThanOrEqualsInput() interface{}
	NumberNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList
	NumberNotInInput() interface{}
	NumberNotInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList
	NumberNotInRangeInput() interface{}
	StringBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList
	StringBeginsWithInput() interface{}
	StringContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList
	StringContainsInput() interface{}
	StringEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList
	StringEndsWithInput() interface{}
	StringIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList
	StringInInput() interface{}
	StringNotBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList
	StringNotBeginsWithInput() interface{}
	StringNotContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList
	StringNotContainsInput() interface{}
	StringNotEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList
	StringNotEndsWithInput() interface{}
	StringNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) BoolEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
	_jsii_.Get(
		j,
		"boolEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) BoolEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionAdvancedFilter {
	var returns *EventgridSystemTopicEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNotNull() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList
	_jsii_.Get(
		j,
		"isNotNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNotNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNotNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefined() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	_jsii_.Get(
		j,
		"isNullOrUndefined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefinedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullOrUndefinedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList
	_jsii_.Get(
		j,
		"numberGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	_jsii_.Get(
		j,
		"numberGreaterThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList
	_jsii_.Get(
		j,
		"numberIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList
	_jsii_.Get(
		j,
		"numberInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList
	_jsii_.Get(
		j,
		"numberLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	_jsii_.Get(
		j,
		"numberLessThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList
	_jsii_.Get(
		j,
		"numberNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList
	_jsii_.Get(
		j,
		"numberNotInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList
	_jsii_.Get(
		j,
		"stringBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList
	_jsii_.Get(
		j,
		"stringContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList
	_jsii_.Get(
		j,
		"stringEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList
	_jsii_.Get(
		j,
		"stringIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList
	_jsii_.Get(
		j,
		"stringNotBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList
	_jsii_.Get(
		j,
		"stringNotContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList
	_jsii_.Get(
		j,
		"stringNotEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList
	_jsii_.Get(
		j,
		"stringNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionAdvancedFilter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutBoolEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putBoolEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutIsNotNull(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putIsNotNull",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutIsNullOrUndefined(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putIsNullOrUndefined",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThan(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThanOrEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberInRange(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThan(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThanOrEquals(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberNotIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberNotInRange(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putNumberNotInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringBeginsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringContains(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringEndsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotBeginsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotContains(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotEndsWith(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotIn(value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"putStringNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetBoolEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetBoolEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetIsNotNull() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNotNull",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetIsNullOrUndefined() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNullOrUndefined",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringInList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWith struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) Get(index *float64) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#function_id EventgridSystemTopicEventSubscription#function_id}.
	FunctionId *string `field:"required" json:"functionId" yaml:"functionId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#max_events_per_batch EventgridSystemTopicEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#preferred_batch_size_in_kilobytes EventgridSystemTopicEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

type EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) FunctionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) MaxEventsPerBatch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) MaxEventsPerBatchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) PreferredBatchSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) PreferredBatchSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference_Override(e EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetFunctionId(val *string) {
	_jsii_.Set(
		j,
		"functionId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetMaxEventsPerBatch(val *float64) {
	_jsii_.Set(
		j,
		"maxEventsPerBatch",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetPreferredBatchSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"preferredBatchSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ResetMaxEventsPerBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxEventsPerBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ResetPreferredBatchSizeInKilobytes() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredBatchSizeInKilobytes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAzureFunctionEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#name EventgridSystemTopicEventSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#resource_group_name EventgridSystemTopicEventSubscription#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#system_topic EventgridSystemTopicEventSubscription#system_topic}.
	SystemTopic *string `field:"required" json:"systemTopic" yaml:"systemTopic"`
	// advanced_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#advanced_filter EventgridSystemTopicEventSubscription#advanced_filter}
	AdvancedFilter *EventgridSystemTopicEventSubscriptionAdvancedFilter `field:"optional" json:"advancedFilter" yaml:"advancedFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#advanced_filtering_on_arrays_enabled EventgridSystemTopicEventSubscription#advanced_filtering_on_arrays_enabled}.
	AdvancedFilteringOnArraysEnabled interface{} `field:"optional" json:"advancedFilteringOnArraysEnabled" yaml:"advancedFilteringOnArraysEnabled"`
	// azure_function_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#azure_function_endpoint EventgridSystemTopicEventSubscription#azure_function_endpoint}
	AzureFunctionEndpoint *EventgridSystemTopicEventSubscriptionAzureFunctionEndpoint `field:"optional" json:"azureFunctionEndpoint" yaml:"azureFunctionEndpoint"`
	// dead_letter_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#dead_letter_identity EventgridSystemTopicEventSubscription#dead_letter_identity}
	DeadLetterIdentity *EventgridSystemTopicEventSubscriptionDeadLetterIdentity `field:"optional" json:"deadLetterIdentity" yaml:"deadLetterIdentity"`
	// delivery_identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#delivery_identity EventgridSystemTopicEventSubscription#delivery_identity}
	DeliveryIdentity *EventgridSystemTopicEventSubscriptionDeliveryIdentity `field:"optional" json:"deliveryIdentity" yaml:"deliveryIdentity"`
	// delivery_property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#delivery_property EventgridSystemTopicEventSubscription#delivery_property}
	DeliveryProperty interface{} `field:"optional" json:"deliveryProperty" yaml:"deliveryProperty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#event_delivery_schema EventgridSystemTopicEventSubscription#event_delivery_schema}.
	EventDeliverySchema *string `field:"optional" json:"eventDeliverySchema" yaml:"eventDeliverySchema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#eventhub_endpoint_id EventgridSystemTopicEventSubscription#eventhub_endpoint_id}.
	EventhubEndpointId *string `field:"optional" json:"eventhubEndpointId" yaml:"eventhubEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#expiration_time_utc EventgridSystemTopicEventSubscription#expiration_time_utc}.
	ExpirationTimeUtc *string `field:"optional" json:"expirationTimeUtc" yaml:"expirationTimeUtc"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#hybrid_connection_endpoint_id EventgridSystemTopicEventSubscription#hybrid_connection_endpoint_id}.
	HybridConnectionEndpointId *string `field:"optional" json:"hybridConnectionEndpointId" yaml:"hybridConnectionEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#id EventgridSystemTopicEventSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#included_event_types EventgridSystemTopicEventSubscription#included_event_types}.
	IncludedEventTypes *[]*string `field:"optional" json:"includedEventTypes" yaml:"includedEventTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#labels EventgridSystemTopicEventSubscription#labels}.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#retry_policy EventgridSystemTopicEventSubscription#retry_policy}
	RetryPolicy *EventgridSystemTopicEventSubscriptionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#service_bus_queue_endpoint_id EventgridSystemTopicEventSubscription#service_bus_queue_endpoint_id}.
	ServiceBusQueueEndpointId *string `field:"optional" json:"serviceBusQueueEndpointId" yaml:"serviceBusQueueEndpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#service_bus_topic_endpoint_id EventgridSystemTopicEventSubscription#service_bus_topic_endpoint_id}.
	ServiceBusTopicEndpointId *string `field:"optional" json:"serviceBusTopicEndpointId" yaml:"serviceBusTopicEndpointId"`
	// storage_blob_dead_letter_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#storage_blob_dead_letter_destination EventgridSystemTopicEventSubscription#storage_blob_dead_letter_destination}
	StorageBlobDeadLetterDestination *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination `field:"optional" json:"storageBlobDeadLetterDestination" yaml:"storageBlobDeadLetterDestination"`
	// storage_queue_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#storage_queue_endpoint EventgridSystemTopicEventSubscription#storage_queue_endpoint}
	StorageQueueEndpoint *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint `field:"optional" json:"storageQueueEndpoint" yaml:"storageQueueEndpoint"`
	// subject_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#subject_filter EventgridSystemTopicEventSubscription#subject_filter}
	SubjectFilter *EventgridSystemTopicEventSubscriptionSubjectFilter `field:"optional" json:"subjectFilter" yaml:"subjectFilter"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#timeouts EventgridSystemTopicEventSubscription#timeouts}
	Timeouts *EventgridSystemTopicEventSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// webhook_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#webhook_endpoint EventgridSystemTopicEventSubscription#webhook_endpoint}
	WebhookEndpoint *EventgridSystemTopicEventSubscriptionWebhookEndpoint `field:"optional" json:"webhookEndpoint" yaml:"webhookEndpoint"`
}

type EventgridSystemTopicEventSubscriptionDeadLetterIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#type EventgridSystemTopicEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#user_assigned_identity EventgridSystemTopicEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

type EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionDeadLetterIdentity
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionDeadLetterIdentity)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionDeadLetterIdentity {
	var returns *EventgridSystemTopicEventSubscriptionDeadLetterIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) UserAssignedIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) UserAssignedIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference_Override(e EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionDeadLetterIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) SetUserAssignedIdentity(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentity",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) ResetUserAssignedIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetUserAssignedIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeadLetterIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionDeliveryIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#type EventgridSystemTopicEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#user_assigned_identity EventgridSystemTopicEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

type EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionDeliveryIdentity
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionDeliveryIdentity)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionDeliveryIdentity {
	var returns *EventgridSystemTopicEventSubscriptionDeliveryIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) UserAssignedIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) UserAssignedIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAssignedIdentityInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference_Override(e EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionDeliveryIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) SetUserAssignedIdentity(val *string) {
	_jsii_.Set(
		j,
		"userAssignedIdentity",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) ResetUserAssignedIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetUserAssignedIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionDeliveryProperty struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#header_name EventgridSystemTopicEventSubscription#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#type EventgridSystemTopicEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#secret EventgridSystemTopicEventSubscription#secret}.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#source_field EventgridSystemTopicEventSubscription#source_field}.
	SourceField *string `field:"optional" json:"sourceField" yaml:"sourceField"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type EventgridSystemTopicEventSubscriptionDeliveryPropertyList interface {
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
	Get(index *float64) EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionDeliveryPropertyList
type jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionDeliveryPropertyList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EventgridSystemTopicEventSubscriptionDeliveryPropertyList {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionDeliveryPropertyList_Override(e EventgridSystemTopicEventSubscriptionDeliveryPropertyList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) Get(index *float64) EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference {
	var returns EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) HeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) HeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) Secret() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SourceField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SourceFieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference_Override(e EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetHeaderName(val *string) {
	_jsii_.Set(
		j,
		"headerName",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetSecret(val interface{}) {
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetSourceField(val *string) {
	_jsii_.Set(
		j,
		"sourceField",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		e,
		"resetSecret",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ResetSourceField() {
	_jsii_.InvokeVoid(
		e,
		"resetSourceField",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		e,
		"resetValue",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionDeliveryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionRetryPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#event_time_to_live EventgridSystemTopicEventSubscription#event_time_to_live}.
	EventTimeToLive *float64 `field:"required" json:"eventTimeToLive" yaml:"eventTimeToLive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#max_delivery_attempts EventgridSystemTopicEventSubscription#max_delivery_attempts}.
	MaxDeliveryAttempts *float64 `field:"required" json:"maxDeliveryAttempts" yaml:"maxDeliveryAttempts"`
}

type EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionRetryPolicy
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionRetryPolicy)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) EventTimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventTimeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) EventTimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventTimeToLiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionRetryPolicy {
	var returns *EventgridSystemTopicEventSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) MaxDeliveryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) MaxDeliveryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionRetryPolicyOutputReference_Override(e EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetEventTimeToLive(val *float64) {
	_jsii_.Set(
		j,
		"eventTimeToLive",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionRetryPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetMaxDeliveryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maxDeliveryAttempts",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#storage_account_id EventgridSystemTopicEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#storage_blob_container_name EventgridSystemTopicEventSubscription#storage_blob_container_name}.
	StorageBlobContainerName *string `field:"required" json:"storageBlobContainerName" yaml:"storageBlobContainerName"`
}

type EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination {
	var returns *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageBlobContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) StorageBlobContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageBlobContainerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference_Override(e EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestination) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetStorageBlobContainerName(val *string) {
	_jsii_.Set(
		j,
		"storageBlobContainerName",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageBlobDeadLetterDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionStorageQueueEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#queue_name EventgridSystemTopicEventSubscription#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#storage_account_id EventgridSystemTopicEventSubscription#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#queue_message_time_to_live_in_seconds EventgridSystemTopicEventSubscription#queue_message_time_to_live_in_seconds}.
	QueueMessageTimeToLiveInSeconds *float64 `field:"optional" json:"queueMessageTimeToLiveInSeconds" yaml:"queueMessageTimeToLiveInSeconds"`
}

type EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) QueueMessageTimeToLiveInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueMessageTimeToLiveInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) QueueMessageTimeToLiveInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueMessageTimeToLiveInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) QueueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference_Override(e EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionStorageQueueEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetQueueMessageTimeToLiveInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"queueMessageTimeToLiveInSeconds",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetQueueName(val *string) {
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetStorageAccountId(val *string) {
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) ResetQueueMessageTimeToLiveInSeconds() {
	_jsii_.InvokeVoid(
		e,
		"resetQueueMessageTimeToLiveInSeconds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionStorageQueueEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionSubjectFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#case_sensitive EventgridSystemTopicEventSubscription#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#subject_begins_with EventgridSystemTopicEventSubscription#subject_begins_with}.
	SubjectBeginsWith *string `field:"optional" json:"subjectBeginsWith" yaml:"subjectBeginsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#subject_ends_with EventgridSystemTopicEventSubscription#subject_ends_with}.
	SubjectEndsWith *string `field:"optional" json:"subjectEndsWith" yaml:"subjectEndsWith"`
}

type EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionSubjectFilter
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionSubjectFilter)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) CaseSensitive() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) CaseSensitiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionSubjectFilter {
	var returns *EventgridSystemTopicEventSubscriptionSubjectFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SubjectBeginsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SubjectBeginsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SubjectEndsWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SubjectEndsWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionSubjectFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionSubjectFilterOutputReference_Override(e EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetCaseSensitive(val interface{}) {
	_jsii_.Set(
		j,
		"caseSensitive",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionSubjectFilter) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetSubjectBeginsWith(val *string) {
	_jsii_.Set(
		j,
		"subjectBeginsWith",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetSubjectEndsWith(val *string) {
	_jsii_.Set(
		j,
		"subjectEndsWith",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ResetCaseSensitive() {
	_jsii_.InvokeVoid(
		e,
		"resetCaseSensitive",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ResetSubjectBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ResetSubjectEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionSubjectFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#create EventgridSystemTopicEventSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#delete EventgridSystemTopicEventSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#read EventgridSystemTopicEventSubscription#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#update EventgridSystemTopicEventSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type EventgridSystemTopicEventSubscriptionTimeoutsOutputReference interface {
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionTimeoutsOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionTimeoutsOutputReference_Override(e EventgridSystemTopicEventSubscriptionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		e,
		"resetCreate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		e,
		"resetDelete",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		e,
		"resetRead",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EventgridSystemTopicEventSubscriptionWebhookEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#url EventgridSystemTopicEventSubscription#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#active_directory_app_id_or_uri EventgridSystemTopicEventSubscription#active_directory_app_id_or_uri}.
	ActiveDirectoryAppIdOrUri *string `field:"optional" json:"activeDirectoryAppIdOrUri" yaml:"activeDirectoryAppIdOrUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#active_directory_tenant_id EventgridSystemTopicEventSubscription#active_directory_tenant_id}.
	ActiveDirectoryTenantId *string `field:"optional" json:"activeDirectoryTenantId" yaml:"activeDirectoryTenantId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#max_events_per_batch EventgridSystemTopicEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#preferred_batch_size_in_kilobytes EventgridSystemTopicEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

type EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference interface {
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
	InternalValue() *EventgridSystemTopicEventSubscriptionWebhookEndpoint
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionWebhookEndpoint)
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryAppIdOrUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryAppIdOrUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ActiveDirectoryTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionWebhookEndpoint {
	var returns *EventgridSystemTopicEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) MaxEventsPerBatchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEventsPerBatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) PreferredBatchSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preferredBatchSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference_Override(e EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetActiveDirectoryAppIdOrUri(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryAppIdOrUri",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetActiveDirectoryTenantId(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryTenantId",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetInternalValue(val *EventgridSystemTopicEventSubscriptionWebhookEndpoint) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetMaxEventsPerBatch(val *float64) {
	_jsii_.Set(
		j,
		"maxEventsPerBatch",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetPreferredBatchSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"preferredBatchSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryAppIdOrUri() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryAppIdOrUri",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ResetActiveDirectoryTenantId() {
	_jsii_.InvokeVoid(
		e,
		"resetActiveDirectoryTenantId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ResetMaxEventsPerBatch() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxEventsPerBatch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ResetPreferredBatchSizeInKilobytes() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredBatchSizeInKilobytes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionWebhookEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

