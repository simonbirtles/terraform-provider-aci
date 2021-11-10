package physDomP

import ( 
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/meta"
	"regexp"
	"terraform-provider-aci/aci/mo/phys"
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

		ObjectClassName:        "physDomP",
		ObjectNamePrefix:       "phys",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "parent_id",
		OverWriteExisting:      false,
		IsReln:					false,
		IsMandatoryReln:		false,
		CanCreate:				true,
		CanRead:				true,
		CanUpdate:				true,
		CanDelete:				true,
	
		ObjectModelF:           getObjectModel,
		SchemaF:                getObjectSchema,
		FormatObjectNameF:      formatObjectDnName,		
	}
}

func getObjectModel(d *schema.ResourceData) interface{} {
	return new(phys.DomP)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName)
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"parent_id": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Default:	 	"uni",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^uni$"), "Must be set to 'uni'"),
			Description:	"Policy Universe - Leave as default, do not set or must be 'uni'",
		},
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "ACI Physical Domain Name",
		},
		"dn": &schema.Schema{
			Type:     	schema.TypeString,
			Computed:    true,
			Description: "ACI Physical Domain DN.",
		},
		"ownerkey": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"ownertag": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
			ValidateFunc: 	validation.StringLenBetween(0, 64),
		},
	}
}