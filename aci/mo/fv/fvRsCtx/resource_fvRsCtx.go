package fvRsCtx

import (
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/mo/crud"
)

func Resource() *schema.Resource {
	return &schema.Resource{
			Create: ResourceCreate,
			Read:   ResourceRead,
			Update: ResourceUpdate,    
			Delete: ResourceDelete,
			Schema: getObjectSchema(),
	}
}

func ResourceCreate(d *schema.ResourceData, m interface{}) error {
	return crud.ResourceCreate(d, m, getObjectConfig())	
}

func ResourceRead(d *schema.ResourceData, m interface{}) error {
	return crud.ResourceRead(d, m, getObjectConfig() )
}

func ResourceUpdate(d *schema.ResourceData, m interface{}) error {
	return crud.ResourceUpdate(d, m, getObjectConfig() )
}

func ResourceDelete(d *schema.ResourceData, m interface{}) error {
	return crud.ResourceDelete(d, m, getObjectConfig() )
}



