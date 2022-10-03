package hpccache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/hpccache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache azurerm_hpc_cache}.
type HpcCache interface {
	cdktf.TerraformResource
	AutomaticallyRotateKeyToLatestEnabled() interface{}
	SetAutomaticallyRotateKeyToLatestEnabled(val interface{})
	AutomaticallyRotateKeyToLatestEnabledInput() interface{}
	CacheSizeInGb() *float64
	SetCacheSizeInGb(val *float64)
	CacheSizeInGbInput() *float64
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
	DefaultAccessPolicy() HpcCacheDefaultAccessPolicyOutputReference
	DefaultAccessPolicyInput() *HpcCacheDefaultAccessPolicy
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectoryActiveDirectory() HpcCacheDirectoryActiveDirectoryOutputReference
	DirectoryActiveDirectoryInput() *HpcCacheDirectoryActiveDirectory
	DirectoryFlatFile() HpcCacheDirectoryFlatFileOutputReference
	DirectoryFlatFileInput() *HpcCacheDirectoryFlatFile
	DirectoryLdap() HpcCacheDirectoryLdapOutputReference
	DirectoryLdapInput() *HpcCacheDirectoryLdap
	Dns() HpcCacheDnsOutputReference
	DnsInput() *HpcCacheDns
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
	Identity() HpcCacheIdentityOutputReference
	IdentityInput() *HpcCacheIdentity
	IdInput() *string
	KeyVaultKeyId() *string
	SetKeyVaultKeyId(val *string)
	KeyVaultKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MountAddresses() *[]*string
	Mtu() *float64
	SetMtu(val *float64)
	MtuInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NtpServer() *string
	SetNtpServer(val *string)
	NtpServerInput() *string
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
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HpcCacheTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutDefaultAccessPolicy(value *HpcCacheDefaultAccessPolicy)
	PutDirectoryActiveDirectory(value *HpcCacheDirectoryActiveDirectory)
	PutDirectoryFlatFile(value *HpcCacheDirectoryFlatFile)
	PutDirectoryLdap(value *HpcCacheDirectoryLdap)
	PutDns(value *HpcCacheDns)
	PutIdentity(value *HpcCacheIdentity)
	PutTimeouts(value *HpcCacheTimeouts)
	ResetAutomaticallyRotateKeyToLatestEnabled()
	ResetDefaultAccessPolicy()
	ResetDirectoryActiveDirectory()
	ResetDirectoryFlatFile()
	ResetDirectoryLdap()
	ResetDns()
	ResetId()
	ResetIdentity()
	ResetKeyVaultKeyId()
	ResetMtu()
	ResetNtpServer()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for HpcCache
type jsiiProxy_HpcCache struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HpcCache) AutomaticallyRotateKeyToLatestEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticallyRotateKeyToLatestEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) AutomaticallyRotateKeyToLatestEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticallyRotateKeyToLatestEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) CacheSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) CacheSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheSizeInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DefaultAccessPolicy() HpcCacheDefaultAccessPolicyOutputReference {
	var returns HpcCacheDefaultAccessPolicyOutputReference
	_jsii_.Get(
		j,
		"defaultAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DefaultAccessPolicyInput() *HpcCacheDefaultAccessPolicy {
	var returns *HpcCacheDefaultAccessPolicy
	_jsii_.Get(
		j,
		"defaultAccessPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryActiveDirectory() HpcCacheDirectoryActiveDirectoryOutputReference {
	var returns HpcCacheDirectoryActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"directoryActiveDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryActiveDirectoryInput() *HpcCacheDirectoryActiveDirectory {
	var returns *HpcCacheDirectoryActiveDirectory
	_jsii_.Get(
		j,
		"directoryActiveDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryFlatFile() HpcCacheDirectoryFlatFileOutputReference {
	var returns HpcCacheDirectoryFlatFileOutputReference
	_jsii_.Get(
		j,
		"directoryFlatFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryFlatFileInput() *HpcCacheDirectoryFlatFile {
	var returns *HpcCacheDirectoryFlatFile
	_jsii_.Get(
		j,
		"directoryFlatFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryLdap() HpcCacheDirectoryLdapOutputReference {
	var returns HpcCacheDirectoryLdapOutputReference
	_jsii_.Get(
		j,
		"directoryLdap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DirectoryLdapInput() *HpcCacheDirectoryLdap {
	var returns *HpcCacheDirectoryLdap
	_jsii_.Get(
		j,
		"directoryLdapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Dns() HpcCacheDnsOutputReference {
	var returns HpcCacheDnsOutputReference
	_jsii_.Get(
		j,
		"dns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) DnsInput() *HpcCacheDns {
	var returns *HpcCacheDns
	_jsii_.Get(
		j,
		"dnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Identity() HpcCacheIdentityOutputReference {
	var returns HpcCacheIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) IdentityInput() *HpcCacheIdentity {
	var returns *HpcCacheIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) KeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) KeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) MountAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mountAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Mtu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) MtuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) NtpServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ntpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) NtpServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ntpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) Timeouts() HpcCacheTimeoutsOutputReference {
	var returns HpcCacheTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCache) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache azurerm_hpc_cache} Resource.
func NewHpcCache(scope constructs.Construct, id *string, config *HpcCacheConfig) HpcCache {
	_init_.Initialize()

	j := jsiiProxy_HpcCache{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache azurerm_hpc_cache} Resource.
func NewHpcCache_Override(h HpcCache, scope constructs.Construct, id *string, config *HpcCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HpcCache) SetAutomaticallyRotateKeyToLatestEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"automaticallyRotateKeyToLatestEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetCacheSizeInGb(val *float64) {
	_jsii_.Set(
		j,
		"cacheSizeInGb",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetMtu(val *float64) {
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetNtpServer(val *string) {
	_jsii_.Set(
		j,
		"ntpServer",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetSkuName(val *string) {
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HpcCache) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
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
func HpcCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HpcCache_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HpcCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HpcCache) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HpcCache) PutDefaultAccessPolicy(value *HpcCacheDefaultAccessPolicy) {
	_jsii_.InvokeVoid(
		h,
		"putDefaultAccessPolicy",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryActiveDirectory(value *HpcCacheDirectoryActiveDirectory) {
	_jsii_.InvokeVoid(
		h,
		"putDirectoryActiveDirectory",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryFlatFile(value *HpcCacheDirectoryFlatFile) {
	_jsii_.InvokeVoid(
		h,
		"putDirectoryFlatFile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryLdap(value *HpcCacheDirectoryLdap) {
	_jsii_.InvokeVoid(
		h,
		"putDirectoryLdap",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDns(value *HpcCacheDns) {
	_jsii_.InvokeVoid(
		h,
		"putDns",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutIdentity(value *HpcCacheIdentity) {
	_jsii_.InvokeVoid(
		h,
		"putIdentity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutTimeouts(value *HpcCacheTimeouts) {
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) ResetAutomaticallyRotateKeyToLatestEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetAutomaticallyRotateKeyToLatestEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetDefaultAccessPolicy() {
	_jsii_.InvokeVoid(
		h,
		"resetDefaultAccessPolicy",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetDirectoryActiveDirectory() {
	_jsii_.InvokeVoid(
		h,
		"resetDirectoryActiveDirectory",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetDirectoryFlatFile() {
	_jsii_.InvokeVoid(
		h,
		"resetDirectoryFlatFile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetDirectoryLdap() {
	_jsii_.InvokeVoid(
		h,
		"resetDirectoryLdap",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetDns() {
	_jsii_.InvokeVoid(
		h,
		"resetDns",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetIdentity() {
	_jsii_.InvokeVoid(
		h,
		"resetIdentity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		h,
		"resetKeyVaultKeyId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetMtu() {
	_jsii_.InvokeVoid(
		h,
		"resetMtu",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetNtpServer() {
	_jsii_.InvokeVoid(
		h,
		"resetNtpServer",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCache) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#cache_size_in_gb HpcCache#cache_size_in_gb}.
	CacheSizeInGb *float64 `field:"required" json:"cacheSizeInGb" yaml:"cacheSizeInGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#location HpcCache#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#name HpcCache#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#resource_group_name HpcCache#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#sku_name HpcCache#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#subnet_id HpcCache#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#automatically_rotate_key_to_latest_enabled HpcCache#automatically_rotate_key_to_latest_enabled}.
	AutomaticallyRotateKeyToLatestEnabled interface{} `field:"optional" json:"automaticallyRotateKeyToLatestEnabled" yaml:"automaticallyRotateKeyToLatestEnabled"`
	// default_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#default_access_policy HpcCache#default_access_policy}
	DefaultAccessPolicy *HpcCacheDefaultAccessPolicy `field:"optional" json:"defaultAccessPolicy" yaml:"defaultAccessPolicy"`
	// directory_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#directory_active_directory HpcCache#directory_active_directory}
	DirectoryActiveDirectory *HpcCacheDirectoryActiveDirectory `field:"optional" json:"directoryActiveDirectory" yaml:"directoryActiveDirectory"`
	// directory_flat_file block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#directory_flat_file HpcCache#directory_flat_file}
	DirectoryFlatFile *HpcCacheDirectoryFlatFile `field:"optional" json:"directoryFlatFile" yaml:"directoryFlatFile"`
	// directory_ldap block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#directory_ldap HpcCache#directory_ldap}
	DirectoryLdap *HpcCacheDirectoryLdap `field:"optional" json:"directoryLdap" yaml:"directoryLdap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#dns HpcCache#dns}
	Dns *HpcCacheDns `field:"optional" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#id HpcCache#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#identity HpcCache#identity}
	Identity *HpcCacheIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#key_vault_key_id HpcCache#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#mtu HpcCache#mtu}.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#ntp_server HpcCache#ntp_server}.
	NtpServer *string `field:"optional" json:"ntpServer" yaml:"ntpServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#tags HpcCache#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#timeouts HpcCache#timeouts}
	Timeouts *HpcCacheTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type HpcCacheDefaultAccessPolicy struct {
	// access_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#access_rule HpcCache#access_rule}
	AccessRule interface{} `field:"required" json:"accessRule" yaml:"accessRule"`
}

type HpcCacheDefaultAccessPolicyAccessRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#access HpcCache#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#scope HpcCache#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#anonymous_gid HpcCache#anonymous_gid}.
	AnonymousGid *float64 `field:"optional" json:"anonymousGid" yaml:"anonymousGid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#anonymous_uid HpcCache#anonymous_uid}.
	AnonymousUid *float64 `field:"optional" json:"anonymousUid" yaml:"anonymousUid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#filter HpcCache#filter}.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#root_squash_enabled HpcCache#root_squash_enabled}.
	RootSquashEnabled interface{} `field:"optional" json:"rootSquashEnabled" yaml:"rootSquashEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#submount_access_enabled HpcCache#submount_access_enabled}.
	SubmountAccessEnabled interface{} `field:"optional" json:"submountAccessEnabled" yaml:"submountAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#suid_enabled HpcCache#suid_enabled}.
	SuidEnabled interface{} `field:"optional" json:"suidEnabled" yaml:"suidEnabled"`
}

type HpcCacheDefaultAccessPolicyAccessRuleList interface {
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
	Get(index *float64) HpcCacheDefaultAccessPolicyAccessRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDefaultAccessPolicyAccessRuleList
type jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHpcCacheDefaultAccessPolicyAccessRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HpcCacheDefaultAccessPolicyAccessRuleList {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHpcCacheDefaultAccessPolicyAccessRuleList_Override(h HpcCacheDefaultAccessPolicyAccessRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) Get(index *float64) HpcCacheDefaultAccessPolicyAccessRuleOutputReference {
	var returns HpcCacheDefaultAccessPolicyAccessRuleOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDefaultAccessPolicyAccessRuleOutputReference interface {
	cdktf.ComplexObject
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
	AnonymousGid() *float64
	SetAnonymousGid(val *float64)
	AnonymousGidInput() *float64
	AnonymousUid() *float64
	SetAnonymousUid(val *float64)
	AnonymousUidInput() *float64
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
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RootSquashEnabled() interface{}
	SetRootSquashEnabled(val interface{})
	RootSquashEnabledInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SubmountAccessEnabled() interface{}
	SetSubmountAccessEnabled(val interface{})
	SubmountAccessEnabledInput() interface{}
	SuidEnabled() interface{}
	SetSuidEnabled(val interface{})
	SuidEnabledInput() interface{}
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
	ResetAnonymousGid()
	ResetAnonymousUid()
	ResetFilter()
	ResetRootSquashEnabled()
	ResetSubmountAccessEnabled()
	ResetSuidEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDefaultAccessPolicyAccessRuleOutputReference
type jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) RootSquashEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootSquashEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) RootSquashEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootSquashEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SubmountAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submountAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SubmountAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submountAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SuidEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suidEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SuidEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suidEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDefaultAccessPolicyAccessRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HpcCacheDefaultAccessPolicyAccessRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHpcCacheDefaultAccessPolicyAccessRuleOutputReference_Override(h HpcCacheDefaultAccessPolicyAccessRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetAccess(val *string) {
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetAnonymousGid(val *float64) {
	_jsii_.Set(
		j,
		"anonymousGid",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetAnonymousUid(val *float64) {
	_jsii_.Set(
		j,
		"anonymousUid",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetFilter(val *string) {
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetRootSquashEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rootSquashEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetSubmountAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"submountAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetSuidEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"suidEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetAnonymousGid() {
	_jsii_.InvokeVoid(
		h,
		"resetAnonymousGid",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetAnonymousUid() {
	_jsii_.InvokeVoid(
		h,
		"resetAnonymousUid",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		h,
		"resetFilter",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetRootSquashEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetRootSquashEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetSubmountAccessEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetSubmountAccessEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetSuidEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetSuidEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDefaultAccessPolicyOutputReference interface {
	cdktf.ComplexObject
	AccessRule() HpcCacheDefaultAccessPolicyAccessRuleList
	AccessRuleInput() interface{}
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
	InternalValue() *HpcCacheDefaultAccessPolicy
	SetInternalValue(val *HpcCacheDefaultAccessPolicy)
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
	PutAccessRule(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDefaultAccessPolicyOutputReference
type jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) AccessRule() HpcCacheDefaultAccessPolicyAccessRuleList {
	var returns HpcCacheDefaultAccessPolicyAccessRuleList
	_jsii_.Get(
		j,
		"accessRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) AccessRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) InternalValue() *HpcCacheDefaultAccessPolicy {
	var returns *HpcCacheDefaultAccessPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDefaultAccessPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDefaultAccessPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDefaultAccessPolicyOutputReference_Override(h HpcCacheDefaultAccessPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) SetInternalValue(val *HpcCacheDefaultAccessPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) PutAccessRule(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putAccessRule",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDirectoryActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#cache_netbios_name HpcCache#cache_netbios_name}.
	CacheNetbiosName *string `field:"required" json:"cacheNetbiosName" yaml:"cacheNetbiosName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#dns_primary_ip HpcCache#dns_primary_ip}.
	DnsPrimaryIp *string `field:"required" json:"dnsPrimaryIp" yaml:"dnsPrimaryIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#domain_name HpcCache#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#domain_netbios_name HpcCache#domain_netbios_name}.
	DomainNetbiosName *string `field:"required" json:"domainNetbiosName" yaml:"domainNetbiosName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#password HpcCache#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#username HpcCache#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#dns_secondary_ip HpcCache#dns_secondary_ip}.
	DnsSecondaryIp *string `field:"optional" json:"dnsSecondaryIp" yaml:"dnsSecondaryIp"`
}

type HpcCacheDirectoryActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	CacheNetbiosName() *string
	SetCacheNetbiosName(val *string)
	CacheNetbiosNameInput() *string
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
	DnsPrimaryIp() *string
	SetDnsPrimaryIp(val *string)
	DnsPrimaryIpInput() *string
	DnsSecondaryIp() *string
	SetDnsSecondaryIp(val *string)
	DnsSecondaryIpInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainNetbiosName() *string
	SetDomainNetbiosName(val *string)
	DomainNetbiosNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HpcCacheDirectoryActiveDirectory
	SetInternalValue(val *HpcCacheDirectoryActiveDirectory)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetDnsSecondaryIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDirectoryActiveDirectoryOutputReference
type jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CacheNetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNetbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CacheNetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNetbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsPrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrimaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsPrimaryIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrimaryIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsSecondaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSecondaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsSecondaryIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSecondaryIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNetbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNetbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InternalValue() *HpcCacheDirectoryActiveDirectory {
	var returns *HpcCacheDirectoryActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryActiveDirectoryOutputReference_Override(h HpcCacheDirectoryActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetCacheNetbiosName(val *string) {
	_jsii_.Set(
		j,
		"cacheNetbiosName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetDnsPrimaryIp(val *string) {
	_jsii_.Set(
		j,
		"dnsPrimaryIp",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetDnsSecondaryIp(val *string) {
	_jsii_.Set(
		j,
		"dnsSecondaryIp",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetDomainNetbiosName(val *string) {
	_jsii_.Set(
		j,
		"domainNetbiosName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetInternalValue(val *HpcCacheDirectoryActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ResetDnsSecondaryIp() {
	_jsii_.InvokeVoid(
		h,
		"resetDnsSecondaryIp",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDirectoryFlatFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#group_file_uri HpcCache#group_file_uri}.
	GroupFileUri *string `field:"required" json:"groupFileUri" yaml:"groupFileUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#password_file_uri HpcCache#password_file_uri}.
	PasswordFileUri *string `field:"required" json:"passwordFileUri" yaml:"passwordFileUri"`
}

type HpcCacheDirectoryFlatFileOutputReference interface {
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
	GroupFileUri() *string
	SetGroupFileUri(val *string)
	GroupFileUriInput() *string
	InternalValue() *HpcCacheDirectoryFlatFile
	SetInternalValue(val *HpcCacheDirectoryFlatFile)
	PasswordFileUri() *string
	SetPasswordFileUri(val *string)
	PasswordFileUriInput() *string
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

// The jsii proxy struct for HpcCacheDirectoryFlatFileOutputReference
type jsiiProxy_HpcCacheDirectoryFlatFileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GroupFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GroupFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) InternalValue() *HpcCacheDirectoryFlatFile {
	var returns *HpcCacheDirectoryFlatFile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) PasswordFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) PasswordFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryFlatFileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryFlatFileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDirectoryFlatFileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryFlatFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryFlatFileOutputReference_Override(h HpcCacheDirectoryFlatFileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryFlatFileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetGroupFileUri(val *string) {
	_jsii_.Set(
		j,
		"groupFileUri",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetInternalValue(val *HpcCacheDirectoryFlatFile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetPasswordFileUri(val *string) {
	_jsii_.Set(
		j,
		"passwordFileUri",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryFlatFileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDirectoryLdap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#base_dn HpcCache#base_dn}.
	BaseDn *string `field:"required" json:"baseDn" yaml:"baseDn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#server HpcCache#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// bind block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#bind HpcCache#bind}
	Bind *HpcCacheDirectoryLdapBind `field:"optional" json:"bind" yaml:"bind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#certificate_validation_uri HpcCache#certificate_validation_uri}.
	CertificateValidationUri *string `field:"optional" json:"certificateValidationUri" yaml:"certificateValidationUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#download_certificate_automatically HpcCache#download_certificate_automatically}.
	DownloadCertificateAutomatically interface{} `field:"optional" json:"downloadCertificateAutomatically" yaml:"downloadCertificateAutomatically"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#encrypted HpcCache#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
}

type HpcCacheDirectoryLdapBind struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#dn HpcCache#dn}.
	Dn *string `field:"required" json:"dn" yaml:"dn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#password HpcCache#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
}

type HpcCacheDirectoryLdapBindOutputReference interface {
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
	Dn() *string
	SetDn(val *string)
	DnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HpcCacheDirectoryLdapBind
	SetInternalValue(val *HpcCacheDirectoryLdapBind)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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

// The jsii proxy struct for HpcCacheDirectoryLdapBindOutputReference
type jsiiProxy_HpcCacheDirectoryLdapBindOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) Dn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) DnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) InternalValue() *HpcCacheDirectoryLdapBind {
	var returns *HpcCacheDirectoryLdapBind
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryLdapBindOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryLdapBindOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDirectoryLdapBindOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapBindOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryLdapBindOutputReference_Override(h HpcCacheDirectoryLdapBindOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapBindOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetDn(val *string) {
	_jsii_.Set(
		j,
		"dn",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetInternalValue(val *HpcCacheDirectoryLdapBind) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapBindOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDirectoryLdapOutputReference interface {
	cdktf.ComplexObject
	BaseDn() *string
	SetBaseDn(val *string)
	BaseDnInput() *string
	Bind() HpcCacheDirectoryLdapBindOutputReference
	BindInput() *HpcCacheDirectoryLdapBind
	CertificateValidationUri() *string
	SetCertificateValidationUri(val *string)
	CertificateValidationUriInput() *string
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
	DownloadCertificateAutomatically() interface{}
	SetDownloadCertificateAutomatically(val interface{})
	DownloadCertificateAutomaticallyInput() interface{}
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *HpcCacheDirectoryLdap
	SetInternalValue(val *HpcCacheDirectoryLdap)
	Server() *string
	SetServer(val *string)
	ServerInput() *string
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
	PutBind(value *HpcCacheDirectoryLdapBind)
	ResetBind()
	ResetCertificateValidationUri()
	ResetDownloadCertificateAutomatically()
	ResetEncrypted()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDirectoryLdapOutputReference
type jsiiProxy_HpcCacheDirectoryLdapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BaseDn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BaseDnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Bind() HpcCacheDirectoryLdapBindOutputReference {
	var returns HpcCacheDirectoryLdapBindOutputReference
	_jsii_.Get(
		j,
		"bind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BindInput() *HpcCacheDirectoryLdapBind {
	var returns *HpcCacheDirectoryLdapBind
	_jsii_.Get(
		j,
		"bindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CertificateValidationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateValidationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CertificateValidationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateValidationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) DownloadCertificateAutomatically() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadCertificateAutomatically",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) DownloadCertificateAutomaticallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadCertificateAutomaticallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InternalValue() *HpcCacheDirectoryLdap {
	var returns *HpcCacheDirectoryLdap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryLdapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryLdapOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDirectoryLdapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryLdapOutputReference_Override(h HpcCacheDirectoryLdapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetBaseDn(val *string) {
	_jsii_.Set(
		j,
		"baseDn",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetCertificateValidationUri(val *string) {
	_jsii_.Set(
		j,
		"certificateValidationUri",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetDownloadCertificateAutomatically(val interface{}) {
	_jsii_.Set(
		j,
		"downloadCertificateAutomatically",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetInternalValue(val *HpcCacheDirectoryLdap) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) PutBind(value *HpcCacheDirectoryLdapBind) {
	_jsii_.InvokeVoid(
		h,
		"putBind",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetBind() {
	_jsii_.InvokeVoid(
		h,
		"resetBind",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetCertificateValidationUri() {
	_jsii_.InvokeVoid(
		h,
		"resetCertificateValidationUri",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetDownloadCertificateAutomatically() {
	_jsii_.InvokeVoid(
		h,
		"resetDownloadCertificateAutomatically",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		h,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheDns struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#servers HpcCache#servers}.
	Servers *[]*string `field:"required" json:"servers" yaml:"servers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#search_domain HpcCache#search_domain}.
	SearchDomain *string `field:"optional" json:"searchDomain" yaml:"searchDomain"`
}

type HpcCacheDnsOutputReference interface {
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
	InternalValue() *HpcCacheDns
	SetInternalValue(val *HpcCacheDns)
	SearchDomain() *string
	SetSearchDomain(val *string)
	SearchDomainInput() *string
	Servers() *[]*string
	SetServers(val *[]*string)
	ServersInput() *[]*string
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
	ResetSearchDomain()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDnsOutputReference
type jsiiProxy_HpcCacheDnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) InternalValue() *HpcCacheDns {
	var returns *HpcCacheDns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SearchDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SearchDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) Servers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) ServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDnsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheDnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDnsOutputReference_Override(h HpcCacheDnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetInternalValue(val *HpcCacheDns) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetSearchDomain(val *string) {
	_jsii_.Set(
		j,
		"searchDomain",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetServers(val *[]*string) {
	_jsii_.Set(
		j,
		"servers",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDnsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) ResetSearchDomain() {
	_jsii_.InvokeVoid(
		h,
		"resetSearchDomain",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#identity_ids HpcCache#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#type HpcCache#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

type HpcCacheIdentityOutputReference interface {
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
	InternalValue() *HpcCacheIdentity
	SetInternalValue(val *HpcCacheIdentity)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheIdentityOutputReference
type jsiiProxy_HpcCacheIdentityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) IdentityIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) IdentityIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identityIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) InternalValue() *HpcCacheIdentity {
	var returns *HpcCacheIdentity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewHpcCacheIdentityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheIdentityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheIdentityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheIdentityOutputReference_Override(h HpcCacheIdentityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheIdentityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetIdentityIds(val *[]*string) {
	_jsii_.Set(
		j,
		"identityIds",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetInternalValue(val *HpcCacheIdentity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HpcCacheIdentityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheIdentityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HpcCacheTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#create HpcCache#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#delete HpcCache#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#read HpcCache#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hpc_cache#update HpcCache#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type HpcCacheTimeoutsOutputReference interface {
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

// The jsii proxy struct for HpcCacheTimeoutsOutputReference
type jsiiProxy_HpcCacheTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewHpcCacheTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HpcCacheTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheTimeoutsOutputReference_Override(h HpcCacheTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HpcCacheTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		h,
		"resetDelete",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		h,
		"resetRead",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		h,
		"resetUpdate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

