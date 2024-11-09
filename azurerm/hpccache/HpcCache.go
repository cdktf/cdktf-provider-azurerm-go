// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/hpccache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/hpc_cache azurerm_hpc_cache}.
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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

func (j *jsiiProxy_HpcCache) Count() interface{} {
	var returns interface{}
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/hpc_cache azurerm_hpc_cache} Resource.
func NewHpcCache(scope constructs.Construct, id *string, config *HpcCacheConfig) HpcCache {
	_init_.Initialize()

	if err := validateNewHpcCacheParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCache{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/hpc_cache azurerm_hpc_cache} Resource.
func NewHpcCache_Override(h HpcCache, scope constructs.Construct, id *string, config *HpcCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HpcCache)SetAutomaticallyRotateKeyToLatestEnabled(val interface{}) {
	if err := j.validateSetAutomaticallyRotateKeyToLatestEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticallyRotateKeyToLatestEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetCacheSizeInGb(val *float64) {
	if err := j.validateSetCacheSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheSizeInGb",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetKeyVaultKeyId(val *string) {
	if err := j.validateSetKeyVaultKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetMtu(val *float64) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetNtpServer(val *string) {
	if err := j.validateSetNtpServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ntpServer",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HpcCache)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a HpcCache resource upon running "cdktf plan <stack-name>".
func HpcCache_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHpcCache_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
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
func HpcCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCache_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HpcCache_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCache_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HpcCache_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCache_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hpcCache.HpcCache",
		"isTerraformResource",
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

func (h *jsiiProxy_HpcCache) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HpcCache) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HpcCache) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
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
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HpcCache) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HpcCache) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HpcCache) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HpcCache) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HpcCache) PutDefaultAccessPolicy(value *HpcCacheDefaultAccessPolicy) {
	if err := h.validatePutDefaultAccessPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDefaultAccessPolicy",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryActiveDirectory(value *HpcCacheDirectoryActiveDirectory) {
	if err := h.validatePutDirectoryActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDirectoryActiveDirectory",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryFlatFile(value *HpcCacheDirectoryFlatFile) {
	if err := h.validatePutDirectoryFlatFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDirectoryFlatFile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDirectoryLdap(value *HpcCacheDirectoryLdap) {
	if err := h.validatePutDirectoryLdapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDirectoryLdap",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutDns(value *HpcCacheDns) {
	if err := h.validatePutDnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDns",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutIdentity(value *HpcCacheIdentity) {
	if err := h.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putIdentity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCache) PutTimeouts(value *HpcCacheTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (h *jsiiProxy_HpcCache) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCache) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
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

