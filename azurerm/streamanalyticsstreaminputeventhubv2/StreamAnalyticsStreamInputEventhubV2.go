// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsstreaminputeventhubv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/streamanalyticsstreaminputeventhubv2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/stream_analytics_stream_input_eventhub_v2 azurerm_stream_analytics_stream_input_eventhub_v2}.
type StreamAnalyticsStreamInputEventhubV2 interface {
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
	EventhubConsumerGroupName() *string
	SetEventhubConsumerGroupName(val *string)
	EventhubConsumerGroupNameInput() *string
	EventhubName() *string
	SetEventhubName(val *string)
	EventhubNameInput() *string
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
	PartitionKey() *string
	SetPartitionKey(val *string)
	PartitionKeyInput() *string
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
	Serialization() StreamAnalyticsStreamInputEventhubV2SerializationOutputReference
	SerializationInput() *StreamAnalyticsStreamInputEventhubV2Serialization
	ServicebusNamespace() *string
	SetServicebusNamespace(val *string)
	ServicebusNamespaceInput() *string
	SharedAccessPolicyKey() *string
	SetSharedAccessPolicyKey(val *string)
	SharedAccessPolicyKeyInput() *string
	SharedAccessPolicyName() *string
	SetSharedAccessPolicyName(val *string)
	SharedAccessPolicyNameInput() *string
	StreamAnalyticsJobId() *string
	SetStreamAnalyticsJobId(val *string)
	StreamAnalyticsJobIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StreamAnalyticsStreamInputEventhubV2TimeoutsOutputReference
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
	PutSerialization(value *StreamAnalyticsStreamInputEventhubV2Serialization)
	PutTimeouts(value *StreamAnalyticsStreamInputEventhubV2Timeouts)
	ResetAuthenticationMode()
	ResetEventhubConsumerGroupName()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitionKey()
	ResetSharedAccessPolicyKey()
	ResetSharedAccessPolicyName()
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

// The jsii proxy struct for StreamAnalyticsStreamInputEventhubV2
type jsiiProxy_StreamAnalyticsStreamInputEventhubV2 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) AuthenticationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) AuthenticationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) EventhubConsumerGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) EventhubConsumerGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubConsumerGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) EventhubName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) EventhubNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) PartitionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) PartitionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partitionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Serialization() StreamAnalyticsStreamInputEventhubV2SerializationOutputReference {
	var returns StreamAnalyticsStreamInputEventhubV2SerializationOutputReference
	_jsii_.Get(
		j,
		"serialization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SerializationInput() *StreamAnalyticsStreamInputEventhubV2Serialization {
	var returns *StreamAnalyticsStreamInputEventhubV2Serialization
	_jsii_.Get(
		j,
		"serializationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ServicebusNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ServicebusNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicebusNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SharedAccessPolicyKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SharedAccessPolicyKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SharedAccessPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SharedAccessPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedAccessPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) StreamAnalyticsJobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsJobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) StreamAnalyticsJobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsJobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) Timeouts() StreamAnalyticsStreamInputEventhubV2TimeoutsOutputReference {
	var returns StreamAnalyticsStreamInputEventhubV2TimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/stream_analytics_stream_input_eventhub_v2 azurerm_stream_analytics_stream_input_eventhub_v2} Resource.
func NewStreamAnalyticsStreamInputEventhubV2(scope constructs.Construct, id *string, config *StreamAnalyticsStreamInputEventhubV2Config) StreamAnalyticsStreamInputEventhubV2 {
	_init_.Initialize()

	if err := validateNewStreamAnalyticsStreamInputEventhubV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamAnalyticsStreamInputEventhubV2{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/stream_analytics_stream_input_eventhub_v2 azurerm_stream_analytics_stream_input_eventhub_v2} Resource.
func NewStreamAnalyticsStreamInputEventhubV2_Override(s StreamAnalyticsStreamInputEventhubV2, scope constructs.Construct, id *string, config *StreamAnalyticsStreamInputEventhubV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetAuthenticationMode(val *string) {
	if err := j.validateSetAuthenticationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationMode",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetEventhubConsumerGroupName(val *string) {
	if err := j.validateSetEventhubConsumerGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubConsumerGroupName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetEventhubName(val *string) {
	if err := j.validateSetEventhubNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetPartitionKey(val *string) {
	if err := j.validateSetPartitionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionKey",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetServicebusNamespace(val *string) {
	if err := j.validateSetServicebusNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicebusNamespace",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetSharedAccessPolicyKey(val *string) {
	if err := j.validateSetSharedAccessPolicyKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessPolicyKey",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetSharedAccessPolicyName(val *string) {
	if err := j.validateSetSharedAccessPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedAccessPolicyName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsStreamInputEventhubV2)SetStreamAnalyticsJobId(val *string) {
	if err := j.validateSetStreamAnalyticsJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamAnalyticsJobId",
		val,
	)
}

// Generates CDKTF code for importing a StreamAnalyticsStreamInputEventhubV2 resource upon running "cdktf plan <stack-name>".
func StreamAnalyticsStreamInputEventhubV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStreamAnalyticsStreamInputEventhubV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
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
func StreamAnalyticsStreamInputEventhubV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsStreamInputEventhubV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsStreamInputEventhubV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsStreamInputEventhubV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsStreamInputEventhubV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsStreamInputEventhubV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StreamAnalyticsStreamInputEventhubV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.streamAnalyticsStreamInputEventhubV2.StreamAnalyticsStreamInputEventhubV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) PutSerialization(value *StreamAnalyticsStreamInputEventhubV2Serialization) {
	if err := s.validatePutSerializationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSerialization",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) PutTimeouts(value *StreamAnalyticsStreamInputEventhubV2Timeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetAuthenticationMode() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthenticationMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetEventhubConsumerGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetEventhubConsumerGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetPartitionKey() {
	_jsii_.InvokeVoid(
		s,
		"resetPartitionKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetSharedAccessPolicyKey() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAccessPolicyKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetSharedAccessPolicyName() {
	_jsii_.InvokeVoid(
		s,
		"resetSharedAccessPolicyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsStreamInputEventhubV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

