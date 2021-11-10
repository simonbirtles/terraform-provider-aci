package vzConsSubjLbl
//
// The filter for the subject of a service contract. 
// A subject represents a sub-application running behind an endpoint group, such as an exchange server. 
// A subject is parented by the contract. 
//
import ( 
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/utils"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/mo/vz"
	"fmt"
)

func getObjectConfig() *meta.ManagedObjectMeta {

	return &meta.ManagedObjectMeta {

		ObjectClassName:        "vzConsSubjLbl",
		ObjectNamePrefix:       "conssubjlbl",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "parent_id",
		OverWriteExisting:      false,
		IsReln:					true,
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
	return new(vz.ConsSubjLbl)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"parent_id": &schema.Schema{
			Type:        		schema.TypeString,
			Required:    		true,
			ForceNew: 			true,
			Description: 		"ACI Consumer Subject Label - Containing EPG ID (DN)",
			ValidateFunc: 		validation.StringLenBetween(1, 255),
		},
		"desc": &schema.Schema{
			Type:     			schema.TypeString,
			Optional: 			true,
			Description: 		"ACI Consumer Label - Description.",
			ValidateFunc:		validation.StringLenBetween(1, 128),
		},
		"dn": &schema.Schema{
			Type:        		schema.TypeString,
			Computed:    		true,
			Description: 		"ACI Consumer Subject Label - DN",
		},	
		"is_complement": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Consumer Subject Label - EPG CoS Value",
			Default: 			"no",
			ValidateFunc:		utils.ValidateList([]string{"no","yes"}),
		},	
		"name": &schema.Schema{
			Type:     			schema.TypeString,
			Required: 			true,
			ForceNew: 			true,
			Description: 		"ACI Consumer Subject Label - Name",
			ValidateFunc:		validation.StringLenBetween(1, 64),
		},
		"name_alias": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Consumer Subject Label - Name Alias.",
			ValidateFunc: 		validation.StringLenBetween(0, 128),
		},
		"ownerkey": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Consumer Subject Label - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 		validation.StringLenBetween(0, 128),
		},
		"ownertag": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Consumer Subject Label - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
			ValidateFunc: 		validation.StringLenBetween(0, 64),
		},
		"tag": &schema.Schema{
			Type:        		schema.TypeString,
			Optional:    		true,
			Description: 		"ACI Consumer Subject Label  - Color Tag",
		},	
	}
}