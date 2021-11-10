package query

import (
	//"errors"
	"fmt"
	"log"
	"encoding/json"
	"github.com/hashicorp/terraform/helper/schema"
	"aci"
	"time"
	"strconv"
)

func Data() *schema.Resource {
	return &schema.Resource{

		Read: data_mo_query,

		Schema: map[string]*schema.Schema{
			"mo_dn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The managed object to be queried.",
			},
			"query_target": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "This parameter restricts the scope of the query (self, children, subtree)",
			},
			"target_subtree_class": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "This parameter specifies which object classes are to be considered when the query-target option is used with a scope other than self . You can specify multiple desired object types as a comma-separated list with no spaces.",
			},
			"query_target_filter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "This parameter specifies a logical filter to be applied to the response. This statement can be used by itself or applied after the query-target statement.",
			},
			"rsp_subtree": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "For objects returned, this option specifies whether child objects are included in the response. (no, children, full)",
			},
			"rsp_subtree_class": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "When child objects are to be returned, this statement specifies that only child objects of the specified object class are included in the respons",
			},
			"rsp_subtree_filter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "When child objects are to be returned, this statement specifies a logical filter to be applied to the child objects.",
			},
			"rsp_subtree_include": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "When child objects are to be returned, this statement specifies additional contained objects or options to be included in the response. (audit-logs, event-logs, faults, fault-records, health, health-records, relations, stats, tasks. [count, no-scoped, required ]]    )",
			},
			"rsp_prop_include": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "This parameter specifies what type of properties should be included in the response when the rsp-subtree option is used with an argument other than no. (all, naming-only, config-only)",
			},
			"order_by": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:	 "",
				Description: "sort the query response by one or more properties of a class, and you can specify the direction of the order using (asc | desc).",
			},
			"expected_result_count": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Default:	-1,
				Description: "Required Query result count. ( -1 One or more result (default) | 0 Must return zero results | 1 Must return one result )",
			},
			"result_data": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Query results.",
			},
			"result_count": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Query result count",
			},
		},
	}
}


func data_mo_query(d *schema.ResourceData, m interface{}) error {

	log.Printf("[DEBUG] data_mo_query")

	apicCookie, ok := m.(map[string]interface{})["APIC-cookie"].(string) 
	if !ok{
		return fmt.Errorf("Unable to extract apic-cookie from ResourceData")
	}

	apics, ok := m.(map[string]interface{})["APIC"].([]string)
	if !ok{
		return fmt.Errorf("Unable to extract apic list from ResourceData")
	}
	
	path := fmt.Sprintf("mo/%s", d.Get("mo_dn").(string))

	var info = new(aci.ApicGetInfo)
	info.ApicClient.Cookie = apicCookie
	info.ApicClient.ApicHosts = apics
	info.Path = path
	info.Delay = 0

	// no need to check for user entries as we default to "" in the schema and string len check in the 
	// api to exclude param if ""	
	info.Filter.Query_target = d.Get("query_target").(string)
	info.Filter.Target_subtree_class = d.Get("target_subtree_class").(string)
	info.Filter.Query_target_filter = d.Get("query_target_filter").(string)
	info.Filter.Rsp_subtree = d.Get("rsp_subtree").(string)
	info.Filter.Rsp_subtree_class = d.Get("rsp_subtree_class").(string)
	info.Filter.Rsp_subtree_filter = d.Get("rsp_subtree_filter").(string)
	info.Filter.Rsp_subtree_include = d.Get("rsp_subtree_include").(string)
	info.Filter.Rsp_prop_include = d.Get("rsp_prop_include").(string)
	info.Filter.Order_by = d.Get("order_by").(string)
	
	log.Printf("[DEBUG] data_mo_query: %s", info.Filter)

	// run query
	response_byte_data, err := aci.Get(info)
	if err != nil {
		err_msg := fmt.Sprintf("aci data_mo_query Failed - %v\nRESPONSE: %v", err, response_byte_data)
		log.Printf("[DEBUG] %s", err_msg)
		return fmt.Errorf(err_msg)
	}

	log.Printf("[DEBUG] Response Payload: [%s]", response_byte_data)

	// from raw APIC response []byte string to json 
	// so json output has attributes names/case/etc as APIC definition
	var json_data interface{}
	err = json.Unmarshal(response_byte_data, &json_data)
	if err != nil {
		return err
	}

	record_count, ok := (json_data).(map[string]interface{})["totalCount"]
	if !ok {
		return  fmt.Errorf("APIC Context Request failed with with malformed response payload") 
	}

	// ( -1 Ignore result count (default) | 0 Must return zero results | 'X' A number >=1 )
	expected_result_count := d.Get("expected_result_count").(int)
	count, err := strconv.Atoi(record_count.(string))
	log.Printf("[DEBUG] data_mo_query: Expected Results: [%d], Actual Results: [%d]", expected_result_count, count)

	d.Set("result_count", record_count)
	d.Set("result_data",  string(response_byte_data)  )
	d.SetId(time.Now().String())

	if(expected_result_count == -1 && count >= 1 )	{
		// expected one or more results and got one or more 	
		return nil
	}

	if(expected_result_count == 0 && count == 0){
		// we expected and got zero results
		return nil
	}

	if(expected_result_count == 1 && count == 1){
		// we expected 1 result and got the same
		return nil
	}

	// if we got to here then none of the above tests passed and we have an unexpected/unwatned result count so 
	// set ID to nil to cause terraform to error out
	d.SetId("")
	return fmt.Errorf("ACI MO Query Error: Expected [%d] and actual [%d] result counts not as planned. Check query or expected result count parameter.", expected_result_count, count)
}
