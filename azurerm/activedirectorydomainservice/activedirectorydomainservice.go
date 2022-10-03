package activedirectorydomainservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/activedirectorydomainservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service azurerm_active_directory_domain_service}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_ActiveDirectoryDomainService) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service azurerm_active_directory_domain_service} Resource.
func NewActiveDirectoryDomainService(scope constructs.Construct, id *string, config *ActiveDirectoryDomainServiceConfig) ActiveDirectoryDomainService {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainService{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service azurerm_active_directory_domain_service} Resource.
func NewActiveDirectoryDomainService_Override(a ActiveDirectoryDomainService, scope constructs.Construct, id *string, config *ActiveDirectoryDomainServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetDomainConfigurationType(val *string) {
	_jsii_.Set(
		j,
		"domainConfigurationType",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetFilteredSyncEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"filteredSyncEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainService) SetTags(val *map[string]*string) {
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
func ActiveDirectoryDomainService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainService",
		"isConstruct",
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

func (a *jsiiProxy_ActiveDirectoryDomainService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutInitialReplicaSet(value *ActiveDirectoryDomainServiceInitialReplicaSet) {
	_jsii_.InvokeVoid(
		a,
		"putInitialReplicaSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutNotifications(value *ActiveDirectoryDomainServiceNotifications) {
	_jsii_.InvokeVoid(
		a,
		"putNotifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutSecureLdap(value *ActiveDirectoryDomainServiceSecureLdap) {
	_jsii_.InvokeVoid(
		a,
		"putSecureLdap",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutSecurity(value *ActiveDirectoryDomainServiceSecurity) {
	_jsii_.InvokeVoid(
		a,
		"putSecurity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainService) PutTimeouts(value *ActiveDirectoryDomainServiceTimeouts) {
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

type ActiveDirectoryDomainServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#domain_name ActiveDirectoryDomainService#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// initial_replica_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#initial_replica_set ActiveDirectoryDomainService#initial_replica_set}
	InitialReplicaSet *ActiveDirectoryDomainServiceInitialReplicaSet `field:"required" json:"initialReplicaSet" yaml:"initialReplicaSet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#location ActiveDirectoryDomainService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#name ActiveDirectoryDomainService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#resource_group_name ActiveDirectoryDomainService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#sku ActiveDirectoryDomainService#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#domain_configuration_type ActiveDirectoryDomainService#domain_configuration_type}.
	DomainConfigurationType *string `field:"optional" json:"domainConfigurationType" yaml:"domainConfigurationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#filtered_sync_enabled ActiveDirectoryDomainService#filtered_sync_enabled}.
	FilteredSyncEnabled interface{} `field:"optional" json:"filteredSyncEnabled" yaml:"filteredSyncEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#id ActiveDirectoryDomainService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#notifications ActiveDirectoryDomainService#notifications}
	Notifications *ActiveDirectoryDomainServiceNotifications `field:"optional" json:"notifications" yaml:"notifications"`
	// secure_ldap block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#secure_ldap ActiveDirectoryDomainService#secure_ldap}
	SecureLdap *ActiveDirectoryDomainServiceSecureLdap `field:"optional" json:"secureLdap" yaml:"secureLdap"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#security ActiveDirectoryDomainService#security}
	Security *ActiveDirectoryDomainServiceSecurity `field:"optional" json:"security" yaml:"security"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#tags ActiveDirectoryDomainService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#timeouts ActiveDirectoryDomainService#timeouts}
	Timeouts *ActiveDirectoryDomainServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type ActiveDirectoryDomainServiceInitialReplicaSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#subnet_id ActiveDirectoryDomainService#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

type ActiveDirectoryDomainServiceInitialReplicaSetOutputReference interface {
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
	DomainControllerIpAddresses() *[]*string
	ExternalAccessIpAddress() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *ActiveDirectoryDomainServiceInitialReplicaSet
	SetInternalValue(val *ActiveDirectoryDomainServiceInitialReplicaSet)
	Location() *string
	ServiceStatus() *string
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

// The jsii proxy struct for ActiveDirectoryDomainServiceInitialReplicaSetOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) DomainControllerIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainControllerIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ExternalAccessIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalAccessIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) InternalValue() *ActiveDirectoryDomainServiceInitialReplicaSet {
	var returns *ActiveDirectoryDomainServiceInitialReplicaSet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ServiceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceInitialReplicaSetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceInitialReplicaSetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceInitialReplicaSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceInitialReplicaSetOutputReference_Override(a ActiveDirectoryDomainServiceInitialReplicaSetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceInitialReplicaSetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetInternalValue(val *ActiveDirectoryDomainServiceInitialReplicaSet) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceInitialReplicaSetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ActiveDirectoryDomainServiceNotifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#additional_recipients ActiveDirectoryDomainService#additional_recipients}.
	AdditionalRecipients *[]*string `field:"optional" json:"additionalRecipients" yaml:"additionalRecipients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#notify_dc_admins ActiveDirectoryDomainService#notify_dc_admins}.
	NotifyDcAdmins interface{} `field:"optional" json:"notifyDcAdmins" yaml:"notifyDcAdmins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#notify_global_admins ActiveDirectoryDomainService#notify_global_admins}.
	NotifyGlobalAdmins interface{} `field:"optional" json:"notifyGlobalAdmins" yaml:"notifyGlobalAdmins"`
}

type ActiveDirectoryDomainServiceNotificationsOutputReference interface {
	cdktf.ComplexObject
	AdditionalRecipients() *[]*string
	SetAdditionalRecipients(val *[]*string)
	AdditionalRecipientsInput() *[]*string
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
	InternalValue() *ActiveDirectoryDomainServiceNotifications
	SetInternalValue(val *ActiveDirectoryDomainServiceNotifications)
	NotifyDcAdmins() interface{}
	SetNotifyDcAdmins(val interface{})
	NotifyDcAdminsInput() interface{}
	NotifyGlobalAdmins() interface{}
	SetNotifyGlobalAdmins(val interface{})
	NotifyGlobalAdminsInput() interface{}
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
	ResetAdditionalRecipients()
	ResetNotifyDcAdmins()
	ResetNotifyGlobalAdmins()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveDirectoryDomainServiceNotificationsOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) AdditionalRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) AdditionalRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) InternalValue() *ActiveDirectoryDomainServiceNotifications {
	var returns *ActiveDirectoryDomainServiceNotifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) NotifyDcAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyDcAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) NotifyDcAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyDcAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) NotifyGlobalAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyGlobalAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) NotifyGlobalAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyGlobalAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceNotificationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceNotificationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceNotificationsOutputReference_Override(a ActiveDirectoryDomainServiceNotificationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetAdditionalRecipients(val *[]*string) {
	_jsii_.Set(
		j,
		"additionalRecipients",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetInternalValue(val *ActiveDirectoryDomainServiceNotifications) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetNotifyDcAdmins(val interface{}) {
	_jsii_.Set(
		j,
		"notifyDcAdmins",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetNotifyGlobalAdmins(val interface{}) {
	_jsii_.Set(
		j,
		"notifyGlobalAdmins",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ResetAdditionalRecipients() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalRecipients",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ResetNotifyDcAdmins() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifyDcAdmins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ResetNotifyGlobalAdmins() {
	_jsii_.InvokeVoid(
		a,
		"resetNotifyGlobalAdmins",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceNotificationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ActiveDirectoryDomainServiceSecureLdap struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#enabled ActiveDirectoryDomainService#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#pfx_certificate ActiveDirectoryDomainService#pfx_certificate}.
	PfxCertificate *string `field:"required" json:"pfxCertificate" yaml:"pfxCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#pfx_certificate_password ActiveDirectoryDomainService#pfx_certificate_password}.
	PfxCertificatePassword *string `field:"required" json:"pfxCertificatePassword" yaml:"pfxCertificatePassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#external_access_enabled ActiveDirectoryDomainService#external_access_enabled}.
	ExternalAccessEnabled interface{} `field:"optional" json:"externalAccessEnabled" yaml:"externalAccessEnabled"`
}

type ActiveDirectoryDomainServiceSecureLdapOutputReference interface {
	cdktf.ComplexObject
	CertificateExpiry() *string
	CertificateThumbprint() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExternalAccessEnabled() interface{}
	SetExternalAccessEnabled(val interface{})
	ExternalAccessEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ActiveDirectoryDomainServiceSecureLdap
	SetInternalValue(val *ActiveDirectoryDomainServiceSecureLdap)
	PfxCertificate() *string
	SetPfxCertificate(val *string)
	PfxCertificateInput() *string
	PfxCertificatePassword() *string
	SetPfxCertificatePassword(val *string)
	PfxCertificatePasswordInput() *string
	PublicCertificate() *string
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
	ResetExternalAccessEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveDirectoryDomainServiceSecureLdapOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CertificateExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ExternalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ExternalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InternalValue() *ActiveDirectoryDomainServiceSecureLdap {
	var returns *ActiveDirectoryDomainServiceSecureLdap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PublicCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceSecureLdapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceSecureLdapOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecureLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceSecureLdapOutputReference_Override(a ActiveDirectoryDomainServiceSecureLdapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecureLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetExternalAccessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"externalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetInternalValue(val *ActiveDirectoryDomainServiceSecureLdap) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetPfxCertificate(val *string) {
	_jsii_.Set(
		j,
		"pfxCertificate",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetPfxCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"pfxCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ResetExternalAccessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAccessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ActiveDirectoryDomainServiceSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#kerberos_armoring_enabled ActiveDirectoryDomainService#kerberos_armoring_enabled}.
	KerberosArmoringEnabled interface{} `field:"optional" json:"kerberosArmoringEnabled" yaml:"kerberosArmoringEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#kerberos_rc4_encryption_enabled ActiveDirectoryDomainService#kerberos_rc4_encryption_enabled}.
	KerberosRc4EncryptionEnabled interface{} `field:"optional" json:"kerberosRc4EncryptionEnabled" yaml:"kerberosRc4EncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#ntlm_v1_enabled ActiveDirectoryDomainService#ntlm_v1_enabled}.
	NtlmV1Enabled interface{} `field:"optional" json:"ntlmV1Enabled" yaml:"ntlmV1Enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#sync_kerberos_passwords ActiveDirectoryDomainService#sync_kerberos_passwords}.
	SyncKerberosPasswords interface{} `field:"optional" json:"syncKerberosPasswords" yaml:"syncKerberosPasswords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#sync_ntlm_passwords ActiveDirectoryDomainService#sync_ntlm_passwords}.
	SyncNtlmPasswords interface{} `field:"optional" json:"syncNtlmPasswords" yaml:"syncNtlmPasswords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#sync_on_prem_passwords ActiveDirectoryDomainService#sync_on_prem_passwords}.
	SyncOnPremPasswords interface{} `field:"optional" json:"syncOnPremPasswords" yaml:"syncOnPremPasswords"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#tls_v1_enabled ActiveDirectoryDomainService#tls_v1_enabled}.
	TlsV1Enabled interface{} `field:"optional" json:"tlsV1Enabled" yaml:"tlsV1Enabled"`
}

type ActiveDirectoryDomainServiceSecurityOutputReference interface {
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
	InternalValue() *ActiveDirectoryDomainServiceSecurity
	SetInternalValue(val *ActiveDirectoryDomainServiceSecurity)
	KerberosArmoringEnabled() interface{}
	SetKerberosArmoringEnabled(val interface{})
	KerberosArmoringEnabledInput() interface{}
	KerberosRc4EncryptionEnabled() interface{}
	SetKerberosRc4EncryptionEnabled(val interface{})
	KerberosRc4EncryptionEnabledInput() interface{}
	NtlmV1Enabled() interface{}
	SetNtlmV1Enabled(val interface{})
	NtlmV1EnabledInput() interface{}
	SyncKerberosPasswords() interface{}
	SetSyncKerberosPasswords(val interface{})
	SyncKerberosPasswordsInput() interface{}
	SyncNtlmPasswords() interface{}
	SetSyncNtlmPasswords(val interface{})
	SyncNtlmPasswordsInput() interface{}
	SyncOnPremPasswords() interface{}
	SetSyncOnPremPasswords(val interface{})
	SyncOnPremPasswordsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsV1Enabled() interface{}
	SetTlsV1Enabled(val interface{})
	TlsV1EnabledInput() interface{}
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
	ResetKerberosArmoringEnabled()
	ResetKerberosRc4EncryptionEnabled()
	ResetNtlmV1Enabled()
	ResetSyncKerberosPasswords()
	ResetSyncNtlmPasswords()
	ResetSyncOnPremPasswords()
	ResetTlsV1Enabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveDirectoryDomainServiceSecurityOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InternalValue() *ActiveDirectoryDomainServiceSecurity {
	var returns *ActiveDirectoryDomainServiceSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosArmoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosArmoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosArmoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosArmoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosRc4EncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosRc4EncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) KerberosRc4EncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosRc4EncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) NtlmV1Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ntlmV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) NtlmV1EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ntlmV1EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncKerberosPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncKerberosPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncKerberosPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncKerberosPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncNtlmPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncNtlmPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncNtlmPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncNtlmPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncOnPremPasswords() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnPremPasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SyncOnPremPasswordsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncOnPremPasswordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TlsV1Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsV1Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) TlsV1EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsV1EnabledInput",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceSecurityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceSecurityOutputReference_Override(a ActiveDirectoryDomainServiceSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetInternalValue(val *ActiveDirectoryDomainServiceSecurity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetKerberosArmoringEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"kerberosArmoringEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetKerberosRc4EncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"kerberosRc4EncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetNtlmV1Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"ntlmV1Enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetSyncKerberosPasswords(val interface{}) {
	_jsii_.Set(
		j,
		"syncKerberosPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetSyncNtlmPasswords(val interface{}) {
	_jsii_.Set(
		j,
		"syncNtlmPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetSyncOnPremPasswords(val interface{}) {
	_jsii_.Set(
		j,
		"syncOnPremPasswords",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) SetTlsV1Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"tlsV1Enabled",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetKerberosArmoringEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosArmoringEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetKerberosRc4EncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetKerberosRc4EncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetNtlmV1Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetNtlmV1Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncKerberosPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncKerberosPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncNtlmPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncNtlmPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetSyncOnPremPasswords() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncOnPremPasswords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ResetTlsV1Enabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsV1Enabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ActiveDirectoryDomainServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#create ActiveDirectoryDomainService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#delete ActiveDirectoryDomainService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#read ActiveDirectoryDomainService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/active_directory_domain_service#update ActiveDirectoryDomainService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ActiveDirectoryDomainServiceTimeoutsOutputReference interface {
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

// The jsii proxy struct for ActiveDirectoryDomainServiceTimeoutsOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceTimeoutsOutputReference_Override(a ActiveDirectoryDomainServiceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		a,
		"resetRead",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

