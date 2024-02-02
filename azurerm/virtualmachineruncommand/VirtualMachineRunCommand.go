// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommand

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/virtualmachineruncommand/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_machine_run_command azurerm_virtual_machine_run_command}.
type VirtualMachineRunCommand interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ErrorBlobManagedIdentity() VirtualMachineRunCommandErrorBlobManagedIdentityOutputReference
	ErrorBlobManagedIdentityInput() *VirtualMachineRunCommandErrorBlobManagedIdentity
	ErrorBlobUri() *string
	SetErrorBlobUri(val *string)
	ErrorBlobUriInput() *string
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
	InstanceView() VirtualMachineRunCommandInstanceViewList
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
	OutputBlobManagedIdentity() VirtualMachineRunCommandOutputBlobManagedIdentityOutputReference
	OutputBlobManagedIdentityInput() *VirtualMachineRunCommandOutputBlobManagedIdentity
	OutputBlobUri() *string
	SetOutputBlobUri(val *string)
	OutputBlobUriInput() *string
	Parameter() VirtualMachineRunCommandParameterList
	ParameterInput() interface{}
	ProtectedParameter() VirtualMachineRunCommandProtectedParameterList
	ProtectedParameterInput() interface{}
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
	RunAsPassword() *string
	SetRunAsPassword(val *string)
	RunAsPasswordInput() *string
	RunAsUser() *string
	SetRunAsUser(val *string)
	RunAsUserInput() *string
	Source() VirtualMachineRunCommandSourceOutputReference
	SourceInput() *VirtualMachineRunCommandSource
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachineRunCommandTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineId() *string
	SetVirtualMachineId(val *string)
	VirtualMachineIdInput() *string
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
	PutErrorBlobManagedIdentity(value *VirtualMachineRunCommandErrorBlobManagedIdentity)
	PutOutputBlobManagedIdentity(value *VirtualMachineRunCommandOutputBlobManagedIdentity)
	PutParameter(value interface{})
	PutProtectedParameter(value interface{})
	PutSource(value *VirtualMachineRunCommandSource)
	PutTimeouts(value *VirtualMachineRunCommandTimeouts)
	ResetErrorBlobManagedIdentity()
	ResetErrorBlobUri()
	ResetId()
	ResetOutputBlobManagedIdentity()
	ResetOutputBlobUri()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameter()
	ResetProtectedParameter()
	ResetRunAsPassword()
	ResetRunAsUser()
	ResetTags()
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

// The jsii proxy struct for VirtualMachineRunCommand
type jsiiProxy_VirtualMachineRunCommand struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachineRunCommand) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ErrorBlobManagedIdentity() VirtualMachineRunCommandErrorBlobManagedIdentityOutputReference {
	var returns VirtualMachineRunCommandErrorBlobManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"errorBlobManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ErrorBlobManagedIdentityInput() *VirtualMachineRunCommandErrorBlobManagedIdentity {
	var returns *VirtualMachineRunCommandErrorBlobManagedIdentity
	_jsii_.Get(
		j,
		"errorBlobManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ErrorBlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorBlobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ErrorBlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorBlobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) InstanceView() VirtualMachineRunCommandInstanceViewList {
	var returns VirtualMachineRunCommandInstanceViewList
	_jsii_.Get(
		j,
		"instanceView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) OutputBlobManagedIdentity() VirtualMachineRunCommandOutputBlobManagedIdentityOutputReference {
	var returns VirtualMachineRunCommandOutputBlobManagedIdentityOutputReference
	_jsii_.Get(
		j,
		"outputBlobManagedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) OutputBlobManagedIdentityInput() *VirtualMachineRunCommandOutputBlobManagedIdentity {
	var returns *VirtualMachineRunCommandOutputBlobManagedIdentity
	_jsii_.Get(
		j,
		"outputBlobManagedIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) OutputBlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBlobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) OutputBlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBlobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Parameter() VirtualMachineRunCommandParameterList {
	var returns VirtualMachineRunCommandParameterList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ProtectedParameter() VirtualMachineRunCommandProtectedParameterList {
	var returns VirtualMachineRunCommandProtectedParameterList
	_jsii_.Get(
		j,
		"protectedParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) ProtectedParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) RunAsPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) RunAsPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) RunAsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) RunAsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Source() VirtualMachineRunCommandSourceOutputReference {
	var returns VirtualMachineRunCommandSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) SourceInput() *VirtualMachineRunCommandSource {
	var returns *VirtualMachineRunCommandSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) Timeouts() VirtualMachineRunCommandTimeoutsOutputReference {
	var returns VirtualMachineRunCommandTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineRunCommand) VirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_machine_run_command azurerm_virtual_machine_run_command} Resource.
func NewVirtualMachineRunCommand(scope constructs.Construct, id *string, config *VirtualMachineRunCommandConfig) VirtualMachineRunCommand {
	_init_.Initialize()

	if err := validateNewVirtualMachineRunCommandParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineRunCommand{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/virtual_machine_run_command azurerm_virtual_machine_run_command} Resource.
func NewVirtualMachineRunCommand_Override(v VirtualMachineRunCommand, scope constructs.Construct, id *string, config *VirtualMachineRunCommandConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetErrorBlobUri(val *string) {
	if err := j.validateSetErrorBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorBlobUri",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetOutputBlobUri(val *string) {
	if err := j.validateSetOutputBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputBlobUri",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetRunAsPassword(val *string) {
	if err := j.validateSetRunAsPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsPassword",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetRunAsUser(val *string) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineRunCommand)SetVirtualMachineId(val *string) {
	if err := j.validateSetVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineId",
		val,
	)
}

// Generates CDKTF code for importing a VirtualMachineRunCommand resource upon running "cdktf plan <stack-name>".
func VirtualMachineRunCommand_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualMachineRunCommand_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
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
func VirtualMachineRunCommand_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineRunCommand_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineRunCommand_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineRunCommand_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineRunCommand_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineRunCommand_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachineRunCommand_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.virtualMachineRunCommand.VirtualMachineRunCommand",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutErrorBlobManagedIdentity(value *VirtualMachineRunCommandErrorBlobManagedIdentity) {
	if err := v.validatePutErrorBlobManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putErrorBlobManagedIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutOutputBlobManagedIdentity(value *VirtualMachineRunCommandOutputBlobManagedIdentity) {
	if err := v.validatePutOutputBlobManagedIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putOutputBlobManagedIdentity",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutParameter(value interface{}) {
	if err := v.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putParameter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutProtectedParameter(value interface{}) {
	if err := v.validatePutProtectedParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putProtectedParameter",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutSource(value *VirtualMachineRunCommandSource) {
	if err := v.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSource",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) PutTimeouts(value *VirtualMachineRunCommandTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetErrorBlobManagedIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetErrorBlobManagedIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetErrorBlobUri() {
	_jsii_.InvokeVoid(
		v,
		"resetErrorBlobUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetOutputBlobManagedIdentity() {
	_jsii_.InvokeVoid(
		v,
		"resetOutputBlobManagedIdentity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetOutputBlobUri() {
	_jsii_.InvokeVoid(
		v,
		"resetOutputBlobUri",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetParameter() {
	_jsii_.InvokeVoid(
		v,
		"resetParameter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetProtectedParameter() {
	_jsii_.InvokeVoid(
		v,
		"resetProtectedParameter",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetRunAsPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetRunAsPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		v,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineRunCommand) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineRunCommand) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

