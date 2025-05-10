// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/systemcentervirtualmachinemanagervirtualmachineinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance azurerm_system_center_virtual_machine_manager_virtual_machine_instance}.
type SystemCenterVirtualMachineManagerVirtualMachineInstance interface {
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
	CustomLocationId() *string
	SetCustomLocationId(val *string)
	CustomLocationIdInput() *string
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
	Hardware() SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference
	HardwareInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware
	Id() *string
	SetId(val *string)
	IdInput() *string
	Infrastructure() SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructureOutputReference
	InfrastructureInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkInterface() SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OperatingSystem() SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystemOutputReference
	OperatingSystemInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem
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
	ScopedResourceId() *string
	SetScopedResourceId(val *string)
	ScopedResourceIdInput() *string
	StorageDisk() SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskList
	StorageDiskInput() interface{}
	SystemCenterVirtualMachineManagerAvailabilitySetIds() *[]*string
	SetSystemCenterVirtualMachineManagerAvailabilitySetIds(val *[]*string)
	SystemCenterVirtualMachineManagerAvailabilitySetIdsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeoutsOutputReference
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
	PutHardware(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware)
	PutInfrastructure(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure)
	PutNetworkInterface(value interface{})
	PutOperatingSystem(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem)
	PutStorageDisk(value interface{})
	PutTimeouts(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeouts)
	ResetHardware()
	ResetId()
	ResetNetworkInterface()
	ResetOperatingSystem()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageDisk()
	ResetSystemCenterVirtualMachineManagerAvailabilitySetIds()
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

// The jsii proxy struct for SystemCenterVirtualMachineManagerVirtualMachineInstance
type jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) CustomLocationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) CustomLocationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLocationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Hardware() SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceHardwareOutputReference
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) HardwareInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware {
	var returns *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware
	_jsii_.Get(
		j,
		"hardwareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Infrastructure() SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructureOutputReference {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructureOutputReference
	_jsii_.Get(
		j,
		"infrastructure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) InfrastructureInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure {
	var returns *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure
	_jsii_.Get(
		j,
		"infrastructureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) NetworkInterface() SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceList {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) OperatingSystem() SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystemOutputReference {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystemOutputReference
	_jsii_.Get(
		j,
		"operatingSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) OperatingSystemInput() *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem {
	var returns *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem
	_jsii_.Get(
		j,
		"operatingSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ScopedResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopedResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ScopedResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopedResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) StorageDisk() SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskList {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskList
	_jsii_.Get(
		j,
		"storageDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) StorageDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) SystemCenterVirtualMachineManagerAvailabilitySetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemCenterVirtualMachineManagerAvailabilitySetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) SystemCenterVirtualMachineManagerAvailabilitySetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"systemCenterVirtualMachineManagerAvailabilitySetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) Timeouts() SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeoutsOutputReference {
	var returns SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance azurerm_system_center_virtual_machine_manager_virtual_machine_instance} Resource.
func NewSystemCenterVirtualMachineManagerVirtualMachineInstance(scope constructs.Construct, id *string, config *SystemCenterVirtualMachineManagerVirtualMachineInstanceConfig) SystemCenterVirtualMachineManagerVirtualMachineInstance {
	_init_.Initialize()

	if err := validateNewSystemCenterVirtualMachineManagerVirtualMachineInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/system_center_virtual_machine_manager_virtual_machine_instance azurerm_system_center_virtual_machine_manager_virtual_machine_instance} Resource.
func NewSystemCenterVirtualMachineManagerVirtualMachineInstance_Override(s SystemCenterVirtualMachineManagerVirtualMachineInstance, scope constructs.Construct, id *string, config *SystemCenterVirtualMachineManagerVirtualMachineInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetCustomLocationId(val *string) {
	if err := j.validateSetCustomLocationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLocationId",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetScopedResourceId(val *string) {
	if err := j.validateSetScopedResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopedResourceId",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance)SetSystemCenterVirtualMachineManagerAvailabilitySetIds(val *[]*string) {
	if err := j.validateSetSystemCenterVirtualMachineManagerAvailabilitySetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemCenterVirtualMachineManagerAvailabilitySetIds",
		val,
	)
}

// Generates CDKTF code for importing a SystemCenterVirtualMachineManagerVirtualMachineInstance resource upon running "cdktf plan <stack-name>".
func SystemCenterVirtualMachineManagerVirtualMachineInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSystemCenterVirtualMachineManagerVirtualMachineInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
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
func SystemCenterVirtualMachineManagerVirtualMachineInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSystemCenterVirtualMachineManagerVirtualMachineInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SystemCenterVirtualMachineManagerVirtualMachineInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSystemCenterVirtualMachineManagerVirtualMachineInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SystemCenterVirtualMachineManagerVirtualMachineInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSystemCenterVirtualMachineManagerVirtualMachineInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SystemCenterVirtualMachineManagerVirtualMachineInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutHardware(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceHardware) {
	if err := s.validatePutHardwareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHardware",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutInfrastructure(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceInfrastructure) {
	if err := s.validatePutInfrastructureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfrastructure",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutNetworkInterface(value interface{}) {
	if err := s.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutOperatingSystem(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceOperatingSystem) {
	if err := s.validatePutOperatingSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOperatingSystem",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutStorageDisk(value interface{}) {
	if err := s.validatePutStorageDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorageDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) PutTimeouts(value *SystemCenterVirtualMachineManagerVirtualMachineInstanceTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetHardware() {
	_jsii_.InvokeVoid(
		s,
		"resetHardware",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetNetworkInterface() {
	_jsii_.InvokeVoid(
		s,
		"resetNetworkInterface",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetOperatingSystem() {
	_jsii_.InvokeVoid(
		s,
		"resetOperatingSystem",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetStorageDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetSystemCenterVirtualMachineManagerAvailabilitySetIds() {
	_jsii_.InvokeVoid(
		s,
		"resetSystemCenterVirtualMachineManagerAvailabilitySetIds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

