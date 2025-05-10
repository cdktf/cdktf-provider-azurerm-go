// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcareservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/healthcareservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/healthcare_service azurerm_healthcare_service}.
type HealthcareService interface {
	cdktf.TerraformResource
	AccessPolicyObjectIds() *[]*string
	SetAccessPolicyObjectIds(val *[]*string)
	AccessPolicyObjectIdsInput() *[]*string
	AuthenticationConfiguration() HealthcareServiceAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() *HealthcareServiceAuthenticationConfiguration
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigurationExportStorageAccountName() *string
	SetConfigurationExportStorageAccountName(val *string)
	ConfigurationExportStorageAccountNameInput() *string
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
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
	Identity() HealthcareServiceIdentityOutputReference
	IdentityInput() *HealthcareServiceIdentity
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
	PutAuthenticationConfiguration(value *HealthcareServiceAuthenticationConfiguration)
	PutCorsConfiguration(value *HealthcareServiceCorsConfiguration)
	PutIdentity(value *HealthcareServiceIdentity)
	PutTimeouts(value *HealthcareServiceTimeouts)
	ResetAccessPolicyObjectIds()
	ResetAuthenticationConfiguration()
	ResetConfigurationExportStorageAccountName()
	ResetCorsConfiguration()
	ResetCosmosdbKeyVaultKeyVersionlessId()
	ResetCosmosdbThroughput()
	ResetId()
	ResetIdentity()
	ResetKind()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
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

func (j *jsiiProxy_HealthcareService) ConfigurationExportStorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationExportStorageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) ConfigurationExportStorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationExportStorageAccountNameInput",
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

func (j *jsiiProxy_HealthcareService) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_HealthcareService) Identity() HealthcareServiceIdentityOutputReference {
	var returns HealthcareServiceIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareService) IdentityInput() *HealthcareServiceIdentity {
	var returns *HealthcareServiceIdentity
	_jsii_.Get(
		j,
		"identityInput",
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


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/healthcare_service azurerm_healthcare_service} Resource.
func NewHealthcareService(scope constructs.Construct, id *string, config *HealthcareServiceConfig) HealthcareService {
	_init_.Initialize()

	if err := validateNewHealthcareServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/healthcare_service azurerm_healthcare_service} Resource.
func NewHealthcareService_Override(h HealthcareService, scope constructs.Construct, id *string, config *HealthcareServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareService)SetAccessPolicyObjectIds(val *[]*string) {
	if err := j.validateSetAccessPolicyObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicyObjectIds",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetConfigurationExportStorageAccountName(val *string) {
	if err := j.validateSetConfigurationExportStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationExportStorageAccountName",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetCosmosdbKeyVaultKeyVersionlessId(val *string) {
	if err := j.validateSetCosmosdbKeyVaultKeyVersionlessIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbKeyVaultKeyVersionlessId",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetCosmosdbThroughput(val *float64) {
	if err := j.validateSetCosmosdbThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbThroughput",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HealthcareService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a HealthcareService resource upon running "cdktf plan <stack-name>".
func HealthcareService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateHealthcareService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
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
func HealthcareService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.healthcareService.HealthcareService",
		"isTerraformResource",
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

func (h *jsiiProxy_HealthcareService) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HealthcareService) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HealthcareService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcareService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcareService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcareService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcareService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcareService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcareService) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcareService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcareService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HealthcareService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthcareService) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareService) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthcareService) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareService) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareService) PutAuthenticationConfiguration(value *HealthcareServiceAuthenticationConfiguration) {
	if err := h.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) PutCorsConfiguration(value *HealthcareServiceCorsConfiguration) {
	if err := h.validatePutCorsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putCorsConfiguration",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) PutIdentity(value *HealthcareServiceIdentity) {
	if err := h.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putIdentity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareService) PutTimeouts(value *HealthcareServiceTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
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

func (h *jsiiProxy_HealthcareService) ResetConfigurationExportStorageAccountName() {
	_jsii_.InvokeVoid(
		h,
		"resetConfigurationExportStorageAccountName",
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

func (h *jsiiProxy_HealthcareService) ResetIdentity() {
	_jsii_.InvokeVoid(
		h,
		"resetIdentity",
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

func (h *jsiiProxy_HealthcareService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
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

