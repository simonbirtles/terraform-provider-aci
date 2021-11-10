package crud

import (
	"fmt"
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/internal/utils"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/internal/api"
)

func DataRead(d *schema.ResourceData, m interface{}, config *meta.ManagedObjectMeta) error {

	log.Printf("[DEBUG] Reading Data Resource - %s", config.ObjectClassName)

	// Gets the MO attributes from the APIC into a map[string]string
	//ParentDnName := config.ObjectParentFieldName
	//DnPrefix := config["DnPrefix"].(string)
	//ObjectNameAttribute := config.ObjectNameFieldName
	ObjectNameFunc := config.FormatObjectNameF
	object_name := ObjectNameFunc(d, config.ObjectNamePrefix, d.Get(config.ObjectNameFieldName).(string))
	dn := fmt.Sprintf("%s/%s", d.Get(config.ObjectParentFieldName), object_name ) 

	data, err := api.Read(m, dn)
	if err != nil {
		d.SetId("")
		return nil
	}

	attrs, err := utils.ExtractMOAttributes(&data, config.ObjectClassName)
	if err != nil {
		d.SetId("")
		return nil
	}	

	mo_struct := config.ObjectModelF(d)  
	key_map := utils.MapApicKeyToSchemaKey(mo_struct)
	utils.SetResourceDataValues(d, attrs, key_map)

	parent_dn, err := utils.GetMOParentDn(dn)
	if err != nil {
		return err
	}
	d.Set("dn", dn)
	d.Set(config.ObjectParentFieldName, parent_dn)
	d.SetId(dn)
	return nil
}

