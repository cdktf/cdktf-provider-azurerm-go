package healthcareservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/healthcareservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service azurerm_healthcare_service}.
type HealthcareService interface {
	cdktf.TerraformResource
	AccessPolicyObjectIds() *[]*string
	SetAccessPolicyObjectIds(val *[]*string)
	AccessPolicyObjectIdsInput() *[]*string
	AuthenticationConfiguration() HealthcareServiceAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() *HealthcareServiceAuthenticationConfiguration
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfiguration() HealthcareServiceCorsConfigurationOutputReference
	CorsConfigurationInput() *HealthcareServiceCorsConfiguration
	CosmosdbKeyVaultKeyVersionlessId() *string
	SetCosmosdbKeyVaultKeyVersionlessId(val *string)
	CosmosdbKeyVaultKeyVersionlessIdInput() *string
	CosmosdbThroughput() *float64
	SetCosmosdbThroughput(val *float64)
	CosmosdbThroughputInput() *float64
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	IdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HealthcareServiceTimeoutsOutputReference
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
	PutAuthenticationConfiguration(value *HealthcareServiceAuthenticationConfiguration)
	PutCorsConfiguration(value *HealthcareServiceCorsConfiguration)
	PutTimeouts(value *HealthcareServiceTimeouts)
	ResetAccessPolicyObjectIds()
	ResetAuthenticationConfiguration()
	ResetCorsConfiguration()
	ResetCosmosdbKeyVaultKeyVersionlessId()
	ResetCosmosdbThroughput()
	ResetId()
	ResetKind()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
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

// The jsii proxy struct for HealthcareService
type jsiiProxy_HealthcareService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HealthcareService) AccessPolicyObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessPolicyObjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) AccessPolicyObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessPolicyObjectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) AuthenticationConfiguration() HealthcareServiceAuthenticationConfigurationOutputReference {
	var returns HealthcareServiceAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) AuthenticationConfigurationInput() *HealthcareServiceAuthenticationConfiguration {
	var returns *HealthcareServiceAuthenticationConfiguration
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CorsConfiguration() HealthcareServiceCorsConfigurationOutputReference {
	var returns HealthcareServiceCorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CorsConfigurationInput() *HealthcareServiceCorsConfiguration {
	var returns *HealthcareServiceCorsConfiguration
	_jsii_.Get(
		j,
		"corsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CosmosdbKeyVaultKeyVersionlessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbKeyVaultKeyVersionlessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CosmosdbKeyVaultKeyVersionlessIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbKeyVaultKeyVersionlessIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CosmosdbThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cosmosdbThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) CosmosdbThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cosmosdbThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) Timeouts() HealthcareServiceTimeoutsOutputReference {
	var returns HealthcareServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service azurerm_healthcare_service} Resource.
func NewHealthcareService(scope constructs.Construct, id *string, config *HealthcareServiceConfig) HealthcareService {
	_init_.Initialize()

	j := jsiiProxy_HealthcareService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service azurerm_healthcare_service} Resource.
func NewHealthcareService_Override(h HealthcareService, scope constructs.Construct, id *string, config *HealthcareServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareService) SetAccessPolicyObjectIds(val *[]*string) {
	_jsii_.Set(
		j,
		"accessPolicyObjectIds",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetCosmosdbKeyVaultKeyVersionlessId(val *string) {
	_jsii_.Set(
		j,
		"cosmosdbKeyVaultKeyVersionlessId",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetCosmosdbThroughput(val *float64) {
	_jsii_.Set(
		j,
		"cosmosdbThroughput",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetPublicNetworkAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HealthcareService) SetTags(val *map[string]*string) {
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
func HealthcareService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HealthcareService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthcareService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HealthcareService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareService) PutAuthenticationConfiguration(value *HealthcareServiceAuthenticationConfiguration) {
	_jsii_.InvokeVoid(
		h,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) PutCorsConfiguration(value *HealthcareServiceCorsConfiguration) {
	_jsii_.InvokeVoid(
		h,
		"putCorsConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) PutTimeouts(value *HealthcareServiceTimeouts) {
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) ResetAccessPolicyObjectIds() {
	_jsii_.InvokeVoid(
		h,
		"resetAccessPolicyObjectIds",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		h,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetCorsConfiguration() {
	_jsii_.InvokeVoid(
		h,
		"resetCorsConfiguration",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetCosmosdbKeyVaultKeyVersionlessId() {
	_jsii_.InvokeVoid(
		h,
		"resetCosmosdbKeyVaultKeyVersionlessId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetCosmosdbThroughput() {
	_jsii_.InvokeVoid(
		h,
		"resetCosmosdbThroughput",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetKind() {
	_jsii_.InvokeVoid(
		h,
		"resetKind",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HealthcareServiceAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#audience HealthcareService#audience}.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#authority HealthcareService#authority}.
	Authority *string `field:"optional" json:"authority" yaml:"authority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#smart_proxy_enabled HealthcareService#smart_proxy_enabled}.
	SmartProxyEnabled interface{} `field:"optional" json:"smartProxyEnabled" yaml:"smartProxyEnabled"`
}

type HealthcareServiceAuthenticationConfigurationOutputReference interface {
	cdktf.ComplexObject
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
	Authority() *string
	SetAuthority(val *string)
	AuthorityInput() *string
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
	InternalValue() *HealthcareServiceAuthenticationConfiguration
	SetInternalValue(val *HealthcareServiceAuthenticationConfiguration)
	SmartProxyEnabled() interface{}
	SetSmartProxyEnabled(val interface{})
	SmartProxyEnabledInput() interface{}
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
	ResetAudience()
	ResetAuthority()
	ResetSmartProxyEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcareServiceAuthenticationConfigurationOutputReference
type jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) Authority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) AuthorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) InternalValue() *HealthcareServiceAuthenticationConfiguration {
	var returns *HealthcareServiceAuthenticationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SmartProxyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smartProxyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SmartProxyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smartProxyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcareServiceAuthenticationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareServiceAuthenticationConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareServiceAuthenticationConfigurationOutputReference_Override(h HealthcareServiceAuthenticationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetAudience(val *string) {
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetAuthority(val *string) {
	_jsii_.Set(
		j,
		"authority",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetInternalValue(val *HealthcareServiceAuthenticationConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetSmartProxyEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"smartProxyEnabled",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		h,
		"resetAudience",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ResetAuthority() {
	_jsii_.InvokeVoid(
		h,
		"resetAuthority",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ResetSmartProxyEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetSmartProxyEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceAuthenticationConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HealthcareServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#location HealthcareService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#name HealthcareService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#resource_group_name HealthcareService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#access_policy_object_ids HealthcareService#access_policy_object_ids}.
	AccessPolicyObjectIds *[]*string `field:"optional" json:"accessPolicyObjectIds" yaml:"accessPolicyObjectIds"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#authentication_configuration HealthcareService#authentication_configuration}
	AuthenticationConfiguration *HealthcareServiceAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// cors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#cors_configuration HealthcareService#cors_configuration}
	CorsConfiguration *HealthcareServiceCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#cosmosdb_key_vault_key_versionless_id HealthcareService#cosmosdb_key_vault_key_versionless_id}.
	CosmosdbKeyVaultKeyVersionlessId *string `field:"optional" json:"cosmosdbKeyVaultKeyVersionlessId" yaml:"cosmosdbKeyVaultKeyVersionlessId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#cosmosdb_throughput HealthcareService#cosmosdb_throughput}.
	CosmosdbThroughput *float64 `field:"optional" json:"cosmosdbThroughput" yaml:"cosmosdbThroughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#id HealthcareService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#kind HealthcareService#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#public_network_access_enabled HealthcareService#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#tags HealthcareService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#timeouts HealthcareService#timeouts}
	Timeouts *HealthcareServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type HealthcareServiceCorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#allow_credentials HealthcareService#allow_credentials}.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#allowed_headers HealthcareService#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#allowed_methods HealthcareService#allowed_methods}.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#allowed_origins HealthcareService#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#max_age_in_seconds HealthcareService#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"optional" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

type HealthcareServiceCorsConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllowCredentials() interface{}
	SetAllowCredentials(val interface{})
	AllowCredentialsInput() interface{}
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	InternalValue() *HealthcareServiceCorsConfiguration
	SetInternalValue(val *HealthcareServiceCorsConfiguration)
	MaxAgeInSeconds() *float64
	SetMaxAgeInSeconds(val *float64)
	MaxAgeInSecondsInput() *float64
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
	ResetAllowCredentials()
	ResetAllowedHeaders()
	ResetAllowedMethods()
	ResetAllowedOrigins()
	ResetMaxAgeInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcareServiceCorsConfigurationOutputReference
type jsiiProxy_HealthcareServiceCorsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) InternalValue() *HealthcareServiceCorsConfiguration {
	var returns *HealthcareServiceCorsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) MaxAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) MaxAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcareServiceCorsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareServiceCorsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HealthcareServiceCorsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceCorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareServiceCorsConfigurationOutputReference_Override(h HealthcareServiceCorsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceCorsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetAllowCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"allowCredentials",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetAllowedHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetAllowedOrigins(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetInternalValue(val *HealthcareServiceCorsConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetMaxAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ResetAllowCredentials() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowCredentials",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ResetAllowedHeaders() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowedHeaders",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		h,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ResetMaxAgeInSeconds() {
	_jsii_.InvokeVoid(
		h,
		"resetMaxAgeInSeconds",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceCorsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HealthcareServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#create HealthcareService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#delete HealthcareService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#read HealthcareService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/healthcare_service#update HealthcareService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type HealthcareServiceTimeoutsOutputReference interface {
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

// The jsii proxy struct for HealthcareServiceTimeoutsOutputReference
type jsiiProxy_HealthcareServiceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewHealthcareServiceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareServiceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HealthcareServiceTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareServiceTimeoutsOutputReference_Override(h HealthcareServiceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HealthcareServiceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		h,
		"resetDelete",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		h,
		"resetRead",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		h,
		"resetUpdate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareServiceTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

