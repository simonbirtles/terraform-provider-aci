package vzRsSubjFiltAtt
//
// The filter for the subject of a service contract. 
// A subject represents a sub-application running behind an endpoint group, such as an exchange server. 
// A subject is parented by the contract. 
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

		ObjectClassName:        "vzRsSubjFiltAtt",
		ObjectNamePrefix:       "rssubjFiltAtt",
		ObjectNameFieldName:	"filter_name",
		ObjectParentFieldName:  "contract_subject_id",
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
	return new(vz.RsSubjFiltAtt)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"contract_subject_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Contract Subject Filter Reln - Containing Tenant ID (DN)",
			//ValidateFunc: 	validation.StringLenBetween(1, 255),
		},
		"directives": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject Filter Reln - Enable Logging",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^none|log$"), "Valid entires are: ''|log"),
			Default:		"",
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Contract Subject Filter Reln - DN",
		},	
		"filter_name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Contract Subject Filter Reln - vzFilter relative name",
			ValidateFunc: 	validation.StringLenBetween(1, 64),
		},
	}
}