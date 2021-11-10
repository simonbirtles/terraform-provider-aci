# ACI Managed Object Raw Query

These data sources allow raw ACI managed object (mo) and class queries using the APIC query paramaters.

# Usage

### Data Sources

#### Managed Object Query
Queries the ACI managed objects.

`mo_dn` - (Required) The managed object distinguished name (DN) to be queried. This is in the raw format with object type prefix. E.g. uni/tn-TEN_INTERNATIONAL/ap-AP_OTN_ALL

`query_target` - (Optional) This parameter restricts the scope of the query (self, children, subtree)

`target_subtree_class` - (Optional) This parameter specifies which object classes are to be considered when the query-target option is used with scope other than self . You can specify multiple desired object types as a comma-separated list with no spaces.

`query_target_filter` - (Optional) This parameter specifies a logical filter to be applied to the response. This statement can be used by itself or applied after the query-target statement.

`rsp_subtree` - (Optional) For objects returned, this option specifies whether child objects are included in the response. (no, children, full)

`rsp_subtree_class` - (Optional) When child objects are to be returned, this statement specifies that only child objects of the specified object class are included in the response.

`rsp_subtree_filter` - (Optional) When child objects are to be returned, this statement specifies a logical filter to be applied to the child objects.

`rsp_subtree_include` - (Optional) When child objects are to be returned, this statement specifies additional contained objects or options to bincluded in the response. (audit-logs, event-logs, faults, fault-records, health, health-records, relations, stats, tasks. [count, no-scoped, required ]]    )

`rsp_prop_include` - (Optional) This parameter specifies what type of properties should be included in the response when the rsp-subtree option iused with an argument other than no. (all, naming-only, config-only)

`order_by` - (Optional) Sort the query response by one or more properties of a class, you can specify the direction of the order using (asc | desc).

`expected_result_count` - (Optional) If the returned result count does not match this value, the provider will return an error during the plan phase to prevent further processing. Useful if you require a single result or at least one result, or want 0 results to validate the query does not return a result. ( -1 One or more result (default) | 0 Must return zero results | 1 Must return one result )

`result_data` - (Computed) Query results.

`result_count` - (Computed) Query result count.



```
variable "interfaces" {
    default = [
          {
            "interface_id": 0,
            "name": "primary",
            "port_group": "TEN_INTERNATIONAL|AP_OTN_ALL|EPG_OTN_VLAN100"
          },
          {
            "interface_id": 1,
            "name": "backup",
            "port_group": "TEN_INTERNATIONAL|AP_OTN_ALL|EPG_OTN_VLAN101"
          },
          {
            "interface_id": 2,
            "name": "management",
            "port_group": "TEN_INTERNATIONAL|AP_OTN_ALL|EPG_OTN_VLAN102"
          }
    ]
}

# ACI Get EPG Bridge Domain
data "aci_raw_query_mo" "epg_bridge_domain" {
    count                       = length(var.interfaces)
    mo_dn                       = format("uni/tn-%s/ap-%s/epg-%s", split("|", var.interfaces[count.index].port_group)[0] , 
                                                                   split("|", var.interfaces[count.index].port_group)[1], 
                                                                   split("|", var.interfaces[count.index].port_group)[2])  
    query_target         		= "children"
	target_subtree_class 		= "fvRsBd"
	query_target_filter  		= ""
	rsp_subtree          		= ""
	rsp_subtree_class    		= ""
	rsp_subtree_filter   		= ""
	rsp_subtree_include  		= ""
	rsp_prop_include                = ""
	order_by             		= ""
}

# ACI Get EPG Bridge Domain Subnets
data "aci_raw_query_mo" "bridge_domain_ip_gateways" {
    count                       = length(local.all_bds) #length(var.networks)
    mo_dn                       = local.all_bds[count.index]  
    query_target         		= "children"
	target_subtree_class 		= "fvSubnet"
	query_target_filter  		= ""
	rsp_subtree          		= ""
	rsp_subtree_class    		= ""
	rsp_subtree_filter   		= ""
	rsp_subtree_include  		= ""
	rsp_prop_include                = ""
	order_by             		= ""
}


locals {

    #
    # Gather Bridge Domain Info
    #
    json_data_bd = [ 
        for data in data.aci_raw_query_mo.epg_bridge_domain:
        jsondecode(data.result_data)["imdata"]
    ]

    all_bds = [
        for data in local.json_data_bd:
            data[0].fvRsBd.attributes.tDn
    ]


    #
    # Gather Subnet Info
    #
    json_data_subnet = [ 
        for data in data.aci_raw_query_mo.bridge_domain_ip_gateways:
        jsondecode(data.result_data)["imdata"]
    ]

    all_subnets = [
        for data in local.json_data_subnet:
        [
            for subnet in data:
            subnet.fvSubnet.attributes.ip
        ]
    ]

}

output "bd_data" {
    value = local.all_bds
}


output "subnet_data" {
    value = local.all_subnets
}

output "tdn" {
    value = [
    for int in var.interfaces:
    format("uni/tn-%s/ap-%s/epg-%s", split("|", int.port_group)[0] , 
                                     split("|", int.port_group)[1], 
                                     split("|", int.port_group)[2])
    ]
}
```

The output of the above example.
```
Outputs:

bd_data = [
  "uni/tn-TEN_INTERNATIONAL/BD-BD_OTN_VLAN100",
  "uni/tn-TEN_INTERNATIONAL/BD-BD_OTN_VLAN101",
  "uni/tn-TEN_INTERNATIONAL/BD-BD_OTN_VLAN102",
]
subnet_data = [
  [
    "10.243.176.206/28",
  ],
  [
    "10.243.176.214/29",
  ],
  [
    "10.243.176.222/29",
  ],
]
tdn = [
  "uni/tn-TEN_INTERNATIONAL/ap-AP_OTN_ALL/epg-EPG_OTN_VLAN100",
  "uni/tn-TEN_INTERNATIONAL/ap-AP_OTN_ALL/epg-EPG_OTN_VLAN101",
  "uni/tn-TEN_INTERNATIONAL/ap-AP_OTN_ALL/epg-EPG_OTN_VLAN102",
]

```