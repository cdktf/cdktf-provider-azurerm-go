package labservicelab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/labservicelab/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/lab_service_lab azurerm_lab_service_lab}.
type LabServiceLab interface {
	cdktf.TerraformResource
	AutoShutdown() LabServiceLabAutoShutdownOutputReference
	AutoShutdownInput() *LabServiceLabAutoShutdown
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionSetting() LabServiceLabConnectionSettingOutputReference
	ConnectionSettingInput() *LabServiceLabConnectionSetting
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	LabPlanId() *string
	SetLabPlanId(val *string)
	LabPlanIdInput() *string
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
	Network() LabServiceLabNetworkOutputReference
	NetworkInput() *LabServiceLabNetwork
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
	Roster() LabServiceLabRosterOutputReference
	RosterInput() *LabServiceLabRoster
	Security() LabServiceLabSecurityOutputReference
	SecurityInput() *LabServiceLabSecurity
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LabServiceLabTimeoutsOutputReference
	TimeoutsInput() interface{}
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	VirtualMachine() LabServiceLabVirtualMachineOutputReference
	VirtualMachineInput() *LabServiceLabVirtualMachine
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
	PutAutoShutdown(value *LabServiceLabAutoShutdown)
	PutConnectionSetting(value *LabServiceLabConnectionSetting)
	PutNetwork(value *LabServiceLabNetwork)
	PutRoster(value *LabServiceLabRoster)
	PutSecurity(value *LabServiceLabSecurity)
	PutTimeouts(value *LabServiceLabTimeouts)
	PutVirtualMachine(value *LabServiceLabVirtualMachine)
	ResetAutoShutdown()
	ResetConnectionSetting()
	ResetDescription()
	ResetId()
	ResetLabPlanId()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRoster()
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

// The jsii proxy struct for LabServiceLab
type jsiiProxy_LabServiceLab struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LabServiceLab) AutoShutdown() LabServiceLabAutoShutdownOutputReference {
	var returns LabServiceLabAutoShutdownOutputReference
	_jsii_.Get(
		j,
		"autoShutdown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) AutoShutdownInput() *LabServiceLabAutoShutdown {
	var returns *LabServiceLabAutoShutdown
	_jsii_.Get(
		j,
		"autoShutdownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ConnectionSetting() LabServiceLabConnectionSettingOutputReference {
	var returns LabServiceLabConnectionSettingOutputReference
	_jsii_.Get(
		j,
		"connectionSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ConnectionSettingInput() *LabServiceLabConnectionSetting {
	var returns *LabServiceLabConnectionSetting
	_jsii_.Get(
		j,
		"connectionSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) LabPlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labPlanId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) LabPlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labPlanIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Network() LabServiceLabNetworkOutputReference {
	var returns LabServiceLabNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) NetworkInput() *LabServiceLabNetwork {
	var returns *LabServiceLabNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Roster() LabServiceLabRosterOutputReference {
	var returns LabServiceLabRosterOutputReference
	_jsii_.Get(
		j,
		"roster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) RosterInput() *LabServiceLabRoster {
	var returns *LabServiceLabRoster
	_jsii_.Get(
		j,
		"rosterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Security() LabServiceLabSecurityOutputReference {
	var returns LabServiceLabSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) SecurityInput() *LabServiceLabSecurity {
	var returns *LabServiceLabSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Timeouts() LabServiceLabTimeoutsOutputReference {
	var returns LabServiceLabTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) VirtualMachine() LabServiceLabVirtualMachineOutputReference {
	var returns LabServiceLabVirtualMachineOutputReference
	_jsii_.Get(
		j,
		"virtualMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLab) VirtualMachineInput() *LabServiceLabVirtualMachine {
	var returns *LabServiceLabVirtualMachine
	_jsii_.Get(
		j,
		"virtualMachineInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/lab_service_lab azurerm_lab_service_lab} Resource.
func NewLabServiceLab(scope constructs.Construct, id *string, config *LabServiceLabConfig) LabServiceLab {
	_init_.Initialize()

	if err := validateNewLabServiceLabParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LabServiceLab{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/lab_service_lab azurerm_lab_service_lab} Resource.
func NewLabServiceLab_Override(l LabServiceLab, scope constructs.Construct, id *string, config *LabServiceLabConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LabServiceLab)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetLabPlanId(val *string) {
	if err := j.validateSetLabPlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labPlanId",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LabServiceLab)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
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
func LabServiceLab_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLabServiceLab_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LabServiceLab_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLabServiceLab_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LabServiceLab_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLabServiceLab_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LabServiceLab_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLab",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LabServiceLab) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LabServiceLab) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LabServiceLab) PutAutoShutdown(value *LabServiceLabAutoShutdown) {
	if err := l.validatePutAutoShutdownParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAutoShutdown",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutConnectionSetting(value *LabServiceLabConnectionSetting) {
	if err := l.validatePutConnectionSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConnectionSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutNetwork(value *LabServiceLabNetwork) {
	if err := l.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNetwork",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutRoster(value *LabServiceLabRoster) {
	if err := l.validatePutRosterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRoster",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutSecurity(value *LabServiceLabSecurity) {
	if err := l.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSecurity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutTimeouts(value *LabServiceLabTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) PutVirtualMachine(value *LabServiceLabVirtualMachine) {
	if err := l.validatePutVirtualMachineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVirtualMachine",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLab) ResetAutoShutdown() {
	_jsii_.InvokeVoid(
		l,
		"resetAutoShutdown",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetConnectionSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetConnectionSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetLabPlanId() {
	_jsii_.InvokeVoid(
		l,
		"resetLabPlanId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetNetwork() {
	_jsii_.InvokeVoid(
		l,
		"resetNetwork",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetRoster() {
	_jsii_.InvokeVoid(
		l,
		"resetRoster",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLab) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLab) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

