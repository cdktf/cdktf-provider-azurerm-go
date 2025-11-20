// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancekubernetescluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesCluster",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backupDatasourceParameters", GoGetter: "BackupDatasourceParameters"},
			_jsii_.MemberProperty{JsiiProperty: "backupDatasourceParametersInput", GoGetter: "BackupDatasourceParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "backupPolicyId", GoGetter: "BackupPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "backupPolicyIdInput", GoGetter: "BackupPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesClusterId", GoGetter: "KubernetesClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesClusterIdInput", GoGetter: "KubernetesClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "protectionState", GoGetter: "ProtectionState"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putBackupDatasourceParameters", GoMethod: "PutBackupDatasourceParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackupDatasourceParameters", GoMethod: "ResetBackupDatasourceParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotResourceGroupName", GoGetter: "SnapshotResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotResourceGroupNameInput", GoGetter: "SnapshotResourceGroupNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vaultId", GoGetter: "VaultId"},
			_jsii_.MemberProperty{JsiiProperty: "vaultIdInput", GoGetter: "VaultIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataProtectionBackupInstanceKubernetesCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterScopedResourcesEnabled", GoGetter: "ClusterScopedResourcesEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "clusterScopedResourcesEnabledInput", GoGetter: "ClusterScopedResourcesEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedNamespaces", GoGetter: "ExcludedNamespaces"},
			_jsii_.MemberProperty{JsiiProperty: "excludedNamespacesInput", GoGetter: "ExcludedNamespacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludedResourceTypes", GoGetter: "ExcludedResourceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "excludedResourceTypesInput", GoGetter: "ExcludedResourceTypesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includedNamespaces", GoGetter: "IncludedNamespaces"},
			_jsii_.MemberProperty{JsiiProperty: "includedNamespacesInput", GoGetter: "IncludedNamespacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "includedResourceTypes", GoGetter: "IncludedResourceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "includedResourceTypesInput", GoGetter: "IncludedResourceTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labelSelectors", GoGetter: "LabelSelectors"},
			_jsii_.MemberProperty{JsiiProperty: "labelSelectorsInput", GoGetter: "LabelSelectorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterScopedResourcesEnabled", GoMethod: "ResetClusterScopedResourcesEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedNamespaces", GoMethod: "ResetExcludedNamespaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedResourceTypes", GoMethod: "ResetExcludedResourceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedNamespaces", GoMethod: "ResetIncludedNamespaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludedResourceTypes", GoMethod: "ResetIncludedResourceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabelSelectors", GoMethod: "ResetLabelSelectors"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeSnapshotEnabled", GoMethod: "ResetVolumeSnapshotEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSnapshotEnabled", GoGetter: "VolumeSnapshotEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSnapshotEnabledInput", GoGetter: "VolumeSnapshotEnabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterConfig",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterTimeouts",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesClusterTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterTimeoutsOutputReference",
		reflect.TypeOf((*DataProtectionBackupInstanceKubernetesClusterTimeoutsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataProtectionBackupInstanceKubernetesClusterTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
