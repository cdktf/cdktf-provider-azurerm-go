// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/workloadssapsinglenodevirtualinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/workloads_sap_single_node_virtual_instance azurerm_workloads_sap_single_node_virtual_instance}.
type WorkloadsSapSingleNodeVirtualInstance interface {
	cdktf.TerraformResource
	AppLocation() *string
	SetAppLocation(val *string)
	AppLocationInput() *string
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
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
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
	Identity() WorkloadsSapSingleNodeVirtualInstanceIdentityOutputReference
	IdentityInput() *WorkloadsSapSingleNodeVirtualInstanceIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedResourceGroupName() *string
	SetManagedResourceGroupName(val *string)
	ManagedResourceGroupNameInput() *string
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
	SapFqdn() *string
	SetSapFqdn(val *string)
	SapFqdnInput() *string
	SapProduct() *string
	SetSapProduct(val *string)
	SapProductInput() *string
	SingleServerConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference
	SingleServerConfigurationInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WorkloadsSapSingleNodeVirtualInstanceTimeoutsOutputReference
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
	PutIdentity(value *WorkloadsSapSingleNodeVirtualInstanceIdentity)
	PutSingleServerConfiguration(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration)
	PutTimeouts(value *WorkloadsSapSingleNodeVirtualInstanceTimeouts)
	ResetId()
	ResetIdentity()
	ResetManagedResourceGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for WorkloadsSapSingleNodeVirtualInstance
type jsiiProxy_WorkloadsSapSingleNodeVirtualInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) AppLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) AppLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Identity() WorkloadsSapSingleNodeVirtualInstanceIdentityOutputReference {
	var returns WorkloadsSapSingleNodeVirtualInstanceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) IdentityInput() *WorkloadsSapSingleNodeVirtualInstanceIdentity {
	var returns *WorkloadsSapSingleNodeVirtualInstanceIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ManagedResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ManagedResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedResourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SapFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SapFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SapProduct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SapProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapProductInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SingleServerConfiguration() WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference {
	var returns WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationOutputReference
	_jsii_.Get(
		j,
		"singleServerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SingleServerConfigurationInput() *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration {
	var returns *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration
	_jsii_.Get(
		j,
		"singleServerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) Timeouts() WorkloadsSapSingleNodeVirtualInstanceTimeoutsOutputReference {
	var returns WorkloadsSapSingleNodeVirtualInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/workloads_sap_single_node_virtual_instance azurerm_workloads_sap_single_node_virtual_instance} Resource.
func NewWorkloadsSapSingleNodeVirtualInstance(scope constructs.Construct, id *string, config *WorkloadsSapSingleNodeVirtualInstanceConfig) WorkloadsSapSingleNodeVirtualInstance {
	_init_.Initialize()

	if err := validateNewWorkloadsSapSingleNodeVirtualInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadsSapSingleNodeVirtualInstance{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/workloads_sap_single_node_virtual_instance azurerm_workloads_sap_single_node_virtual_instance} Resource.
func NewWorkloadsSapSingleNodeVirtualInstance_Override(w WorkloadsSapSingleNodeVirtualInstance, scope constructs.Construct, id *string, config *WorkloadsSapSingleNodeVirtualInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetAppLocation(val *string) {
	if err := j.validateSetAppLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLocation",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetEnvironment(val *string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetManagedResourceGroupName(val *string) {
	if err := j.validateSetManagedResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedResourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetSapFqdn(val *string) {
	if err := j.validateSetSapFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapFqdn",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetSapProduct(val *string) {
	if err := j.validateSetSapProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapProduct",
		val,
	)
}

func (j *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a WorkloadsSapSingleNodeVirtualInstance resource upon running "cdktf plan <stack-name>".
func WorkloadsSapSingleNodeVirtualInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWorkloadsSapSingleNodeVirtualInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
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
func WorkloadsSapSingleNodeVirtualInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkloadsSapSingleNodeVirtualInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkloadsSapSingleNodeVirtualInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkloadsSapSingleNodeVirtualInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkloadsSapSingleNodeVirtualInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkloadsSapSingleNodeVirtualInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkloadsSapSingleNodeVirtualInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.workloadsSapSingleNodeVirtualInstance.WorkloadsSapSingleNodeVirtualInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) PutIdentity(value *WorkloadsSapSingleNodeVirtualInstanceIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) PutSingleServerConfiguration(value *WorkloadsSapSingleNodeVirtualInstanceSingleServerConfiguration) {
	if err := w.validatePutSingleServerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSingleServerConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) PutTimeouts(value *WorkloadsSapSingleNodeVirtualInstanceTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetManagedResourceGroupName() {
	_jsii_.InvokeVoid(
		w,
		"resetManagedResourceGroupName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadsSapSingleNodeVirtualInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

