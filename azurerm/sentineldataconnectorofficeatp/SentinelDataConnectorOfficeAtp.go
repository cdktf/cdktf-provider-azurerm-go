package sentineldataconnectorofficeatp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/sentineldataconnectorofficeatp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp azurerm_sentinel_data_connector_office_atp}.
type SentinelDataConnectorOfficeAtp interface {
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
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
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
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SentinelDataConnectorOfficeAtpTimeoutsOutputReference
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
	PutTimeouts(value *SentinelDataConnectorOfficeAtpTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTenantId()
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

// The jsii proxy struct for SentinelDataConnectorOfficeAtp
type jsiiProxy_SentinelDataConnectorOfficeAtp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) Timeouts() SentinelDataConnectorOfficeAtpTimeoutsOutputReference {
	var returns SentinelDataConnectorOfficeAtpTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp azurerm_sentinel_data_connector_office_atp} Resource.
func NewSentinelDataConnectorOfficeAtp(scope constructs.Construct, id *string, config *SentinelDataConnectorOfficeAtpConfig) SentinelDataConnectorOfficeAtp {
	_init_.Initialize()

	if err := validateNewSentinelDataConnectorOfficeAtpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelDataConnectorOfficeAtp{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_data_connector_office_atp azurerm_sentinel_data_connector_office_atp} Resource.
func NewSentinelDataConnectorOfficeAtp_Override(s SentinelDataConnectorOfficeAtp, scope constructs.Construct, id *string, config *SentinelDataConnectorOfficeAtpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetLogAnalyticsWorkspaceId(val *string) {
	if err := j.validateSetLogAnalyticsWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelDataConnectorOfficeAtp)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
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
func SentinelDataConnectorOfficeAtp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelDataConnectorOfficeAtp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelDataConnectorOfficeAtp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelDataConnectorOfficeAtp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelDataConnectorOfficeAtp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelDataConnectorOfficeAtp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelDataConnectorOfficeAtp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelDataConnectorOfficeAtp.SentinelDataConnectorOfficeAtp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) PutTimeouts(value *SentinelDataConnectorOfficeAtpTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ResetTenantId() {
	_jsii_.InvokeVoid(
		s,
		"resetTenantId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelDataConnectorOfficeAtp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
