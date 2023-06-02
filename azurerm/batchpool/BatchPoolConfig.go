package batchpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchPoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#account_name BatchPool#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#node_agent_sku_id BatchPool#node_agent_sku_id}.
	NodeAgentSkuId *string `field:"required" json:"nodeAgentSkuId" yaml:"nodeAgentSkuId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#resource_group_name BatchPool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// storage_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#storage_image_reference BatchPool#storage_image_reference}
	StorageImageReference *BatchPoolStorageImageReference `field:"required" json:"storageImageReference" yaml:"storageImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#vm_size BatchPool#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// auto_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#auto_scale BatchPool#auto_scale}
	AutoScale *BatchPoolAutoScale `field:"optional" json:"autoScale" yaml:"autoScale"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#certificate BatchPool#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#container_configuration BatchPool#container_configuration}
	ContainerConfiguration *BatchPoolContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#data_disks BatchPool#data_disks}
	DataDisks interface{} `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#disk_encryption BatchPool#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#display_name BatchPool#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#extensions BatchPool#extensions}
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// fixed_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#fixed_scale BatchPool#fixed_scale}
	FixedScale *BatchPoolFixedScale `field:"optional" json:"fixedScale" yaml:"fixedScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#id BatchPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#identity BatchPool#identity}
	Identity *BatchPoolIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#inter_node_communication BatchPool#inter_node_communication}.
	InterNodeCommunication *string `field:"optional" json:"interNodeCommunication" yaml:"interNodeCommunication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#license_type BatchPool#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#max_tasks_per_node BatchPool#max_tasks_per_node}.
	MaxTasksPerNode *float64 `field:"optional" json:"maxTasksPerNode" yaml:"maxTasksPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#metadata BatchPool#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// mount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#mount BatchPool#mount}
	Mount interface{} `field:"optional" json:"mount" yaml:"mount"`
	// network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#network_configuration BatchPool#network_configuration}
	NetworkConfiguration *BatchPoolNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// node_placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#node_placement BatchPool#node_placement}
	NodePlacement interface{} `field:"optional" json:"nodePlacement" yaml:"nodePlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#os_disk_placement BatchPool#os_disk_placement}.
	OsDiskPlacement *string `field:"optional" json:"osDiskPlacement" yaml:"osDiskPlacement"`
	// start_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#start_task BatchPool#start_task}
	StartTask *BatchPoolStartTask `field:"optional" json:"startTask" yaml:"startTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#stop_pending_resize_operation BatchPool#stop_pending_resize_operation}.
	StopPendingResizeOperation interface{} `field:"optional" json:"stopPendingResizeOperation" yaml:"stopPendingResizeOperation"`
	// task_scheduling_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#task_scheduling_policy BatchPool#task_scheduling_policy}
	TaskSchedulingPolicy interface{} `field:"optional" json:"taskSchedulingPolicy" yaml:"taskSchedulingPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#timeouts BatchPool#timeouts}
	Timeouts *BatchPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#user_accounts BatchPool#user_accounts}
	UserAccounts interface{} `field:"optional" json:"userAccounts" yaml:"userAccounts"`
	// windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/batch_pool#windows BatchPool#windows}
	Windows interface{} `field:"optional" json:"windows" yaml:"windows"`
}

