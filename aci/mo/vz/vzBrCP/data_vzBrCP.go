package vzBrCP

import (
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/mo/crud"
)

func Data() *schema.Resource {
	return &schema.Resource{
		Read: dataRead,
		Schema: getObjectSchema(),
	}
}

func dataRead(d *schema.ResourceData, m interface{}) error {
	return crud.DataRead(d, m, getObjectConfig())
}


