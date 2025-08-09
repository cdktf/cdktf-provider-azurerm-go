// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/paloaltolocalrulestackrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule azurerm_palo_alto_local_rulestack_rule}.
type PaloAltoLocalRulestackRule interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	Applications() *[]*string
	SetApplications(val *[]*string)
	ApplicationsInput() *[]*string
	AuditComment() *string
	SetAuditComment(val *string)
	AuditCommentInput() *string
	Category() PaloAltoLocalRulestackRuleCategoryOutputReference
	CategoryInput() *PaloAltoLocalRulestackRuleCategory
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
	DecryptionRuleType() *string
	SetDecryptionRuleType(val *string)
	DecryptionRuleTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Destination() PaloAltoLocalRulestackRuleDestinationOutputReference
	DestinationInput() *PaloAltoLocalRulestackRuleDestination
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	InspectionCertificateId() *string
	SetInspectionCertificateId(val *string)
	InspectionCertificateIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingEnabled() interface{}
	SetLoggingEnabled(val interface{})
	LoggingEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NegateDestination() interface{}
	SetNegateDestination(val interface{})
	NegateDestinationInput() interface{}
	NegateSource() interface{}
	SetNegateSource(val interface{})
	NegateSourceInput() interface{}
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	ProtocolPorts() *[]*string
	SetProtocolPorts(val *[]*string)
	ProtocolPortsInput() *[]*string
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
	RulestackId() *string
	SetRulestackId(val *string)
	RulestackIdInput() *string
	Source() PaloAltoLocalRulestackRuleSourceOutputReference
	SourceInput() *PaloAltoLocalRulestackRuleSource
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PaloAltoLocalRulestackRuleTimeoutsOutputReference
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
	PutCategory(value *PaloAltoLocalRulestackRuleCategory)
	PutDestination(value *PaloAltoLocalRulestackRuleDestination)
	PutSource(value *PaloAltoLocalRulestackRuleSource)
	PutTimeouts(value *PaloAltoLocalRulestackRuleTimeouts)
	ResetAuditComment()
	ResetCategory()
	ResetDecryptionRuleType()
	ResetDescription()
	ResetEnabled()
	ResetId()
	ResetInspectionCertificateId()
	ResetLoggingEnabled()
	ResetNegateDestination()
	ResetNegateSource()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtocol()
	ResetProtocolPorts()
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

// The jsii proxy struct for PaloAltoLocalRulestackRule
type jsiiProxy_PaloAltoLocalRulestackRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) AuditComment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditComment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) AuditCommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditCommentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Category() PaloAltoLocalRulestackRuleCategoryOutputReference {
	var returns PaloAltoLocalRulestackRuleCategoryOutputReference
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) CategoryInput() *PaloAltoLocalRulestackRuleCategory {
	var returns *PaloAltoLocalRulestackRuleCategory
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) DecryptionRuleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decryptionRuleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) DecryptionRuleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"decryptionRuleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Destination() PaloAltoLocalRulestackRuleDestinationOutputReference {
	var returns PaloAltoLocalRulestackRuleDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) DestinationInput() *PaloAltoLocalRulestackRuleDestination {
	var returns *PaloAltoLocalRulestackRuleDestination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) InspectionCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectionCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) InspectionCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectionCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) LoggingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) LoggingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) NegateDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) NegateDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) NegateSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) NegateSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negateSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ProtocolPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) ProtocolPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) RulestackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulestackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) RulestackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulestackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Source() PaloAltoLocalRulestackRuleSourceOutputReference {
	var returns PaloAltoLocalRulestackRuleSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) SourceInput() *PaloAltoLocalRulestackRuleSource {
	var returns *PaloAltoLocalRulestackRuleSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) Timeouts() PaloAltoLocalRulestackRuleTimeoutsOutputReference {
	var returns PaloAltoLocalRulestackRuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule azurerm_palo_alto_local_rulestack_rule} Resource.
func NewPaloAltoLocalRulestackRule(scope constructs.Construct, id *string, config *PaloAltoLocalRulestackRuleConfig) PaloAltoLocalRulestackRule {
	_init_.Initialize()

	if err := validateNewPaloAltoLocalRulestackRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoLocalRulestackRule{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/palo_alto_local_rulestack_rule azurerm_palo_alto_local_rulestack_rule} Resource.
func NewPaloAltoLocalRulestackRule_Override(p PaloAltoLocalRulestackRule, scope constructs.Construct, id *string, config *PaloAltoLocalRulestackRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetApplications(val *[]*string) {
	if err := j.validateSetApplicationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetAuditComment(val *string) {
	if err := j.validateSetAuditCommentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditComment",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetDecryptionRuleType(val *string) {
	if err := j.validateSetDecryptionRuleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decryptionRuleType",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetInspectionCertificateId(val *string) {
	if err := j.validateSetInspectionCertificateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectionCertificateId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetLoggingEnabled(val interface{}) {
	if err := j.validateSetLoggingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingEnabled",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetNegateDestination(val interface{}) {
	if err := j.validateSetNegateDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negateDestination",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetNegateSource(val interface{}) {
	if err := j.validateSetNegateSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negateSource",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetProtocolPorts(val *[]*string) {
	if err := j.validateSetProtocolPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolPorts",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetRulestackId(val *string) {
	if err := j.validateSetRulestackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulestackId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a PaloAltoLocalRulestackRule resource upon running "cdktf plan <stack-name>".
func PaloAltoLocalRulestackRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePaloAltoLocalRulestackRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
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
func PaloAltoLocalRulestackRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoLocalRulestackRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoLocalRulestackRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoLocalRulestackRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoLocalRulestackRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoLocalRulestackRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PaloAltoLocalRulestackRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) PutCategory(value *PaloAltoLocalRulestackRuleCategory) {
	if err := p.validatePutCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCategory",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) PutDestination(value *PaloAltoLocalRulestackRuleDestination) {
	if err := p.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDestination",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) PutSource(value *PaloAltoLocalRulestackRuleSource) {
	if err := p.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSource",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) PutTimeouts(value *PaloAltoLocalRulestackRuleTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetAuditComment() {
	_jsii_.InvokeVoid(
		p,
		"resetAuditComment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetCategory() {
	_jsii_.InvokeVoid(
		p,
		"resetCategory",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetDecryptionRuleType() {
	_jsii_.InvokeVoid(
		p,
		"resetDecryptionRuleType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetDescription() {
	_jsii_.InvokeVoid(
		p,
		"resetDescription",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetInspectionCertificateId() {
	_jsii_.InvokeVoid(
		p,
		"resetInspectionCertificateId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetLoggingEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetLoggingEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetNegateDestination() {
	_jsii_.InvokeVoid(
		p,
		"resetNegateDestination",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetNegateSource() {
	_jsii_.InvokeVoid(
		p,
		"resetNegateSource",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetProtocol() {
	_jsii_.InvokeVoid(
		p,
		"resetProtocol",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetProtocolPorts() {
	_jsii_.InvokeVoid(
		p,
		"resetProtocolPorts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

