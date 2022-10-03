package dataazurermstorageaccountblobcontainersas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/dataazurermstorageaccountblobcontainersas/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas} Data Source.
func NewDataAzurermStorageAccountBlobContainerSas(scope constructs.Construct, id *string, config *DataAzurermStorageAccountBlobContainerSasConfig) DataAzurermStorageAccountBlobContainerSas {
	_init_.Initialize()

	j := jsiiProxy_DataAzurermStorageAccountBlobContainerSas{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas azurerm_storage_account_blob_container_sas} Data Source.
func NewDataAzurermStorageAccountBlobContainerSas_Override(d DataAzurermStorageAccountBlobContainerSas, scope constructs.Construct, id *string, config *DataAzurermStorageAccountBlobContainerSasConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetCacheControl(val *string) {
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetConnectionString(val *string) {
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetContainerName(val *string) {
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetContentDisposition(val *string) {
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetContentEncoding(val *string) {
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetContentLanguage(val *string) {
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetExpiry(val *string) {
	_jsii_.Set(
		j,
		"expiry",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetHttpsOnly(val interface{}) {
	_jsii_.Set(
		j,
		"httpsOnly",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetIpAddress(val *string) {
	_jsii_.Set(
		j,
		"ipAddress",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) SetStart(val *string) {
	_jsii_.Set(
		j,
		"start",
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
func DataAzurermStorageAccountBlobContainerSas_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSas",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) PutPermissions(value *DataAzurermStorageAccountBlobContainerSasPermissions) {
	_jsii_.InvokeVoid(
		d,
		"putPermissions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSas) PutTimeouts(value *DataAzurermStorageAccountBlobContainerSasTimeouts) {
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

type DataAzurermStorageAccountBlobContainerSasConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#connection_string DataAzurermStorageAccountBlobContainerSas#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#container_name DataAzurermStorageAccountBlobContainerSas#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#expiry DataAzurermStorageAccountBlobContainerSas#expiry}.
	Expiry *string `field:"required" json:"expiry" yaml:"expiry"`
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#permissions DataAzurermStorageAccountBlobContainerSas#permissions}
	Permissions *DataAzurermStorageAccountBlobContainerSasPermissions `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#start DataAzurermStorageAccountBlobContainerSas#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#cache_control DataAzurermStorageAccountBlobContainerSas#cache_control}.
	CacheControl *string `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#content_disposition DataAzurermStorageAccountBlobContainerSas#content_disposition}.
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#content_encoding DataAzurermStorageAccountBlobContainerSas#content_encoding}.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#content_language DataAzurermStorageAccountBlobContainerSas#content_language}.
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#content_type DataAzurermStorageAccountBlobContainerSas#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#https_only DataAzurermStorageAccountBlobContainerSas#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#id DataAzurermStorageAccountBlobContainerSas#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#ip_address DataAzurermStorageAccountBlobContainerSas#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#timeouts DataAzurermStorageAccountBlobContainerSas#timeouts}
	Timeouts *DataAzurermStorageAccountBlobContainerSasTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DataAzurermStorageAccountBlobContainerSasPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#add DataAzurermStorageAccountBlobContainerSas#add}.
	Add interface{} `field:"required" json:"add" yaml:"add"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#create DataAzurermStorageAccountBlobContainerSas#create}.
	Create interface{} `field:"required" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#delete DataAzurermStorageAccountBlobContainerSas#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#list DataAzurermStorageAccountBlobContainerSas#list}.
	List interface{} `field:"required" json:"list" yaml:"list"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#read DataAzurermStorageAccountBlobContainerSas#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#write DataAzurermStorageAccountBlobContainerSas#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
}

type DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference interface {
	cdktf.ComplexObject
	Add() interface{}
	SetAdd(val interface{})
	AddInput() interface{}
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
	Create() interface{}
	SetCreate(val interface{})
	CreateInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() interface{}
	SetDelete(val interface{})
	DeleteInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DataAzurermStorageAccountBlobContainerSasPermissions
	SetInternalValue(val *DataAzurermStorageAccountBlobContainerSasPermissions)
	List() interface{}
	SetList(val interface{})
	ListInput() interface{}
	Read() interface{}
	SetRead(val interface{})
	ReadInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Write() interface{}
	SetWrite(val interface{})
	WriteInput() interface{}
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

// The jsii proxy struct for DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference
type jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Add() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"add",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) AddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Create() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) CreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Delete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) DeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) InternalValue() *DataAzurermStorageAccountBlobContainerSasPermissions {
	var returns *DataAzurermStorageAccountBlobContainerSasPermissions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) List() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"list",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Read() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Write() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"write",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) WriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeInput",
		&returns,
	)
	return returns
}


func NewDataAzurermStorageAccountBlobContainerSasPermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAzurermStorageAccountBlobContainerSasPermissionsOutputReference_Override(d DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetAdd(val interface{}) {
	_jsii_.Set(
		j,
		"add",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetCreate(val interface{}) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetDelete(val interface{}) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetInternalValue(val *DataAzurermStorageAccountBlobContainerSasPermissions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetList(val interface{}) {
	_jsii_.Set(
		j,
		"list",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetRead(val interface{}) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) SetWrite(val interface{}) {
	_jsii_.Set(
		j,
		"write",
		val,
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasPermissionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAzurermStorageAccountBlobContainerSasTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/storage_account_blob_container_sas#read DataAzurermStorageAccountBlobContainerSas#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

type DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference interface {
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
	ResetRead()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference
type jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference_Override(d DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageAccountBlobContainerSas.DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageAccountBlobContainerSasTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

