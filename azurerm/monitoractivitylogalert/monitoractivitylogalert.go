package monitoractivitylogalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/monitoractivitylogalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert azurerm_monitor_activity_log_alert}.
type MonitorActivityLogAlert interface {
	cdktf.TerraformResource
	Action() MonitorActivityLogAlertActionList
	ActionInput() interface{}
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
	Criteria() MonitorActivityLogAlertCriteriaOutputReference
	CriteriaInput() *MonitorActivityLogAlertCriteria
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MonitorActivityLogAlertTimeoutsOutputReference
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
	PutAction(value interface{})
	PutCriteria(value *MonitorActivityLogAlertCriteria)
	PutTimeouts(value *MonitorActivityLogAlertTimeouts)
	ResetAction()
	ResetDescription()
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for MonitorActivityLogAlert
type jsiiProxy_MonitorActivityLogAlert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MonitorActivityLogAlert) Action() MonitorActivityLogAlertActionList {
	var returns MonitorActivityLogAlertActionList
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Criteria() MonitorActivityLogAlertCriteriaOutputReference {
	var returns MonitorActivityLogAlertCriteriaOutputReference
	_jsii_.Get(
		j,
		"criteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) CriteriaInput() *MonitorActivityLogAlertCriteria {
	var returns *MonitorActivityLogAlertCriteria
	_jsii_.Get(
		j,
		"criteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) Timeouts() MonitorActivityLogAlertTimeoutsOutputReference {
	var returns MonitorActivityLogAlertTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlert) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert azurerm_monitor_activity_log_alert} Resource.
func NewMonitorActivityLogAlert(scope constructs.Construct, id *string, config *MonitorActivityLogAlertConfig) MonitorActivityLogAlert {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlert{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert azurerm_monitor_activity_log_alert} Resource.
func NewMonitorActivityLogAlert_Override(m MonitorActivityLogAlert, scope constructs.Construct, id *string, config *MonitorActivityLogAlertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlert",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlert) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
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
func MonitorActivityLogAlert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MonitorActivityLogAlert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) PutAction(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putAction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) PutCriteria(value *MonitorActivityLogAlertCriteria) {
	_jsii_.InvokeVoid(
		m,
		"putCriteria",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) PutTimeouts(value *MonitorActivityLogAlertTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetAction() {
	_jsii_.InvokeVoid(
		m,
		"resetAction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#action_group_id MonitorActivityLogAlert#action_group_id}.
	ActionGroupId *string `field:"required" json:"actionGroupId" yaml:"actionGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#webhook_properties MonitorActivityLogAlert#webhook_properties}.
	WebhookProperties *map[string]*string `field:"optional" json:"webhookProperties" yaml:"webhookProperties"`
}

type MonitorActivityLogAlertActionList interface {
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
	Get(index *float64) MonitorActivityLogAlertActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertActionList
type jsiiProxy_MonitorActivityLogAlertActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorActivityLogAlertActionList {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertActionList_Override(m MonitorActivityLogAlertActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionList) Get(index *float64) MonitorActivityLogAlertActionOutputReference {
	var returns MonitorActivityLogAlertActionOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertActionOutputReference interface {
	cdktf.ComplexObject
	ActionGroupId() *string
	SetActionGroupId(val *string)
	ActionGroupIdInput() *string
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebhookProperties() *map[string]*string
	SetWebhookProperties(val *map[string]*string)
	WebhookPropertiesInput() *map[string]*string
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
	ResetWebhookProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertActionOutputReference
type jsiiProxy_MonitorActivityLogAlertActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ActionGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ActionGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) WebhookProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webhookProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) WebhookPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"webhookPropertiesInput",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorActivityLogAlertActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertActionOutputReference_Override(m MonitorActivityLogAlertActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetActionGroupId(val *string) {
	_jsii_.Set(
		j,
		"actionGroupId",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertActionOutputReference) SetWebhookProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"webhookProperties",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ResetWebhookProperties() {
	_jsii_.InvokeVoid(
		m,
		"resetWebhookProperties",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertConfig struct {
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
	// criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#criteria MonitorActivityLogAlert#criteria}
	Criteria *MonitorActivityLogAlertCriteria `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#name MonitorActivityLogAlert#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_group_name MonitorActivityLogAlert#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#scopes MonitorActivityLogAlert#scopes}.
	Scopes *[]*string `field:"required" json:"scopes" yaml:"scopes"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#action MonitorActivityLogAlert#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#description MonitorActivityLogAlert#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#enabled MonitorActivityLogAlert#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#id MonitorActivityLogAlert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#tags MonitorActivityLogAlert#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#timeouts MonitorActivityLogAlert#timeouts}
	Timeouts *MonitorActivityLogAlertTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MonitorActivityLogAlertCriteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#category MonitorActivityLogAlert#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#caller MonitorActivityLogAlert#caller}.
	Caller *string `field:"optional" json:"caller" yaml:"caller"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#level MonitorActivityLogAlert#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#operation_name MonitorActivityLogAlert#operation_name}.
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#recommendation_category MonitorActivityLogAlert#recommendation_category}.
	RecommendationCategory *string `field:"optional" json:"recommendationCategory" yaml:"recommendationCategory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#recommendation_impact MonitorActivityLogAlert#recommendation_impact}.
	RecommendationImpact *string `field:"optional" json:"recommendationImpact" yaml:"recommendationImpact"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#recommendation_type MonitorActivityLogAlert#recommendation_type}.
	RecommendationType *string `field:"optional" json:"recommendationType" yaml:"recommendationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_group MonitorActivityLogAlert#resource_group}.
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// resource_health block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_health MonitorActivityLogAlert#resource_health}
	ResourceHealth interface{} `field:"optional" json:"resourceHealth" yaml:"resourceHealth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_id MonitorActivityLogAlert#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_provider MonitorActivityLogAlert#resource_provider}.
	ResourceProvider *string `field:"optional" json:"resourceProvider" yaml:"resourceProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#resource_type MonitorActivityLogAlert#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// service_health block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#service_health MonitorActivityLogAlert#service_health}
	ServiceHealth interface{} `field:"optional" json:"serviceHealth" yaml:"serviceHealth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#status MonitorActivityLogAlert#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#sub_status MonitorActivityLogAlert#sub_status}.
	SubStatus *string `field:"optional" json:"subStatus" yaml:"subStatus"`
}

type MonitorActivityLogAlertCriteriaOutputReference interface {
	cdktf.ComplexObject
	Caller() *string
	SetCaller(val *string)
	CallerInput() *string
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
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
	InternalValue() *MonitorActivityLogAlertCriteria
	SetInternalValue(val *MonitorActivityLogAlertCriteria)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	RecommendationCategory() *string
	SetRecommendationCategory(val *string)
	RecommendationCategoryInput() *string
	RecommendationImpact() *string
	SetRecommendationImpact(val *string)
	RecommendationImpactInput() *string
	RecommendationType() *string
	SetRecommendationType(val *string)
	RecommendationTypeInput() *string
	ResourceGroup() *string
	SetResourceGroup(val *string)
	ResourceGroupInput() *string
	ResourceHealth() MonitorActivityLogAlertCriteriaResourceHealthList
	ResourceHealthInput() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceProvider() *string
	SetResourceProvider(val *string)
	ResourceProviderInput() *string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	ServiceHealth() MonitorActivityLogAlertCriteriaServiceHealthList
	ServiceHealthInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubStatus() *string
	SetSubStatus(val *string)
	SubStatusInput() *string
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
	PutResourceHealth(value interface{})
	PutServiceHealth(value interface{})
	ResetCaller()
	ResetLevel()
	ResetOperationName()
	ResetRecommendationCategory()
	ResetRecommendationImpact()
	ResetRecommendationType()
	ResetResourceGroup()
	ResetResourceHealth()
	ResetResourceId()
	ResetResourceProvider()
	ResetResourceType()
	ResetServiceHealth()
	ResetStatus()
	ResetSubStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaOutputReference
type jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Caller() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CallerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InternalValue() *MonitorActivityLogAlertCriteria {
	var returns *MonitorActivityLogAlertCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationImpact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationImpactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationImpactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceHealth() MonitorActivityLogAlertCriteriaResourceHealthList {
	var returns MonitorActivityLogAlertCriteriaResourceHealthList
	_jsii_.Get(
		j,
		"resourceHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ServiceHealth() MonitorActivityLogAlertCriteriaServiceHealthList {
	var returns MonitorActivityLogAlertCriteriaServiceHealthList
	_jsii_.Get(
		j,
		"serviceHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ServiceHealthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActivityLogAlertCriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaOutputReference_Override(m MonitorActivityLogAlertCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetCaller(val *string) {
	_jsii_.Set(
		j,
		"caller",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetCategory(val *string) {
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetInternalValue(val *MonitorActivityLogAlertCriteria) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetLevel(val *string) {
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetOperationName(val *string) {
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetRecommendationCategory(val *string) {
	_jsii_.Set(
		j,
		"recommendationCategory",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetRecommendationImpact(val *string) {
	_jsii_.Set(
		j,
		"recommendationImpact",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetRecommendationType(val *string) {
	_jsii_.Set(
		j,
		"recommendationType",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetResourceGroup(val *string) {
	_jsii_.Set(
		j,
		"resourceGroup",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetResourceProvider(val *string) {
	_jsii_.Set(
		j,
		"resourceProvider",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetResourceType(val *string) {
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetSubStatus(val *string) {
	_jsii_.Set(
		j,
		"subStatus",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) PutResourceHealth(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putResourceHealth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) PutServiceHealth(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putServiceHealth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetCaller() {
	_jsii_.InvokeVoid(
		m,
		"resetCaller",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		m,
		"resetOperationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationCategory() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationCategory",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationImpact() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationImpact",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationType() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceHealth() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceHealth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceProvider() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceProvider",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetServiceHealth() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceHealth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetSubStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetSubStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertCriteriaResourceHealth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#current MonitorActivityLogAlert#current}.
	Current *[]*string `field:"optional" json:"current" yaml:"current"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#previous MonitorActivityLogAlert#previous}.
	Previous *[]*string `field:"optional" json:"previous" yaml:"previous"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#reason MonitorActivityLogAlert#reason}.
	Reason *[]*string `field:"optional" json:"reason" yaml:"reason"`
}

type MonitorActivityLogAlertCriteriaResourceHealthList interface {
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
	Get(index *float64) MonitorActivityLogAlertCriteriaResourceHealthOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaResourceHealthList
type jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaResourceHealthList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorActivityLogAlertCriteriaResourceHealthList {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaResourceHealthList_Override(m MonitorActivityLogAlertCriteriaResourceHealthList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) Get(index *float64) MonitorActivityLogAlertCriteriaResourceHealthOutputReference {
	var returns MonitorActivityLogAlertCriteriaResourceHealthOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertCriteriaResourceHealthOutputReference interface {
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
	Current() *[]*string
	SetCurrent(val *[]*string)
	CurrentInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Previous() *[]*string
	SetPrevious(val *[]*string)
	PreviousInput() *[]*string
	Reason() *[]*string
	SetReason(val *[]*string)
	ReasonInput() *[]*string
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
	ResetCurrent()
	ResetPrevious()
	ResetReason()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaResourceHealthOutputReference
type jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Current() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"current",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) CurrentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"currentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Previous() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) PreviousInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Reason() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ReasonInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"reasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaResourceHealthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorActivityLogAlertCriteriaResourceHealthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaResourceHealthOutputReference_Override(m MonitorActivityLogAlertCriteriaResourceHealthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaResourceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetCurrent(val *[]*string) {
	_jsii_.Set(
		j,
		"current",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetPrevious(val *[]*string) {
	_jsii_.Set(
		j,
		"previous",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetReason(val *[]*string) {
	_jsii_.Set(
		j,
		"reason",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetCurrent() {
	_jsii_.InvokeVoid(
		m,
		"resetCurrent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetPrevious() {
	_jsii_.InvokeVoid(
		m,
		"resetPrevious",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ResetReason() {
	_jsii_.InvokeVoid(
		m,
		"resetReason",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaResourceHealthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertCriteriaServiceHealth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#events MonitorActivityLogAlert#events}.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#locations MonitorActivityLogAlert#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#services MonitorActivityLogAlert#services}.
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

type MonitorActivityLogAlertCriteriaServiceHealthList interface {
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
	Get(index *float64) MonitorActivityLogAlertCriteriaServiceHealthOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaServiceHealthList
type jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaServiceHealthList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MonitorActivityLogAlertCriteriaServiceHealthList {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaServiceHealthList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaServiceHealthList_Override(m MonitorActivityLogAlertCriteriaServiceHealthList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaServiceHealthList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) Get(index *float64) MonitorActivityLogAlertCriteriaServiceHealthOutputReference {
	var returns MonitorActivityLogAlertCriteriaServiceHealthOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertCriteriaServiceHealthOutputReference interface {
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
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locations() *[]*string
	SetLocations(val *[]*string)
	LocationsInput() *[]*string
	Services() *[]*string
	SetServices(val *[]*string)
	ServicesInput() *[]*string
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
	ResetEvents()
	ResetLocations()
	ResetServices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaServiceHealthOutputReference
type jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) LocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) Services() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"services",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaServiceHealthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorActivityLogAlertCriteriaServiceHealthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaServiceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaServiceHealthOutputReference_Override(m MonitorActivityLogAlertCriteriaServiceHealthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaServiceHealthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetLocations(val *[]*string) {
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetServices(val *[]*string) {
	_jsii_.Set(
		j,
		"services",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		m,
		"resetEvents",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		m,
		"resetLocations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ResetServices() {
	_jsii_.InvokeVoid(
		m,
		"resetServices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaServiceHealthOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MonitorActivityLogAlertTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#create MonitorActivityLogAlert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#delete MonitorActivityLogAlert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#read MonitorActivityLogAlert#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/monitor_activity_log_alert#update MonitorActivityLogAlert#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MonitorActivityLogAlertTimeoutsOutputReference interface {
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

// The jsii proxy struct for MonitorActivityLogAlertTimeoutsOutputReference
type jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActivityLogAlertTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertTimeoutsOutputReference_Override(m MonitorActivityLogAlertTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

