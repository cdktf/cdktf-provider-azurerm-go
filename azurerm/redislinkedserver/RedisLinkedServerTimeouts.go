package redislinkedserver


type RedisLinkedServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_linked_server#create RedisLinkedServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_linked_server#delete RedisLinkedServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_linked_server#read RedisLinkedServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/redis_linked_server#update RedisLinkedServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
