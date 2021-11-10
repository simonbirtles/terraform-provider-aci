Package Decoder
====
There are several abbreviations used in the names of classes in the ACI object model. 

Listed below are some descriptions of commonly used abbreviations.

aaa: authentication, authorization, accounting
ac: atomic counters
actrl: access control
actrlcap: access control capability
adcom: appliance director communication
aib: adjacency information base
arp: address resolution protocol
bgp: border gateway protocol
callhome: Cisco smart call home services
cap: capability
cdp: Cisco discovery protocol
cnw: node cluster
comm: communication policy
comp: compute
compat: compatibility
condition: health policy
config: configuration policy
coop: Council of Oracles protocol
copp: control plane policing policy: contains set of rules describing policer rates
ctrlr: controller
ctx: context
datetime: date/time policy
dbg: debug
dbgac: debug atomic counters
dbgexp: debug export policy
dhcp: dynamic host configuration protocol
dhcptlv: dynamic host configuration protocol type length value
dhcptlvpol: dynamic host configuration protocol type length value policy
dns: domain name service
draw: graph visualization for GUI
epm: endpoint manager
eqpt: equipment
eqptcap: equipment capability
eqptcapacity: equipment capacity
eqptdiag: equipment diagnostics
eqptdiagp: equipment diagnostics policy
ethpm: ethernet policy manager
event: event policy
extnw: external network
fabric: fabric
fault: fault policy, counters
file: file path, config import/export policy
firmware: firmware
fmcast: fabric multicast
fsm: finite state machine
fv: fabric virtualization
fvns: fabric virtualization namespace
fvtopo: fabric virtualization topology
geo: geolocation
glean: glean adjacency
ha: high availability
health: health score
hvs: hypervisors virtual switch
icmp: internet control protocol
icmpv4: internet control protocol version 4
icmpv6: internet control protocol version 6
ident: identity
igmp: internet group management protocol
igmpsnoop: internet group management protocol snooping
im: interface manager module
imginstall: image install
infra: infrastructure
ip: internet protocol
ipv4: internet protocol version 4
ipv6: internet protocol version 6
isis: intermediate system to intermediate system
isistlv: intermediate system to intermediate system type length value
l1: layer 1
l1cap: layer 1 capability
l2: layer 2
l2cap: layer 2 capability
l2ext: layer 2 external
l3: layer 3
l3cap: layer 3 capability
l3ext: layer 3 external
l3vm: Layer 3 Virtual Machine
lacp: link aggregation protocol
lbp: load balancing policy
leqpt: loose equipment (unmanaged nodes, not in the fabric)
lldp: link layer discovery protocol
lldptlv: link layer discovery protocol type length value
lldptlvpol: link layer discovery protocol type length value policy
maint: maintenance
mcast: multicast
mcp: master control processor
memory: memory statistics
mgmt: management
mo: managed object
mock: mock **(objects used on the simulator mostly for showing stats/faults/etc)**
mon: monitoring
monitor: monitor (SPAN)
naming: abstract for objects with names
nd: neighbor discovery
nw: network
oam: ethernet operations, administrations and management
observer: observer for statistics, fault, state, health, logs/history
opflex: OpFlex
os: operating system
ospf: open shortest path first
pc: port channel
pcons: **generated and used by internal processes**
phys: physical domain profile
ping: ping execution and results
pki: public key infrastructure
pol: policy definition
policer: traffic policing (rate limiting)
pool: object pool
pres: **generated and used by internal processes**
proc: system load, cpu, and memory utilization statistics
psu: power supply unit policy
qos: quality of service policy
qosm: qos statistics
qosp: qos/ 802.1p
rbqm: debugging
regress: regression
reln: **generated and used by internal processes**
repl: **generated and used by internal processes**
res: **generated and used by internal processes**
rib: routing information base
rmon: remote network monitoring/ interface stats/counters
rpm: route policy map
rtcom: route control community list
rtctrl: route control
rtextcom: router extended community
rtflt: route filter
rtleak: route leak
rtmap: RPM route map
rtpfx: route prefix list
rtregcom: route regular community list
rtsum: route summarization address/policy
satm: satellite manager
snmp: simple network management protocol
span: switched port analyzer
stats: statistics collection policies
statstore: statistics data holders
stormctrl: storm control (traffic suppression) policy
stp: spanning tree protocol definitions and policy
sts: Service Tag Switching (used for services insertion)
svccore: core policy
svi: switched virtual interface/ routed VLAN interface
synthetic: synthetic objects (for testing)
sysdebug: system debug
sysfile: system files
syshist: system cards reset records/history
syslog: syslog policy
sysmgr: system manager (firmware, supervisor, system states, etc)
sysmgrp: container for cores policy & abstract class for all qos policy definitions
tag: alias (use descriptive name for dn), tags (group multiple objects by a descriptive name)
task: task execution, instance, and result
test: abstract class for test rule, subject, and result
testinfralab: test infrastructure
tlv: type, length, value system structures
top: system task manager for processor activity
topoctrl: topology control policy (sharding, fabric LB, fabric VxLan, etc)
traceroute: traceroute execution and results
traceroutep: traceroute end points
trig: triggering policy
tunnel: tunneling
uribv4: ipv4 unicast routing information base entity
vlan: vlan instances
vlanmgr: vlan manager control plane
vmm: virtual machine manager (controller, vmm policy and definitions)
vns: virtual network service (L4-L7 policy and definitions)
vpc: virtual port channel (vpc policy and definitions)
vsvc: service labels (provider/consumer)
vtap: translated address of external node (NATed IP of service node)
vxlan: Virtually extensible LAN definitions
vz: virtual zones (former name of the policy controls) i.e. Contracts

Model Naming Schemes
====
Rs: Relationship source
Rt: Relationship target
Ag: Aggregated stats
BrCP: Binary Contract Profile