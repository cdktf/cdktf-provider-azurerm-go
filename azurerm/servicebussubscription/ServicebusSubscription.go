// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebussubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/servicebussubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/servicebus_subscription azurerm_servicebus_subscription}.
type ServicebusSubscription interface {
	cdktf.TerraformResource
	AutoDeleteOnIdle() *string
	SetAutoDeleteOnIdle(val *string)
	AutoDeleteOnIdleInput() *string
	BatchedOperationsEnabled() interface{}
	SetBatchedOperationsEnabled(val interface{})
	BatchedOperationsEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientScopedSubscription() ServicebusSubscriptionClientScopedSubscriptionOutputReference
	ClientScopedSubscriptionEnabled() interface{}
	SetClientScopedSubscriptionEnabled(val interface{})
	ClientScopedSubscriptionEnabledInput() interface{}
	ClientScopedSubscriptionInput() *ServicebusSubscriptionClientScopedSubscription
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
	DeadLetteringOnFilterEvaluationError() interface{}
	SetDeadLetteringOnFilterEvaluationError(val interface{})
	DeadLetteringOnFilterEvaluationErrorInput() interface{}
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
	Timeouts() ServicebusSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TopicId() *string
	SetTopicId(val *string)
	TopicIdInput() *string
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
	PutClientScopedSubscription(value *ServicebusSubscriptionClientScopedSubscription)
	PutTimeouts(value *ServicebusSubscriptionTimeouts)
	ResetAutoDeleteOnIdle()
	ResetBatchedOperationsEnabled()
	ResetClientScopedSubscription()
	ResetClientScopedSubscriptionEnabled()
	ResetDeadLetteringOnFilterEvaluationError()
	ResetDeadLetteringOnMessageExpiration()
	ResetDefaultMessageTtl()
	ResetForwardDeadLetteredMessagesTo()
	ResetForwardTo()
	ResetId()
	ResetLockDuration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ServicebusSubscription
type jsiiProxy_ServicebusSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicebusSubscription) AutoDeleteOnIdle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) AutoDeleteOnIdleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) BatchedOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) BatchedOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ClientScopedSubscription() ServicebusSubscriptionClientScopedSubscriptionOutputReference {
	var returns ServicebusSubscriptionClientScopedSubscriptionOutputReference
	_jsii_.Get(
		j,
		"clientScopedSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ClientScopedSubscriptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientScopedSubscriptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ClientScopedSubscriptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientScopedSubscriptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ClientScopedSubscriptionInput() *ServicebusSubscriptionClientScopedSubscription {
	var returns *ServicebusSubscriptionClientScopedSubscription
	_jsii_.Get(
		j,
		"clientScopedSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DeadLetteringOnFilterEvaluationError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnFilterEvaluationError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DeadLetteringOnFilterEvaluationErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnFilterEvaluationErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DeadLetteringOnMessageExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnMessageExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DeadLetteringOnMessageExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetteringOnMessageExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DefaultMessageTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DefaultMessageTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ForwardDeadLetteredMessagesTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardDeadLetteredMessagesTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ForwardDeadLetteredMessagesToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardDeadLetteredMessagesToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ForwardTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) ForwardToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) LockDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) LockDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) MaxDeliveryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) MaxDeliveryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) RequiresSession() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresSession",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) RequiresSessionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresSessionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) Timeouts() ServicebusSubscriptionTimeoutsOutputReference {
	var returns ServicebusSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) TopicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/servicebus_subscription azurerm_servicebus_subscription} Resource.
func NewServicebusSubscription(scope constructs.Construct, id *string, config *ServicebusSubscriptionConfig) ServicebusSubscription {
	_init_.Initialize()

	if err := validateNewServicebusSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicebusSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/servicebus_subscription azurerm_servicebus_subscription} Resource.
func NewServicebusSubscription_Override(s ServicebusSubscription, scope constructs.Construct, id *string, config *ServicebusSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetAutoDeleteOnIdle(val *string) {
	if err := j.validateSetAutoDeleteOnIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetBatchedOperationsEnabled(val interface{}) {
	if err := j.validateSetBatchedOperationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchedOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetClientScopedSubscriptionEnabled(val interface{}) {
	if err := j.validateSetClientScopedSubscriptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientScopedSubscriptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetDeadLetteringOnFilterEvaluationError(val interface{}) {
	if err := j.validateSetDeadLetteringOnFilterEvaluationErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadLetteringOnFilterEvaluationError",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetDeadLetteringOnMessageExpiration(val interface{}) {
	if err := j.validateSetDeadLetteringOnMessageExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deadLetteringOnMessageExpiration",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetDefaultMessageTtl(val *string) {
	if err := j.validateSetDefaultMessageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetForwardDeadLetteredMessagesTo(val *string) {
	if err := j.validateSetForwardDeadLetteredMessagesToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardDeadLetteredMessagesTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetForwardTo(val *string) {
	if err := j.validateSetForwardToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetLockDuration(val *string) {
	if err := j.validateSetLockDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockDuration",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetMaxDeliveryCount(val *float64) {
	if err := j.validateSetMaxDeliveryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDeliveryCount",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetRequiresSession(val interface{}) {
	if err := j.validateSetRequiresSessionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresSession",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription)SetTopicId(val *string) {
	if err := j.validateSetTopicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicId",
		val,
	)
}

// Generates CDKTF code for importing a ServicebusSubscription resource upon running "cdktf plan <stack-name>".
func ServicebusSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicebusSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
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
func ServicebusSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicebusSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicebusSubscription) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicebusSubscription) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServicebusSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServicebusSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServicebusSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServicebusSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServicebusSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServicebusSubscription) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServicebusSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServicebusSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicebusSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusSubscription) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicebusSubscription) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusSubscription) PutClientScopedSubscription(value *ServicebusSubscriptionClientScopedSubscription) {
	if err := s.validatePutClientScopedSubscriptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putClientScopedSubscription",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusSubscription) PutTimeouts(value *ServicebusSubscriptionTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetAutoDeleteOnIdle() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoDeleteOnIdle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetBatchedOperationsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchedOperationsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetClientScopedSubscription() {
	_jsii_.InvokeVoid(
		s,
		"resetClientScopedSubscription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetClientScopedSubscriptionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetClientScopedSubscriptionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetDeadLetteringOnFilterEvaluationError() {
	_jsii_.InvokeVoid(
		s,
		"resetDeadLetteringOnFilterEvaluationError",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetDeadLetteringOnMessageExpiration() {
	_jsii_.InvokeVoid(
		s,
		"resetDeadLetteringOnMessageExpiration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetDefaultMessageTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultMessageTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetForwardDeadLetteredMessagesTo() {
	_jsii_.InvokeVoid(
		s,
		"resetForwardDeadLetteredMessagesTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetForwardTo() {
	_jsii_.InvokeVoid(
		s,
		"resetForwardTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetLockDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetLockDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetRequiresSession() {
	_jsii_.InvokeVoid(
		s,
		"resetRequiresSession",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

