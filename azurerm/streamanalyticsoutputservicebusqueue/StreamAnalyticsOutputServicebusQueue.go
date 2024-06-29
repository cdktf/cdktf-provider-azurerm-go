// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsoutputservicebusqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/streamanalyticsoutputservicebusqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/stream_analytics_output_servicebus_queue azurerm_stream_analytics_output_servicebus_queue}.
type StreamAnalyticsOutputServicebusQueue interface {
	cdktf.TerraformResource
	AuthenticationMode() *string
	SetAuthenticationMode(val *string)
	AuthenticationModeInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PropertyColumns() *[]*string
	SetPropertyColumns(val *[]*string)
	PropertyColumnsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QueueName() *string
	SetQueueName(val *string)
	QueueNameInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Serialization() StreamAnalyticsOutputServicebusQueueSerializationOutputReference
	SerializationInput() *StreamAnalyticsOutputServicebusQueueSerialization
	ServicebusNamespace() *string
	SetServicebusNamespace(val *string)
	ServicebusNamespaceInput() *string
	SharedAccessPolicyKey() *string
	SetSharedAccessPolicyKey(val *string)
	SharedAccessPolicyKeyInput() *string
	SharedAccessPolicyName() *string
	SetSharedAccessPolicyName(val *string)
	SharedAccessPolicyNameInput() *string
	StreamAnalyticsJobName() *string
	SetStreamAnalyticsJobName(val *string)
	StreamAnalyticsJobNameInput() *string
	SystemPropertyColumns() *map[string]*string
	SetSystemPropertyColumns(val *map[string]*string)
	SystemPropertyColumnsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StreamAnalyticsOutputServicebusQueueTimeoutsOutputReference
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
	PutSerialization(value *StreamAnalyticsOutputServicebusQueueSerialization)
	PutTimeouts(value *StreamAnalyticsOutputServicebusQueueTimeouts)
	ResetAuthenticationMode()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPropertyColumns()
	ResetSharedAccessPolicyKey()
	ResetSharedAccessPolicyName()
	ResetSystemPropertyColumns()
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

// The jsii proxy struct for StreamAnalyticsOutputServicebusQueue
type jsiiProxy_StreamAnalyticsOutputServicebusQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) AuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) AuthenticationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) PropertyColumns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"propertyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) PropertyColumnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"propertyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) QueueName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) QueueNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Serialization() StreamAnalyticsOutputServicebusQueueSerializationOutputReference {
	var returns StreamAnalyticsOutputServicebusQueueSerializationOutputReference
	_jsii_.Get(
		j,
		"serialization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SerializationInput() *StreamAnalyticsOutputServicebusQueueSerialization {
	var returns *StreamAnalyticsOutputServicebusQueueSerialization
	_jsii_.Get(
		j,
		"serializationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ServicebusNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ServicebusNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SharedAccessPolicyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SharedAccessPolicyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SharedAccessPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SharedAccessPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) StreamAnalyticsJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) StreamAnalyticsJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SystemPropertyColumns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemPropertyColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SystemPropertyColumnsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemPropertyColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) Timeouts() StreamAnalyticsOutputServicebusQueueTimeoutsOutputReference {
	var returns StreamAnalyticsOutputServicebusQueueTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/stream_analytics_output_servicebus_queue azurerm_stream_analytics_output_servicebus_queue} Resource.
func NewStreamAnalyticsOutputServicebusQueue(scope constructs.Construct, id *string, config *StreamAnalyticsOutputServicebusQueueConfig) StreamAnalyticsOutputServicebusQueue {
	_init_.Initialize()

	if err := validateNewStreamAnalyticsOutputServicebusQueueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamAnalyticsOutputServicebusQueue{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/stream_analytics_output_servicebus_queue azurerm_stream_analytics_output_servicebus_queue} Resource.
func NewStreamAnalyticsOutputServicebusQueue_Override(s StreamAnalyticsOutputServicebusQueue, scope constructs.Construct, id *string, config *StreamAnalyticsOutputServicebusQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetAuthenticationMode(val *string) {
	if err := j.validateSetAuthenticationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMode",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetPropertyColumns(val *[]*string) {
	if err := j.validateSetPropertyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propertyColumns",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetQueueName(val *string) {
	if err := j.validateSetQueueNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetServicebusNamespace(val *string) {
	if err := j.validateSetServicebusNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicebusNamespace",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetSharedAccessPolicyKey(val *string) {
	if err := j.validateSetSharedAccessPolicyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessPolicyKey",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetSharedAccessPolicyName(val *string) {
	if err := j.validateSetSharedAccessPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessPolicyName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetStreamAnalyticsJobName(val *string) {
	if err := j.validateSetStreamAnalyticsJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamAnalyticsJobName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsOutputServicebusQueue)SetSystemPropertyColumns(val *map[string]*string) {
	if err := j.validateSetSystemPropertyColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPropertyColumns",
		val,
	)
}

// Generates CDKTF code for importing a StreamAnalyticsOutputServicebusQueue resource upon running "cdktf plan <stack-name>".
func StreamAnalyticsOutputServicebusQueue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStreamAnalyticsOutputServicebusQueue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
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
func StreamAnalyticsOutputServicebusQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsOutputServicebusQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsOutputServicebusQueue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsOutputServicebusQueue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsOutputServicebusQueue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsOutputServicebusQueue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StreamAnalyticsOutputServicebusQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.streamAnalyticsOutputServicebusQueue.StreamAnalyticsOutputServicebusQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) PutSerialization(value *StreamAnalyticsOutputServicebusQueueSerialization) {
	if err := s.validatePutSerializationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSerialization",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) PutTimeouts(value *StreamAnalyticsOutputServicebusQueueTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetAuthenticationMode() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticationMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetPropertyColumns() {
	_jsii_.InvokeVoid(
		s,
		"resetPropertyColumns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetSharedAccessPolicyKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAccessPolicyKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetSharedAccessPolicyName() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAccessPolicyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetSystemPropertyColumns() {
	_jsii_.InvokeVoid(
		s,
		"resetSystemPropertyColumns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsOutputServicebusQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

