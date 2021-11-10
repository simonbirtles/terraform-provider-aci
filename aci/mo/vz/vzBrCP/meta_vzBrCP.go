package vzBrCP
//
// A contract is a logical container for the subjects which relate to the filters that govern the rules for communication between endpoint groups (EPGs).
// Without a contract, the default forwarding policy is to not allow any communication between EPGs but all communication within an EPG is allowed. 
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

		ObjectClassName:        "vzBrCP",
		ObjectNamePrefix:       "brc",
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
	return new(vz.BrCP)
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
			Description: 	"ACI Contract - Containing Tenant ID (DN)",
			//ValidateFunc: 	validation.StringLenBetween(1, 255),
		},

		"name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid Length 1-64, Valid characters are: [a-zA-Z0-9_.-]+"),
			Description: 	"ACI Contract - Name",
		},
		"desc": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - Description",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Contract - DN",
		},	

		"name_alias": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - Name Alias ",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid Length 1-64, Valid characters are: [a-zA-Z0-9_.-]+"),
		},
		"ownerkey": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"ownertag": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - A tag for enabling clients to add their own data. For example, to indicate who created this object.",
			ValidateFunc: 	validation.StringLenBetween(0, 64),
		},
		"priority_level": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - QoS Level.",
			Default:		"unspecified",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|level3|level2|level1$"), "Valid entried are: application-profile|tenant|context|global"),
		},
		"contract_scope": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile ",
			Default:		"context",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^application-profile|tenant|context|global$"), "Valid entried are: application-profile|tenant|context|global"),
		},
		"target_dscp": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Contract - TThe target differentiated services code point (DSCP). CS0, CS1, AFXY, etc",
			Default:		"unspecified",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified$"), "Valid entried are: CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified"),
		},	
	}
}