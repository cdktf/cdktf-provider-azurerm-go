package siterecoveryhypervreplicationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/siterecoveryhypervreplicationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy azurerm_site_recovery_hyperv_replication_policy}.
type SiteRecoveryHypervReplicationPolicy interface {
	cdktf.TerraformResource
	ApplicationConsistentSnapshotFrequencyInHours() *float64
	SetApplicationConsistentSnapshotFrequencyInHours(val *float64)
	ApplicationConsistentSnapshotFrequencyInHoursInput() *float64
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
	// Experimental.
	RawOverrides() interface{}
	RecoveryPointRetentionInHours() *float64
	SetRecoveryPointRetentionInHours(val *float64)
	RecoveryPointRetentionInHoursInput() *float64
	RecoveryVaultId() *string
	SetRecoveryVaultId(val *string)
	RecoveryVaultIdInput() *string
	ReplicationIntervalInSeconds() *float64
	SetReplicationIntervalInSeconds(val *float64)
	ReplicationIntervalInSecondsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SiteRecoveryHypervReplicationPolicyTimeoutsOutputReference
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
	PutTimeouts(value *SiteRecoveryHypervReplicationPolicyTimeouts)
	ResetId()
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

// The jsii proxy struct for SiteRecoveryHypervReplicationPolicy
type jsiiProxy_SiteRecoveryHypervReplicationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ApplicationConsistentSnapshotFrequencyInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationConsistentSnapshotFrequencyInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ApplicationConsistentSnapshotFrequencyInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationConsistentSnapshotFrequencyInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) RecoveryPointRetentionInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryPointRetentionInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) RecoveryPointRetentionInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recoveryPointRetentionInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) RecoveryVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) RecoveryVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recoveryVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ReplicationIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ReplicationIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicationIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) Timeouts() SiteRecoveryHypervReplicationPolicyTimeoutsOutputReference {
	var returns SiteRecoveryHypervReplicationPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy azurerm_site_recovery_hyperv_replication_policy} Resource.
func NewSiteRecoveryHypervReplicationPolicy(scope constructs.Construct, id *string, config *SiteRecoveryHypervReplicationPolicyConfig) SiteRecoveryHypervReplicationPolicy {
	_init_.Initialize()

	if err := validateNewSiteRecoveryHypervReplicationPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SiteRecoveryHypervReplicationPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_hyperv_replication_policy azurerm_site_recovery_hyperv_replication_policy} Resource.
func NewSiteRecoveryHypervReplicationPolicy_Override(s SiteRecoveryHypervReplicationPolicy, scope constructs.Construct, id *string, config *SiteRecoveryHypervReplicationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetApplicationConsistentSnapshotFrequencyInHours(val *float64) {
	if err := j.validateSetApplicationConsistentSnapshotFrequencyInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationConsistentSnapshotFrequencyInHours",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetRecoveryPointRetentionInHours(val *float64) {
	if err := j.validateSetRecoveryPointRetentionInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryPointRetentionInHours",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetRecoveryVaultId(val *string) {
	if err := j.validateSetRecoveryVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recoveryVaultId",
		val,
	)
}

func (j *jsiiProxy_SiteRecoveryHypervReplicationPolicy)SetReplicationIntervalInSeconds(val *float64) {
	if err := j.validateSetReplicationIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationIntervalInSeconds",
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
func SiteRecoveryHypervReplicationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryHypervReplicationPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryHypervReplicationPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryHypervReplicationPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SiteRecoveryHypervReplicationPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSiteRecoveryHypervReplicationPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SiteRecoveryHypervReplicationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.siteRecoveryHypervReplicationPolicy.SiteRecoveryHypervReplicationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) PutTimeouts(value *SiteRecoveryHypervReplicationPolicyTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SiteRecoveryHypervReplicationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
