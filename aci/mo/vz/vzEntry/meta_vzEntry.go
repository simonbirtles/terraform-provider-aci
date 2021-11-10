package vzEntry
//
// A filter entry is a combination of network traffic classification properties. 
// A filter entry is parented by the filter (vzFilter), which can encapsulate multiple filters.
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

		ObjectClassName:        "vzEntry",
		ObjectNamePrefix:       "e",
		ObjectNameFieldName:	"name",
		ObjectParentFieldName:  "filter_group_id",
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
	return new(vz.Entry)
}

// specific formatting for object name i.e. subnet has [..] brackets around ip for name
func formatObjectDnName(d *schema.ResourceData, ObjectNamePrefix string, ObjectName string) string {
	return fmt.Sprintf("%s-%s", ObjectNamePrefix, ObjectName )
}

func getObjectSchema() map[string]*schema.Schema {

	return map[string]*schema.Schema{

		"filter_group_id": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			Description: 	"ACI Filter Entry - Containing Filter Group (vzFilter) ID (DN)",
			//ValidateFunc: 	validation.StringLenBetween(1, 255),
		},

		"apply_to_frag": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - applyToFrag",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^yes|no$"), "Valid entires are: yes|no"),
			Default:		"no",
		},
		"arp_operation": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - ARP opcodes ",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|req|reply$"), "Valid entires are: unspecified|req|reply"),
			Default:		"unspecified",
		},
		"dest_from_port": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Destination From Port (0-65535) ",
			ValidateFunc: 	ValidateIPPort(),
			Default:		"unspecified",
		},
		"dest_to_port": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Destination To Port (0-65535) ",
			ValidateFunc: 	ValidateIPPort(),
			Default:		"unspecified",
		},
		"desc": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Enty - Description",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"dn": &schema.Schema{
			Type:        	schema.TypeString,
			Computed:    	true,
			Description: 	"ACI Filter Entry - DN",
		},	
		"ethernet_type": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Ether Type",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|ipv4|trill|arp|ipv6|mpls_ucast|mac_security|fcoe|ip$"), "Valid entires are: unspecified|ipv4|trill|arp|ipv6|mpls_ucast|mac_security|fcoe|ip"),
			Default:		"unspecified",
		},
		"icmpv4_type": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - ICMPv4 Type",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^echo-rep|dst-unreach|src-quench|echo|time-exceeded|unspecified$"), "Valid entires are: echo-rep|dst-unreach|src-quench|echo|time-exceeded|unspecified"),
			Default:		"unspecified",
		},
		"icmpv6_type": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - ICMPv6 Type",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|dst-unreach|time-exceeded|echo-req|echo-rep|nbr-solicit|nbr-advert|redirect$"), "Valid entires are: unspecified|dst-unreach|time-exceeded|echo-req|echo-rep|nbr-solicit|nbr-advert|redirect"),
			Default:		"unspecified",
		},
		"match_dscp": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - vzFilter relative name",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified$"), "Valid entried are: CS0|CS1|AF11|AF12|AF13|CS2|AF21|AF22|AF23|CS3|AF31|AF32|AF33|CS4|AF41|AF42|AF43|CS5|VA|EF|CS6|CS7|unspecified"),
			Default:		"unspecified",
		},
		"name": &schema.Schema{
			Type:        	schema.TypeString,
			Required:    	true,
			ForceNew: 		true,
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_.-]+$"), "Valid Length 1-64, Valid characters are: [a-zA-Z0-9_.-]+"),
			Description: 	"ACI Filter Entry - Name",
		},
		"name_alias": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Arbitary key for enabling clients to own their data for entity correlation.",
			ValidateFunc: 	validation.StringLenBetween(0, 128),
		},
		"l3_ip_protocol": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - L3 IP Protocol",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^unspecified|icmp|igmp|tcp|egp|igp|udp|icmpv6|eigrp|ospfigp|pim|l2tp$"), "Valid entires are: unspecified|icmp|igmp|tcp|egp|igp|udp|icmpv6|eigrp|ospfigp|pim|l2tp"),
			Default:		"unspecified",
		},
		"source_from_port": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Source From Port (0-65535) ",
			ValidateFunc: 	ValidateIPPort(),
			Default:		"unspecified",
		},
		"source_to_port": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Source To Port (0-65535) ",
			ValidateFunc: 	ValidateIPPort(),
			Default:		"unspecified",
		},
		"stateful": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - Enabled stateful entries for AVS firewall",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^yes|no$"), "Valid entires are: yes|no"),
			Default:		"no",
		},
		"tcp_session_flags": &schema.Schema{
			Type:        	schema.TypeString,
			Optional:    	true,
			Description: 	"ACI Filter Entry - TCP Session Rules",
			ValidateFunc: 	validation.StringMatch(regexp.MustCompile("^|unspecified|est|syn|ack|fin|rst$"), "Valid entires are: unspecified|est|syn|ack|fin|rst"),
			Default:		"",
		},
	}
}

// Validation function to be moved to seperate package.
//
func ValidateIPPort() schema.SchemaValidateFunc {
 
	return func(i interface{}, k string) ([]string, []error){

		v, ok := i.(string)
		if !ok {
			return nil, []error{fmt.Errorf("expected type of %s to be string", k)}
		}

		// check for integer string of 20,25,53,80,110,443,554 and reject
		r := regexp.MustCompile("^([2][0]|[2][5]|[5][3]|[8][0]|[1][1][0]|[4][4][3]|[5][5][4])$")
		if ok := r.MatchString(v); ok {
			return nil, []error{fmt.Errorf(ip_port_warning, v)}
		}

		r = regexp.MustCompile("^(ftp-data|smtp|dns|http|pop3|https|rstp|unspecified|[0-9]|[1-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-6][0-9][0-9][0-9][0-5])$")
		if ok = r.MatchString(v); !ok {
			return nil, []error{fmt.Errorf(ip_port_warning, v)}
		}

		return nil, nil
	}
}


const ip_port_warning = `

ENTERED VALUE: [%s]

The entered value must be a string value in the range 0-65535 with the 
following exceptions, i.e. "321" or "8080".

EXCEPTIONS:
The APIC automatically converts the following integers to named string literals and will
cause Terraform to update on every run from named literal to the configured integer in 
the TF file. 

20 	ftp-data
25 	smtp
53 	dns
80 	http
110	pop3
443	https
554	rtsp

For these ports, use the string literal and not an integer, for all other ports
use the integer value as a string, i.e. 8080 as "8080"

`