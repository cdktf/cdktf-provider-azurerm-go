// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/servicebusqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_queue azurerm_servicebus_queue}.
type ServicebusQueue interface {
	cdktf.TerraformResource
	AutoDeleteOnIdle() *string
	SetAutoDeleteOnIdle(val *string)
	AutoDeleteOnIdleInput() *string
	BatchedOperationsEnabled() interface{}
	SetBatchedOperationsEnabled(val interface{})
	BatchedOperationsEnabledInput() interface{}
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
	ExpressEnabled() interface{}
	SetExpressEnabled(val interface{})
	ExpressEnabledInput() interface{}
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
	PartitioningEnabled() interface{}
	SetPartitioningEnabled(val interface{})
	PartitioningEnabledInput() interface{}
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
	PutTimeouts(value *ServicebusQueueTimeouts)
	ResetAutoDeleteOnIdle()
	ResetBatchedOperationsEnabled()
	ResetDeadLetteringOnMessageExpiration()
	ResetDefaultMessageTtl()
	ResetDuplicateDetectionHistoryTimeWindow()
	ResetExpressEnabled()
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
	ResetPartitioningEnabled()
	ResetRequiresDuplicateDetection()
	ResetRequiresSession()
	ResetStatus()
	ResetTimeouts()
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

func (j *jsiiProxy_ServicebusQueue) BatchedOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) BatchedOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabledInput",
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

func (j *jsiiProxy_ServicebusQueue) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ServicebusQueue) ExpressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) ExpressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressEnabledInput",
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

func (j *jsiiProxy_ServicebusQueue) PartitioningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitioningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusQueue) PartitioningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitioningEnabledInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_queue azurerm_servicebus_queue} Resource.
func NewServicebusQueue(scope constructs.Construct, id *string, config *ServicebusQueueConfig) ServicebusQueue {
	_init_.Initialize()

	if err := validateNewServicebusQueueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicebusQueue{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_queue azurerm_servicebus_queue} Resource.
func NewServicebusQueue_Override(s ServicebusQueue, scope constructs.Construct, id *string, config *ServicebusQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetAutoDeleteOnIdle(val *string) {
	if err := j.validateSetAutoDeleteOnIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetBatchedOperationsEnabled(val interface{}) {
	if err := j.validateSetBatchedOperationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchedOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetDeadLetteringOnMessageExpiration(val interface{}) {
	if err := j.validateSetDeadLetteringOnMessageExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadLetteringOnMessageExpiration",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetDefaultMessageTtl(val *string) {
	if err := j.validateSetDefaultMessageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetDuplicateDetectionHistoryTimeWindow(val *string) {
	if err := j.validateSetDuplicateDetectionHistoryTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duplicateDetectionHistoryTimeWindow",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetExpressEnabled(val interface{}) {
	if err := j.validateSetExpressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetForwardDeadLetteredMessagesTo(val *string) {
	if err := j.validateSetForwardDeadLetteredMessagesToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardDeadLetteredMessagesTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetForwardTo(val *string) {
	if err := j.validateSetForwardToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetLockDuration(val *string) {
	if err := j.validateSetLockDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockDuration",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetMaxDeliveryCount(val *float64) {
	if err := j.validateSetMaxDeliveryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDeliveryCount",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetMaxMessageSizeInKilobytes(val *float64) {
	if err := j.validateSetMaxMessageSizeInKilobytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMessageSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetMaxSizeInMegabytes(val *float64) {
	if err := j.validateSetMaxSizeInMegabytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetNamespaceId(val *string) {
	if err := j.validateSetNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetPartitioningEnabled(val interface{}) {
	if err := j.validateSetPartitioningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitioningEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetRequiresDuplicateDetection(val interface{}) {
	if err := j.validateSetRequiresDuplicateDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresDuplicateDetection",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetRequiresSession(val interface{}) {
	if err := j.validateSetRequiresSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresSession",
		val,
	)
}

func (j *jsiiProxy_ServicebusQueue)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Generates CDKTF code for importing a ServicebusQueue resource upon running "cdktf plan <stack-name>".
func ServicebusQueue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicebusQueue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
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
func ServicebusQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusQueue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusQueue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusQueue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusQueue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusQueue.ServicebusQueue",
		"isTerraformResource",
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

func (s *jsiiProxy_ServicebusQueue) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicebusQueue) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicebusQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusQueue) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicebusQueue) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusQueue) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusQueue) PutTimeouts(value *ServicebusQueueTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (s *jsiiProxy_ServicebusQueue) ResetBatchedOperationsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchedOperationsEnabled",
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

func (s *jsiiProxy_ServicebusQueue) ResetExpressEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetExpressEnabled",
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

func (s *jsiiProxy_ServicebusQueue) ResetPartitioningEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPartitioningEnabled",
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

func (s *jsiiProxy_ServicebusQueue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusQueue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
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

