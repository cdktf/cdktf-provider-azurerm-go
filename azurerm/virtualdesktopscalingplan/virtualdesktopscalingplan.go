package virtualdesktopscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/virtualdesktopscalingplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan azurerm_virtual_desktop_scaling_plan}.
type VirtualDesktopScalingPlan interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExclusionTag() *string
	SetExclusionTag(val *string)
	ExclusionTagInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	FriendlyName() *string
	SetFriendlyName(val *string)
	FriendlyNameInput() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostPool() VirtualDesktopScalingPlanHostPoolList
	HostPoolInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	Schedule() VirtualDesktopScalingPlanScheduleList
	ScheduleInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualDesktopScalingPlanTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutHostPool(value interface{})
	PutSchedule(value interface{})
	PutTimeouts(value *VirtualDesktopScalingPlanTimeouts)
	ResetDescription()
	ResetExclusionTag()
	ResetFriendlyName()
	ResetHostPool()
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

// The jsii proxy struct for VirtualDesktopScalingPlan
type jsiiProxy_VirtualDesktopScalingPlan struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ExclusionTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ExclusionTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) FriendlyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) FriendlyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) HostPool() VirtualDesktopScalingPlanHostPoolList {
	var returns VirtualDesktopScalingPlanHostPoolList
	_jsii_.Get(
		j,
		"hostPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) HostPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Schedule() VirtualDesktopScalingPlanScheduleList {
	var returns VirtualDesktopScalingPlanScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) Timeouts() VirtualDesktopScalingPlanTimeoutsOutputReference {
	var returns VirtualDesktopScalingPlanTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan azurerm_virtual_desktop_scaling_plan} Resource.
func NewVirtualDesktopScalingPlan(scope constructs.Construct, id *string, config *VirtualDesktopScalingPlanConfig) VirtualDesktopScalingPlan {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlan{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlan",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan azurerm_virtual_desktop_scaling_plan} Resource.
func NewVirtualDesktopScalingPlan_Override(v VirtualDesktopScalingPlan, scope constructs.Construct, id *string, config *VirtualDesktopScalingPlanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlan",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetExclusionTag(val *string) {
	_jsii_.Set(
		j,
		"exclusionTag",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetFriendlyName(val *string) {
	_jsii_.Set(
		j,
		"friendlyName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlan) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
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
func VirtualDesktopScalingPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualDesktopScalingPlan_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlan",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) PutHostPool(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putHostPool",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) PutSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		v,
		"putSchedule",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) PutTimeouts(value *VirtualDesktopScalingPlanTimeouts) {
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetDescription() {
	_jsii_.InvokeVoid(
		v,
		"resetDescription",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetExclusionTag() {
	_jsii_.InvokeVoid(
		v,
		"resetExclusionTag",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetFriendlyName() {
	_jsii_.InvokeVoid(
		v,
		"resetFriendlyName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetHostPool() {
	_jsii_.InvokeVoid(
		v,
		"resetHostPool",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlan) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopScalingPlanConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#location VirtualDesktopScalingPlan#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#name VirtualDesktopScalingPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#resource_group_name VirtualDesktopScalingPlan#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#schedule VirtualDesktopScalingPlan#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#time_zone VirtualDesktopScalingPlan#time_zone}.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#description VirtualDesktopScalingPlan#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#exclusion_tag VirtualDesktopScalingPlan#exclusion_tag}.
	ExclusionTag *string `field:"optional" json:"exclusionTag" yaml:"exclusionTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#friendly_name VirtualDesktopScalingPlan#friendly_name}.
	FriendlyName *string `field:"optional" json:"friendlyName" yaml:"friendlyName"`
	// host_pool block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#host_pool VirtualDesktopScalingPlan#host_pool}
	HostPool interface{} `field:"optional" json:"hostPool" yaml:"hostPool"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#id VirtualDesktopScalingPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#tags VirtualDesktopScalingPlan#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#timeouts VirtualDesktopScalingPlan#timeouts}
	Timeouts *VirtualDesktopScalingPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type VirtualDesktopScalingPlanHostPool struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#hostpool_id VirtualDesktopScalingPlan#hostpool_id}.
	HostpoolId *string `field:"required" json:"hostpoolId" yaml:"hostpoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#scaling_plan_enabled VirtualDesktopScalingPlan#scaling_plan_enabled}.
	ScalingPlanEnabled interface{} `field:"required" json:"scalingPlanEnabled" yaml:"scalingPlanEnabled"`
}

type VirtualDesktopScalingPlanHostPoolList interface {
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
	Get(index *float64) VirtualDesktopScalingPlanHostPoolOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopScalingPlanHostPoolList
type jsiiProxy_VirtualDesktopScalingPlanHostPoolList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanHostPoolList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualDesktopScalingPlanHostPoolList {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlanHostPoolList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanHostPoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanHostPoolList_Override(v VirtualDesktopScalingPlanHostPoolList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanHostPoolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) Get(index *float64) VirtualDesktopScalingPlanHostPoolOutputReference {
	var returns VirtualDesktopScalingPlanHostPoolOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopScalingPlanHostPoolOutputReference interface {
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
	HostpoolId() *string
	SetHostpoolId(val *string)
	HostpoolIdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ScalingPlanEnabled() interface{}
	SetScalingPlanEnabled(val interface{})
	ScalingPlanEnabledInput() interface{}
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

// The jsii proxy struct for VirtualDesktopScalingPlanHostPoolOutputReference
type jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) HostpoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostpoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) HostpoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostpoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ScalingPlanEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPlanEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ScalingPlanEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingPlanEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanHostPoolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualDesktopScalingPlanHostPoolOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanHostPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanHostPoolOutputReference_Override(v VirtualDesktopScalingPlanHostPoolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanHostPoolOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetHostpoolId(val *string) {
	_jsii_.Set(
		j,
		"hostpoolId",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetScalingPlanEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"scalingPlanEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanHostPoolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopScalingPlanSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#days_of_week VirtualDesktopScalingPlan#days_of_week}.
	DaysOfWeek *[]*string `field:"required" json:"daysOfWeek" yaml:"daysOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#name VirtualDesktopScalingPlan#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#off_peak_load_balancing_algorithm VirtualDesktopScalingPlan#off_peak_load_balancing_algorithm}.
	OffPeakLoadBalancingAlgorithm *string `field:"required" json:"offPeakLoadBalancingAlgorithm" yaml:"offPeakLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#off_peak_start_time VirtualDesktopScalingPlan#off_peak_start_time}.
	OffPeakStartTime *string `field:"required" json:"offPeakStartTime" yaml:"offPeakStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#peak_load_balancing_algorithm VirtualDesktopScalingPlan#peak_load_balancing_algorithm}.
	PeakLoadBalancingAlgorithm *string `field:"required" json:"peakLoadBalancingAlgorithm" yaml:"peakLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#peak_start_time VirtualDesktopScalingPlan#peak_start_time}.
	PeakStartTime *string `field:"required" json:"peakStartTime" yaml:"peakStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_capacity_threshold_percent VirtualDesktopScalingPlan#ramp_down_capacity_threshold_percent}.
	RampDownCapacityThresholdPercent *float64 `field:"required" json:"rampDownCapacityThresholdPercent" yaml:"rampDownCapacityThresholdPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_force_logoff_users VirtualDesktopScalingPlan#ramp_down_force_logoff_users}.
	RampDownForceLogoffUsers interface{} `field:"required" json:"rampDownForceLogoffUsers" yaml:"rampDownForceLogoffUsers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_load_balancing_algorithm VirtualDesktopScalingPlan#ramp_down_load_balancing_algorithm}.
	RampDownLoadBalancingAlgorithm *string `field:"required" json:"rampDownLoadBalancingAlgorithm" yaml:"rampDownLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_minimum_hosts_percent VirtualDesktopScalingPlan#ramp_down_minimum_hosts_percent}.
	RampDownMinimumHostsPercent *float64 `field:"required" json:"rampDownMinimumHostsPercent" yaml:"rampDownMinimumHostsPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_notification_message VirtualDesktopScalingPlan#ramp_down_notification_message}.
	RampDownNotificationMessage *string `field:"required" json:"rampDownNotificationMessage" yaml:"rampDownNotificationMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_start_time VirtualDesktopScalingPlan#ramp_down_start_time}.
	RampDownStartTime *string `field:"required" json:"rampDownStartTime" yaml:"rampDownStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_stop_hosts_when VirtualDesktopScalingPlan#ramp_down_stop_hosts_when}.
	RampDownStopHostsWhen *string `field:"required" json:"rampDownStopHostsWhen" yaml:"rampDownStopHostsWhen"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_down_wait_time_minutes VirtualDesktopScalingPlan#ramp_down_wait_time_minutes}.
	RampDownWaitTimeMinutes *float64 `field:"required" json:"rampDownWaitTimeMinutes" yaml:"rampDownWaitTimeMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_up_load_balancing_algorithm VirtualDesktopScalingPlan#ramp_up_load_balancing_algorithm}.
	RampUpLoadBalancingAlgorithm *string `field:"required" json:"rampUpLoadBalancingAlgorithm" yaml:"rampUpLoadBalancingAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_up_start_time VirtualDesktopScalingPlan#ramp_up_start_time}.
	RampUpStartTime *string `field:"required" json:"rampUpStartTime" yaml:"rampUpStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_up_capacity_threshold_percent VirtualDesktopScalingPlan#ramp_up_capacity_threshold_percent}.
	RampUpCapacityThresholdPercent *float64 `field:"optional" json:"rampUpCapacityThresholdPercent" yaml:"rampUpCapacityThresholdPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#ramp_up_minimum_hosts_percent VirtualDesktopScalingPlan#ramp_up_minimum_hosts_percent}.
	RampUpMinimumHostsPercent *float64 `field:"optional" json:"rampUpMinimumHostsPercent" yaml:"rampUpMinimumHostsPercent"`
}

type VirtualDesktopScalingPlanScheduleList interface {
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
	Get(index *float64) VirtualDesktopScalingPlanScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopScalingPlanScheduleList
type jsiiProxy_VirtualDesktopScalingPlanScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) VirtualDesktopScalingPlanScheduleList {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlanScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanScheduleList_Override(v VirtualDesktopScalingPlanScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleList) Get(index *float64) VirtualDesktopScalingPlanScheduleOutputReference {
	var returns VirtualDesktopScalingPlanScheduleOutputReference

	_jsii_.Invoke(
		v,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopScalingPlanScheduleOutputReference interface {
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OffPeakLoadBalancingAlgorithm() *string
	SetOffPeakLoadBalancingAlgorithm(val *string)
	OffPeakLoadBalancingAlgorithmInput() *string
	OffPeakStartTime() *string
	SetOffPeakStartTime(val *string)
	OffPeakStartTimeInput() *string
	PeakLoadBalancingAlgorithm() *string
	SetPeakLoadBalancingAlgorithm(val *string)
	PeakLoadBalancingAlgorithmInput() *string
	PeakStartTime() *string
	SetPeakStartTime(val *string)
	PeakStartTimeInput() *string
	RampDownCapacityThresholdPercent() *float64
	SetRampDownCapacityThresholdPercent(val *float64)
	RampDownCapacityThresholdPercentInput() *float64
	RampDownForceLogoffUsers() interface{}
	SetRampDownForceLogoffUsers(val interface{})
	RampDownForceLogoffUsersInput() interface{}
	RampDownLoadBalancingAlgorithm() *string
	SetRampDownLoadBalancingAlgorithm(val *string)
	RampDownLoadBalancingAlgorithmInput() *string
	RampDownMinimumHostsPercent() *float64
	SetRampDownMinimumHostsPercent(val *float64)
	RampDownMinimumHostsPercentInput() *float64
	RampDownNotificationMessage() *string
	SetRampDownNotificationMessage(val *string)
	RampDownNotificationMessageInput() *string
	RampDownStartTime() *string
	SetRampDownStartTime(val *string)
	RampDownStartTimeInput() *string
	RampDownStopHostsWhen() *string
	SetRampDownStopHostsWhen(val *string)
	RampDownStopHostsWhenInput() *string
	RampDownWaitTimeMinutes() *float64
	SetRampDownWaitTimeMinutes(val *float64)
	RampDownWaitTimeMinutesInput() *float64
	RampUpCapacityThresholdPercent() *float64
	SetRampUpCapacityThresholdPercent(val *float64)
	RampUpCapacityThresholdPercentInput() *float64
	RampUpLoadBalancingAlgorithm() *string
	SetRampUpLoadBalancingAlgorithm(val *string)
	RampUpLoadBalancingAlgorithmInput() *string
	RampUpMinimumHostsPercent() *float64
	SetRampUpMinimumHostsPercent(val *float64)
	RampUpMinimumHostsPercentInput() *float64
	RampUpStartTime() *string
	SetRampUpStartTime(val *string)
	RampUpStartTimeInput() *string
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
	ResetRampUpCapacityThresholdPercent()
	ResetRampUpMinimumHostsPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopScalingPlanScheduleOutputReference
type jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownCapacityThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownCapacityThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownCapacityThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownCapacityThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownForceLogoffUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rampDownForceLogoffUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownForceLogoffUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rampDownForceLogoffUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownMinimumHostsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownMinimumHostsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownMinimumHostsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownMinimumHostsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownNotificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownNotificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownNotificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownNotificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStopHostsWhen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStopHostsWhen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStopHostsWhenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStopHostsWhenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownWaitTimeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownWaitTimeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownWaitTimeMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownWaitTimeMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpCapacityThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpCapacityThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpCapacityThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpCapacityThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpMinimumHostsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpMinimumHostsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpMinimumHostsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpMinimumHostsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualDesktopScalingPlanScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanScheduleOutputReference_Override(v VirtualDesktopScalingPlanScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetDaysOfWeek(val *[]*string) {
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetOffPeakLoadBalancingAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"offPeakLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetOffPeakStartTime(val *string) {
	_jsii_.Set(
		j,
		"offPeakStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetPeakLoadBalancingAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"peakLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetPeakStartTime(val *string) {
	_jsii_.Set(
		j,
		"peakStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownCapacityThresholdPercent(val *float64) {
	_jsii_.Set(
		j,
		"rampDownCapacityThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownForceLogoffUsers(val interface{}) {
	_jsii_.Set(
		j,
		"rampDownForceLogoffUsers",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownLoadBalancingAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"rampDownLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownMinimumHostsPercent(val *float64) {
	_jsii_.Set(
		j,
		"rampDownMinimumHostsPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownNotificationMessage(val *string) {
	_jsii_.Set(
		j,
		"rampDownNotificationMessage",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownStartTime(val *string) {
	_jsii_.Set(
		j,
		"rampDownStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownStopHostsWhen(val *string) {
	_jsii_.Set(
		j,
		"rampDownStopHostsWhen",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampDownWaitTimeMinutes(val *float64) {
	_jsii_.Set(
		j,
		"rampDownWaitTimeMinutes",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampUpCapacityThresholdPercent(val *float64) {
	_jsii_.Set(
		j,
		"rampUpCapacityThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampUpLoadBalancingAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"rampUpLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampUpMinimumHostsPercent(val *float64) {
	_jsii_.Set(
		j,
		"rampUpMinimumHostsPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetRampUpStartTime(val *string) {
	_jsii_.Set(
		j,
		"rampUpStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ResetRampUpCapacityThresholdPercent() {
	_jsii_.InvokeVoid(
		v,
		"resetRampUpCapacityThresholdPercent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ResetRampUpMinimumHostsPercent() {
	_jsii_.InvokeVoid(
		v,
		"resetRampUpMinimumHostsPercent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type VirtualDesktopScalingPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#create VirtualDesktopScalingPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#delete VirtualDesktopScalingPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#read VirtualDesktopScalingPlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_scaling_plan#update VirtualDesktopScalingPlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type VirtualDesktopScalingPlanTimeoutsOutputReference interface {
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

// The jsii proxy struct for VirtualDesktopScalingPlanTimeoutsOutputReference
type jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualDesktopScalingPlanTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanTimeoutsOutputReference_Override(v VirtualDesktopScalingPlanTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		v,
		"resetCreate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		v,
		"resetDelete",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		v,
		"resetRead",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		v,
		"resetUpdate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

