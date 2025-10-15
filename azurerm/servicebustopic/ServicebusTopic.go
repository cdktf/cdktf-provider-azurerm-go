// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebustopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/servicebustopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_topic azurerm_servicebus_topic}.
type ServicebusTopic interface {
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SupportOrdering() interface{}
	SetSupportOrdering(val interface{})
	SupportOrderingInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServicebusTopicTimeoutsOutputReference
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
	PutTimeouts(value *ServicebusTopicTimeouts)
	ResetAutoDeleteOnIdle()
	ResetBatchedOperationsEnabled()
	ResetDefaultMessageTtl()
	ResetDuplicateDetectionHistoryTimeWindow()
	ResetExpressEnabled()
	ResetId()
	ResetMaxMessageSizeInKilobytes()
	ResetMaxSizeInMegabytes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartitioningEnabled()
	ResetRequiresDuplicateDetection()
	ResetStatus()
	ResetSupportOrdering()
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

// The jsii proxy struct for ServicebusTopic
type jsiiProxy_ServicebusTopic struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServicebusTopic) AutoDeleteOnIdle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) AutoDeleteOnIdleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteOnIdleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) BatchedOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) BatchedOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchedOperationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) DefaultMessageTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) DefaultMessageTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultMessageTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) DuplicateDetectionHistoryTimeWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duplicateDetectionHistoryTimeWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) DuplicateDetectionHistoryTimeWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duplicateDetectionHistoryTimeWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) ExpressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) ExpressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) MaxMessageSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) MaxMessageSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) MaxSizeInMegabytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) MaxSizeInMegabytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) NamespaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) PartitioningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitioningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) PartitioningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitioningEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) RequiresDuplicateDetection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresDuplicateDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) RequiresDuplicateDetectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiresDuplicateDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) SupportOrdering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportOrdering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) SupportOrderingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"supportOrderingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) Timeouts() ServicebusTopicTimeoutsOutputReference {
	var returns ServicebusTopicTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_topic azurerm_servicebus_topic} Resource.
func NewServicebusTopic(scope constructs.Construct, id *string, config *ServicebusTopicConfig) ServicebusTopic {
	_init_.Initialize()

	if err := validateNewServicebusTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicebusTopic{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/servicebus_topic azurerm_servicebus_topic} Resource.
func NewServicebusTopic_Override(s ServicebusTopic, scope constructs.Construct, id *string, config *ServicebusTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetAutoDeleteOnIdle(val *string) {
	if err := j.validateSetAutoDeleteOnIdleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetBatchedOperationsEnabled(val interface{}) {
	if err := j.validateSetBatchedOperationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchedOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetDefaultMessageTtl(val *string) {
	if err := j.validateSetDefaultMessageTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetDuplicateDetectionHistoryTimeWindow(val *string) {
	if err := j.validateSetDuplicateDetectionHistoryTimeWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duplicateDetectionHistoryTimeWindow",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetExpressEnabled(val interface{}) {
	if err := j.validateSetExpressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetMaxMessageSizeInKilobytes(val *float64) {
	if err := j.validateSetMaxMessageSizeInKilobytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMessageSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetMaxSizeInMegabytes(val *float64) {
	if err := j.validateSetMaxSizeInMegabytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetNamespaceId(val *string) {
	if err := j.validateSetNamespaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetPartitioningEnabled(val interface{}) {
	if err := j.validateSetPartitioningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitioningEnabled",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetRequiresDuplicateDetection(val interface{}) {
	if err := j.validateSetRequiresDuplicateDetectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresDuplicateDetection",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic)SetSupportOrdering(val interface{}) {
	if err := j.validateSetSupportOrderingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportOrdering",
		val,
	)
}

// Generates CDKTF code for importing a ServicebusTopic resource upon running "cdktf plan <stack-name>".
func ServicebusTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServicebusTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
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
func ServicebusTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServicebusTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServicebusTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServicebusTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServicebusTopic) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServicebusTopic) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServicebusTopic) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServicebusTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServicebusTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServicebusTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServicebusTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServicebusTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServicebusTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServicebusTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServicebusTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServicebusTopic) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServicebusTopic) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServicebusTopic) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusTopic) PutTimeouts(value *ServicebusTopicTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetAutoDeleteOnIdle() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoDeleteOnIdle",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetBatchedOperationsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBatchedOperationsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetDefaultMessageTtl() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultMessageTtl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetDuplicateDetectionHistoryTimeWindow() {
	_jsii_.InvokeVoid(
		s,
		"resetDuplicateDetectionHistoryTimeWindow",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetExpressEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetExpressEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetMaxMessageSizeInKilobytes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxMessageSizeInKilobytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetMaxSizeInMegabytes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxSizeInMegabytes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetPartitioningEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetPartitioningEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetRequiresDuplicateDetection() {
	_jsii_.InvokeVoid(
		s,
		"resetRequiresDuplicateDetection",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetSupportOrdering() {
	_jsii_.InvokeVoid(
		s,
		"resetSupportOrdering",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

