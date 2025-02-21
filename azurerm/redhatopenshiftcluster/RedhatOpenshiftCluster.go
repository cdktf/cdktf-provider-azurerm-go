// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redhatopenshiftcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/redhatopenshiftcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/redhat_openshift_cluster azurerm_redhat_openshift_cluster}.
type RedhatOpenshiftCluster interface {
	cdktf.TerraformResource
	ApiServerProfile() RedhatOpenshiftClusterApiServerProfileOutputReference
	ApiServerProfileInput() *RedhatOpenshiftClusterApiServerProfile
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterProfile() RedhatOpenshiftClusterClusterProfileOutputReference
	ClusterProfileInput() *RedhatOpenshiftClusterClusterProfile
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConsoleUrl() *string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	IngressProfile() RedhatOpenshiftClusterIngressProfileOutputReference
	IngressProfileInput() *RedhatOpenshiftClusterIngressProfile
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MainProfile() RedhatOpenshiftClusterMainProfileOutputReference
	MainProfileInput() *RedhatOpenshiftClusterMainProfile
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkProfile() RedhatOpenshiftClusterNetworkProfileOutputReference
	NetworkProfileInput() *RedhatOpenshiftClusterNetworkProfile
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
	ServicePrincipal() RedhatOpenshiftClusterServicePrincipalOutputReference
	ServicePrincipalInput() *RedhatOpenshiftClusterServicePrincipal
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RedhatOpenshiftClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkerProfile() RedhatOpenshiftClusterWorkerProfileOutputReference
	WorkerProfileInput() *RedhatOpenshiftClusterWorkerProfile
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
	PutApiServerProfile(value *RedhatOpenshiftClusterApiServerProfile)
	PutClusterProfile(value *RedhatOpenshiftClusterClusterProfile)
	PutIngressProfile(value *RedhatOpenshiftClusterIngressProfile)
	PutMainProfile(value *RedhatOpenshiftClusterMainProfile)
	PutNetworkProfile(value *RedhatOpenshiftClusterNetworkProfile)
	PutServicePrincipal(value *RedhatOpenshiftClusterServicePrincipal)
	PutTimeouts(value *RedhatOpenshiftClusterTimeouts)
	PutWorkerProfile(value *RedhatOpenshiftClusterWorkerProfile)
	ResetId()
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

// The jsii proxy struct for RedhatOpenshiftCluster
type jsiiProxy_RedhatOpenshiftCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ApiServerProfile() RedhatOpenshiftClusterApiServerProfileOutputReference {
	var returns RedhatOpenshiftClusterApiServerProfileOutputReference
	_jsii_.Get(
		j,
		"apiServerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ApiServerProfileInput() *RedhatOpenshiftClusterApiServerProfile {
	var returns *RedhatOpenshiftClusterApiServerProfile
	_jsii_.Get(
		j,
		"apiServerProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ClusterProfile() RedhatOpenshiftClusterClusterProfileOutputReference {
	var returns RedhatOpenshiftClusterClusterProfileOutputReference
	_jsii_.Get(
		j,
		"clusterProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ClusterProfileInput() *RedhatOpenshiftClusterClusterProfile {
	var returns *RedhatOpenshiftClusterClusterProfile
	_jsii_.Get(
		j,
		"clusterProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) IngressProfile() RedhatOpenshiftClusterIngressProfileOutputReference {
	var returns RedhatOpenshiftClusterIngressProfileOutputReference
	_jsii_.Get(
		j,
		"ingressProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) IngressProfileInput() *RedhatOpenshiftClusterIngressProfile {
	var returns *RedhatOpenshiftClusterIngressProfile
	_jsii_.Get(
		j,
		"ingressProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) MainProfile() RedhatOpenshiftClusterMainProfileOutputReference {
	var returns RedhatOpenshiftClusterMainProfileOutputReference
	_jsii_.Get(
		j,
		"mainProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) MainProfileInput() *RedhatOpenshiftClusterMainProfile {
	var returns *RedhatOpenshiftClusterMainProfile
	_jsii_.Get(
		j,
		"mainProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) NetworkProfile() RedhatOpenshiftClusterNetworkProfileOutputReference {
	var returns RedhatOpenshiftClusterNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) NetworkProfileInput() *RedhatOpenshiftClusterNetworkProfile {
	var returns *RedhatOpenshiftClusterNetworkProfile
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ServicePrincipal() RedhatOpenshiftClusterServicePrincipalOutputReference {
	var returns RedhatOpenshiftClusterServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) ServicePrincipalInput() *RedhatOpenshiftClusterServicePrincipal {
	var returns *RedhatOpenshiftClusterServicePrincipal
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) Timeouts() RedhatOpenshiftClusterTimeoutsOutputReference {
	var returns RedhatOpenshiftClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) WorkerProfile() RedhatOpenshiftClusterWorkerProfileOutputReference {
	var returns RedhatOpenshiftClusterWorkerProfileOutputReference
	_jsii_.Get(
		j,
		"workerProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedhatOpenshiftCluster) WorkerProfileInput() *RedhatOpenshiftClusterWorkerProfile {
	var returns *RedhatOpenshiftClusterWorkerProfile
	_jsii_.Get(
		j,
		"workerProfileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/redhat_openshift_cluster azurerm_redhat_openshift_cluster} Resource.
func NewRedhatOpenshiftCluster(scope constructs.Construct, id *string, config *RedhatOpenshiftClusterConfig) RedhatOpenshiftCluster {
	_init_.Initialize()

	if err := validateNewRedhatOpenshiftClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedhatOpenshiftCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/redhat_openshift_cluster azurerm_redhat_openshift_cluster} Resource.
func NewRedhatOpenshiftCluster_Override(r RedhatOpenshiftCluster, scope constructs.Construct, id *string, config *RedhatOpenshiftClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_RedhatOpenshiftCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a RedhatOpenshiftCluster resource upon running "cdktf plan <stack-name>".
func RedhatOpenshiftCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRedhatOpenshiftCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
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
func RedhatOpenshiftCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedhatOpenshiftCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedhatOpenshiftCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedhatOpenshiftCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedhatOpenshiftCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedhatOpenshiftCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedhatOpenshiftCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.redhatOpenshiftCluster.RedhatOpenshiftCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutApiServerProfile(value *RedhatOpenshiftClusterApiServerProfile) {
	if err := r.validatePutApiServerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putApiServerProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutClusterProfile(value *RedhatOpenshiftClusterClusterProfile) {
	if err := r.validatePutClusterProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putClusterProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutIngressProfile(value *RedhatOpenshiftClusterIngressProfile) {
	if err := r.validatePutIngressProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIngressProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutMainProfile(value *RedhatOpenshiftClusterMainProfile) {
	if err := r.validatePutMainProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMainProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutNetworkProfile(value *RedhatOpenshiftClusterNetworkProfile) {
	if err := r.validatePutNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutServicePrincipal(value *RedhatOpenshiftClusterServicePrincipal) {
	if err := r.validatePutServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServicePrincipal",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutTimeouts(value *RedhatOpenshiftClusterTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) PutWorkerProfile(value *RedhatOpenshiftClusterWorkerProfile) {
	if err := r.validatePutWorkerProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putWorkerProfile",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedhatOpenshiftCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedhatOpenshiftCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

