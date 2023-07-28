package redisenterprisedatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/redisenterprisedatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/redis_enterprise_database azurerm_redis_enterprise_database}.
type RedisEnterpriseDatabase interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientProtocol() *string
	SetClientProtocol(val *string)
	ClientProtocolInput() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusteringPolicy() *string
	SetClusteringPolicy(val *string)
	ClusteringPolicyInput() *string
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
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedDatabaseGroupNickname() *string
	SetLinkedDatabaseGroupNickname(val *string)
	LinkedDatabaseGroupNicknameInput() *string
	LinkedDatabaseId() *[]*string
	SetLinkedDatabaseId(val *[]*string)
	LinkedDatabaseIdInput() *[]*string
	Module() RedisEnterpriseDatabaseModuleList
	ModuleInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PrimaryAccessKey() *string
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
	SecondaryAccessKey() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RedisEnterpriseDatabaseTimeoutsOutputReference
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
	PutModule(value interface{})
	PutTimeouts(value *RedisEnterpriseDatabaseTimeouts)
	ResetClientProtocol()
	ResetClusteringPolicy()
	ResetEvictionPolicy()
	ResetId()
	ResetLinkedDatabaseGroupNickname()
	ResetLinkedDatabaseId()
	ResetModule()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetResourceGroupName()
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

// The jsii proxy struct for RedisEnterpriseDatabase
type jsiiProxy_RedisEnterpriseDatabase struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedisEnterpriseDatabase) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClientProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClientProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClusteringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ClusteringPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusteringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) LinkedDatabaseGroupNickname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedDatabaseGroupNickname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) LinkedDatabaseGroupNicknameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedDatabaseGroupNicknameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) LinkedDatabaseId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) LinkedDatabaseIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Module() RedisEnterpriseDatabaseModuleList {
	var returns RedisEnterpriseDatabaseModuleList
	_jsii_.Get(
		j,
		"module",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ModuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) PrimaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) SecondaryAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) Timeouts() RedisEnterpriseDatabaseTimeoutsOutputReference {
	var returns RedisEnterpriseDatabaseTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisEnterpriseDatabase) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/redis_enterprise_database azurerm_redis_enterprise_database} Resource.
func NewRedisEnterpriseDatabase(scope constructs.Construct, id *string, config *RedisEnterpriseDatabaseConfig) RedisEnterpriseDatabase {
	_init_.Initialize()

	if err := validateNewRedisEnterpriseDatabaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisEnterpriseDatabase{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/redis_enterprise_database azurerm_redis_enterprise_database} Resource.
func NewRedisEnterpriseDatabase_Override(r RedisEnterpriseDatabase, scope constructs.Construct, id *string, config *RedisEnterpriseDatabaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetClientProtocol(val *string) {
	if err := j.validateSetClientProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientProtocol",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetClusteringPolicy(val *string) {
	if err := j.validateSetClusteringPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusteringPolicy",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetLinkedDatabaseGroupNickname(val *string) {
	if err := j.validateSetLinkedDatabaseGroupNicknameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedDatabaseGroupNickname",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetLinkedDatabaseId(val *[]*string) {
	if err := j.validateSetLinkedDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedDatabaseId",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedisEnterpriseDatabase)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func RedisEnterpriseDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisEnterpriseDatabase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedisEnterpriseDatabase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisEnterpriseDatabase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedisEnterpriseDatabase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisEnterpriseDatabase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedisEnterpriseDatabase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.redisEnterpriseDatabase.RedisEnterpriseDatabase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) PutModule(value interface{}) {
	if err := r.validatePutModuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putModule",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) PutTimeouts(value *RedisEnterpriseDatabaseTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetClientProtocol() {
	_jsii_.InvokeVoid(
		r,
		"resetClientProtocol",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetClusteringPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetClusteringPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetLinkedDatabaseGroupNickname() {
	_jsii_.InvokeVoid(
		r,
		"resetLinkedDatabaseGroupNickname",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetLinkedDatabaseId() {
	_jsii_.InvokeVoid(
		r,
		"resetLinkedDatabaseId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetModule() {
	_jsii_.InvokeVoid(
		r,
		"resetModule",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetResourceGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisEnterpriseDatabase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisEnterpriseDatabase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

