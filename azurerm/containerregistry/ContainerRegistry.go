// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistry

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/containerregistry/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_registry azurerm_container_registry}.
type ContainerRegistry interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	AdminPassword() *string
	AdminUsername() *string
	AnonymousPullEnabled() interface{}
	SetAnonymousPullEnabled(val interface{})
	AnonymousPullEnabledInput() interface{}
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
	DataEndpointEnabled() interface{}
	SetDataEndpointEnabled(val interface{})
	DataEndpointEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Encryption() ContainerRegistryEncryptionList
	EncryptionInput() interface{}
	ExportPolicyEnabled() interface{}
	SetExportPolicyEnabled(val interface{})
	ExportPolicyEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Georeplications() ContainerRegistryGeoreplicationsList
	GeoreplicationsInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() ContainerRegistryIdentityOutputReference
	IdentityInput() *ContainerRegistryIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoginServer() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkRuleBypassOption() *string
	SetNetworkRuleBypassOption(val *string)
	NetworkRuleBypassOptionInput() *string
	NetworkRuleSet() ContainerRegistryNetworkRuleSetList
	NetworkRuleSetInput() interface{}
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
	QuarantinePolicyEnabled() interface{}
	SetQuarantinePolicyEnabled(val interface{})
	QuarantinePolicyEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RetentionPolicyInDays() *float64
	SetRetentionPolicyInDays(val *float64)
	RetentionPolicyInDaysInput() *float64
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerRegistryTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustPolicyEnabled() interface{}
	SetTrustPolicyEnabled(val interface{})
	TrustPolicyEnabledInput() interface{}
	ZoneRedundancyEnabled() interface{}
	SetZoneRedundancyEnabled(val interface{})
	ZoneRedundancyEnabledInput() interface{}
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
	PutEncryption(value interface{})
	PutGeoreplications(value interface{})
	PutIdentity(value *ContainerRegistryIdentity)
	PutNetworkRuleSet(value interface{})
	PutTimeouts(value *ContainerRegistryTimeouts)
	ResetAdminEnabled()
	ResetAnonymousPullEnabled()
	ResetDataEndpointEnabled()
	ResetEncryption()
	ResetExportPolicyEnabled()
	ResetGeoreplications()
	ResetId()
	ResetIdentity()
	ResetNetworkRuleBypassOption()
	ResetNetworkRuleSet()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetQuarantinePolicyEnabled()
	ResetRetentionPolicyInDays()
	ResetTags()
	ResetTimeouts()
	ResetTrustPolicyEnabled()
	ResetZoneRedundancyEnabled()
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

// The jsii proxy struct for ContainerRegistry
type jsiiProxy_ContainerRegistry struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerRegistry) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) AnonymousPullEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousPullEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) AnonymousPullEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousPullEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) DataEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) DataEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Encryption() ContainerRegistryEncryptionList {
	var returns ContainerRegistryEncryptionList
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) EncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ExportPolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportPolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ExportPolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportPolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Georeplications() ContainerRegistryGeoreplicationsList {
	var returns ContainerRegistryGeoreplicationsList
	_jsii_.Get(
		j,
		"georeplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) GeoreplicationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"georeplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Identity() ContainerRegistryIdentityOutputReference {
	var returns ContainerRegistryIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) IdentityInput() *ContainerRegistryIdentity {
	var returns *ContainerRegistryIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) LoginServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) NetworkRuleBypassOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkRuleBypassOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) NetworkRuleBypassOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkRuleBypassOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) NetworkRuleSet() ContainerRegistryNetworkRuleSetList {
	var returns ContainerRegistryNetworkRuleSetList
	_jsii_.Get(
		j,
		"networkRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) NetworkRuleSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) QuarantinePolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quarantinePolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) QuarantinePolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"quarantinePolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) RetentionPolicyInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) RetentionPolicyInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPolicyInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) Timeouts() ContainerRegistryTimeoutsOutputReference {
	var returns ContainerRegistryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TrustPolicyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustPolicyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) TrustPolicyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustPolicyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ZoneRedundancyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundancyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerRegistry) ZoneRedundancyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneRedundancyEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_registry azurerm_container_registry} Resource.
func NewContainerRegistry(scope constructs.Construct, id *string, config *ContainerRegistryConfig) ContainerRegistry {
	_init_.Initialize()

	if err := validateNewContainerRegistryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerRegistry{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/container_registry azurerm_container_registry} Resource.
func NewContainerRegistry_Override(c ContainerRegistry, scope constructs.Construct, id *string, config *ContainerRegistryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetAnonymousPullEnabled(val interface{}) {
	if err := j.validateSetAnonymousPullEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousPullEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetDataEndpointEnabled(val interface{}) {
	if err := j.validateSetDataEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetExportPolicyEnabled(val interface{}) {
	if err := j.validateSetExportPolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetNetworkRuleBypassOption(val *string) {
	if err := j.validateSetNetworkRuleBypassOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkRuleBypassOption",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetQuarantinePolicyEnabled(val interface{}) {
	if err := j.validateSetQuarantinePolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quarantinePolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetRetentionPolicyInDays(val *float64) {
	if err := j.validateSetRetentionPolicyInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPolicyInDays",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetTrustPolicyEnabled(val interface{}) {
	if err := j.validateSetTrustPolicyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustPolicyEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerRegistry)SetZoneRedundancyEnabled(val interface{}) {
	if err := j.validateSetZoneRedundancyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneRedundancyEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ContainerRegistry resource upon running "cdktf plan <stack-name>".
func ContainerRegistry_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerRegistry_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
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
func ContainerRegistry_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistry_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerRegistry_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistry_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerRegistry_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerRegistry_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerRegistry_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.containerRegistry.ContainerRegistry",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerRegistry) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerRegistry) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerRegistry) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerRegistry) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerRegistry) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerRegistry) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerRegistry) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerRegistry) PutEncryption(value interface{}) {
	if err := c.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistry) PutGeoreplications(value interface{}) {
	if err := c.validatePutGeoreplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGeoreplications",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistry) PutIdentity(value *ContainerRegistryIdentity) {
	if err := c.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIdentity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistry) PutNetworkRuleSet(value interface{}) {
	if err := c.validatePutNetworkRuleSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkRuleSet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistry) PutTimeouts(value *ContainerRegistryTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetAnonymousPullEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAnonymousPullEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetDataEndpointEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDataEndpointEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetExportPolicyEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetExportPolicyEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetGeoreplications() {
	_jsii_.InvokeVoid(
		c,
		"resetGeoreplications",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetIdentity() {
	_jsii_.InvokeVoid(
		c,
		"resetIdentity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetNetworkRuleBypassOption() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkRuleBypassOption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetNetworkRuleSet() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkRuleSet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetQuarantinePolicyEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetQuarantinePolicyEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetRetentionPolicyInDays() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionPolicyInDays",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetTrustPolicyEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustPolicyEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) ResetZoneRedundancyEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetZoneRedundancyEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerRegistry) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerRegistry) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

