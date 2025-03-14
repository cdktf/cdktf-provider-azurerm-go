// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesclusterextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/arckubernetesclusterextension/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/arc_kubernetes_cluster_extension azurerm_arc_kubernetes_cluster_extension}.
type ArcKubernetesClusterExtension interface {
	cdktf.TerraformResource
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
	Identity() ArcKubernetesClusterExtensionIdentityOutputReference
	IdentityInput() *ArcKubernetesClusterExtensionIdentity
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
	Timeouts() ArcKubernetesClusterExtensionTimeoutsOutputReference
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
	PutIdentity(value *ArcKubernetesClusterExtensionIdentity)
	PutTimeouts(value *ArcKubernetesClusterExtensionTimeouts)
	ResetConfigurationProtectedSettings()
	ResetConfigurationSettings()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for ArcKubernetesClusterExtension
type jsiiProxy_ArcKubernetesClusterExtension struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ConfigurationProtectedSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationProtectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ConfigurationProtectedSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationProtectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ConfigurationSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ConfigurationSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"configurationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ExtensionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ExtensionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Identity() ArcKubernetesClusterExtensionIdentityOutputReference {
	var returns ArcKubernetesClusterExtensionIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) IdentityInput() *ArcKubernetesClusterExtensionIdentity {
	var returns *ArcKubernetesClusterExtensionIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ReleaseNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ReleaseNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ReleaseTrain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTrain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) ReleaseTrainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseTrainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TargetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TargetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Timeouts() ArcKubernetesClusterExtensionTimeoutsOutputReference {
	var returns ArcKubernetesClusterExtensionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesClusterExtension) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/arc_kubernetes_cluster_extension azurerm_arc_kubernetes_cluster_extension} Resource.
func NewArcKubernetesClusterExtension(scope constructs.Construct, id *string, config *ArcKubernetesClusterExtensionConfig) ArcKubernetesClusterExtension {
	_init_.Initialize()

	if err := validateNewArcKubernetesClusterExtensionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcKubernetesClusterExtension{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/arc_kubernetes_cluster_extension azurerm_arc_kubernetes_cluster_extension} Resource.
func NewArcKubernetesClusterExtension_Override(a ArcKubernetesClusterExtension, scope constructs.Construct, id *string, config *ArcKubernetesClusterExtensionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetConfigurationProtectedSettings(val *map[string]*string) {
	if err := j.validateSetConfigurationProtectedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationProtectedSettings",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetConfigurationSettings(val *map[string]*string) {
	if err := j.validateSetConfigurationSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationSettings",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetExtensionType(val *string) {
	if err := j.validateSetExtensionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionType",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetReleaseNamespace(val *string) {
	if err := j.validateSetReleaseNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseNamespace",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetReleaseTrain(val *string) {
	if err := j.validateSetReleaseTrainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseTrain",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetTargetNamespace(val *string) {
	if err := j.validateSetTargetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetNamespace",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesClusterExtension)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a ArcKubernetesClusterExtension resource upon running "cdktf plan <stack-name>".
func ArcKubernetesClusterExtension_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateArcKubernetesClusterExtension_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
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
func ArcKubernetesClusterExtension_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArcKubernetesClusterExtension_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArcKubernetesClusterExtension_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArcKubernetesClusterExtension_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArcKubernetesClusterExtension_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArcKubernetesClusterExtension_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArcKubernetesClusterExtension_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.arcKubernetesClusterExtension.ArcKubernetesClusterExtension",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesClusterExtension) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) PutIdentity(value *ArcKubernetesClusterExtensionIdentity) {
	if err := a.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) PutTimeouts(value *ArcKubernetesClusterExtensionTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetConfigurationProtectedSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationProtectedSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetConfigurationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetReleaseNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetReleaseNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetReleaseTrain() {
	_jsii_.InvokeVoid(
		a,
		"resetReleaseTrain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetTargetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ResetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesClusterExtension) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

