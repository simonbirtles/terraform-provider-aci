package crud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"terraform-provider-aci/aci/internal/utils"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/internal/api"
	"log"
)

func ResourceCreate(d *schema.ResourceData, m interface{}, config *meta.ManagedObjectMeta) error {

	log.Printf("[DEBUG] Creating New Resource - %s", config.ObjectClassName)

	mo_struct := config.ObjectModelF(d)  
	parent_dn := d.Get(config.ObjectParentFieldName).(string)
	ObjectNameFunc := config.FormatObjectNameF
	// maybe do a GetOk and f not found config.ObjectNameFieldName - cause panic with good error message
	object_name := ObjectNameFunc(d, config.ObjectNamePrefix, d.Get(config.ObjectNameFieldName).(string))

	attributes, err := utils.CreateMOAttributesPayload(d, mo_struct, config)
    if err != nil {
        return err
	}
	
	dn, err := api.Create(
		parent_dn, 
		config.ObjectClassName, 
		object_name,
		attributes,
		config.OverWriteExisting,		
		d,
		m)
	
	if err != nil {
		return err
	}

	d.Set("dn", dn)
	d.SetId(dn)
	return ResourceRead(d, m, config)
}

func ResourceRead(d *schema.ResourceData, m interface{}, config *meta.ManagedObjectMeta) error {

	log.Printf("[DEBUG] Reading Resource - %s", config.ObjectClassName)
	// Gets the MO attributes from the APIC into a map[string]string
	dn := d.Id()
	data, err := api.Read(m, dn)
	if err != nil {
		d.SetId("")
		return nil
	}

	// convert the raw APIC response payload to a map[string]interface{} / map[string]string
	// also check for response errors in err
	apic_attrs, err := utils.ExtractMOAttributes(&data, config.ObjectClassName)
	if err != nil {
		d.SetId("")
		return nil
	}	

	mo_struct := config.ObjectModelF(d)  
	key_map := utils.MapApicKeyToSchemaKey(mo_struct)

	// using MO attributes in map[string]interface{} / map[string]string
	// with key in APIC MO attribute name/case, set the schema.ResourceData.attributes.*.value
	utils.SetResourceDataValues(d, apic_attrs, key_map)

	parent_dn, err := utils.GetMOParentDn(dn)
	if err != nil {
		return err
	}

	d.Set(config.ObjectParentFieldName, parent_dn)
	return nil
}

func ResourceUpdate(d *schema.ResourceData, m interface{}, config *meta.ManagedObjectMeta) error {

	log.Printf("[DEBUG] Updating Resource - %s", config.ObjectClassName)
	hasChange := false
	
	// get pointer to new struct representing this managed object
	// we use this for reflection only and to avoid having to maintain seperate
	// structs and maps
	bds := config.ObjectModelF(d) 
	// use the struct to get a list of the keys/fields in the struct
	keys := utils.GetMOSchemaKeys(bds)
	// for each key in the struct, use to check if TF user has changed this key from original set by user
	for _, key := range keys {
		log.Printf("[DEBUG] [UPDATE] Checking key change - %s : %s", config.ObjectClassName, key)
		hasChange = hasChange ||  d.HasChange(key) 
	}

	// if any attr has change, call update with all values if updated or not.
	if hasChange {		
		mo_struct := config.ObjectModelF(d) 
		attributes, err := utils.CreateMOAttributesPayload(d, mo_struct, config)
		if err != nil {
			return err
		}

		parent_dn := d.Get(config.ObjectParentFieldName).(string)	
		err = api.Update(
			parent_dn, 
			config.ObjectClassName,
			attributes,
			d,
			m)

		if err != nil {
			return err
		}
	}
	return ResourceRead(d, m, config)
}

func ResourceDelete(d *schema.ResourceData, m interface{}, config *meta.ManagedObjectMeta) error {

	log.Printf("[DEBUG] Deleting Resource - %s", config.ObjectClassName)

	if !config.CanDelete {
		log.Printf("[INFO] Attempt to delete undeletable resource, skipping - %s", d.Id())
		return nil
	}

	err := api.Delete(m, d.Id())
	if err != nil {
		return err
	}
	return nil
}





