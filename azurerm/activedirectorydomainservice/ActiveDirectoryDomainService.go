// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/activedirectorydomainservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/active_directory_domain_service azurerm_active_directory_domain_service}.
type ActiveDirectoryDomainService interface {
	cdktf.TerraformResource
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentId() *string
	DomainConfigurationType() *string
	SetDomainConfigurationType(val *string)
	DomainConfigurationTypeInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	FilteredSyncEnabled() interface{}
	SetFilteredSyncEnabled(val interface{})
	FilteredSyncEnabledInput() interface{}
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
	InitialReplicaSet() ActiveDirectoryDomainServiceInitialReplicaSetOutputReference
	InitialReplicaSetInput() *ActiveDirectoryDomainServiceInitialReplicaSet
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
	Notifications() ActiveDirectoryDomainServiceNotificationsOutputReference
	NotificationsInput() *ActiveDirectoryDomainServiceNotifications
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
	ResourceId() *string
	SecureLdap() ActiveDirectoryDomainServiceSecureLdapOutputReference
	SecureLdapInput() *ActiveDirectoryDomainServiceSecureLdap
	Security() ActiveDirectoryDomainServiceSecurityOutputReference
	SecurityInput() *ActiveDirectoryDomainServiceSecurity
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SyncOwner() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ActiveDirectoryDomainServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *float64
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
	PutInitialReplicaSet(value *ActiveDirectoryDomainServiceInitialReplicaSet)
	PutNotifications(value *ActiveDirectoryDomainServiceNotifications)
	PutSecureLdap(value *ActiveDirectoryDomainServiceSecureLdap)
	PutSecurity(value *ActiveDirectoryDomainServiceSecurity)
	PutTimeouts(value *ActiveDirectoryDomainServiceTimeouts)
	ResetDomainConfigurationType()
	ResetFilteredSyncEnabled()
	ResetId()
	ResetNotifications()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecureLdap()
	ResetSecurity()
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

// The jsii proxy struct for ActiveDirectoryDomainService
type jsiiProxy_ActiveDirectoryDomainService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ActiveDirectoryDomainService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DomainConfigurationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DomainConfigurationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainConfigurationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) FilteredSyncEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filteredSyncEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) FilteredSyncEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filteredSyncEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) InitialReplicaSet() ActiveDirectoryDomainServiceInitialReplicaSetOutputReference {
	var returns ActiveDirectoryDomainServiceInitialReplicaSetOutputReference
	_jsii_.Get(
		j,
		"initialReplicaSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) InitialReplicaSetInput() *ActiveDirectoryDomainServiceInitialReplicaSet {
	var returns *ActiveDirectoryDomainServiceInitialReplicaSet
	_jsii_.Get(
		j,
		"initialReplicaSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Notifications() ActiveDirectoryDomainServiceNotificationsOutputReference {
	var returns ActiveDirectoryDomainServiceNotificationsOutputReference
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) NotificationsInput() *ActiveDirectoryDomainServiceNotifications {
	var returns *ActiveDirectoryDomainServiceNotifications
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SecureLdap() ActiveDirectoryDomainServiceSecureLdapOutputReference {
	var returns ActiveDirectoryDomainServiceSecureLdapOutputReference
	_jsii_.Get(
		j,
		"secureLdap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SecureLdapInput() *ActiveDirectoryDomainServiceSecureLdap {
	var returns *ActiveDirectoryDomainServiceSecureLdap
	_jsii_.Get(
		j,
		"secureLdapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Security() ActiveDirectoryDomainServiceSecurityOutputReference {
	var returns ActiveDirectoryDomainServiceSecurityOutputReference
	_jsii_.Get(
		j,
		"security",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SecurityInput() *ActiveDirectoryDomainServiceSecurity {
	var returns *ActiveDirectoryDomainServiceSecurity
	_jsii_.Get(
		j,
		"securityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SyncOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syncOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Timeouts() ActiveDirectoryDomainServiceTimeoutsOutputReference {
	var returns ActiveDirectoryDomainServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainService) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/active_directory_domain_service azurerm_active_directory_domain_service} Resource.
func NewActiveDirectoryDomainService(scope constructs.Construct, id *string, config *ActiveDirectoryDomainServiceConfig) ActiveDirectoryDomainService {
	_init_.Initialize()

	if err := validateNewActiveDirectoryDomainServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveDirectoryDomainService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/active_directory_domain_service azurerm_active_directory_domain_service} Resource.
func NewActiveDirectoryDomainService_Override(a ActiveDirectoryDomainService, scope constructs.Construct, id *string, config *ActiveDirectoryDomainServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetDomainConfigurationType(val *string) {
	if err := j.validateSetDomainConfigurationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainConfigurationType",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetFilteredSyncEnabled(val interface{}) {
	if err := j.validateSetFilteredSyncEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filteredSyncEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a ActiveDirectoryDomainService resource upon running "cdktf plan <stack-name>".
func ActiveDirectoryDomainService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
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
func ActiveDirectoryDomainService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActiveDirectoryDomainService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActiveDirectoryDomainService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveDirectoryDomainService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ActiveDirectoryDomainService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutInitialReplicaSet(value *ActiveDirectoryDomainServiceInitialReplicaSet) {
	if err := a.validatePutInitialReplicaSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialReplicaSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutNotifications(value *ActiveDirectoryDomainServiceNotifications) {
	if err := a.validatePutNotificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutSecureLdap(value *ActiveDirectoryDomainServiceSecureLdap) {
	if err := a.validatePutSecureLdapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecureLdap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutSecurity(value *ActiveDirectoryDomainServiceSecurity) {
	if err := a.validatePutSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutTimeouts(value *ActiveDirectoryDomainServiceTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetDomainConfigurationType() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainConfigurationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetFilteredSyncEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetFilteredSyncEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetNotifications() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetSecureLdap() {
	_jsii_.InvokeVoid(
		a,
		"resetSecureLdap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

