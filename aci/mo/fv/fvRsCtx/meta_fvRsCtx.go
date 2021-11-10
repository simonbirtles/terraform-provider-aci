package fvRsCtx

import ( 
	"github.com/hashicorp/terraform/helper/schema"
	//"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/mo/fv"
	"fmt"
)

/*
ClassName is the ACI Managed Object Abstract Class Name
DnPrefix is the ACI Managed Object Relative Name Prefix
ParentDnName is the Terraform Provider Schema Name Used To Identify The Parent Object, internally this is the ACI MO Parent object but we may 
			 call it something different in the schema to make it clear to TF users
ObjectNameAttribute is the Schema Name that is the relative name that ACI uses primarly for the naming of an object relativly.
OverWriteExisting is used when an object exists by default such as a rsctx when a BD is created but we need to manage the object, so 
			 this flag stop the check of existing mo and takes management of it and makes the create function like a update and take management func.
*/

func getObjectConfig() *meta.ManagedObjectMeta {

	return &meta.ManagedObjectMeta {

		ObjectClassName:        "fvRsCtx",
		ObjectNamePrefix:       "rsctx",
		ObjectNameFieldName:	"ctx_name",
		ObjectParentFieldName:  "parent_id",
		OverWriteExisting:      true,
		IsReln:					true,
		IsMandatoryReln:		true,
		CanCreate:				false,
		CanRead	:				true,
		CanUpdate:				true,
		CanDelete:				false,
	
		ObjectModelF:           getObjectModel,
		SchemaF :               getObjectSchema,
		FormatObjectNameF:      formatObjectDnName,		
	}
}

func getObjectModel(d *schema.ResourceData) interface{} {
	return new(fv.RsCtx)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s", ObjectNamePrefix)			// return 'rsctx' only
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"parent_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 	 	true,
			Description: 	"ACI Managed Object Parent ID (DN)",
			//ValidateFunc:	validation.StringLenBetween(1, 255),
		},
		"dn": {
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"The DN (ID) of the ACI EPG Bridge Domain relationship managed object.",
		},
		"ownerkey": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Context - Arbitary key for enabling clients to own their data for entity correlation.",
		},
		"ownertag": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Context - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
		},
		"ctx_name": {
			Type:         	schema.TypeString,
			Required:      	true,
			Description:   	"The Context name is used by this reln MO.",
		},
	}
}