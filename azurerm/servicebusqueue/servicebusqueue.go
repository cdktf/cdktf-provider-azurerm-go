package servicebusqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicebusqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue azurerm_servicebus_queue}.
type ServicebusQueue interface {
	cdktf.TerraformResource
	AutoDeleteOnIdle() *string
	SetAutoDeleteOnIdle(val *string)
	AutoDeleteOnIdleInput() *string
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
	DeadLetteringOnMessageExpiration() interface{}
	SetDeadLetteringOnMessageExpiration(val interface{})
	DeadLetteringOnMessageExpirationInput() interface{}
	DefaultMessageTtl() *string
	SetDefaultMessageTtl(val *string)
	DefaultMessageTtlInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DuplicateDetectionHistoryTimeWindow() *string
	SetDuplicateDetectionHistoryTimeWindow(val *string)
	DuplicateDetectionHistoryTimeWindowInput() *string
	EnableBatchedOperations() interface{}
	SetEnableBatchedOperations(val interface{})
	EnableBatchedOperationsInput() interface{}
	EnableExpress() interface{}
	SetEnableExpress(val interface{})
	EnableExpressInput() interface{}
	EnablePartitioning() interface{}
	SetEnablePartitioning(val interface{})
	EnablePartitioningInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	ForwardDeadLetteredMessagesTo() *string
	SetForwardDeadLetteredMessagesTo(val *string)
	ForwardDeadLetteredMessagesToInput() *string
	ForwardTo() *string
	SetForwardTo(val *string)
	ForwardToInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LockDuration() *string
	SetLockDuration(val *string)
	LockDurationInput() *string
	MaxDeliveryCount() *float64
	SetMaxDeliveryCount(val *float64)
	MaxDeliveryCountInput() *float64
	MaxMessageSizeInKilobytes() *float64
	SetMaxMessageSizeInKilobytes(val *float64)
	MaxMessageSizeInKilobytesInput() *float64
	MaxSizeInMegabytes() *float64
	SetMaxSizeInMegabytes(val *float64)
	MaxSizeInMegabytesInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceId() *string
	SetNamespaceId(val *string)
	NamespaceIdInput() *string
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
	RequiresDuplicateDetection() interface{}
	SetRequiresDuplicateDetection(val interface{})
	RequiresDuplicateDetectionInput() interface{}
	RequiresSession() interface{}
	SetRequiresSession(val interface{})
	RequiresSessionInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServicebusQueueTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *ServicebusQueueTimeouts)
	ResetAutoDeleteOnIdle()
	ResetDeadLetteringOnMessageExpiration()
	ResetDefaultMessageTtl()
	ResetDuplicateDetectionHistoryTimeWindow()
	ResetEnableBatchedOperations()
	ResetEnableExpress()
	ResetEnablePartitioning()
	ResetForwardDeadLetteredMessagesTo()
	ResetForwardTo()
	ResetId()
	ResetLockDuration()
	ResetMaxDeliveryCount()
	ResetMaxMessageSizeInKilobytes()
	ResetMaxSizeInMegabytes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequiresDuplicateDetection()
	ResetRequiresSession()
	ResetStatus()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServicebusQueue
type jsiiProxy_ServicebusQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicebusQueue) AutoDeleteOnIdle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) AutoDeleteOnIdleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DeadLetteringOnMessageExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnMessageExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DeadLetteringOnMessageExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnMessageExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DefaultMessageTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DefaultMessageTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DuplicateDetectionHistoryTimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duplicateDetectionHistoryTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) DuplicateDetectionHistoryTimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duplicateDetectionHistoryTimeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnableBatchedOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnableBatchedOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnableExpress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExpress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnableExpressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExpressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnablePartitioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) EnablePartitioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ForwardDeadLetteredMessagesTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardDeadLetteredMessagesTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ForwardDeadLetteredMessagesToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardDeadLetteredMessagesToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ForwardTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ForwardToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) LockDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) LockDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxDeliveryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxDeliveryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxMessageSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxMessageSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxSizeInMegabytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) MaxSizeInMegabytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) NamespaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) RequiresDuplicateDetection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresDuplicateDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) RequiresDuplicateDetectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresDuplicateDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) RequiresSession() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) RequiresSessionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) Timeouts() ServicebusQueueTimeoutsOutputReference {
	var returns ServicebusQueueTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue azurerm_servicebus_queue} Resource.
func NewServicebusQueue(scope constructs.Construct, id *string, config *ServicebusQueueConfig) ServicebusQueue {
	_init_.Initialize()

	j := jsiiProxy_ServicebusQueue{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue azurerm_servicebus_queue} Resource.
func NewServicebusQueue_Override(s ServicebusQueue, scope constructs.Construct, id *string, config *ServicebusQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetAutoDeleteOnIdle(val *string) {
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetDeadLetteringOnMessageExpiration(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetteringOnMessageExpiration",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetDefaultMessageTtl(val *string) {
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetDuplicateDetectionHistoryTimeWindow(val *string) {
	_jsii_.Set(
		j,
		"duplicateDetectionHistoryTimeWindow",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetEnableBatchedOperations(val interface{}) {
	_jsii_.Set(
		j,
		"enableBatchedOperations",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetEnableExpress(val interface{}) {
	_jsii_.Set(
		j,
		"enableExpress",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetEnablePartitioning(val interface{}) {
	_jsii_.Set(
		j,
		"enablePartitioning",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetForwardDeadLetteredMessagesTo(val *string) {
	_jsii_.Set(
		j,
		"forwardDeadLetteredMessagesTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetForwardTo(val *string) {
	_jsii_.Set(
		j,
		"forwardTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetLockDuration(val *string) {
	_jsii_.Set(
		j,
		"lockDuration",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetMaxDeliveryCount(val *float64) {
	_jsii_.Set(
		j,
		"maxDeliveryCount",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetMaxMessageSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"maxMessageSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetMaxSizeInMegabytes(val *float64) {
	_jsii_.Set(
		j,
		"maxSizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetNamespaceId(val *string) {
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetRequiresDuplicateDetection(val interface{}) {
	_jsii_.Set(
		j,
		"requiresDuplicateDetection",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetRequiresSession(val interface{}) {
	_jsii_.Set(
		j,
		"requiresSession",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
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
func ServicebusQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicebusQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicebusQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusQueue) PutTimeouts(value *ServicebusQueueTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetAutoDeleteOnIdle() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoDeleteOnIdle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetDeadLetteringOnMessageExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetDeadLetteringOnMessageExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetDefaultMessageTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultMessageTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetDuplicateDetectionHistoryTimeWindow() {
	_jsii_.InvokeVoid(
		s,
		"resetDuplicateDetectionHistoryTimeWindow",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetEnableBatchedOperations() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableBatchedOperations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetEnableExpress() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableExpress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetEnablePartitioning() {
	_jsii_.InvokeVoid(
		s,
		"resetEnablePartitioning",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetForwardDeadLetteredMessagesTo() {
	_jsii_.InvokeVoid(
		s,
		"resetForwardDeadLetteredMessagesTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetForwardTo() {
	_jsii_.InvokeVoid(
		s,
		"resetForwardTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetLockDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLockDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetMaxDeliveryCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDeliveryCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetMaxMessageSizeInKilobytes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxMessageSizeInKilobytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetMaxSizeInMegabytes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSizeInMegabytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetRequiresDuplicateDetection() {
	_jsii_.InvokeVoid(
		s,
		"resetRequiresDuplicateDetection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetRequiresSession() {
	_jsii_.InvokeVoid(
		s,
		"resetRequiresSession",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServicebusQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#name ServicebusQueue#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#namespace_id ServicebusQueue#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#auto_delete_on_idle ServicebusQueue#auto_delete_on_idle}.
	AutoDeleteOnIdle *string `field:"optional" json:"autoDeleteOnIdle" yaml:"autoDeleteOnIdle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#dead_lettering_on_message_expiration ServicebusQueue#dead_lettering_on_message_expiration}.
	DeadLetteringOnMessageExpiration interface{} `field:"optional" json:"deadLetteringOnMessageExpiration" yaml:"deadLetteringOnMessageExpiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#default_message_ttl ServicebusQueue#default_message_ttl}.
	DefaultMessageTtl *string `field:"optional" json:"defaultMessageTtl" yaml:"defaultMessageTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#duplicate_detection_history_time_window ServicebusQueue#duplicate_detection_history_time_window}.
	DuplicateDetectionHistoryTimeWindow *string `field:"optional" json:"duplicateDetectionHistoryTimeWindow" yaml:"duplicateDetectionHistoryTimeWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#enable_batched_operations ServicebusQueue#enable_batched_operations}.
	EnableBatchedOperations interface{} `field:"optional" json:"enableBatchedOperations" yaml:"enableBatchedOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#enable_express ServicebusQueue#enable_express}.
	EnableExpress interface{} `field:"optional" json:"enableExpress" yaml:"enableExpress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#enable_partitioning ServicebusQueue#enable_partitioning}.
	EnablePartitioning interface{} `field:"optional" json:"enablePartitioning" yaml:"enablePartitioning"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#forward_dead_lettered_messages_to ServicebusQueue#forward_dead_lettered_messages_to}.
	ForwardDeadLetteredMessagesTo *string `field:"optional" json:"forwardDeadLetteredMessagesTo" yaml:"forwardDeadLetteredMessagesTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#forward_to ServicebusQueue#forward_to}.
	ForwardTo *string `field:"optional" json:"forwardTo" yaml:"forwardTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#id ServicebusQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#lock_duration ServicebusQueue#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#max_delivery_count ServicebusQueue#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"optional" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#max_message_size_in_kilobytes ServicebusQueue#max_message_size_in_kilobytes}.
	MaxMessageSizeInKilobytes *float64 `field:"optional" json:"maxMessageSizeInKilobytes" yaml:"maxMessageSizeInKilobytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#max_size_in_megabytes ServicebusQueue#max_size_in_megabytes}.
	MaxSizeInMegabytes *float64 `field:"optional" json:"maxSizeInMegabytes" yaml:"maxSizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#requires_duplicate_detection ServicebusQueue#requires_duplicate_detection}.
	RequiresDuplicateDetection interface{} `field:"optional" json:"requiresDuplicateDetection" yaml:"requiresDuplicateDetection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#requires_session ServicebusQueue#requires_session}.
	RequiresSession interface{} `field:"optional" json:"requiresSession" yaml:"requiresSession"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#status ServicebusQueue#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#timeouts ServicebusQueue#timeouts}
	Timeouts *ServicebusQueueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ServicebusQueueTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#create ServicebusQueue#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#delete ServicebusQueue#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#read ServicebusQueue#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_queue#update ServicebusQueue#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServicebusQueueTimeoutsOutputReference interface {
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

// The jsii proxy struct for ServicebusQueueTimeoutsOutputReference
type jsiiProxy_ServicebusQueueTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServicebusQueueTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusQueueTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicebusQueueTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueueTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusQueueTimeoutsOutputReference_Override(s ServicebusQueueTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueueTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueueTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueueTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

