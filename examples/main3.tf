#
# Terraform ACI Provider Validation

#
# Providers
#
provider "aci" {
    username = "${var.aci_username}"
    password = "${var.aci_password}"
    apic = "${var.aci_apic}"
    allow_unverified_ssl = "${var.aci_ignore_ssl}"
    sync_delay = 600
}
// TENANT
//////////////////////////////////////////
data "aci_fvTenant" "prod_engineering_tenant" {
    name    = "TEN_PROD_ENGINEERING"
}

// APPLICATION PROFILE
//////////////////////////////////////////
resource "aci_fvAp" "app_prod_engineering" {
    tenant_id               = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name                    = "AP_PROD_ENGINEERING"
    ownertag                = "managed-TERRAFORM"
}
resource "aci_tagInst" "tag_app_prod_engineering_tf" {
    parent_id   = "${aci_fvAp.app_prod_engineering.id}"
    name        = "managed-TERRAFORM"
}
resource "aci_tagInst" "tag_app_prod_engineering_project" {
    parent_id   = "${aci_fvAp.app_prod_engineering.id}"
    name        = "project-PROD_ENGINEERING"
}

// EPG 01
//////////////////////////////////////////
resource "aci_fvAEPg" "epg_01" {
    ap_id                   = "${aci_fvAp.app_prod_engineering.id}"
    name                    = "EPG_PROD_ENGINEERING_01"
    preferred_group_member  = "include"
    intra_epg_isolation     = "enforced"
    forward_control         = ""
}
resource "aci_fvRsBd" "epg_01_bd" {
    parent_id           = "${aci_fvAEPg.epg_01.id}"
    bridge_domain_name  = "${data.aci_fvBD.bd_read.name}"
}
resource "aci_tagInst" "tag_epg_prod_engineering_sh_tf" {
    parent_id   = "${aci_fvAEPg.epg_01.id}"
    name        = "managed-TERRAFORM"
}
resource "aci_tagInst" "tag_epg_prod_engineering_sh_project" {
    parent_id   = "${aci_fvAEPg.epg_01.id}"
    name        = "project-PROD_ENGINEERING"
}

// EPG - 02
//////////////////////////////////////////
resource "aci_fvAEPg" "epg_02" {
    ap_id                   = "${aci_fvAp.app_prod_engineering.id}"
    name                    = "EPG_PROD_ENGINEERING_02"
    preferred_group_member  = "include"
    intra_epg_isolation     = "enforced"
    forward_control         = ""
}
resource "aci_fvRsBd" "epg_02_bd" {
    parent_id           = "${aci_fvAEPg.epg_02.id}"
    bridge_domain_name  = "${data.aci_fvBD.bd_read.name}"
}
resource "aci_tagInst" "tag_epg_prod_engineering_02_tf" {
    parent_id   = "${aci_fvAEPg.epg_02.id}"
    name        = "managed-TERRAFORM"
}
resource "aci_tagInst" "tag_epg_prod_engineering_02_project" {
    parent_id   = "${aci_fvAEPg.epg_02.id}"
    name        = "project-PROD_ENGINEERING"
}


// BRIDGE DOMAINS
//////////////////////////////////////////
variable "bridge_domain" {
    default = "BD_ENGINEERING_IC"
}

data "aci_fvBD" "bd_read" {
    tenant_id   = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name        = "BD_PROD_ENGINEERING"
}

resource "aci_fvBD" "bd_c" {
    tenant_id                   = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    #tenant_id                  = "uni/tn-TEN_PROD_ENGINEERING"
    name                        = "${var.bridge_domain}"
    desc                        = "Test BD For Terraform"
    ip_learning                 = "yes"
    //arp_flooding                = "yes"
    limit_ip_learn_to_subnet    = "no"
    bridge_domain_mac           = "00:22:BD:F8:19:FC"
}

resource "aci_fvRsCtx" "rsctx_bdc_prod_engineering2" {
    parent_id           = "${aci_fvBD.bd_c.id}"                   # reln_from
    ctx_name            = "${aci_fvCtx.vrf_prod_engineering2.name}"         # reln_to
}

// SUBNETS 
//////////////////////////////////////////
resource "aci_fvSubnet" "subnet_prod_engineering_bd" {
    parent_id   = "${aci_fvBD.bd_c.id}"
    ip          = "192.168.34.254/24"
    scope       = "private,shared"
    ctrl        = "nd"
    virtual     = "no"
}

resource "aci_tagInst" "tag_tf" {
    parent_id   = "${aci_fvBD.bd_c.id}"
    name        = "managed-TERRAFORM"
}

resource "aci_tagInst" "tag_project" {
    parent_id   = "${aci_fvBD.bd_c.id}"
    name        = "project-PROD_ENGINEERING"
}

// VRF / CONTEXT
//////////////////////////////////////////
resource "aci_fvCtx" "vrf_prod_engineering2" {
    tenant_id                   = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name                        = "VRF_ACI_TF"
    ownertag                    = "Terraform"
    policy_control_direction    = "egress"
}

resource "aci_tagInst" "tag_tf_vrf_prod_engineering2" {
    parent_id   = "${aci_fvCtx.vrf_prod_engineering2.id}"
    name        = "managed-TERRAFORM"
}

data "aci_fvCtx" "soc_vrf" {
    tenant_id           = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name                = "VRF_PROD_ENGINEERING"
}

// PHYS DOM PROFILE
//////////////////////////////////////////
data "aci_phys_dom" "phys_dom_profile" {
    name = "PHYSDOM_DC_EAST_SERVERS"
}

// VMM DOM PROFILE
//////////////////////////////////////////
data "aci_vmm_dom" "vmm_dom_profile" {
    name = "VMM_VMW_DVS_01"
}

// CONTRACT
//////////////////////////////////////////
resource "aci_vzBrCP" "contract_a" {
    tenant_id       = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name            = "CNT_IC_01"
    desc            = "Engineering IC"
    ownerkey        = "Terraform"
    contract_scope  = "tenant"
}
// CONTRACT SUBJECT
//////////////////////////////////////////
resource "aci_vzSubj" "contract_a_subj_one" {
    contract_id             = "${aci_vzBrCP.contract_a.id}"
    name                    = "SUBJ_CNT_B"
    desc                    = "TF Subject B"
    priority_level          = "level2"
    reverse_filter_ports    = "yes"
    consumer_subject_match  = "AtmostOne"
    provider_subject_match  = "AtmostOne"
}

// CONTRACT SUBJECT FILTER RELN
/////////////////////////////////////////
resource "aci_vzRsSubjFiltAtt" "contract_a_filter_http" {
    contract_subject_id = "${aci_vzSubj.contract_a_subj_one.id}"
    filter_name      = "FILTER_HTTP"
}
resource "aci_vzRsSubjFiltAtt" "contract_a_filter_ssh" {
    contract_subject_id = "${aci_vzSubj.contract_a_subj_one.id}"
    filter_name      = "FILTER_SSH"
}

// FILTER GROUP vzFilter
/////////////////////////////////////////
resource "aci_vzFilter" "filter_group_web" {
    tenant_id   = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name        = "FILTER_HTTP"
    name_alias  = "alias_name_http"
}

resource "aci_vzFilter" "filter_group_mgmt" {
    tenant_id   = "${data.aci_fvTenant.prod_engineering_tenant.id}"
    name        = "FILTER_SSH"
    name_alias  = "alias_name_ssh"
}

// FILTER ENTRIES
/////////////////////////////////////////
resource "aci_vzEntry" "filter_entry_one" {
    filter_group_id     = "${aci_vzFilter.filter_group_web.id}"
    name                = "HTTPS"
    desc                = "Generic HTTPS Filter"
    ethernet_type       = "ipv4"
    l3_ip_protocol      = "tcp"
    dest_from_port      = "http"
    dest_to_port        = "http"
    tcp_session_flags   = ""
   // stateful            = "yes"
}

resource "aci_vzEntry" "filter_entry_two" {
    filter_group_id     = "${aci_vzFilter.filter_group_mgmt.id}"
    name                = "SSH"
    desc                = "Generic SSH Filter"
    ethernet_type       = "ipv4"
    l3_ip_protocol      = "tcp"
    dest_from_port      = "22"
    dest_to_port        = "22"
    tcp_session_flags   = ""
   // stateful            = "yes"
}

// EPG CONTRACT RS - CONSUMER [EPG]
/////////////////////////////////////////
resource "aci_fvRsCons" "epg_consumed_contract" {
    epg_id              = "${aci_fvAEPg.epg_01.id}"
    contract_name       = "${aci_vzBrCP.contract_a.name}"
}

// EPG CONTRACT RS - PROVIDER [EPG]
/////////////////////////////////////////
resource "aci_fvRsProv" "epg_provided_contract" {
    epg_id              = "${aci_fvAEPg.epg_02.id}"
    contract_name       = "${aci_vzBrCP.contract_a.name}"
}



// EPG PROVIDER LABEL - [EPG]
/////////////////////////////////////////
resource "aci_vzProvLbl" "epg_contract_provider_label" {
    parent_id       = "${aci_fvAEPg.epg_02.id}"
    name            = "PROVLBL_EPG_02_RED"
    tag             = "red"
}
// EPG CONSUMER LABEL - [EPG]
/////////////////////////////////////////
resource "aci_vzConsLbl" "epg_contract_consumer_label" {
    parent_id       = "${aci_fvAEPg.epg_01.id}"
    name            = "PROVLBL_EPG_02_RED"
    tag             = "red"
}
/*
// EPG PROVIDER SUBJECT LABEL - [EPG]
/////////////////////////////////////////
resource "aci_vzProvSubjLbl" "epg_contract_provider_subject_label" {
    parent_id       = "${aci_fvAEPg.epg_01.id}"
    name            = "PROVLBLSUBJ_EPG"
    tag             = "red"
}
// EPG PROVIDER SUBJECT LABEL - [SUBJ]
/////////////////////////////////////////
resource "aci_vzProvSubjLbl" "epg_contract_provider_subject_label_subj" {
    parent_id       = "${aci_vzSubj.contract_a_subj_one.id}"
    name            = "PROVLBLSUBJ_EPG"
    tag             = "red"
}
*/



/*
// EPG CONSUMER SUBJECT LABEL - [EPG]
/////////////////////////////////////////
resource "aci_vzConsSubjLbl" "epg_contract_consumer_subject_label" {
    parent_id       = "${aci_fvAEPg.epg_01.id}"
    name            = "CONSLBLSUBJ_EPG"
    tag             = "blue"

}
*/
// EPG VMM VMware DOMAIN RS
/////////////////////////////////////////
resource "aci_fvRsDomAtt" "epg_rs_vmware_dom" {
    epg_id              = "${aci_fvAEPg.epg_01.id}"
    domain_profile_id   = "${data.aci_vmm_dom.vmm_dom_profile.dn}"   
}

// EPG VMM PHYSICAL DOMAIN RS
/////////////////////////////////////////
resource "aci_fvRsDomAtt" "epg_rs_phys_dom" {
    epg_id              = "${aci_fvAEPg.epg_01.id}"
    domain_profile_id   = "${data.aci_phys_dom.phys_dom_profile.dn}"
}




// OUTPUTS
//////////////////////////////////////////
output "prod_engineering_tenant" {
    value = "${data.aci_fvTenant.prod_engineering_tenant.id}"
}
output "prod_engineering_bd" {
    value = "${data.aci_fvBD.bd_read.id}"
}
output "physical_domain_dn" {
    value = "${data.aci_phys_dom.phys_dom_profile.dn}"
}
output "vmm_domain_dn" {
    value = "${data.aci_vmm_dom.vmm_dom_profile.dn}"
}
output "contract_a_dn" {
    value = "${aci_vzBrCP.contract_a.dn}"
}
output "port_group_seperator" {
    value = "${aci_fvRsDomAtt.epg_rs_vmware_dom.delimiter}"
}
output "vmware_port_group" {
    value = "${local.vmware_portgroup_name_sh}"
}

// LOCALS
//////////////////////////////////////////
locals {
    vmware_portgroup_name_sh = "${format("%s%s%s%s%s", data.aci_fvTenant.prod_engineering_tenant.name,    
                                                        aci_fvRsDomAtt.epg_rs_vmware_dom.delimiter , 
                                                        aci_fvAp.app_prod_engineering.name, 
                                                        aci_fvRsDomAtt.epg_rs_vmware_dom.delimiter, 
                                                        aci_fvAEPg.epg_01.name )}"
}



