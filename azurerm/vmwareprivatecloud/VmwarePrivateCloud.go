// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/vmwareprivatecloud/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/vmware_private_cloud azurerm_vmware_private_cloud}.
type VmwarePrivateCloud interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Circuit() VmwarePrivateCloudCircuitList
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HcxCloudManagerEndpoint() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternetConnectionEnabled() interface{}
	SetInternetConnectionEnabled(val interface{})
	InternetConnectionEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementCluster() VmwarePrivateCloudManagementClusterOutputReference
	ManagementClusterInput() *VmwarePrivateCloudManagementCluster
	ManagementSubnetCidr() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSubnetCidr() *string
	SetNetworkSubnetCidr(val *string)
	NetworkSubnetCidrInput() *string
	// The tree node.
	Node() constructs.Node
	NsxtCertificateThumbprint() *string
	NsxtManagerEndpoint() *string
	NsxtPassword() *string
	SetNsxtPassword(val *string)
	NsxtPasswordInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisioningSubnetCidr() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VmwarePrivateCloudTimeoutsOutputReference
	TimeoutsInput() interface{}
	VcenterCertificateThumbprint() *string
	VcenterPassword() *string
	SetVcenterPassword(val *string)
	VcenterPasswordInput() *string
	VcsaEndpoint() *string
	VmotionSubnetCidr() *string
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
	PutManagementCluster(value *VmwarePrivateCloudManagementCluster)
	PutTimeouts(value *VmwarePrivateCloudTimeouts)
	ResetId()
	ResetInternetConnectionEnabled()
	ResetNsxtPassword()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetVcenterPassword()
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

// The jsii proxy struct for VmwarePrivateCloud
type jsiiProxy_VmwarePrivateCloud struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VmwarePrivateCloud) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Circuit() VmwarePrivateCloudCircuitList {
	var returns VmwarePrivateCloudCircuitList
	_jsii_.Get(
		j,
		"circuit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) HcxCloudManagerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcxCloudManagerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) InternetConnectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetConnectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) InternetConnectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetConnectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ManagementCluster() VmwarePrivateCloudManagementClusterOutputReference {
	var returns VmwarePrivateCloudManagementClusterOutputReference
	_jsii_.Get(
		j,
		"managementCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ManagementClusterInput() *VmwarePrivateCloudManagementCluster {
	var returns *VmwarePrivateCloudManagementCluster
	_jsii_.Get(
		j,
		"managementClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ManagementSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NetworkSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NetworkSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NsxtCertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NsxtManagerEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtManagerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NsxtPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) NsxtPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxtPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ProvisioningSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) Timeouts() VmwarePrivateCloudTimeoutsOutputReference {
	var returns VmwarePrivateCloudTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) VcenterCertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) VcenterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) VcenterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) VcsaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcsaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwarePrivateCloud) VmotionSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmotionSubnetCidr",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/vmware_private_cloud azurerm_vmware_private_cloud} Resource.
func NewVmwarePrivateCloud(scope constructs.Construct, id *string, config *VmwarePrivateCloudConfig) VmwarePrivateCloud {
	_init_.Initialize()

	if err := validateNewVmwarePrivateCloudParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwarePrivateCloud{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/vmware_private_cloud azurerm_vmware_private_cloud} Resource.
func NewVmwarePrivateCloud_Override(v VmwarePrivateCloud, scope constructs.Construct, id *string, config *VmwarePrivateCloudConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetInternetConnectionEnabled(val interface{}) {
	if err := j.validateSetInternetConnectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetConnectionEnabled",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetNetworkSubnetCidr(val *string) {
	if err := j.validateSetNetworkSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetNsxtPassword(val *string) {
	if err := j.validateSetNsxtPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxtPassword",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VmwarePrivateCloud)SetVcenterPassword(val *string) {
	if err := j.validateSetVcenterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcenterPassword",
		val,
	)
}

// Generates CDKTF code for importing a VmwarePrivateCloud resource upon running "cdktf plan <stack-name>".
func VmwarePrivateCloud_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateVmwarePrivateCloud_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
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
func VmwarePrivateCloud_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVmwarePrivateCloud_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VmwarePrivateCloud_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVmwarePrivateCloud_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VmwarePrivateCloud_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVmwarePrivateCloud_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VmwarePrivateCloud_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.vmwarePrivateCloud.VmwarePrivateCloud",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmwarePrivateCloud) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmwarePrivateCloud) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwarePrivateCloud) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) PutManagementCluster(value *VmwarePrivateCloudManagementCluster) {
	if err := v.validatePutManagementClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putManagementCluster",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) PutTimeouts(value *VmwarePrivateCloudTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetInternetConnectionEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetInternetConnectionEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetNsxtPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetNsxtPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) ResetVcenterPassword() {
	_jsii_.InvokeVoid(
		v,
		"resetVcenterPassword",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwarePrivateCloud) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwarePrivateCloud) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

