package vzSubj
//
// A subject is a sub-application running behind an endpoint group (for example, an Exchange server). 
// A subject is parented by the contract, which can encapsulate multiple subjects. 
// An endpoint group associated to a contract is providing one or more subjects or is communicating with the subject as a peer entity. 
// An endpoint group always associates with a subject and defines rules under the association for consuming/providing/peer-to-peer communications to that subject. 
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

		ObjectClassName:        "vzSubj",
		ObjectNamePrefix:       "subj",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "contract_id",
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
	return new(vz.Subj)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		// parent DN/ID
		"contract_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Contract Subject - Parent Contract ID (DN)",
			//ValidateFunc: 	validation.StringLenBetween(1, 255),
		},
		// Object Specific Fields		
		"consumer_subject_match": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - The subject match criteria across consumers.  ",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^All|AtleastOne|AtmostOne|None$"), "Valid entires are: All|AtleastOne|AtmostOne|None"),
			Default:		"AtleastOne",
		},
		"desc": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - Description",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Contract Subject - DN",
		},	
		"name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid Length 1-64, Valid characters are: [a-zA-Z0-9_.-]+"),
			Description: 	"ACI Contract Subject - Name",
		},
		"name_alias": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"priority_level": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - QoS Level.",
			Default:		"unspecified",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|level3|level2|level1$"), "Valid entried are: application-profile|tenant|context|global"),
		},
		"provider_subject_match": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - The subject match criteria across providers.",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^All|AtleastOne|AtmostOne|None$"), "Valid entires are: All|AtleastOne|AtmostOne|None"),
			Default:		"AtleastOne",
		},
		"reverse_filter_ports": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile ",
			Default:		"yes",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^yes|no$"), "Valid entried are: yes|no"),
		},
		"target_dscp": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract Subject - TThe target differentiated services code point (DSCP). CS0, CS1, AFXY, etc",
			Default:		"unspecified",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified$"), "Valid entried are: CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified"),
		},	
	}
}