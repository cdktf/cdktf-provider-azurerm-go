package automationsoftwareupdateconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/automationsoftwareupdateconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration azurerm_automation_software_update_configuration}.
type AutomationSoftwareUpdateConfiguration interface {
	cdktf.TerraformResource
	AutomationAccountId() *string
	SetAutomationAccountId(val *string)
	AutomationAccountIdInput() *string
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
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	ErrorCode() *string
	ErrorMeesage() *string
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
	Linux() AutomationSoftwareUpdateConfigurationLinuxList
	LinuxInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NonAzureComputerNames() *[]*string
	SetNonAzureComputerNames(val *[]*string)
	NonAzureComputerNamesInput() *[]*string
	OperatingSystem() *string
	SetOperatingSystem(val *string)
	OperatingSystemInput() *string
	PostTask() AutomationSoftwareUpdateConfigurationPostTaskList
	PostTaskInput() interface{}
	PreTask() AutomationSoftwareUpdateConfigurationPreTaskList
	PreTaskInput() interface{}
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
	Schedule() AutomationSoftwareUpdateConfigurationScheduleList
	ScheduleInput() interface{}
	Target() AutomationSoftwareUpdateConfigurationTargetOutputReference
	TargetInput() *AutomationSoftwareUpdateConfigurationTarget
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AutomationSoftwareUpdateConfigurationTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineIds() *[]*string
	SetVirtualMachineIds(val *[]*string)
	VirtualMachineIdsInput() *[]*string
	Windows() AutomationSoftwareUpdateConfigurationWindowsOutputReference
	WindowsInput() *AutomationSoftwareUpdateConfigurationWindows
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
	PutLinux(value interface{})
	PutPostTask(value interface{})
	PutPreTask(value interface{})
	PutSchedule(value interface{})
	PutTarget(value *AutomationSoftwareUpdateConfigurationTarget)
	PutTimeouts(value *AutomationSoftwareUpdateConfigurationTimeouts)
	PutWindows(value *AutomationSoftwareUpdateConfigurationWindows)
	ResetDuration()
	ResetId()
	ResetLinux()
	ResetNonAzureComputerNames()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostTask()
	ResetPreTask()
	ResetSchedule()
	ResetTarget()
	ResetTimeouts()
	ResetVirtualMachineIds()
	ResetWindows()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AutomationSoftwareUpdateConfiguration
type jsiiProxy_AutomationSoftwareUpdateConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) AutomationAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) AutomationAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automationAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ErrorCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ErrorMeesage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMeesage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Linux() AutomationSoftwareUpdateConfigurationLinuxList {
	var returns AutomationSoftwareUpdateConfigurationLinuxList
	_jsii_.Get(
		j,
		"linux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) LinuxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) NonAzureComputerNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonAzureComputerNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) NonAzureComputerNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonAzureComputerNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) OperatingSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) OperatingSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PostTask() AutomationSoftwareUpdateConfigurationPostTaskList {
	var returns AutomationSoftwareUpdateConfigurationPostTaskList
	_jsii_.Get(
		j,
		"postTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PostTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PreTask() AutomationSoftwareUpdateConfigurationPreTaskList {
	var returns AutomationSoftwareUpdateConfigurationPreTaskList
	_jsii_.Get(
		j,
		"preTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PreTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Schedule() AutomationSoftwareUpdateConfigurationScheduleList {
	var returns AutomationSoftwareUpdateConfigurationScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Target() AutomationSoftwareUpdateConfigurationTargetOutputReference {
	var returns AutomationSoftwareUpdateConfigurationTargetOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) TargetInput() *AutomationSoftwareUpdateConfigurationTarget {
	var returns *AutomationSoftwareUpdateConfigurationTarget
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Timeouts() AutomationSoftwareUpdateConfigurationTimeoutsOutputReference {
	var returns AutomationSoftwareUpdateConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) VirtualMachineIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualMachineIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) VirtualMachineIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"virtualMachineIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Windows() AutomationSoftwareUpdateConfigurationWindowsOutputReference {
	var returns AutomationSoftwareUpdateConfigurationWindowsOutputReference
	_jsii_.Get(
		j,
		"windows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) WindowsInput() *AutomationSoftwareUpdateConfigurationWindows {
	var returns *AutomationSoftwareUpdateConfigurationWindows
	_jsii_.Get(
		j,
		"windowsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration azurerm_automation_software_update_configuration} Resource.
func NewAutomationSoftwareUpdateConfiguration(scope constructs.Construct, id *string, config *AutomationSoftwareUpdateConfigurationConfig) AutomationSoftwareUpdateConfiguration {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration azurerm_automation_software_update_configuration} Resource.
func NewAutomationSoftwareUpdateConfiguration_Override(a AutomationSoftwareUpdateConfiguration, scope constructs.Construct, id *string, config *AutomationSoftwareUpdateConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetAutomationAccountId(val *string) {
	_jsii_.Set(
		j,
		"automationAccountId",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetDuration(val *string) {
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetNonAzureComputerNames(val *[]*string) {
	_jsii_.Set(
		j,
		"nonAzureComputerNames",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetOperatingSystem(val *string) {
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) SetVirtualMachineIds(val *[]*string) {
	_jsii_.Set(
		j,
		"virtualMachineIds",
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
func AutomationSoftwareUpdateConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AutomationSoftwareUpdateConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutLinux(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putLinux",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutPostTask(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putPostTask",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutPreTask(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putPreTask",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutTarget(value *AutomationSoftwareUpdateConfigurationTarget) {
	_jsii_.InvokeVoid(
		a,
		"putTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutTimeouts(value *AutomationSoftwareUpdateConfigurationTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutWindows(value *AutomationSoftwareUpdateConfigurationWindows) {
	_jsii_.InvokeVoid(
		a,
		"putWindows",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetLinux() {
	_jsii_.InvokeVoid(
		a,
		"resetLinux",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetNonAzureComputerNames() {
	_jsii_.InvokeVoid(
		a,
		"resetNonAzureComputerNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetPostTask() {
	_jsii_.InvokeVoid(
		a,
		"resetPostTask",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetPreTask() {
	_jsii_.InvokeVoid(
		a,
		"resetPreTask",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetVirtualMachineIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualMachineIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetWindows() {
	_jsii_.InvokeVoid(
		a,
		"resetWindows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#automation_account_id AutomationSoftwareUpdateConfiguration#automation_account_id}.
	AutomationAccountId *string `field:"required" json:"automationAccountId" yaml:"automationAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#name AutomationSoftwareUpdateConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#operating_system AutomationSoftwareUpdateConfiguration#operating_system}.
	OperatingSystem *string `field:"required" json:"operatingSystem" yaml:"operatingSystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#duration AutomationSoftwareUpdateConfiguration#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#id AutomationSoftwareUpdateConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// linux block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#linux AutomationSoftwareUpdateConfiguration#linux}
	Linux interface{} `field:"optional" json:"linux" yaml:"linux"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#non_azure_computer_names AutomationSoftwareUpdateConfiguration#non_azure_computer_names}.
	NonAzureComputerNames *[]*string `field:"optional" json:"nonAzureComputerNames" yaml:"nonAzureComputerNames"`
	// post_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#post_task AutomationSoftwareUpdateConfiguration#post_task}
	PostTask interface{} `field:"optional" json:"postTask" yaml:"postTask"`
	// pre_task block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#pre_task AutomationSoftwareUpdateConfiguration#pre_task}
	PreTask interface{} `field:"optional" json:"preTask" yaml:"preTask"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#schedule AutomationSoftwareUpdateConfiguration#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#target AutomationSoftwareUpdateConfiguration#target}
	Target *AutomationSoftwareUpdateConfigurationTarget `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#timeouts AutomationSoftwareUpdateConfiguration#timeouts}
	Timeouts *AutomationSoftwareUpdateConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#virtual_machine_ids AutomationSoftwareUpdateConfiguration#virtual_machine_ids}.
	VirtualMachineIds *[]*string `field:"optional" json:"virtualMachineIds" yaml:"virtualMachineIds"`
	// windows block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#windows AutomationSoftwareUpdateConfiguration#windows}
	Windows *AutomationSoftwareUpdateConfigurationWindows `field:"optional" json:"windows" yaml:"windows"`
}

type AutomationSoftwareUpdateConfigurationLinux struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#classification_included AutomationSoftwareUpdateConfiguration#classification_included}.
	ClassificationIncluded *string `field:"optional" json:"classificationIncluded" yaml:"classificationIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#excluded_packages AutomationSoftwareUpdateConfiguration#excluded_packages}.
	ExcludedPackages *[]*string `field:"optional" json:"excludedPackages" yaml:"excludedPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#included_packages AutomationSoftwareUpdateConfiguration#included_packages}.
	IncludedPackages *[]*string `field:"optional" json:"includedPackages" yaml:"includedPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#reboot AutomationSoftwareUpdateConfiguration#reboot}.
	Reboot *string `field:"optional" json:"reboot" yaml:"reboot"`
}

type AutomationSoftwareUpdateConfigurationLinuxList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationLinuxOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationLinuxList
type jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationLinuxList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationLinuxList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationLinuxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationLinuxList_Override(a AutomationSoftwareUpdateConfigurationLinuxList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationLinuxList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) Get(index *float64) AutomationSoftwareUpdateConfigurationLinuxOutputReference {
	var returns AutomationSoftwareUpdateConfigurationLinuxOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationLinuxOutputReference interface {
	cdktf.ComplexObject
	ClassificationIncluded() *string
	SetClassificationIncluded(val *string)
	ClassificationIncludedInput() *string
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
	ExcludedPackages() *[]*string
	SetExcludedPackages(val *[]*string)
	ExcludedPackagesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedPackages() *[]*string
	SetIncludedPackages(val *[]*string)
	IncludedPackagesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Reboot() *string
	SetReboot(val *string)
	RebootInput() *string
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
	ResetClassificationIncluded()
	ResetExcludedPackages()
	ResetIncludedPackages()
	ResetReboot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationLinuxOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ClassificationIncluded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ClassificationIncludedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ExcludedPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ExcludedPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) IncludedPackages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) IncludedPackagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) Reboot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) RebootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationLinuxOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationLinuxOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationLinuxOutputReference_Override(a AutomationSoftwareUpdateConfigurationLinuxOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationLinuxOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetClassificationIncluded(val *string) {
	_jsii_.Set(
		j,
		"classificationIncluded",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetExcludedPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedPackages",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetIncludedPackages(val *[]*string) {
	_jsii_.Set(
		j,
		"includedPackages",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetReboot(val *string) {
	_jsii_.Set(
		j,
		"reboot",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ResetClassificationIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetClassificationIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ResetExcludedPackages() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedPackages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ResetIncludedPackages() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludedPackages",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ResetReboot() {
	_jsii_.InvokeVoid(
		a,
		"resetReboot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationLinuxOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationPostTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#parameters AutomationSoftwareUpdateConfiguration#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#source AutomationSoftwareUpdateConfiguration#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

type AutomationSoftwareUpdateConfigurationPostTaskList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationPostTaskOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationPostTaskList
type jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationPostTaskList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationPostTaskList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPostTaskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationPostTaskList_Override(a AutomationSoftwareUpdateConfigurationPostTaskList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPostTaskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) Get(index *float64) AutomationSoftwareUpdateConfigurationPostTaskOutputReference {
	var returns AutomationSoftwareUpdateConfigurationPostTaskOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationPostTaskOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	ResetParameters()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationPostTaskOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationPostTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationPostTaskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPostTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationPostTaskOutputReference_Override(a AutomationSoftwareUpdateConfigurationPostTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPostTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPostTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationPreTask struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#parameters AutomationSoftwareUpdateConfiguration#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#source AutomationSoftwareUpdateConfiguration#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

type AutomationSoftwareUpdateConfigurationPreTaskList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationPreTaskOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationPreTaskList
type jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationPreTaskList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationPreTaskList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPreTaskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationPreTaskList_Override(a AutomationSoftwareUpdateConfigurationPreTaskList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPreTaskList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) Get(index *float64) AutomationSoftwareUpdateConfigurationPreTaskOutputReference {
	var returns AutomationSoftwareUpdateConfigurationPreTaskOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationPreTaskOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	ResetParameters()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationPreTaskOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationPreTaskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationPreTaskOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPreTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationPreTaskOutputReference_Override(a AutomationSoftwareUpdateConfigurationPreTaskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationPreTaskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetSource(val *string) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationPreTaskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#advanced_month_days AutomationSoftwareUpdateConfiguration#advanced_month_days}.
	AdvancedMonthDays *[]*float64 `field:"optional" json:"advancedMonthDays" yaml:"advancedMonthDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#advanced_week_days AutomationSoftwareUpdateConfiguration#advanced_week_days}.
	AdvancedWeekDays *[]*string `field:"optional" json:"advancedWeekDays" yaml:"advancedWeekDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#description AutomationSoftwareUpdateConfiguration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#expiry_time AutomationSoftwareUpdateConfiguration#expiry_time}.
	ExpiryTime *string `field:"optional" json:"expiryTime" yaml:"expiryTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#expiry_time_offset_minutes AutomationSoftwareUpdateConfiguration#expiry_time_offset_minutes}.
	ExpiryTimeOffsetMinutes *float64 `field:"optional" json:"expiryTimeOffsetMinutes" yaml:"expiryTimeOffsetMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#frequency AutomationSoftwareUpdateConfiguration#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#interval AutomationSoftwareUpdateConfiguration#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#is_enabled AutomationSoftwareUpdateConfiguration#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// monthly_occurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#monthly_occurrence AutomationSoftwareUpdateConfiguration#monthly_occurrence}
	MonthlyOccurrence interface{} `field:"optional" json:"monthlyOccurrence" yaml:"monthlyOccurrence"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#next_run AutomationSoftwareUpdateConfiguration#next_run}.
	NextRun *string `field:"optional" json:"nextRun" yaml:"nextRun"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#next_run_offset_minutes AutomationSoftwareUpdateConfiguration#next_run_offset_minutes}.
	NextRunOffsetMinutes *float64 `field:"optional" json:"nextRunOffsetMinutes" yaml:"nextRunOffsetMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#start_time AutomationSoftwareUpdateConfiguration#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#start_time_offset_minutes AutomationSoftwareUpdateConfiguration#start_time_offset_minutes}.
	StartTimeOffsetMinutes *float64 `field:"optional" json:"startTimeOffsetMinutes" yaml:"startTimeOffsetMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#time_zone AutomationSoftwareUpdateConfiguration#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

type AutomationSoftwareUpdateConfigurationScheduleList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationScheduleList
type jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationScheduleList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationScheduleList_Override(a AutomationSoftwareUpdateConfigurationScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) Get(index *float64) AutomationSoftwareUpdateConfigurationScheduleOutputReference {
	var returns AutomationSoftwareUpdateConfigurationScheduleOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrence struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#day AutomationSoftwareUpdateConfiguration#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#occurrence AutomationSoftwareUpdateConfiguration#occurrence}.
	Occurrence *float64 `field:"required" json:"occurrence" yaml:"occurrence"`
}

type AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList
type jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList_Override(a AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) Get(index *float64) AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference {
	var returns AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference interface {
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
	Day() *string
	SetDay(val *string)
	DayInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Occurrence() *float64
	SetOccurrence(val *float64)
	OccurrenceInput() *float64
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

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) Day() *string {
	var returns *string
	_jsii_.Get(
		j,
		"day",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) DayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) Occurrence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"occurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) OccurrenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"occurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference_Override(a AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetDay(val *string) {
	_jsii_.Set(
		j,
		"day",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetOccurrence(val *float64) {
	_jsii_.Set(
		j,
		"occurrence",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationScheduleOutputReference interface {
	cdktf.ComplexObject
	AdvancedMonthDays() *[]*float64
	SetAdvancedMonthDays(val *[]*float64)
	AdvancedMonthDaysInput() *[]*float64
	AdvancedWeekDays() *[]*string
	SetAdvancedWeekDays(val *[]*string)
	AdvancedWeekDaysInput() *[]*string
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
	CreationTime() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExpiryTime() *string
	SetExpiryTime(val *string)
	ExpiryTimeInput() *string
	ExpiryTimeOffsetMinutes() *float64
	SetExpiryTimeOffsetMinutes(val *float64)
	ExpiryTimeOffsetMinutesInput() *float64
	// Experimental.
	Fqn() *string
	Frequency() *string
	SetFrequency(val *string)
	FrequencyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IsEnabled() interface{}
	SetIsEnabled(val interface{})
	IsEnabledInput() interface{}
	LastModifiedTime() *string
	MonthlyOccurrence() AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList
	MonthlyOccurrenceInput() interface{}
	NextRun() *string
	SetNextRun(val *string)
	NextRunInput() *string
	NextRunOffsetMinutes() *float64
	SetNextRunOffsetMinutes(val *float64)
	NextRunOffsetMinutesInput() *float64
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	StartTimeOffsetMinutes() *float64
	SetStartTimeOffsetMinutes(val *float64)
	StartTimeOffsetMinutesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutMonthlyOccurrence(value interface{})
	ResetAdvancedMonthDays()
	ResetAdvancedWeekDays()
	ResetDescription()
	ResetExpiryTime()
	ResetExpiryTimeOffsetMinutes()
	ResetFrequency()
	ResetInterval()
	ResetIsEnabled()
	ResetMonthlyOccurrence()
	ResetNextRun()
	ResetNextRunOffsetMinutes()
	ResetStartTime()
	ResetStartTimeOffsetMinutes()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationScheduleOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedMonthDays() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"advancedMonthDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedMonthDaysInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"advancedMonthDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedWeekDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advancedWeekDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) AdvancedWeekDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"advancedWeekDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiryTimeOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ExpiryTimeOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expiryTimeOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Frequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) FrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) IsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) MonthlyOccurrence() AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList {
	var returns AutomationSoftwareUpdateConfigurationScheduleMonthlyOccurrenceList
	_jsii_.Get(
		j,
		"monthlyOccurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) MonthlyOccurrenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthlyOccurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nextRunOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) NextRunOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nextRunOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeOffsetMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeOffsetMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) StartTimeOffsetMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeOffsetMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationScheduleOutputReference_Override(a AutomationSoftwareUpdateConfigurationScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetAdvancedMonthDays(val *[]*float64) {
	_jsii_.Set(
		j,
		"advancedMonthDays",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetAdvancedWeekDays(val *[]*string) {
	_jsii_.Set(
		j,
		"advancedWeekDays",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetExpiryTime(val *string) {
	_jsii_.Set(
		j,
		"expiryTime",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetExpiryTimeOffsetMinutes(val *float64) {
	_jsii_.Set(
		j,
		"expiryTimeOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetFrequency(val *string) {
	_jsii_.Set(
		j,
		"frequency",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetIsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"isEnabled",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetNextRun(val *string) {
	_jsii_.Set(
		j,
		"nextRun",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetNextRunOffsetMinutes(val *float64) {
	_jsii_.Set(
		j,
		"nextRunOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetStartTimeOffsetMinutes(val *float64) {
	_jsii_.Set(
		j,
		"startTimeOffsetMinutes",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) PutMonthlyOccurrence(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putMonthlyOccurrence",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetAdvancedMonthDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedMonthDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetAdvancedWeekDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedWeekDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetExpiryTime() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiryTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetExpiryTimeOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiryTimeOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetFrequency() {
	_jsii_.InvokeVoid(
		a,
		"resetFrequency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetIsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetMonthlyOccurrence() {
	_jsii_.InvokeVoid(
		a,
		"resetMonthlyOccurrence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetNextRun() {
	_jsii_.InvokeVoid(
		a,
		"resetNextRun",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetNextRunOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetNextRunOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetStartTimeOffsetMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTimeOffsetMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTarget struct {
	// azure_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#azure_query AutomationSoftwareUpdateConfiguration#azure_query}
	AzureQuery interface{} `field:"optional" json:"azureQuery" yaml:"azureQuery"`
	// non_azure_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#non_azure_query AutomationSoftwareUpdateConfiguration#non_azure_query}
	NonAzureQuery interface{} `field:"optional" json:"nonAzureQuery" yaml:"nonAzureQuery"`
}

type AutomationSoftwareUpdateConfigurationTargetAzureQuery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#locations AutomationSoftwareUpdateConfiguration#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#scope AutomationSoftwareUpdateConfiguration#scope}.
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tag_filter AutomationSoftwareUpdateConfiguration#tag_filter}.
	TagFilter *string `field:"optional" json:"tagFilter" yaml:"tagFilter"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tags AutomationSoftwareUpdateConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

type AutomationSoftwareUpdateConfigurationTargetAzureQueryList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetAzureQueryList
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationTargetAzureQueryList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryList_Override(a AutomationSoftwareUpdateConfigurationTargetAzureQueryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) Get(index *float64) AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference {
	var returns AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Locations() *[]*string
	SetLocations(val *[]*string)
	LocationsInput() *[]*string
	Scope() *[]*string
	SetScope(val *[]*string)
	ScopeInput() *[]*string
	TagFilter() *string
	SetTagFilter(val *string)
	TagFilterInput() *string
	Tags() AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList
	TagsInput() interface{}
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
	PutTags(value interface{})
	ResetLocations()
	ResetScope()
	ResetTagFilter()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) LocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) Scope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) TagFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) TagFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) Tags() AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList {
	var returns AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference_Override(a AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetLocations(val *[]*string) {
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetScope(val *[]*string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetTagFilter(val *string) {
	_jsii_.Set(
		j,
		"tagFilter",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) PutTags(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		a,
		"resetLocations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ResetTagFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetTagFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetAzureQueryTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#tag AutomationSoftwareUpdateConfiguration#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#values AutomationSoftwareUpdateConfiguration#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

type AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList_Override(a AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) Get(index *float64) AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference {
	var returns AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference_Override(a AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetTag(val *string) {
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) SetValues(val *[]*string) {
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetAzureQueryTagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetNonAzureQuery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#function_alias AutomationSoftwareUpdateConfiguration#function_alias}.
	FunctionAlias *string `field:"optional" json:"functionAlias" yaml:"functionAlias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#workspace_id AutomationSoftwareUpdateConfiguration#workspace_id}.
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

type AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList interface {
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
	Get(index *float64) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryList_Override(a AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Get(index *float64) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference {
	var returns AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference interface {
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
	FunctionAlias() *string
	SetFunctionAlias(val *string)
	FunctionAliasInput() *string
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
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	ResetFunctionAlias()
	ResetWorkspaceId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) FunctionAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) FunctionAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference_Override(a AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetFunctionAlias(val *string) {
	_jsii_.Set(
		j,
		"functionAlias",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ResetFunctionAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetNonAzureQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTargetOutputReference interface {
	cdktf.ComplexObject
	AzureQuery() AutomationSoftwareUpdateConfigurationTargetAzureQueryList
	AzureQueryInput() interface{}
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
	InternalValue() *AutomationSoftwareUpdateConfigurationTarget
	SetInternalValue(val *AutomationSoftwareUpdateConfigurationTarget)
	NonAzureQuery() AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList
	NonAzureQueryInput() interface{}
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
	PutAzureQuery(value interface{})
	PutNonAzureQuery(value interface{})
	ResetAzureQuery()
	ResetNonAzureQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTargetOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) AzureQuery() AutomationSoftwareUpdateConfigurationTargetAzureQueryList {
	var returns AutomationSoftwareUpdateConfigurationTargetAzureQueryList
	_jsii_.Get(
		j,
		"azureQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) AzureQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) InternalValue() *AutomationSoftwareUpdateConfigurationTarget {
	var returns *AutomationSoftwareUpdateConfigurationTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) NonAzureQuery() AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList {
	var returns AutomationSoftwareUpdateConfigurationTargetNonAzureQueryList
	_jsii_.Get(
		j,
		"nonAzureQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) NonAzureQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonAzureQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomationSoftwareUpdateConfigurationTargetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTargetOutputReference_Override(a AutomationSoftwareUpdateConfigurationTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) SetInternalValue(val *AutomationSoftwareUpdateConfigurationTarget) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) PutAzureQuery(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putAzureQuery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) PutNonAzureQuery(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putNonAzureQuery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ResetAzureQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ResetNonAzureQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetNonAzureQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#create AutomationSoftwareUpdateConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#delete AutomationSoftwareUpdateConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#read AutomationSoftwareUpdateConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#update AutomationSoftwareUpdateConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type AutomationSoftwareUpdateConfigurationTimeoutsOutputReference interface {
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

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationTimeoutsOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomationSoftwareUpdateConfigurationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationTimeoutsOutputReference_Override(a AutomationSoftwareUpdateConfigurationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AutomationSoftwareUpdateConfigurationWindows struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#classification_included AutomationSoftwareUpdateConfiguration#classification_included}.
	ClassificationIncluded *string `field:"optional" json:"classificationIncluded" yaml:"classificationIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#classifications_included AutomationSoftwareUpdateConfiguration#classifications_included}.
	ClassificationsIncluded *[]*string `field:"optional" json:"classificationsIncluded" yaml:"classificationsIncluded"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#excluded_knowledge_base_numbers AutomationSoftwareUpdateConfiguration#excluded_knowledge_base_numbers}.
	ExcludedKnowledgeBaseNumbers *[]*string `field:"optional" json:"excludedKnowledgeBaseNumbers" yaml:"excludedKnowledgeBaseNumbers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#included_knowledge_base_numbers AutomationSoftwareUpdateConfiguration#included_knowledge_base_numbers}.
	IncludedKnowledgeBaseNumbers *[]*string `field:"optional" json:"includedKnowledgeBaseNumbers" yaml:"includedKnowledgeBaseNumbers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_software_update_configuration#reboot AutomationSoftwareUpdateConfiguration#reboot}.
	Reboot *string `field:"optional" json:"reboot" yaml:"reboot"`
}

type AutomationSoftwareUpdateConfigurationWindowsOutputReference interface {
	cdktf.ComplexObject
	ClassificationIncluded() *string
	SetClassificationIncluded(val *string)
	ClassificationIncludedInput() *string
	ClassificationsIncluded() *[]*string
	SetClassificationsIncluded(val *[]*string)
	ClassificationsIncludedInput() *[]*string
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
	ExcludedKnowledgeBaseNumbers() *[]*string
	SetExcludedKnowledgeBaseNumbers(val *[]*string)
	ExcludedKnowledgeBaseNumbersInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedKnowledgeBaseNumbers() *[]*string
	SetIncludedKnowledgeBaseNumbers(val *[]*string)
	IncludedKnowledgeBaseNumbersInput() *[]*string
	InternalValue() *AutomationSoftwareUpdateConfigurationWindows
	SetInternalValue(val *AutomationSoftwareUpdateConfigurationWindows)
	Reboot() *string
	SetReboot(val *string)
	RebootInput() *string
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
	ResetClassificationIncluded()
	ResetClassificationsIncluded()
	ResetExcludedKnowledgeBaseNumbers()
	ResetIncludedKnowledgeBaseNumbers()
	ResetReboot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomationSoftwareUpdateConfigurationWindowsOutputReference
type jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationIncluded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationIncludedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationsIncluded() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsIncluded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ClassificationsIncludedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classificationsIncludedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ExcludedKnowledgeBaseNumbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedKnowledgeBaseNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ExcludedKnowledgeBaseNumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedKnowledgeBaseNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) IncludedKnowledgeBaseNumbers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedKnowledgeBaseNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) IncludedKnowledgeBaseNumbersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedKnowledgeBaseNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InternalValue() *AutomationSoftwareUpdateConfigurationWindows {
	var returns *AutomationSoftwareUpdateConfigurationWindows
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Reboot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reboot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) RebootInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutomationSoftwareUpdateConfigurationWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomationSoftwareUpdateConfigurationWindowsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomationSoftwareUpdateConfigurationWindowsOutputReference_Override(a AutomationSoftwareUpdateConfigurationWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfigurationWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetClassificationIncluded(val *string) {
	_jsii_.Set(
		j,
		"classificationIncluded",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetClassificationsIncluded(val *[]*string) {
	_jsii_.Set(
		j,
		"classificationsIncluded",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetExcludedKnowledgeBaseNumbers(val *[]*string) {
	_jsii_.Set(
		j,
		"excludedKnowledgeBaseNumbers",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetIncludedKnowledgeBaseNumbers(val *[]*string) {
	_jsii_.Set(
		j,
		"includedKnowledgeBaseNumbers",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetInternalValue(val *AutomationSoftwareUpdateConfigurationWindows) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetReboot(val *string) {
	_jsii_.Set(
		j,
		"reboot",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetClassificationIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetClassificationIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetClassificationsIncluded() {
	_jsii_.InvokeVoid(
		a,
		"resetClassificationsIncluded",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetExcludedKnowledgeBaseNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedKnowledgeBaseNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetIncludedKnowledgeBaseNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludedKnowledgeBaseNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ResetReboot() {
	_jsii_.InvokeVoid(
		a,
		"resetReboot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfigurationWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

