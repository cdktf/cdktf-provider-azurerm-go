// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustocluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/kustocluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/kusto_cluster azurerm_kusto_cluster}.
type KustoCluster interface {
	cdktf.TerraformResource
	AllowedFqdns() *[]*string
	SetAllowedFqdns(val *[]*string)
	AllowedFqdnsInput() *[]*string
	AllowedIpRanges() *[]*string
	SetAllowedIpRanges(val *[]*string)
	AllowedIpRangesInput() *[]*string
	AutoStopEnabled() interface{}
	SetAutoStopEnabled(val interface{})
	AutoStopEnabledInput() interface{}
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
	DataIngestionUri() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskEncryptionEnabled() interface{}
	SetDiskEncryptionEnabled(val interface{})
	DiskEncryptionEnabledInput() interface{}
	DoubleEncryptionEnabled() interface{}
	SetDoubleEncryptionEnabled(val interface{})
	DoubleEncryptionEnabledInput() interface{}
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
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
	Identity() KustoClusterIdentityOutputReference
	IdentityInput() *KustoClusterIdentity
	IdInput() *string
	LanguageExtensions() *[]*string
	SetLanguageExtensions(val *[]*string)
	LanguageExtensionsInput() *[]*string
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
	OptimizedAutoScale() KustoClusterOptimizedAutoScaleOutputReference
	OptimizedAutoScaleInput() *KustoClusterOptimizedAutoScale
	OutboundNetworkAccessRestricted() interface{}
	SetOutboundNetworkAccessRestricted(val interface{})
	OutboundNetworkAccessRestrictedInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicIpType() *string
	SetPublicIpType(val *string)
	PublicIpTypeInput() *string
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	PurgeEnabled() interface{}
	SetPurgeEnabled(val interface{})
	PurgeEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Sku() KustoClusterSkuOutputReference
	SkuInput() *KustoClusterSku
	StreamingIngestionEnabled() interface{}
	SetStreamingIngestionEnabled(val interface{})
	StreamingIngestionEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustedExternalTenants() *[]*string
	SetTrustedExternalTenants(val *[]*string)
	TrustedExternalTenantsInput() *[]*string
	Uri() *string
	VirtualNetworkConfiguration() KustoClusterVirtualNetworkConfigurationOutputReference
	VirtualNetworkConfigurationInput() *KustoClusterVirtualNetworkConfiguration
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	PutIdentity(value *KustoClusterIdentity)
	PutOptimizedAutoScale(value *KustoClusterOptimizedAutoScale)
	PutSku(value *KustoClusterSku)
	PutTimeouts(value *KustoClusterTimeouts)
	PutVirtualNetworkConfiguration(value *KustoClusterVirtualNetworkConfiguration)
	ResetAllowedFqdns()
	ResetAllowedIpRanges()
	ResetAutoStopEnabled()
	ResetDiskEncryptionEnabled()
	ResetDoubleEncryptionEnabled()
	ResetEngine()
	ResetId()
	ResetIdentity()
	ResetLanguageExtensions()
	ResetOptimizedAutoScale()
	ResetOutboundNetworkAccessRestricted()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicIpType()
	ResetPublicNetworkAccessEnabled()
	ResetPurgeEnabled()
	ResetStreamingIngestionEnabled()
	ResetTags()
	ResetTimeouts()
	ResetTrustedExternalTenants()
	ResetVirtualNetworkConfiguration()
	ResetZones()
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

// The jsii proxy struct for KustoCluster
type jsiiProxy_KustoCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KustoCluster) AllowedFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) AllowedFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) AllowedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) AllowedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) AutoStopEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStopEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) AutoStopEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStopEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DataIngestionUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataIngestionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DiskEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DiskEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DoubleEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doubleEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) DoubleEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doubleEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Identity() KustoClusterIdentityOutputReference {
	var returns KustoClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) IdentityInput() *KustoClusterIdentity {
	var returns *KustoClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) LanguageExtensions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"languageExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) LanguageExtensionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"languageExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) OptimizedAutoScale() KustoClusterOptimizedAutoScaleOutputReference {
	var returns KustoClusterOptimizedAutoScaleOutputReference
	_jsii_.Get(
		j,
		"optimizedAutoScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) OptimizedAutoScaleInput() *KustoClusterOptimizedAutoScale {
	var returns *KustoClusterOptimizedAutoScale
	_jsii_.Get(
		j,
		"optimizedAutoScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) OutboundNetworkAccessRestricted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestricted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) OutboundNetworkAccessRestrictedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outboundNetworkAccessRestrictedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PublicIpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PublicIpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PurgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) PurgeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Sku() KustoClusterSkuOutputReference {
	var returns KustoClusterSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) SkuInput() *KustoClusterSku {
	var returns *KustoClusterSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) StreamingIngestionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingIngestionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) StreamingIngestionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingIngestionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Timeouts() KustoClusterTimeoutsOutputReference {
	var returns KustoClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TrustedExternalTenants() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedExternalTenants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) TrustedExternalTenantsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedExternalTenantsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) VirtualNetworkConfiguration() KustoClusterVirtualNetworkConfigurationOutputReference {
	var returns KustoClusterVirtualNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"virtualNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) VirtualNetworkConfigurationInput() *KustoClusterVirtualNetworkConfiguration {
	var returns *KustoClusterVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"virtualNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCluster) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/kusto_cluster azurerm_kusto_cluster} Resource.
func NewKustoCluster(scope constructs.Construct, id *string, config *KustoClusterConfig) KustoCluster {
	_init_.Initialize()

	if err := validateNewKustoClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KustoCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/kusto_cluster azurerm_kusto_cluster} Resource.
func NewKustoCluster_Override(k KustoCluster, scope constructs.Construct, id *string, config *KustoClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoCluster)SetAllowedFqdns(val *[]*string) {
	if err := j.validateSetAllowedFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedFqdns",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetAllowedIpRanges(val *[]*string) {
	if err := j.validateSetAllowedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIpRanges",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetAutoStopEnabled(val interface{}) {
	if err := j.validateSetAutoStopEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStopEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetDiskEncryptionEnabled(val interface{}) {
	if err := j.validateSetDiskEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetDoubleEncryptionEnabled(val interface{}) {
	if err := j.validateSetDoubleEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doubleEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetLanguageExtensions(val *[]*string) {
	if err := j.validateSetLanguageExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageExtensions",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetOutboundNetworkAccessRestricted(val interface{}) {
	if err := j.validateSetOutboundNetworkAccessRestrictedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundNetworkAccessRestricted",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetPublicIpType(val *string) {
	if err := j.validateSetPublicIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpType",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetPurgeEnabled(val interface{}) {
	if err := j.validateSetPurgeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetStreamingIngestionEnabled(val interface{}) {
	if err := j.validateSetStreamingIngestionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingIngestionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetTrustedExternalTenants(val *[]*string) {
	if err := j.validateSetTrustedExternalTenantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedExternalTenants",
		val,
	)
}

func (j *jsiiProxy_KustoCluster)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a KustoCluster resource upon running "cdktf plan <stack-name>".
func KustoCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKustoCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
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
func KustoCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoCluster) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KustoCluster) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KustoCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KustoCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KustoCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KustoCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KustoCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KustoCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KustoCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KustoCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KustoCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KustoCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KustoCluster) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KustoCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KustoCluster) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KustoCluster) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoCluster) PutIdentity(value *KustoClusterIdentity) {
	if err := k.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutOptimizedAutoScale(value *KustoClusterOptimizedAutoScale) {
	if err := k.validatePutOptimizedAutoScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putOptimizedAutoScale",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutSku(value *KustoClusterSku) {
	if err := k.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSku",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutTimeouts(value *KustoClusterTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutVirtualNetworkConfiguration(value *KustoClusterVirtualNetworkConfiguration) {
	if err := k.validatePutVirtualNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putVirtualNetworkConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) ResetAllowedFqdns() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedFqdns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetAllowedIpRanges() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedIpRanges",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetAutoStopEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetAutoStopEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetDiskEncryptionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetDiskEncryptionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetDoubleEncryptionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetDoubleEncryptionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetEngine() {
	_jsii_.InvokeVoid(
		k,
		"resetEngine",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetLanguageExtensions() {
	_jsii_.InvokeVoid(
		k,
		"resetLanguageExtensions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetOptimizedAutoScale() {
	_jsii_.InvokeVoid(
		k,
		"resetOptimizedAutoScale",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetOutboundNetworkAccessRestricted() {
	_jsii_.InvokeVoid(
		k,
		"resetOutboundNetworkAccessRestricted",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetPublicIpType() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicIpType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetPurgeEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPurgeEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetStreamingIngestionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetStreamingIngestionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetTrustedExternalTenants() {
	_jsii_.InvokeVoid(
		k,
		"resetTrustedExternalTenants",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetVirtualNetworkConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetVirtualNetworkConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) ResetZones() {
	_jsii_.InvokeVoid(
		k,
		"resetZones",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

