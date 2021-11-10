package vzFilter
//
// A filter policy is a GROUP of resolvable filter entries. 
// Each filter entry is a combination of network traffic classification properties.
//
import ( 
	"regexp"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"terraform-provider-aci/aci/internal/meta"
	"terraform-provider-aci/aci/mo/vz"
	"fmt"
)

func getObjectConfig() *meta.ManagedObjectMeta {

	return &meta.ManagedObjectMeta {

		ObjectClassName:        "vzFilter",
		ObjectNamePrefix:       "flt",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "tenant_id",
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
	return new(vz.Filter)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"tenant_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Filter - Containing Tenant ID (DN)",
			//ValidateFunc: 	validation.StringLenBetween(1, 255),
		},
		"desc": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter - Description",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Filter - DN",
		},			
		"name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid Length 1-64, Valid characters are: [a-zA-Z0-9_.-]+"),
			Description: 	"ACI Filter - Name",
		},
		"name_alias": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
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