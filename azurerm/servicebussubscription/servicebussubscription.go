package servicebussubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicebussubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription azurerm_servicebus_subscription}.
type ServicebusSubscription interface {
	cdktf.TerraformResource
	AutoDeleteOnIdle() *string
	SetAutoDeleteOnIdle(val *string)
	AutoDeleteOnIdleInput() *string
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	EnableBatchedOperations() interface{}
	SetEnableBatchedOperations(val interface{})
	EnableBatchedOperationsInput() interface{}
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
	PutClientScopedSubscription(value *ServicebusSubscriptionClientScopedSubscription)
	PutTimeouts(value *ServicebusSubscriptionTimeouts)
	ResetAutoDeleteOnIdle()
	ResetClientScopedSubscription()
	ResetClientScopedSubscriptionEnabled()
	ResetDeadLetteringOnFilterEvaluationError()
	ResetDeadLetteringOnMessageExpiration()
	ResetDefaultMessageTtl()
	ResetEnableBatchedOperations()
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

func (j *jsiiProxy_ServicebusSubscription) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_ServicebusSubscription) EnableBatchedOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscription) EnableBatchedOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperationsInput",
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription azurerm_servicebus_subscription} Resource.
func NewServicebusSubscription(scope constructs.Construct, id *string, config *ServicebusSubscriptionConfig) ServicebusSubscription {
	_init_.Initialize()

	j := jsiiProxy_ServicebusSubscription{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription azurerm_servicebus_subscription} Resource.
func NewServicebusSubscription_Override(s ServicebusSubscription, scope constructs.Construct, id *string, config *ServicebusSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetAutoDeleteOnIdle(val *string) {
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetClientScopedSubscriptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clientScopedSubscriptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetDeadLetteringOnFilterEvaluationError(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetteringOnFilterEvaluationError",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetDeadLetteringOnMessageExpiration(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetteringOnMessageExpiration",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetDefaultMessageTtl(val *string) {
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetEnableBatchedOperations(val interface{}) {
	_jsii_.Set(
		j,
		"enableBatchedOperations",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetForwardDeadLetteredMessagesTo(val *string) {
	_jsii_.Set(
		j,
		"forwardDeadLetteredMessagesTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetForwardTo(val *string) {
	_jsii_.Set(
		j,
		"forwardTo",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetLockDuration(val *string) {
	_jsii_.Set(
		j,
		"lockDuration",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetMaxDeliveryCount(val *float64) {
	_jsii_.Set(
		j,
		"maxDeliveryCount",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetRequiresSession(val interface{}) {
	_jsii_.Set(
		j,
		"requiresSession",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscription) SetTopicId(val *string) {
	_jsii_.Set(
		j,
		"topicId",
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
func ServicebusSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscription",
		"isConstruct",
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

func (s *jsiiProxy_ServicebusSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusSubscription) PutClientScopedSubscription(value *ServicebusSubscriptionClientScopedSubscription) {
	_jsii_.InvokeVoid(
		s,
		"putClientScopedSubscription",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusSubscription) PutTimeouts(value *ServicebusSubscriptionTimeouts) {
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

func (s *jsiiProxy_ServicebusSubscription) ResetEnableBatchedOperations() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableBatchedOperations",
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

type ServicebusSubscriptionClientScopedSubscription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#client_id ServicebusSubscription#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#is_client_scoped_subscription_shareable ServicebusSubscription#is_client_scoped_subscription_shareable}.
	IsClientScopedSubscriptionShareable interface{} `field:"optional" json:"isClientScopedSubscriptionShareable" yaml:"isClientScopedSubscriptionShareable"`
}

type ServicebusSubscriptionClientScopedSubscriptionOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	InternalValue() *ServicebusSubscriptionClientScopedSubscription
	SetInternalValue(val *ServicebusSubscriptionClientScopedSubscription)
	IsClientScopedSubscriptionDurable() cdktf.IResolvable
	IsClientScopedSubscriptionShareable() interface{}
	SetIsClientScopedSubscriptionShareable(val interface{})
	IsClientScopedSubscriptionShareableInput() interface{}
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
	ResetClientId()
	ResetIsClientScopedSubscriptionShareable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicebusSubscriptionClientScopedSubscriptionOutputReference
type jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) InternalValue() *ServicebusSubscriptionClientScopedSubscription {
	var returns *ServicebusSubscriptionClientScopedSubscription
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) IsClientScopedSubscriptionDurable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isClientScopedSubscriptionDurable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) IsClientScopedSubscriptionShareable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isClientScopedSubscriptionShareable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) IsClientScopedSubscriptionShareableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isClientScopedSubscriptionShareableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServicebusSubscriptionClientScopedSubscriptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusSubscriptionClientScopedSubscriptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscriptionClientScopedSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusSubscriptionClientScopedSubscriptionOutputReference_Override(s ServicebusSubscriptionClientScopedSubscriptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscriptionClientScopedSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetInternalValue(val *ServicebusSubscriptionClientScopedSubscription) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetIsClientScopedSubscriptionShareable(val interface{}) {
	_jsii_.Set(
		j,
		"isClientScopedSubscriptionShareable",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		s,
		"resetClientId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ResetIsClientScopedSubscriptionShareable() {
	_jsii_.InvokeVoid(
		s,
		"resetIsClientScopedSubscriptionShareable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionClientScopedSubscriptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServicebusSubscriptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#max_delivery_count ServicebusSubscription#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"required" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#name ServicebusSubscription#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#topic_id ServicebusSubscription#topic_id}.
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#auto_delete_on_idle ServicebusSubscription#auto_delete_on_idle}.
	AutoDeleteOnIdle *string `field:"optional" json:"autoDeleteOnIdle" yaml:"autoDeleteOnIdle"`
	// client_scoped_subscription block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#client_scoped_subscription ServicebusSubscription#client_scoped_subscription}
	ClientScopedSubscription *ServicebusSubscriptionClientScopedSubscription `field:"optional" json:"clientScopedSubscription" yaml:"clientScopedSubscription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#client_scoped_subscription_enabled ServicebusSubscription#client_scoped_subscription_enabled}.
	ClientScopedSubscriptionEnabled interface{} `field:"optional" json:"clientScopedSubscriptionEnabled" yaml:"clientScopedSubscriptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#dead_lettering_on_filter_evaluation_error ServicebusSubscription#dead_lettering_on_filter_evaluation_error}.
	DeadLetteringOnFilterEvaluationError interface{} `field:"optional" json:"deadLetteringOnFilterEvaluationError" yaml:"deadLetteringOnFilterEvaluationError"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#dead_lettering_on_message_expiration ServicebusSubscription#dead_lettering_on_message_expiration}.
	DeadLetteringOnMessageExpiration interface{} `field:"optional" json:"deadLetteringOnMessageExpiration" yaml:"deadLetteringOnMessageExpiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#default_message_ttl ServicebusSubscription#default_message_ttl}.
	DefaultMessageTtl *string `field:"optional" json:"defaultMessageTtl" yaml:"defaultMessageTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#enable_batched_operations ServicebusSubscription#enable_batched_operations}.
	EnableBatchedOperations interface{} `field:"optional" json:"enableBatchedOperations" yaml:"enableBatchedOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#forward_dead_lettered_messages_to ServicebusSubscription#forward_dead_lettered_messages_to}.
	ForwardDeadLetteredMessagesTo *string `field:"optional" json:"forwardDeadLetteredMessagesTo" yaml:"forwardDeadLetteredMessagesTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#forward_to ServicebusSubscription#forward_to}.
	ForwardTo *string `field:"optional" json:"forwardTo" yaml:"forwardTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#id ServicebusSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#lock_duration ServicebusSubscription#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#requires_session ServicebusSubscription#requires_session}.
	RequiresSession interface{} `field:"optional" json:"requiresSession" yaml:"requiresSession"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#status ServicebusSubscription#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#timeouts ServicebusSubscription#timeouts}
	Timeouts *ServicebusSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ServicebusSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#create ServicebusSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#delete ServicebusSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#read ServicebusSubscription#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#update ServicebusSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServicebusSubscriptionTimeoutsOutputReference interface {
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

// The jsii proxy struct for ServicebusSubscriptionTimeoutsOutputReference
type jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServicebusSubscriptionTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusSubscriptionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusSubscriptionTimeoutsOutputReference_Override(s ServicebusSubscriptionTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusSubscription.ServicebusSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusSubscriptionTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

