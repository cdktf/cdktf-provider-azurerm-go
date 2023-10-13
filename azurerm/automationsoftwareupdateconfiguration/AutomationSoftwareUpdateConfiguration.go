// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/automationsoftwareupdateconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/automation_software_update_configuration azurerm_automation_software_update_configuration}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	ErrorCode() *string
	ErrorMeesage() *string
	ErrorMessage() *string
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
	Linux() AutomationSoftwareUpdateConfigurationLinuxOutputReference
	LinuxInput() *AutomationSoftwareUpdateConfigurationLinux
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
	PostTask() AutomationSoftwareUpdateConfigurationPostTaskOutputReference
	PostTaskInput() *AutomationSoftwareUpdateConfigurationPostTask
	PreTask() AutomationSoftwareUpdateConfigurationPreTaskOutputReference
	PreTaskInput() *AutomationSoftwareUpdateConfigurationPreTask
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
	Schedule() AutomationSoftwareUpdateConfigurationScheduleOutputReference
	ScheduleInput() *AutomationSoftwareUpdateConfigurationSchedule
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
	PutLinux(value *AutomationSoftwareUpdateConfigurationLinux)
	PutPostTask(value *AutomationSoftwareUpdateConfigurationPostTask)
	PutPreTask(value *AutomationSoftwareUpdateConfigurationPreTask)
	PutSchedule(value *AutomationSoftwareUpdateConfigurationSchedule)
	PutTarget(value *AutomationSoftwareUpdateConfigurationTarget)
	PutTimeouts(value *AutomationSoftwareUpdateConfigurationTimeouts)
	PutWindows(value *AutomationSoftwareUpdateConfigurationWindows)
	ResetDuration()
	ResetId()
	ResetLinux()
	ResetNonAzureComputerNames()
	ResetOperatingSystem()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostTask()
	ResetPreTask()
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

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorMessage",
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

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Linux() AutomationSoftwareUpdateConfigurationLinuxOutputReference {
	var returns AutomationSoftwareUpdateConfigurationLinuxOutputReference
	_jsii_.Get(
		j,
		"linux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) LinuxInput() *AutomationSoftwareUpdateConfigurationLinux {
	var returns *AutomationSoftwareUpdateConfigurationLinux
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

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PostTask() AutomationSoftwareUpdateConfigurationPostTaskOutputReference {
	var returns AutomationSoftwareUpdateConfigurationPostTaskOutputReference
	_jsii_.Get(
		j,
		"postTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PostTaskInput() *AutomationSoftwareUpdateConfigurationPostTask {
	var returns *AutomationSoftwareUpdateConfigurationPostTask
	_jsii_.Get(
		j,
		"postTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PreTask() AutomationSoftwareUpdateConfigurationPreTaskOutputReference {
	var returns AutomationSoftwareUpdateConfigurationPreTaskOutputReference
	_jsii_.Get(
		j,
		"preTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) PreTaskInput() *AutomationSoftwareUpdateConfigurationPreTask {
	var returns *AutomationSoftwareUpdateConfigurationPreTask
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

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) Schedule() AutomationSoftwareUpdateConfigurationScheduleOutputReference {
	var returns AutomationSoftwareUpdateConfigurationScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration) ScheduleInput() *AutomationSoftwareUpdateConfigurationSchedule {
	var returns *AutomationSoftwareUpdateConfigurationSchedule
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/automation_software_update_configuration azurerm_automation_software_update_configuration} Resource.
func NewAutomationSoftwareUpdateConfiguration(scope constructs.Construct, id *string, config *AutomationSoftwareUpdateConfigurationConfig) AutomationSoftwareUpdateConfiguration {
	_init_.Initialize()

	if err := validateNewAutomationSoftwareUpdateConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomationSoftwareUpdateConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/automation_software_update_configuration azurerm_automation_software_update_configuration} Resource.
func NewAutomationSoftwareUpdateConfiguration_Override(a AutomationSoftwareUpdateConfiguration, scope constructs.Construct, id *string, config *AutomationSoftwareUpdateConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetAutomationAccountId(val *string) {
	if err := j.validateSetAutomationAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automationAccountId",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetNonAzureComputerNames(val *[]*string) {
	if err := j.validateSetNonAzureComputerNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonAzureComputerNames",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetOperatingSystem(val *string) {
	if err := j.validateSetOperatingSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystem",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AutomationSoftwareUpdateConfiguration)SetVirtualMachineIds(val *[]*string) {
	if err := j.validateSetVirtualMachineIdsParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateAutomationSoftwareUpdateConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationSoftwareUpdateConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationSoftwareUpdateConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AutomationSoftwareUpdateConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAutomationSoftwareUpdateConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.automationSoftwareUpdateConfiguration.AutomationSoftwareUpdateConfiguration",
		"isTerraformResource",
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
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutLinux(value *AutomationSoftwareUpdateConfigurationLinux) {
	if err := a.validatePutLinuxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLinux",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutPostTask(value *AutomationSoftwareUpdateConfigurationPostTask) {
	if err := a.validatePutPostTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostTask",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutPreTask(value *AutomationSoftwareUpdateConfigurationPreTask) {
	if err := a.validatePutPreTaskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPreTask",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutSchedule(value *AutomationSoftwareUpdateConfigurationSchedule) {
	if err := a.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutTarget(value *AutomationSoftwareUpdateConfigurationTarget) {
	if err := a.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutTimeouts(value *AutomationSoftwareUpdateConfigurationTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) PutWindows(value *AutomationSoftwareUpdateConfigurationWindows) {
	if err := a.validatePutWindowsParameters(value); err != nil {
		panic(err)
	}
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

func (a *jsiiProxy_AutomationSoftwareUpdateConfiguration) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		a,
		"resetOperatingSystem",
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

