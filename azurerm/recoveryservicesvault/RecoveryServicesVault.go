// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/recoveryservicesvault/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/recovery_services_vault azurerm_recovery_services_vault}.
type RecoveryServicesVault interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClassicVmwareReplicationEnabled() interface{}
	SetClassicVmwareReplicationEnabled(val interface{})
	ClassicVmwareReplicationEnabledInput() interface{}
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
	CrossRegionRestoreEnabled() interface{}
	SetCrossRegionRestoreEnabled(val interface{})
	CrossRegionRestoreEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encryption() RecoveryServicesVaultEncryptionOutputReference
	EncryptionInput() *RecoveryServicesVaultEncryption
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
	Identity() RecoveryServicesVaultIdentityOutputReference
	IdentityInput() *RecoveryServicesVaultIdentity
	IdInput() *string
	Immutability() *string
	SetImmutability(val *string)
	ImmutabilityInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Monitoring() RecoveryServicesVaultMonitoringOutputReference
	MonitoringInput() *RecoveryServicesVaultMonitoring
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SoftDeleteEnabled() interface{}
	SetSoftDeleteEnabled(val interface{})
	SoftDeleteEnabledInput() interface{}
	StorageModeType() *string
	SetStorageModeType(val *string)
	StorageModeTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RecoveryServicesVaultTimeoutsOutputReference
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutEncryption(value *RecoveryServicesVaultEncryption)
	PutIdentity(value *RecoveryServicesVaultIdentity)
	PutMonitoring(value *RecoveryServicesVaultMonitoring)
	PutTimeouts(value *RecoveryServicesVaultTimeouts)
	ResetClassicVmwareReplicationEnabled()
	ResetCrossRegionRestoreEnabled()
	ResetEncryption()
	ResetId()
	ResetIdentity()
	ResetImmutability()
	ResetMonitoring()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetSoftDeleteEnabled()
	ResetStorageModeType()
	ResetTags()
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

// The jsii proxy struct for RecoveryServicesVault
type jsiiProxy_RecoveryServicesVault struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RecoveryServicesVault) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ClassicVmwareReplicationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classicVmwareReplicationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ClassicVmwareReplicationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classicVmwareReplicationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) CrossRegionRestoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionRestoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) CrossRegionRestoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossRegionRestoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Encryption() RecoveryServicesVaultEncryptionOutputReference {
	var returns RecoveryServicesVaultEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) EncryptionInput() *RecoveryServicesVaultEncryption {
	var returns *RecoveryServicesVaultEncryption
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Identity() RecoveryServicesVaultIdentityOutputReference {
	var returns RecoveryServicesVaultIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) IdentityInput() *RecoveryServicesVaultIdentity {
	var returns *RecoveryServicesVaultIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Immutability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"immutability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ImmutabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"immutabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Monitoring() RecoveryServicesVaultMonitoringOutputReference {
	var returns RecoveryServicesVaultMonitoringOutputReference
	_jsii_.Get(
		j,
		"monitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) MonitoringInput() *RecoveryServicesVaultMonitoring {
	var returns *RecoveryServicesVaultMonitoring
	_jsii_.Get(
		j,
		"monitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) SoftDeleteEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softDeleteEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) SoftDeleteEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"softDeleteEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) StorageModeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageModeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) StorageModeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageModeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) Timeouts() RecoveryServicesVaultTimeoutsOutputReference {
	var returns RecoveryServicesVaultTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RecoveryServicesVault) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/recovery_services_vault azurerm_recovery_services_vault} Resource.
func NewRecoveryServicesVault(scope constructs.Construct, id *string, config *RecoveryServicesVaultConfig) RecoveryServicesVault {
	_init_.Initialize()

	if err := validateNewRecoveryServicesVaultParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecoveryServicesVault{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/recovery_services_vault azurerm_recovery_services_vault} Resource.
func NewRecoveryServicesVault_Override(r RecoveryServicesVault, scope constructs.Construct, id *string, config *RecoveryServicesVaultConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetClassicVmwareReplicationEnabled(val interface{}) {
	if err := j.validateSetClassicVmwareReplicationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classicVmwareReplicationEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetCrossRegionRestoreEnabled(val interface{}) {
	if err := j.validateSetCrossRegionRestoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRegionRestoreEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetImmutability(val *string) {
	if err := j.validateSetImmutabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"immutability",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetSoftDeleteEnabled(val interface{}) {
	if err := j.validateSetSoftDeleteEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softDeleteEnabled",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetStorageModeType(val *string) {
	if err := j.validateSetStorageModeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageModeType",
		val,
	)
}

func (j *jsiiProxy_RecoveryServicesVault)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a RecoveryServicesVault resource upon running "cdktf plan <stack-name>".
func RecoveryServicesVault_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRecoveryServicesVault_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
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
func RecoveryServicesVault_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecoveryServicesVault_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RecoveryServicesVault_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecoveryServicesVault_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RecoveryServicesVault_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRecoveryServicesVault_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RecoveryServicesVault_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.recoveryServicesVault.RecoveryServicesVault",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RecoveryServicesVault) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RecoveryServicesVault) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVault) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RecoveryServicesVault) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RecoveryServicesVault) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RecoveryServicesVault) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVault) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RecoveryServicesVault) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RecoveryServicesVault) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RecoveryServicesVault) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RecoveryServicesVault) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) PutEncryption(value *RecoveryServicesVaultEncryption) {
	if err := r.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEncryption",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) PutIdentity(value *RecoveryServicesVaultIdentity) {
	if err := r.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) PutMonitoring(value *RecoveryServicesVaultMonitoring) {
	if err := r.validatePutMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMonitoring",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) PutTimeouts(value *RecoveryServicesVaultTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetClassicVmwareReplicationEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetClassicVmwareReplicationEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetCrossRegionRestoreEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetCrossRegionRestoreEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetEncryption() {
	_jsii_.InvokeVoid(
		r,
		"resetEncryption",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetImmutability() {
	_jsii_.InvokeVoid(
		r,
		"resetImmutability",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetMonitoring() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoring",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetSoftDeleteEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetSoftDeleteEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetStorageModeType() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageModeType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RecoveryServicesVault) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVault) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVault) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryServicesVault) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

