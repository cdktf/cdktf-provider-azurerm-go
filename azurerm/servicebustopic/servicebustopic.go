package servicebustopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicebustopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic azurerm_servicebus_topic}.
type ServicebusTopic interface {
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
	PutTimeouts(value *ServicebusTopicTimeouts)
	ResetAutoDeleteOnIdle()
	ResetDefaultMessageTtl()
	ResetDuplicateDetectionHistoryTimeWindow()
	ResetEnableBatchedOperations()
	ResetEnableExpress()
	ResetEnablePartitioning()
	ResetId()
	ResetMaxMessageSizeInKilobytes()
	ResetMaxSizeInMegabytes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequiresDuplicateDetection()
	ResetStatus()
	ResetSupportOrdering()
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

func (j *jsiiProxy_ServicebusTopic) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_ServicebusTopic) EnableBatchedOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) EnableBatchedOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBatchedOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) EnableExpress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExpress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) EnableExpressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExpressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) EnablePartitioning() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopic) EnablePartitioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartitioningInput",
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic azurerm_servicebus_topic} Resource.
func NewServicebusTopic(scope constructs.Construct, id *string, config *ServicebusTopicConfig) ServicebusTopic {
	_init_.Initialize()

	j := jsiiProxy_ServicebusTopic{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic azurerm_servicebus_topic} Resource.
func NewServicebusTopic_Override(s ServicebusTopic, scope constructs.Construct, id *string, config *ServicebusTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetAutoDeleteOnIdle(val *string) {
	_jsii_.Set(
		j,
		"autoDeleteOnIdle",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetDefaultMessageTtl(val *string) {
	_jsii_.Set(
		j,
		"defaultMessageTtl",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetDuplicateDetectionHistoryTimeWindow(val *string) {
	_jsii_.Set(
		j,
		"duplicateDetectionHistoryTimeWindow",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetEnableBatchedOperations(val interface{}) {
	_jsii_.Set(
		j,
		"enableBatchedOperations",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetEnableExpress(val interface{}) {
	_jsii_.Set(
		j,
		"enableExpress",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetEnablePartitioning(val interface{}) {
	_jsii_.Set(
		j,
		"enablePartitioning",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetMaxMessageSizeInKilobytes(val *float64) {
	_jsii_.Set(
		j,
		"maxMessageSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetMaxSizeInMegabytes(val *float64) {
	_jsii_.Set(
		j,
		"maxSizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetNamespaceId(val *string) {
	_jsii_.Set(
		j,
		"namespaceId",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetRequiresDuplicateDetection(val interface{}) {
	_jsii_.Set(
		j,
		"requiresDuplicateDetection",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopic) SetSupportOrdering(val interface{}) {
	_jsii_.Set(
		j,
		"supportOrdering",
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
func ServicebusTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		"isConstruct",
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

func (s *jsiiProxy_ServicebusTopic) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServicebusTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopic) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServicebusTopic) PutTimeouts(value *ServicebusTopicTimeouts) {
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

func (s *jsiiProxy_ServicebusTopic) ResetEnableBatchedOperations() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableBatchedOperations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetEnableExpress() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableExpress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopic) ResetEnablePartitioning() {
	_jsii_.InvokeVoid(
		s,
		"resetEnablePartitioning",
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

type ServicebusTopicConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#name ServicebusTopic#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#namespace_id ServicebusTopic#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#auto_delete_on_idle ServicebusTopic#auto_delete_on_idle}.
	AutoDeleteOnIdle *string `field:"optional" json:"autoDeleteOnIdle" yaml:"autoDeleteOnIdle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#default_message_ttl ServicebusTopic#default_message_ttl}.
	DefaultMessageTtl *string `field:"optional" json:"defaultMessageTtl" yaml:"defaultMessageTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#duplicate_detection_history_time_window ServicebusTopic#duplicate_detection_history_time_window}.
	DuplicateDetectionHistoryTimeWindow *string `field:"optional" json:"duplicateDetectionHistoryTimeWindow" yaml:"duplicateDetectionHistoryTimeWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#enable_batched_operations ServicebusTopic#enable_batched_operations}.
	EnableBatchedOperations interface{} `field:"optional" json:"enableBatchedOperations" yaml:"enableBatchedOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#enable_express ServicebusTopic#enable_express}.
	EnableExpress interface{} `field:"optional" json:"enableExpress" yaml:"enableExpress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#enable_partitioning ServicebusTopic#enable_partitioning}.
	EnablePartitioning interface{} `field:"optional" json:"enablePartitioning" yaml:"enablePartitioning"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#id ServicebusTopic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#max_message_size_in_kilobytes ServicebusTopic#max_message_size_in_kilobytes}.
	MaxMessageSizeInKilobytes *float64 `field:"optional" json:"maxMessageSizeInKilobytes" yaml:"maxMessageSizeInKilobytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#max_size_in_megabytes ServicebusTopic#max_size_in_megabytes}.
	MaxSizeInMegabytes *float64 `field:"optional" json:"maxSizeInMegabytes" yaml:"maxSizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#requires_duplicate_detection ServicebusTopic#requires_duplicate_detection}.
	RequiresDuplicateDetection interface{} `field:"optional" json:"requiresDuplicateDetection" yaml:"requiresDuplicateDetection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#status ServicebusTopic#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#support_ordering ServicebusTopic#support_ordering}.
	SupportOrdering interface{} `field:"optional" json:"supportOrdering" yaml:"supportOrdering"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#timeouts ServicebusTopic#timeouts}
	Timeouts *ServicebusTopicTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ServicebusTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#create ServicebusTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#delete ServicebusTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#read ServicebusTopic#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_topic#update ServicebusTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServicebusTopicTimeoutsOutputReference interface {
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

// The jsii proxy struct for ServicebusTopicTimeoutsOutputReference
type jsiiProxy_ServicebusTopicTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServicebusTopicTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServicebusTopicTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServicebusTopicTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopicTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicebusTopicTimeoutsOutputReference_Override(s ServicebusTopicTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopicTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServicebusTopicTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicebusTopicTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

