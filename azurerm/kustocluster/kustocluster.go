package kustocluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kustocluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster azurerm_kusto_cluster}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_KustoCluster) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster azurerm_kusto_cluster} Resource.
func NewKustoCluster(scope constructs.Construct, id *string, config *KustoClusterConfig) KustoCluster {
	_init_.Initialize()

	j := jsiiProxy_KustoCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster azurerm_kusto_cluster} Resource.
func NewKustoCluster_Override(k KustoCluster, scope constructs.Construct, id *string, config *KustoClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoCluster) SetAllowedFqdns(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedFqdns",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetAllowedIpRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedIpRanges",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetAutoStopEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoStopEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetDiskEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"diskEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetDoubleEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"doubleEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetEngine(val *string) {
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetLanguageExtensions(val *[]*string) {
	_jsii_.Set(
		j,
		"languageExtensions",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetOutboundNetworkAccessRestricted(val interface{}) {
	_jsii_.Set(
		j,
		"outboundNetworkAccessRestricted",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetPublicIpType(val *string) {
	_jsii_.Set(
		j,
		"publicIpType",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetPurgeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"purgeEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetStreamingIngestionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"streamingIngestionEnabled",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetTrustedExternalTenants(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedExternalTenants",
		val,
	)
}

func (j *jsiiProxy_KustoCluster) SetZones(val *[]*string) {
	_jsii_.Set(
		j,
		"zones",
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
func KustoCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.kustoCluster.KustoCluster",
		"isConstruct",
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

func (k *jsiiProxy_KustoCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoCluster) PutIdentity(value *KustoClusterIdentity) {
	_jsii_.InvokeVoid(
		k,
		"putIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutOptimizedAutoScale(value *KustoClusterOptimizedAutoScale) {
	_jsii_.InvokeVoid(
		k,
		"putOptimizedAutoScale",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutSku(value *KustoClusterSku) {
	_jsii_.InvokeVoid(
		k,
		"putSku",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutTimeouts(value *KustoClusterTimeouts) {
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCluster) PutVirtualNetworkConfiguration(value *KustoClusterVirtualNetworkConfiguration) {
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

type KustoClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#location KustoCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#name KustoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#resource_group_name KustoCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// sku block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#sku KustoCluster#sku}
	Sku *KustoClusterSku `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#allowed_fqdns KustoCluster#allowed_fqdns}.
	AllowedFqdns *[]*string `field:"optional" json:"allowedFqdns" yaml:"allowedFqdns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#allowed_ip_ranges KustoCluster#allowed_ip_ranges}.
	AllowedIpRanges *[]*string `field:"optional" json:"allowedIpRanges" yaml:"allowedIpRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#auto_stop_enabled KustoCluster#auto_stop_enabled}.
	AutoStopEnabled interface{} `field:"optional" json:"autoStopEnabled" yaml:"autoStopEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#disk_encryption_enabled KustoCluster#disk_encryption_enabled}.
	DiskEncryptionEnabled interface{} `field:"optional" json:"diskEncryptionEnabled" yaml:"diskEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#double_encryption_enabled KustoCluster#double_encryption_enabled}.
	DoubleEncryptionEnabled interface{} `field:"optional" json:"doubleEncryptionEnabled" yaml:"doubleEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#engine KustoCluster#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#id KustoCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#identity KustoCluster#identity}
	Identity *KustoClusterIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#language_extensions KustoCluster#language_extensions}.
	LanguageExtensions *[]*string `field:"optional" json:"languageExtensions" yaml:"languageExtensions"`
	// optimized_auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#optimized_auto_scale KustoCluster#optimized_auto_scale}
	OptimizedAutoScale *KustoClusterOptimizedAutoScale `field:"optional" json:"optimizedAutoScale" yaml:"optimizedAutoScale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#outbound_network_access_restricted KustoCluster#outbound_network_access_restricted}.
	OutboundNetworkAccessRestricted interface{} `field:"optional" json:"outboundNetworkAccessRestricted" yaml:"outboundNetworkAccessRestricted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#public_ip_type KustoCluster#public_ip_type}.
	PublicIpType *string `field:"optional" json:"publicIpType" yaml:"publicIpType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#public_network_access_enabled KustoCluster#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#purge_enabled KustoCluster#purge_enabled}.
	PurgeEnabled interface{} `field:"optional" json:"purgeEnabled" yaml:"purgeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#streaming_ingestion_enabled KustoCluster#streaming_ingestion_enabled}.
	StreamingIngestionEnabled interface{} `field:"optional" json:"streamingIngestionEnabled" yaml:"streamingIngestionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#tags KustoCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#timeouts KustoCluster#timeouts}
	Timeouts *KustoClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#trusted_external_tenants KustoCluster#trusted_external_tenants}.
	TrustedExternalTenants *[]*string `field:"optional" json:"trustedExternalTenants" yaml:"trustedExternalTenants"`
	// virtual_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#virtual_network_configuration KustoCluster#virtual_network_configuration}
	VirtualNetworkConfiguration *KustoClusterVirtualNetworkConfiguration `field:"optional" json:"virtualNetworkConfiguration" yaml:"virtualNetworkConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#zones KustoCluster#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type KustoClusterIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#type KustoCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#identity_ids KustoCluster#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type KustoClusterIdentityOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IdentityIds() *[]*string
	SetIdentityIds(val *[]*string)
	IdentityIdsInput() *[]*string
	InternalValue() *KustoClusterIdentity
	SetInternalValue(val *KustoClusterIdentity)
	PrincipalId() *string
	TenantId() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIdentityIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoClusterIdentityOutputReference
type jsiiProxy_KustoClusterIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) InternalValue() *KustoClusterIdentity {
	var returns *KustoClusterIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewKustoClusterIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoClusterIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoClusterIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoClusterIdentityOutputReference_Override(k KustoClusterIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetInternalValue(val *KustoClusterIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KustoClusterIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		k,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoClusterOptimizedAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#maximum_instances KustoCluster#maximum_instances}.
	MaximumInstances *float64 `field:"required" json:"maximumInstances" yaml:"maximumInstances"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#minimum_instances KustoCluster#minimum_instances}.
	MinimumInstances *float64 `field:"required" json:"minimumInstances" yaml:"minimumInstances"`
}

type KustoClusterOptimizedAutoScaleOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KustoClusterOptimizedAutoScale
	SetInternalValue(val *KustoClusterOptimizedAutoScale)
	MaximumInstances() *float64
	SetMaximumInstances(val *float64)
	MaximumInstancesInput() *float64
	MinimumInstances() *float64
	SetMinimumInstances(val *float64)
	MinimumInstancesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoClusterOptimizedAutoScaleOutputReference
type jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) InternalValue() *KustoClusterOptimizedAutoScale {
	var returns *KustoClusterOptimizedAutoScale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) MaximumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) MaximumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) MinimumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) MinimumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKustoClusterOptimizedAutoScaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoClusterOptimizedAutoScaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterOptimizedAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoClusterOptimizedAutoScaleOutputReference_Override(k KustoClusterOptimizedAutoScaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterOptimizedAutoScaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetInternalValue(val *KustoClusterOptimizedAutoScale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetMaximumInstances(val *float64) {
	_jsii_.Set(
		j,
		"maximumInstances",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetMinimumInstances(val *float64) {
	_jsii_.Set(
		j,
		"minimumInstances",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterOptimizedAutoScaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoClusterSku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#name KustoCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#capacity KustoCluster#capacity}.
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}

type KustoClusterSkuOutputReference interface {
	cdktf.ComplexObject
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *KustoClusterSku
	SetInternalValue(val *KustoClusterSku)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoClusterSkuOutputReference
type jsiiProxy_KustoClusterSkuOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) InternalValue() *KustoClusterSku {
	var returns *KustoClusterSku
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKustoClusterSkuOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoClusterSkuOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoClusterSkuOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterSkuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoClusterSkuOutputReference_Override(k KustoClusterSkuOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterSkuOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetInternalValue(val *KustoClusterSku) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoClusterSkuOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		k,
		"resetCapacity",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterSkuOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#create KustoCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#delete KustoCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#read KustoCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#update KustoCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type KustoClusterTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoClusterTimeoutsOutputReference
type jsiiProxy_KustoClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewKustoClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoClusterTimeoutsOutputReference_Override(k KustoClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KustoClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		k,
		"resetRead",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		k,
		"resetUpdate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type KustoClusterVirtualNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#data_management_public_ip_id KustoCluster#data_management_public_ip_id}.
	DataManagementPublicIpId *string `field:"required" json:"dataManagementPublicIpId" yaml:"dataManagementPublicIpId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#engine_public_ip_id KustoCluster#engine_public_ip_id}.
	EnginePublicIpId *string `field:"required" json:"enginePublicIpId" yaml:"enginePublicIpId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kusto_cluster#subnet_id KustoCluster#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

type KustoClusterVirtualNetworkConfigurationOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataManagementPublicIpId() *string
	SetDataManagementPublicIpId(val *string)
	DataManagementPublicIpIdInput() *string
	EnginePublicIpId() *string
	SetEnginePublicIpId(val *string)
	EnginePublicIpIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KustoClusterVirtualNetworkConfiguration
	SetInternalValue(val *KustoClusterVirtualNetworkConfiguration)
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KustoClusterVirtualNetworkConfigurationOutputReference
type jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) DataManagementPublicIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataManagementPublicIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) DataManagementPublicIpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataManagementPublicIpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) EnginePublicIpId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enginePublicIpId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) EnginePublicIpIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enginePublicIpIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) InternalValue() *KustoClusterVirtualNetworkConfiguration {
	var returns *KustoClusterVirtualNetworkConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKustoClusterVirtualNetworkConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KustoClusterVirtualNetworkConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKustoClusterVirtualNetworkConfigurationOutputReference_Override(k KustoClusterVirtualNetworkConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kustoCluster.KustoClusterVirtualNetworkConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetDataManagementPublicIpId(val *string) {
	_jsii_.Set(
		j,
		"dataManagementPublicIpId",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetEnginePublicIpId(val *string) {
	_jsii_.Set(
		j,
		"enginePublicIpId",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetInternalValue(val *KustoClusterVirtualNetworkConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoClusterVirtualNetworkConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

