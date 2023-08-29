// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebustopic

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopic",
		reflect.TypeOf((*ServicebusTopic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteOnIdle", GoGetter: "AutoDeleteOnIdle"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteOnIdleInput", GoGetter: "AutoDeleteOnIdleInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultMessageTtl", GoGetter: "DefaultMessageTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultMessageTtlInput", GoGetter: "DefaultMessageTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "duplicateDetectionHistoryTimeWindow", GoGetter: "DuplicateDetectionHistoryTimeWindow"},
			_jsii_.MemberProperty{JsiiProperty: "duplicateDetectionHistoryTimeWindowInput", GoGetter: "DuplicateDetectionHistoryTimeWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableBatchedOperations", GoGetter: "EnableBatchedOperations"},
			_jsii_.MemberProperty{JsiiProperty: "enableBatchedOperationsInput", GoGetter: "EnableBatchedOperationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableExpress", GoGetter: "EnableExpress"},
			_jsii_.MemberProperty{JsiiProperty: "enableExpressInput", GoGetter: "EnableExpressInput"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartitioning", GoGetter: "EnablePartitioning"},
			_jsii_.MemberProperty{JsiiProperty: "enablePartitioningInput", GoGetter: "EnablePartitioningInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxMessageSizeInKilobytes", GoGetter: "MaxMessageSizeInKilobytes"},
			_jsii_.MemberProperty{JsiiProperty: "maxMessageSizeInKilobytesInput", GoGetter: "MaxMessageSizeInKilobytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxSizeInMegabytes", GoGetter: "MaxSizeInMegabytes"},
			_jsii_.MemberProperty{JsiiProperty: "maxSizeInMegabytesInput", GoGetter: "MaxSizeInMegabytesInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceId", GoGetter: "NamespaceId"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceIdInput", GoGetter: "NamespaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requiresDuplicateDetection", GoGetter: "RequiresDuplicateDetection"},
			_jsii_.MemberProperty{JsiiProperty: "requiresDuplicateDetectionInput", GoGetter: "RequiresDuplicateDetectionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDeleteOnIdle", GoMethod: "ResetAutoDeleteOnIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultMessageTtl", GoMethod: "ResetDefaultMessageTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDuplicateDetectionHistoryTimeWindow", GoMethod: "ResetDuplicateDetectionHistoryTimeWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableBatchedOperations", GoMethod: "ResetEnableBatchedOperations"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableExpress", GoMethod: "ResetEnableExpress"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePartitioning", GoMethod: "ResetEnablePartitioning"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxMessageSizeInKilobytes", GoMethod: "ResetMaxMessageSizeInKilobytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxSizeInMegabytes", GoMethod: "ResetMaxSizeInMegabytes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiresDuplicateDetection", GoMethod: "ResetRequiresDuplicateDetection"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportOrdering", GoMethod: "ResetSupportOrdering"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "supportOrdering", GoGetter: "SupportOrdering"},
			_jsii_.MemberProperty{JsiiProperty: "supportOrderingInput", GoGetter: "SupportOrderingInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ServicebusTopic{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopicConfig",
		reflect.TypeOf((*ServicebusTopicConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopicTimeouts",
		reflect.TypeOf((*ServicebusTopicTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.servicebusTopic.ServicebusTopicTimeoutsOutputReference",
		reflect.TypeOf((*ServicebusTopicTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_ServicebusTopicTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
