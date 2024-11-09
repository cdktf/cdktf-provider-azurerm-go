// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusterextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/kubernetesclusterextension/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kubernetes_cluster_extension azurerm_kubernetes_cluster_extension}.
type KubernetesClusterExtension interface {
	cdktf.TerraformResource
	AksAssignedIdentity() KubernetesClusterExtensionAksAssignedIdentityList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConfigurationProtectedSettings() *map[string]*string
	SetConfigurationProtectedSettings(val *map[string]*string)
	ConfigurationProtectedSettingsInput() *map[string]*string
	ConfigurationSettings() *map[string]*string
	SetConfigurationSettings(val *map[string]*string)
	ConfigurationSettingsInput() *map[string]*string
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
	CurrentVersion() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExtensionType() *string
	SetExtensionType(val *string)
	ExtensionTypeInput() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Plan() KubernetesClusterExtensionPlanOutputReference
	PlanInput() *KubernetesClusterExtensionPlan
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
	ReleaseNamespace() *string
	SetReleaseNamespace(val *string)
	ReleaseNamespaceInput() *string
	ReleaseTrain() *string
	SetReleaseTrain(val *string)
	ReleaseTrainInput() *string
	TargetNamespace() *string
	SetTargetNamespace(val *string)
	TargetNamespaceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesClusterExtensionTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutPlan(value *KubernetesClusterExtensionPlan)
	PutTimeouts(value *KubernetesClusterExtensionTimeouts)
	ResetConfigurationProtectedSettings()
	ResetConfigurationSettings()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetReleaseNamespace()
	ResetReleaseTrain()
	ResetTargetNamespace()
	ResetTimeouts()
	ResetVersion()
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

// The jsii proxy struct for KubernetesClusterExtension
type jsiiProxy_KubernetesClusterExtension struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesClusterExtension) AksAssignedIdentity() KubernetesClusterExtensionAksAssignedIdentityList {
	var returns KubernetesClusterExtensionAksAssignedIdentityList
	_jsii_.Get(
		j,
		"aksAssignedIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ConfigurationProtectedSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationProtectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ConfigurationProtectedSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationProtectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ConfigurationSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ConfigurationSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ExtensionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ExtensionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Plan() KubernetesClusterExtensionPlanOutputReference {
	var returns KubernetesClusterExtensionPlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) PlanInput() *KubernetesClusterExtensionPlan {
	var returns *KubernetesClusterExtensionPlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ReleaseNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ReleaseNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ReleaseTrain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTrain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) ReleaseTrainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTrainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TargetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TargetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Timeouts() KubernetesClusterExtensionTimeoutsOutputReference {
	var returns KubernetesClusterExtensionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterExtension) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kubernetes_cluster_extension azurerm_kubernetes_cluster_extension} Resource.
func NewKubernetesClusterExtension(scope constructs.Construct, id *string, config *KubernetesClusterExtensionConfig) KubernetesClusterExtension {
	_init_.Initialize()

	if err := validateNewKubernetesClusterExtensionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterExtension{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/kubernetes_cluster_extension azurerm_kubernetes_cluster_extension} Resource.
func NewKubernetesClusterExtension_Override(k KubernetesClusterExtension, scope constructs.Construct, id *string, config *KubernetesClusterExtensionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetConfigurationProtectedSettings(val *map[string]*string) {
	if err := j.validateSetConfigurationProtectedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationProtectedSettings",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetConfigurationSettings(val *map[string]*string) {
	if err := j.validateSetConfigurationSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSettings",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetExtensionType(val *string) {
	if err := j.validateSetExtensionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionType",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetReleaseNamespace(val *string) {
	if err := j.validateSetReleaseNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseNamespace",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetReleaseTrain(val *string) {
	if err := j.validateSetReleaseTrainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseTrain",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetTargetNamespace(val *string) {
	if err := j.validateSetTargetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNamespace",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterExtension)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a KubernetesClusterExtension resource upon running "cdktf plan <stack-name>".
func KubernetesClusterExtension_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesClusterExtension_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
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
func KubernetesClusterExtension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterExtension_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesClusterExtension_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterExtension_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesClusterExtension_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesClusterExtension_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesClusterExtension_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kubernetesClusterExtension.KubernetesClusterExtension",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) PutPlan(value *KubernetesClusterExtensionPlan) {
	if err := k.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPlan",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) PutTimeouts(value *KubernetesClusterExtensionTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetConfigurationProtectedSettings() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigurationProtectedSettings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetConfigurationSettings() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigurationSettings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetPlan() {
	_jsii_.InvokeVoid(
		k,
		"resetPlan",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetReleaseNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetReleaseNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetReleaseTrain() {
	_jsii_.InvokeVoid(
		k,
		"resetReleaseTrain",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetTargetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetTargetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) ResetVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterExtension) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterExtension) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

