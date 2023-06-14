package dataazurermhybridcomputemachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/dataazurermhybridcomputemachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/hybrid_compute_machine azurerm_hybrid_compute_machine}.
type DataAzurermHybridComputeMachine interface {
	cdktf.TerraformDataSource
	AdFqdn() *string
	AgentConfiguration() DataAzurermHybridComputeMachineAgentConfigurationList
	AgentVersion() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientPublicKey() *string
	CloudMetadata() DataAzurermHybridComputeMachineCloudMetadataList
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
	DetectedProperties() cdktf.StringMap
	DisplayName() *string
	DnsFqdn() *string
	DomainName() *string
	ErrorDetails() DataAzurermHybridComputeMachineErrorDetailsList
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
	Identity() DataAzurermHybridComputeMachineIdentityList
	IdInput() *string
	LastStatusChange() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	LocationData() DataAzurermHybridComputeMachineLocationDataList
	MachineFqdn() *string
	MssqlDiscovered() cdktf.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsName() *string
	OsProfile() DataAzurermHybridComputeMachineOsProfileList
	OsSku() *string
	OsType() *string
	OsVersion() *string
	ParentClusterResourceId() *string
	PrivateLinkScopeResourceId() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServiceStatus() DataAzurermHybridComputeMachineServiceStatusList
	Status() *string
	Tags() cdktf.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermHybridComputeMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	VmId() *string
	VmUuid() *string
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
	PutTimeouts(value *DataAzurermHybridComputeMachineTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for DataAzurermHybridComputeMachine
type jsiiProxy_DataAzurermHybridComputeMachine struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) AdFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) AgentConfiguration() DataAzurermHybridComputeMachineAgentConfigurationList {
	var returns DataAzurermHybridComputeMachineAgentConfigurationList
	_jsii_.Get(
		j,
		"agentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ClientPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) CloudMetadata() DataAzurermHybridComputeMachineCloudMetadataList {
	var returns DataAzurermHybridComputeMachineCloudMetadataList
	_jsii_.Get(
		j,
		"cloudMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) DetectedProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"detectedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) DnsFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ErrorDetails() DataAzurermHybridComputeMachineErrorDetailsList {
	var returns DataAzurermHybridComputeMachineErrorDetailsList
	_jsii_.Get(
		j,
		"errorDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Identity() DataAzurermHybridComputeMachineIdentityList {
	var returns DataAzurermHybridComputeMachineIdentityList
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) LastStatusChange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastStatusChange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) LocationData() DataAzurermHybridComputeMachineLocationDataList {
	var returns DataAzurermHybridComputeMachineLocationDataList
	_jsii_.Get(
		j,
		"locationData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) MachineFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) MssqlDiscovered() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mssqlDiscovered",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) OsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) OsProfile() DataAzurermHybridComputeMachineOsProfileList {
	var returns DataAzurermHybridComputeMachineOsProfileList
	_jsii_.Get(
		j,
		"osProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) OsSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) OsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ParentClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentClusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) PrivateLinkScopeResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkScopeResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) ServiceStatus() DataAzurermHybridComputeMachineServiceStatusList {
	var returns DataAzurermHybridComputeMachineServiceStatusList
	_jsii_.Get(
		j,
		"serviceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Tags() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) Timeouts() DataAzurermHybridComputeMachineTimeoutsOutputReference {
	var returns DataAzurermHybridComputeMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) VmId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine) VmUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmUuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/hybrid_compute_machine azurerm_hybrid_compute_machine} Data Source.
func NewDataAzurermHybridComputeMachine(scope constructs.Construct, id *string, config *DataAzurermHybridComputeMachineConfig) DataAzurermHybridComputeMachine {
	_init_.Initialize()

	if err := validateNewDataAzurermHybridComputeMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermHybridComputeMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/data-sources/hybrid_compute_machine azurerm_hybrid_compute_machine} Data Source.
func NewDataAzurermHybridComputeMachine_Override(d DataAzurermHybridComputeMachine, scope constructs.Construct, id *string, config *DataAzurermHybridComputeMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermHybridComputeMachine)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func DataAzurermHybridComputeMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermHybridComputeMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermHybridComputeMachine_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermHybridComputeMachine_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermHybridComputeMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermHybridComputeMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermHybridComputeMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermHybridComputeMachine.DataAzurermHybridComputeMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) PutTimeouts(value *DataAzurermHybridComputeMachineTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermHybridComputeMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

