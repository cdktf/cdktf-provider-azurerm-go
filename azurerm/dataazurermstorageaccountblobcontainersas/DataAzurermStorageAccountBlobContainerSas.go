// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccountblobcontainersas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/dataazurermstorageaccountblobcontainersas/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas}.
type DataAzurermStorageAccountBlobContainerSas interface {
	cdktf.TerraformDataSource
	CacheControl() *string
	SetCacheControl(val *string)
	CacheControlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConnectionString() *string
	SetConnectionString(val *string)
	ConnectionStringInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	ContentDisposition() *string
	SetContentDisposition(val *string)
	ContentDispositionInput() *string
	ContentEncoding() *string
	SetContentEncoding(val *string)
	ContentEncodingInput() *string
	ContentLanguage() *string
	SetContentLanguage(val *string)
	ContentLanguageInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Expiry() *string
	SetExpiry(val *string)
	ExpiryInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpsOnly() interface{}
	SetHttpsOnly(val interface{})
	HttpsOnlyInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddress() *string
	SetIpAddress(val *string)
	IpAddressInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Permissions() DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference
	PermissionsInput() *DataAzurermStorageAccountBlobContainerSasPermissions
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Sas() *string
	Start() *string
	SetStart(val *string)
	StartInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference
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
	PutPermissions(value *DataAzurermStorageAccountBlobContainerSasPermissions)
	PutTimeouts(value *DataAzurermStorageAccountBlobContainerSasTimeouts)
	ResetCacheControl()
	ResetContentDisposition()
	ResetContentEncoding()
	ResetContentLanguage()
	ResetContentType()
	ResetHttpsOnly()
	ResetId()
	ResetIpAddress()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermStorageAccountBlobContainerSas
type jsiiProxy_DataAzurermStorageAccountBlobContainerSas struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) HttpsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) HttpsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) IpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Permissions() DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference {
	var returns DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) PermissionsInput() *DataAzurermStorageAccountBlobContainerSasPermissions {
	var returns *DataAzurermStorageAccountBlobContainerSasPermissions
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Sas() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Start() *string {
	var returns *string
	_jsii_.Get(
		j,
		"start",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) StartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Timeouts() DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference {
	var returns DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas} Data Source.
func NewDataAzurermStorageAccountBlobContainerSas(scope constructs.Construct, id *string, config *DataAzurermStorageAccountBlobContainerSasConfig) DataAzurermStorageAccountBlobContainerSas {
	_init_.Initialize()

	if err := validateNewDataAzurermStorageAccountBlobContainerSasParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermStorageAccountBlobContainerSas{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/data-sources/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas} Data Source.
func NewDataAzurermStorageAccountBlobContainerSas_Override(d DataAzurermStorageAccountBlobContainerSas, scope constructs.Construct, id *string, config *DataAzurermStorageAccountBlobContainerSasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetConnectionString(val *string) {
	if err := j.validateSetConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetExpiry(val *string) {
	if err := j.validateSetExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiry",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetHttpsOnly(val interface{}) {
	if err := j.validateSetHttpsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetIpAddress(val *string) {
	if err := j.validateSetIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas)SetStart(val *string) {
	if err := j.validateSetStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"start",
		val,
	)
}

// Generates CDKTF code for importing a DataAzurermStorageAccountBlobContainerSas resource upon running "cdktf plan <stack-name>".
func DataAzurermStorageAccountBlobContainerSas_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccountBlobContainerSas_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
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
func DataAzurermStorageAccountBlobContainerSas_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccountBlobContainerSas_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermStorageAccountBlobContainerSas_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccountBlobContainerSas_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermStorageAccountBlobContainerSas_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermStorageAccountBlobContainerSas_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermStorageAccountBlobContainerSas_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) PutPermissions(value *DataAzurermStorageAccountBlobContainerSasPermissions) {
	if err := d.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPermissions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) PutTimeouts(value *DataAzurermStorageAccountBlobContainerSasTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetCacheControl() {
	_jsii_.InvokeVoid(
		d,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		d,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		d,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		d,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetContentType() {
	_jsii_.InvokeVoid(
		d,
		"resetContentType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetHttpsOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpsOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetIpAddress() {
	_jsii_.InvokeVoid(
		d,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

