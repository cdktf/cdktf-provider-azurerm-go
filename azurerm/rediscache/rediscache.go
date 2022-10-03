package rediscache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/rediscache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache azurerm_redis_cache}.
type RedisCache interface {
	cdktf.TerraformResource
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableNonSslPort() interface{}
	SetEnableNonSslPort(val interface{})
	EnableNonSslPortInput() interface{}
	Family() *string
	SetFamily(val *string)
	FamilyInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	Id() *string
	SetId(val *string)
	Identity() RedisCacheIdentityOutputReference
	IdentityInput() *RedisCacheIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MinimumTlsVersion() *string
	SetMinimumTlsVersion(val *string)
	MinimumTlsVersionInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PatchSchedule() RedisCachePatchScheduleList
	PatchScheduleInput() interface{}
	Port() *float64
	PrimaryAccessKey() *string
	PrimaryConnectionString() *string
	PrivateStaticIpAddress() *string
	SetPrivateStaticIpAddress(val *string)
	PrivateStaticIpAddressInput() *string
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
	RedisConfiguration() RedisCacheRedisConfigurationOutputReference
	RedisConfigurationInput() *RedisCacheRedisConfiguration
	RedisVersion() *string
	SetRedisVersion(val *string)
	RedisVersionInput() *string
	ReplicasPerMaster() *float64
	SetReplicasPerMaster(val *float64)
	ReplicasPerMasterInput() *float64
	ReplicasPerPrimary() *float64
	SetReplicasPerPrimary(val *float64)
	ReplicasPerPrimaryInput() *float64
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryAccessKey() *string
	SecondaryConnectionString() *string
	ShardCount() *float64
	SetShardCount(val *float64)
	ShardCountInput() *float64
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SslPort() *float64
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantSettings() *map[string]*string
	SetTenantSettings(val *map[string]*string)
	TenantSettingsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RedisCacheTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutIdentity(value *RedisCacheIdentity)
	PutPatchSchedule(value interface{})
	PutRedisConfiguration(value *RedisCacheRedisConfiguration)
	PutTimeouts(value *RedisCacheTimeouts)
	ResetEnableNonSslPort()
	ResetId()
	ResetIdentity()
	ResetMinimumTlsVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatchSchedule()
	ResetPrivateStaticIpAddress()
	ResetPublicNetworkAccessEnabled()
	ResetRedisConfiguration()
	ResetRedisVersion()
	ResetReplicasPerMaster()
	ResetReplicasPerPrimary()
	ResetShardCount()
	ResetSubnetId()
	ResetTags()
	ResetTenantSettings()
	ResetTimeouts()
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

// The jsii proxy struct for RedisCache
type jsiiProxy_RedisCache struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedisCache) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) EnableNonSslPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSslPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) EnableNonSslPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSslPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Identity() RedisCacheIdentityOutputReference {
	var returns RedisCacheIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) IdentityInput() *RedisCacheIdentity {
	var returns *RedisCacheIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) MinimumTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PatchSchedule() RedisCachePatchScheduleList {
	var returns RedisCachePatchScheduleList
	_jsii_.Get(
		j,
		"patchSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PatchScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patchScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PrivateStaticIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateStaticIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PrivateStaticIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateStaticIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) RedisConfiguration() RedisCacheRedisConfigurationOutputReference {
	var returns RedisCacheRedisConfigurationOutputReference
	_jsii_.Get(
		j,
		"redisConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) RedisConfigurationInput() *RedisCacheRedisConfiguration {
	var returns *RedisCacheRedisConfiguration
	_jsii_.Get(
		j,
		"redisConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) RedisVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) RedisVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ReplicasPerMaster() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerMaster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ReplicasPerMasterInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerMasterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ReplicasPerPrimary() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ReplicasPerPrimaryInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ShardCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ShardCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SslPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sslPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TenantSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tenantSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TenantSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tenantSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Timeouts() RedisCacheTimeoutsOutputReference {
	var returns RedisCacheTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCache) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache azurerm_redis_cache} Resource.
func NewRedisCache(scope constructs.Construct, id *string, config *RedisCacheConfig) RedisCache {
	_init_.Initialize()

	j := jsiiProxy_RedisCache{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache azurerm_redis_cache} Resource.
func NewRedisCache_Override(r RedisCache, scope constructs.Construct, id *string, config *RedisCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCache",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedisCache) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetEnableNonSslPort(val interface{}) {
	_jsii_.Set(
		j,
		"enableNonSslPort",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetMinimumTlsVersion(val *string) {
	_jsii_.Set(
		j,
		"minimumTlsVersion",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetPrivateStaticIpAddress(val *string) {
	_jsii_.Set(
		j,
		"privateStaticIpAddress",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetRedisVersion(val *string) {
	_jsii_.Set(
		j,
		"redisVersion",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetReplicasPerMaster(val *float64) {
	_jsii_.Set(
		j,
		"replicasPerMaster",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetReplicasPerPrimary(val *float64) {
	_jsii_.Set(
		j,
		"replicasPerPrimary",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetShardCount(val *float64) {
	_jsii_.Set(
		j,
		"shardCount",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetTenantSettings(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tenantSettings",
		val,
	)
}

func (j *jsiiProxy_RedisCache) SetZones(val *[]*string) {
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
func RedisCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redisCache.RedisCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedisCache_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.redisCache.RedisCache",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedisCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedisCache) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedisCache) PutIdentity(value *RedisCacheIdentity) {
	_jsii_.InvokeVoid(
		r,
		"putIdentity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCache) PutPatchSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"putPatchSchedule",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCache) PutRedisConfiguration(value *RedisCacheRedisConfiguration) {
	_jsii_.InvokeVoid(
		r,
		"putRedisConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCache) PutTimeouts(value *RedisCacheTimeouts) {
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCache) ResetEnableNonSslPort() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableNonSslPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetIdentity() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetMinimumTlsVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetMinimumTlsVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetPatchSchedule() {
	_jsii_.InvokeVoid(
		r,
		"resetPatchSchedule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetPrivateStaticIpAddress() {
	_jsii_.InvokeVoid(
		r,
		"resetPrivateStaticIpAddress",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetRedisConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetRedisConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetRedisVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetRedisVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetReplicasPerMaster() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicasPerMaster",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetReplicasPerPrimary() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicasPerPrimary",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetShardCount() {
	_jsii_.InvokeVoid(
		r,
		"resetShardCount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetSubnetId() {
	_jsii_.InvokeVoid(
		r,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetTenantSettings() {
	_jsii_.InvokeVoid(
		r,
		"resetTenantSettings",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) ResetZones() {
	_jsii_.InvokeVoid(
		r,
		"resetZones",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCache) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCache) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedisCacheConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#capacity RedisCache#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#family RedisCache#family}.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#location RedisCache#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#name RedisCache#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#resource_group_name RedisCache#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#sku_name RedisCache#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#enable_non_ssl_port RedisCache#enable_non_ssl_port}.
	EnableNonSslPort interface{} `field:"optional" json:"enableNonSslPort" yaml:"enableNonSslPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#id RedisCache#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#identity RedisCache#identity}
	Identity *RedisCacheIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#minimum_tls_version RedisCache#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// patch_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#patch_schedule RedisCache#patch_schedule}
	PatchSchedule interface{} `field:"optional" json:"patchSchedule" yaml:"patchSchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#private_static_ip_address RedisCache#private_static_ip_address}.
	PrivateStaticIpAddress *string `field:"optional" json:"privateStaticIpAddress" yaml:"privateStaticIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#public_network_access_enabled RedisCache#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// redis_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#redis_configuration RedisCache#redis_configuration}
	RedisConfiguration *RedisCacheRedisConfiguration `field:"optional" json:"redisConfiguration" yaml:"redisConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#redis_version RedisCache#redis_version}.
	RedisVersion *string `field:"optional" json:"redisVersion" yaml:"redisVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#replicas_per_master RedisCache#replicas_per_master}.
	ReplicasPerMaster *float64 `field:"optional" json:"replicasPerMaster" yaml:"replicasPerMaster"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#replicas_per_primary RedisCache#replicas_per_primary}.
	ReplicasPerPrimary *float64 `field:"optional" json:"replicasPerPrimary" yaml:"replicasPerPrimary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#shard_count RedisCache#shard_count}.
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#subnet_id RedisCache#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#tags RedisCache#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#tenant_settings RedisCache#tenant_settings}.
	TenantSettings *map[string]*string `field:"optional" json:"tenantSettings" yaml:"tenantSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#timeouts RedisCache#timeouts}
	Timeouts *RedisCacheTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#zones RedisCache#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

type RedisCacheIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#type RedisCache#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#identity_ids RedisCache#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

type RedisCacheIdentityOutputReference interface {
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
	InternalValue() *RedisCacheIdentity
	SetInternalValue(val *RedisCacheIdentity)
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

// The jsii proxy struct for RedisCacheIdentityOutputReference
type jsiiProxy_RedisCacheIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) InternalValue() *RedisCacheIdentity {
	var returns *RedisCacheIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) PrincipalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewRedisCacheIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisCacheIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RedisCacheIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisCacheIdentityOutputReference_Override(r RedisCacheIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetInternalValue(val *RedisCacheIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedisCacheIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) ResetIdentityIds() {
	_jsii_.InvokeVoid(
		r,
		"resetIdentityIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedisCachePatchSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#day_of_week RedisCache#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#maintenance_window RedisCache#maintenance_window}.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#start_hour_utc RedisCache#start_hour_utc}.
	StartHourUtc *float64 `field:"optional" json:"startHourUtc" yaml:"startHourUtc"`
}

type RedisCachePatchScheduleList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) RedisCachePatchScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisCachePatchScheduleList
type jsiiProxy_RedisCachePatchScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_RedisCachePatchScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRedisCachePatchScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) RedisCachePatchScheduleList {
	_init_.Initialize()

	j := jsiiProxy_RedisCachePatchScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCachePatchScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRedisCachePatchScheduleList_Override(r RedisCachePatchScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCachePatchScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_RedisCachePatchScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleList) Get(index *float64) RedisCachePatchScheduleOutputReference {
	var returns RedisCachePatchScheduleOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedisCachePatchScheduleOutputReference interface {
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
	DayOfWeek() *string
	SetDayOfWeek(val *string)
	DayOfWeekInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaintenanceWindow() *string
	SetMaintenanceWindow(val *string)
	MaintenanceWindowInput() *string
	StartHourUtc() *float64
	SetStartHourUtc(val *float64)
	StartHourUtcInput() *float64
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
	ResetMaintenanceWindow()
	ResetStartHourUtc()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisCachePatchScheduleOutputReference
type jsiiProxy_RedisCachePatchScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) StartHourUtc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) StartHourUtcInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedisCachePatchScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RedisCachePatchScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RedisCachePatchScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCachePatchScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRedisCachePatchScheduleOutputReference_Override(r RedisCachePatchScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCachePatchScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetDayOfWeek(val *string) {
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetStartHourUtc(val *float64) {
	_jsii_.Set(
		j,
		"startHourUtc",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCachePatchScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) ResetStartHourUtc() {
	_jsii_.InvokeVoid(
		r,
		"resetStartHourUtc",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCachePatchScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedisCacheRedisConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#aof_backup_enabled RedisCache#aof_backup_enabled}.
	AofBackupEnabled interface{} `field:"optional" json:"aofBackupEnabled" yaml:"aofBackupEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#aof_storage_connection_string_0 RedisCache#aof_storage_connection_string_0}.
	AofStorageConnectionString0 *string `field:"optional" json:"aofStorageConnectionString0" yaml:"aofStorageConnectionString0"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#aof_storage_connection_string_1 RedisCache#aof_storage_connection_string_1}.
	AofStorageConnectionString1 *string `field:"optional" json:"aofStorageConnectionString1" yaml:"aofStorageConnectionString1"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#enable_authentication RedisCache#enable_authentication}.
	EnableAuthentication interface{} `field:"optional" json:"enableAuthentication" yaml:"enableAuthentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#maxfragmentationmemory_reserved RedisCache#maxfragmentationmemory_reserved}.
	MaxfragmentationmemoryReserved *float64 `field:"optional" json:"maxfragmentationmemoryReserved" yaml:"maxfragmentationmemoryReserved"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#maxmemory_delta RedisCache#maxmemory_delta}.
	MaxmemoryDelta *float64 `field:"optional" json:"maxmemoryDelta" yaml:"maxmemoryDelta"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#maxmemory_policy RedisCache#maxmemory_policy}.
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#maxmemory_reserved RedisCache#maxmemory_reserved}.
	MaxmemoryReserved *float64 `field:"optional" json:"maxmemoryReserved" yaml:"maxmemoryReserved"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#notify_keyspace_events RedisCache#notify_keyspace_events}.
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#rdb_backup_enabled RedisCache#rdb_backup_enabled}.
	RdbBackupEnabled interface{} `field:"optional" json:"rdbBackupEnabled" yaml:"rdbBackupEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#rdb_backup_frequency RedisCache#rdb_backup_frequency}.
	RdbBackupFrequency *float64 `field:"optional" json:"rdbBackupFrequency" yaml:"rdbBackupFrequency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#rdb_backup_max_snapshot_count RedisCache#rdb_backup_max_snapshot_count}.
	RdbBackupMaxSnapshotCount *float64 `field:"optional" json:"rdbBackupMaxSnapshotCount" yaml:"rdbBackupMaxSnapshotCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#rdb_storage_connection_string RedisCache#rdb_storage_connection_string}.
	RdbStorageConnectionString *string `field:"optional" json:"rdbStorageConnectionString" yaml:"rdbStorageConnectionString"`
}

type RedisCacheRedisConfigurationOutputReference interface {
	cdktf.ComplexObject
	AofBackupEnabled() interface{}
	SetAofBackupEnabled(val interface{})
	AofBackupEnabledInput() interface{}
	AofStorageConnectionString0() *string
	SetAofStorageConnectionString0(val *string)
	AofStorageConnectionString0Input() *string
	AofStorageConnectionString1() *string
	SetAofStorageConnectionString1(val *string)
	AofStorageConnectionString1Input() *string
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
	EnableAuthentication() interface{}
	SetEnableAuthentication(val interface{})
	EnableAuthenticationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RedisCacheRedisConfiguration
	SetInternalValue(val *RedisCacheRedisConfiguration)
	Maxclients() *float64
	MaxfragmentationmemoryReserved() *float64
	SetMaxfragmentationmemoryReserved(val *float64)
	MaxfragmentationmemoryReservedInput() *float64
	MaxmemoryDelta() *float64
	SetMaxmemoryDelta(val *float64)
	MaxmemoryDeltaInput() *float64
	MaxmemoryPolicy() *string
	SetMaxmemoryPolicy(val *string)
	MaxmemoryPolicyInput() *string
	MaxmemoryReserved() *float64
	SetMaxmemoryReserved(val *float64)
	MaxmemoryReservedInput() *float64
	NotifyKeyspaceEvents() *string
	SetNotifyKeyspaceEvents(val *string)
	NotifyKeyspaceEventsInput() *string
	RdbBackupEnabled() interface{}
	SetRdbBackupEnabled(val interface{})
	RdbBackupEnabledInput() interface{}
	RdbBackupFrequency() *float64
	SetRdbBackupFrequency(val *float64)
	RdbBackupFrequencyInput() *float64
	RdbBackupMaxSnapshotCount() *float64
	SetRdbBackupMaxSnapshotCount(val *float64)
	RdbBackupMaxSnapshotCountInput() *float64
	RdbStorageConnectionString() *string
	SetRdbStorageConnectionString(val *string)
	RdbStorageConnectionStringInput() *string
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
	ResetAofBackupEnabled()
	ResetAofStorageConnectionString0()
	ResetAofStorageConnectionString1()
	ResetEnableAuthentication()
	ResetMaxfragmentationmemoryReserved()
	ResetMaxmemoryDelta()
	ResetMaxmemoryPolicy()
	ResetMaxmemoryReserved()
	ResetNotifyKeyspaceEvents()
	ResetRdbBackupEnabled()
	ResetRdbBackupFrequency()
	ResetRdbBackupMaxSnapshotCount()
	ResetRdbStorageConnectionString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisCacheRedisConfigurationOutputReference
type jsiiProxy_RedisCacheRedisConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aofBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aofBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString0() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString0",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString0Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString0Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) AofStorageConnectionString1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aofStorageConnectionString1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) EnableAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) EnableAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InternalValue() *RedisCacheRedisConfiguration {
	var returns *RedisCacheRedisConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Maxclients() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxclients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxfragmentationmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxfragmentationmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxfragmentationmemoryReservedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxfragmentationmemoryReservedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryDelta() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryDelta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryDeltaInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryDeltaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxmemoryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryReserved() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryReserved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) MaxmemoryReservedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxmemoryReservedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) NotifyKeyspaceEvents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) NotifyKeyspaceEventsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notifyKeyspaceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdbBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rdbBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupMaxSnapshotCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupMaxSnapshotCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbBackupMaxSnapshotCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rdbBackupMaxSnapshotCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbStorageConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbStorageConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) RdbStorageConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbStorageConnectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedisCacheRedisConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisCacheRedisConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RedisCacheRedisConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisCacheRedisConfigurationOutputReference_Override(r RedisCacheRedisConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheRedisConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetAofBackupEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"aofBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetAofStorageConnectionString0(val *string) {
	_jsii_.Set(
		j,
		"aofStorageConnectionString0",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetAofStorageConnectionString1(val *string) {
	_jsii_.Set(
		j,
		"aofStorageConnectionString1",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetEnableAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"enableAuthentication",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetInternalValue(val *RedisCacheRedisConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetMaxfragmentationmemoryReserved(val *float64) {
	_jsii_.Set(
		j,
		"maxfragmentationmemoryReserved",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetMaxmemoryDelta(val *float64) {
	_jsii_.Set(
		j,
		"maxmemoryDelta",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetMaxmemoryPolicy(val *string) {
	_jsii_.Set(
		j,
		"maxmemoryPolicy",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetMaxmemoryReserved(val *float64) {
	_jsii_.Set(
		j,
		"maxmemoryReserved",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetNotifyKeyspaceEvents(val *string) {
	_jsii_.Set(
		j,
		"notifyKeyspaceEvents",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetRdbBackupEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rdbBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetRdbBackupFrequency(val *float64) {
	_jsii_.Set(
		j,
		"rdbBackupFrequency",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetRdbBackupMaxSnapshotCount(val *float64) {
	_jsii_.Set(
		j,
		"rdbBackupMaxSnapshotCount",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetRdbStorageConnectionString(val *string) {
	_jsii_.Set(
		j,
		"rdbStorageConnectionString",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCacheRedisConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofBackupEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAofBackupEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofStorageConnectionString0() {
	_jsii_.InvokeVoid(
		r,
		"resetAofStorageConnectionString0",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetAofStorageConnectionString1() {
	_jsii_.InvokeVoid(
		r,
		"resetAofStorageConnectionString1",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetEnableAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxfragmentationmemoryReserved() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxfragmentationmemoryReserved",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryDelta() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryDelta",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetMaxmemoryReserved() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxmemoryReserved",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetNotifyKeyspaceEvents() {
	_jsii_.InvokeVoid(
		r,
		"resetNotifyKeyspaceEvents",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupFrequency() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupFrequency",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbBackupMaxSnapshotCount() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbBackupMaxSnapshotCount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ResetRdbStorageConnectionString() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbStorageConnectionString",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheRedisConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RedisCacheTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#create RedisCache#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#delete RedisCache#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#read RedisCache#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_cache#update RedisCache#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type RedisCacheTimeoutsOutputReference interface {
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

// The jsii proxy struct for RedisCacheTimeoutsOutputReference
type jsiiProxy_RedisCacheTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewRedisCacheTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisCacheTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RedisCacheTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisCacheTimeoutsOutputReference_Override(r RedisCacheTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisCache.RedisCacheTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedisCacheTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		r,
		"resetCreate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		r,
		"resetDelete",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		r,
		"resetRead",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCacheTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

