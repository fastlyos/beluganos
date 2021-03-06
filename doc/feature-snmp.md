# [Feature guide] SNMP

SNMP (Simple Network Management Protocol) is a management protocol of network elements. Beluganos supports SNMP MIB to check interface status or traffic counter.

## Pre-requirements

- The installation is required in advance. Please refer [install.md](install.md) before proceeding.
- The setup of Beluganos is required in advance. Please refer [setup.md](setup.md) before proceeding.
- In case you will use OpenNSL or OF-DPA:
	- The installation of OpenNetworkLinux for your white-box switches is required in advance. Please refer [setup-hardware.md](setup-hardware.md) before proceeding.
	- The setup for ASIC API is required in advance. Please refer [setup-onsl.md](setup-onsl.md) or [setup-ofdpa.md](setup-ofdpa.md).

## Setup

### Install SNMP feature

You don't required additional operation to install SNMP feature, because SNMP features are already installed in [install.md](install.md).

### Start SNMP process

To use SNMP feature, you should start 4 process.

```
$ sudo systemctl start snmpd
$ sudo systemctl start fibsd
$ sudo systemctl start snmpproxyd-trap
$ sudo systemctl start snmpproxyd-mib
```
### Stop SNMP process

```
$ sudo systemctl stop snmpproxyd-mib
$ sudo systemctl stop snmpproxyd-trap
$ sudo systemctl stop fibsd
$ sudo systemctl stop snmpd
```

## Overviews

### SNMP MIB

You can get some stats from Beluganos by SNMP request.

```
$ snmpget -v <version> -c <community-name> <server-address> <oid>
$ snmpwalk -v <version> -c <community-name> <server-address> <oid>
```

The supported version is `2c`, and default settings of community is `public`.

For example, by issuing following commands from Beluganos itself, ifOperStatus will be returned.

```
# example for getting ifOperStatus (interface up or down)
$ snmpwalk -v 2c -c public localhost .1.3.6.1.2.1.2.2.1.8
```

### SNMP trap

When `ifOperStatus` changed, you can get trap notification by Beluganos.  You may configure IP address of trap receiver by editing the files.

```
$ vi /etc/beluganos/snmpproxyd.yaml

snmpproxy:
  default:
  ~~ (snipped) ~~
    trap2map:
      eth1: 1
      eth2: 2
      eth3: 3
      ~~ (snipped) ~~
    trap2sink:
      - addr: 192.168.122.1:161
      - addr: 192.168.122.2:161
```

- `trap2map`: The mapping list between interface name of Linux container (eth1, eth2, ...) and logical interface number recognized by ASIC driver (1, 2, ...).
	- Format: `<LXC-interface-name>: <ASIC-logical-interface-number>`
	- This settings are depend on your hardware. The physical interface number should be changed. Sample configuration is described at the bottom of this document.
- `trap2sink`: The lists of IP address of trap servers.
	- Format: `addr: <Trap-server-address>:<Trap-port>`
	- You can set one or more SNMP trap servers.

After changing, please restart trap process to reflect changes.

```
$ sudo systemctl restart snmpproxyd-trap
```

## Feature Details

### Supported statistics by SNMP MIB

Internal OID is used only Beluganos. In general, only standard OID should be used.

|  Standard OID           |  Internal OID               |  MIB name       | OpenFlow(PortStats) |
|-------------------------|-----------------------------|-----------------|---------------------|
| .1.3.6.1.2.1.2.2.1.8    | .1.3.6.1.4.99999.2.2.1.8    | ifOperStatus    | (No)                |
| .1.3.6.1.2.1.2.2.1.10   | .1.3.6.1.4.99999.2.2.1.10   | ifInOctets      | rx\_bytes           |
| .1.3.6.1.2.1.2.2.1.11   | .1.3.6.1.4.99999.2.2.1.11   | ifInUcastPkts   | rx\_packets         |
| .1.3.6.1.2.1.2.2.1.12   | .1.3.6.1.4.99999.2.2.1.12   | ifInNUcastPkts  | (No)                |
| .1.3.6.1.2.1.2.2.1.13   | .1.3.6.1.4.99999.2.2.1.13   | ifInDiscards    | rx\_dropped         |
| .1.3.6.1.2.1.2.2.1.14   | .1.3.6.1.4.99999.2.2.1.14   | ifInErrors      | rx\_errors          |
| .1.3.6.1.2.1.2.2.1.16   | .1.3.6.1.4.99999.2.2.1.16   | ifOutOctets     | tx\_bytes           |
| .1.3.6.1.2.1.2.2.1.17   | .1.3.6.1.4.99999.2.2.1.17   | ifOutUcastPkts  | tx\_packets         |
| .1.3.6.1.2.1.2.2.1.18   | .1.3.6.1.4.99999.2.2.1.18   | ifOutNUcastPkts | (No)                |
| .1.3.6.1.2.1.2.2.1.19   | .1.3.6.1.4.99999.2.2.1.19   | ifOutDiscards   | tx\_dropped         |
| .1.3.6.1.2.1.2.2.1.20   | .1.3.6.1.4.99999.2.2.1.20   | ifOutErrors     | tx\_errors          |
| .1.3.6.1.2.1.31.1.1.1.1 | .1.3.6.1.4.99999.31.1.1.1.1 | ifName          | ifname on container |

### Architecture of SNMP features

In terms of architecture of SNMP features, there is some difference point compared with legacy routers. The main point of difference is that Beluganos has "proxy" process.

#### SNMP MIB

Beluganos uses NET-SNMP("[snmpd]") to support SNMP MIB feature, but NET-SNMP is not assumed to rewrite the value of standard OID. This is because internal OID (enterprise OID) is used by NET-SNMP in order to extend the statistics feature. Moreover, to support standard MIB, Beluganos has "[snmpproxyd-mib]" component. This is a proxy daemon that convert internal OID to standard OID. Thus, [snmp client] can use standard MIB to get some statistics.

```
                 | <---           Beluganos's host           ---> |

                         port:161          port:8161
 [snmp client] <-+-> [snmpproxyd-mib]  <-> [snmpd(+fibssnmp)]
                                               |
                           /var/lib/beluganos/fibc_stats.yaml
                                               |
                                            [fibsd] <-> [fibcd] <-+-> [OpenNSL]
```

#### SNMP trap

SNMP trap feature is also realized by "NET-SNMP". In [snmpproxyd-trap] component, the conversion of ifindex.

```
                 | <---              Beluganos's host               ---> |
                                         | <---     on container    ---> |

                         port:162
 [trap server] <-+-> [snmpproxyd-trap] <-+- [snmpd]
                                         +- [snmpifmond]
                                            (send interface information)
```

## Appendix

### Appendix A. Advanced configurations

There is four point to configure.

- Daemon
	1. fibsd (Beluganos)
	2. snmpproxyd-mib, snmpproxyd-trap (Beluganos)
	3. snmpd (NET-SNMP)
- Not daemon
	1. fibssnmp (Beluganos)

#### fibsd (Beluganos)

Generally, you do NOT have to change this configurations.

```
$ vi /etc/beluganos/fibsd.conf

FIBC_ADDR=localhost:8080
STATS_PATH=/var/lib/beluganos/fibc_stats.yaml
UPDATE_TIME=5s
```

- `FIBC_ADDR`
	- The fibcd address (DO NOT EDIT)
- `STATS_PATH`
	- The path of data file to save statistics information (DO NOT EDIT)
- `UPDATE_TIME`
	- Update interval (ex: 5s, 1m30s)

The statistics is saved as text file at `STATS_PATH`. The example of text file is here:

```
$ cat /var/lib/beluganos/fibc_stats.yaml

---

port_stats:
- ifName: eth1
  ifInDiscards: 2000
  ifInErrors: 2100
  ifInNUcastPkts: 21000
  ifInOctets: 22000
  ifInUcastPkts: 20000
  ifOperStatus: 1
  ifOutDiscards: 1000
  ifOutErrors: 1100
  ifOutNUcastPkts: 11000
  ifOutOctets: 12000
  ifOutUcastPkts: 10000
  port_no: 1

- ifName: eth2
  ifInDiscards: 2001
  ifInErrors: 2101
  ifInNUcastPkts: 21001
  ifInOctets: 22001
  ifInUcastPkts: 20001
  ifOperStatus: 1
  ifOutDiscards: 1001
  ifOutErrors: 1101
  ifOutNUcastPkts: 11001
  ifOutOctets: 12001
  ifOutUcastPkts: 10001
  port_no: 2

```

Note that Beluganos's module will create `fibc_stats.yaml` automatically. This value is synchronized with hardware's statistics every `UPDATE_TIME` (by default, 5 sec.). Please do not edit this file manually.

#### snmpproxyd (Beluganos)

The setting file of snmpproxyd-mib and snmpproxyd-trap is common.

```
$ vi /etc/beluganos/snmpproxyd.yaml

---

snmpproxy:
  default:
    oidmap:
      - name:  ifOperStatus
        oid:   .1.3.6.1.2.1.2.2.1.8
        local: .1.3.6.1.4.99999.2.2.1.8
      - name:  ifInOctets
        oid:   .1.3.6.1.2.1.2.2.1.10
        local: .1.3.6.1.4.99999.2.2.1.10
      ~~~~ (snipped) ~~~~
    ifmap:
      oidmap:
        min: 0
        max: 1023
      shift:
        min: 1024
        max: 2147483647
    trap2map:
      eth1: 1
      eth2: 2
      eth3: 3
      ~~~~ (snipped) ~~~~
      lxdbr0: 500
    trap2sink:
      - addr: 192.168.122.1:161
      - addr: 192.168.122.2:161
```
- `snmpproxy`
	- `default`
		- `oidmap`
			- The list of conversion internal MIB and standard MIB.
			- Generally, DO NOT EDIT.
		- `ifmap`
			- DO NOT EDIT.
		- `trap2map`
			- The list of conversion ifindex. This depends on hardware. The example is described at the bottom of this documents.
		- `trap2sink`
			- The lists of IP address of trap servers.
			- You can set one or more SNMP trap servers.

#### snmpd (NET-SNMP)

The required settings will be edited automatically. The settings which is setup automatically is described here. Note that `-` means deletion is required, and `+` means addition is required.

##### `/etc/snmp/snmpd.conf`

```
+ pass_persist <OID> /usr/bin/fibssnmp

- # createUser internalUser  MD5 "this is only ever used internally, but still change the password"
+ createUser internalUser  MD5 "this is only ever used internally, but still change the password"

- #trap2sink    localhost public
+ trap2sink    192.169.1.1 public

- defaultMonitors
+ #defaultMonitors

- linkUpDownNotifications
+ #linkUpDownNotifications

+ notificationEvent linkUpTrap linkUp ifIndex ifAdminStatus ifOperStatus
+ monitor -r 10 -o ifName -e linkUpTrap "Generate linkUp" ifOperStatus != 2
+ notificationEvent linkDownTrap linkDown ifindex ifAdminStatus ifOperStatus
+ monitor -r 10 -o ifName -e linkDownTrap "Generate linkDown" ifOperStatus == 2
```

Note: The `monitor -r <interval-in-sec>` is the settings about the interval time of monitoring interface. You may change it.

##### `/etc/snmp/snmp.conf`

```
- mibs :
+ mibs +ALL
```

##### `/lib/systemd/system/snmpd.service`

```
- Environment="MIBS="
+ #Environment="MIBS="

- ExecStart=/usr/sbin/snmpd -Lsd -Lf /dev/null -u Debian-snmp -g Debian-snmp -I -smux,mteTrigger,mteTriggerConf -f
+ ExecStart=/usr/sbin/snmpd -Lsd -Lf /dev/null -u Debian-snmp -g Debian-snmp -I -smux -f
```
##### Reflect changes

At LXC, please enter following:

```
$ sudo systemctl daemon-reload
$ sudo systemctl restart snmpd
```

#### fibssnmp (Beluganos)

Generally, you do NOT have to change this configurations.

```
$ vi  /etc/beluganos/fibssnmp.yaml

---

handlers:
- oid : .1.3.6.1.4.1.99999.1
  name: ifName
  type: string
- oid : .1.3.6.1.4.1.99999.2
  name: ifInOctets
  type: integer
```

- `oid`: internal OID
- `name`: name of OID
- `type`: type of value

### Appendix B. example of `trap2map` configuration

To use SNMP trap feature, you should change `trap2map` configuration at `/etc/beluganos/snmpproxyd.yaml` depend on your hardware. The mapping information is described at [configure-portmapping.md](configure-portmapping.md).

In this chapter, the examples of `trap2map` is described.

#### AS5710-54X, AS5712-54X, AS5812-54X, AS5812-54T / OpenNSL

Note: This configuration is for without breakout cable case. If you configure breakout cable, you should change this configuration.

```
    trap2map:
      eth1: 1
      eth2: 2
      eth3: 3
      eth4: 4
      eth5: 5
      eth6: 6
      eth7: 7
      eth8: 8
      eth9: 9
      eth10: 10
      eth11: 11
      eth12: 12
      eth13: 13
      eth14: 14
      eth15: 15
      eth16: 16
      eth17: 17
      eth18: 18
      eth19: 19
      eth20: 20
      eth21: 21
      eth22: 22
      eth23: 23
      eth24: 24
      eth25: 25
      eth26: 26
      eth27: 27
      eth28: 28
      eth29: 29
      eth30: 30
      eth31: 31
      eth32: 32
      eth33: 33
      eth34: 34
      eth35: 35
      eth36: 36
      eth37: 37
      eth38: 38
      eth39: 39
      eth40: 40
      eth41: 41
      eth42: 42
      eth43: 43
      eth44: 44
      eth45: 45
      eth46: 46
      eth47: 47
      eth48: 48
      eth49: 49
      eth50: 53
      eth51: 57
      eth52: 61
      eth53: 65
      eth54: 69
      lxdbr0: 500
```

#### AS7712-32X, AS7716-32X without breakout / OpenNSL

Note: This configuration is for without breakout cable case. If you configure breakout cable, you should change this configuration.

```
    trap2map:
      eth1: 50
      eth2: 54
      eth3: 58
      eth4: 62
      eth5: 68
      eth6: 72
      eth7: 76
      eth8: 80
      eth9: 34
      eth10: 38
      eth11: 42
      eth12: 46
      eth13: 84
      eth14: 88
      eth15: 92
      eth16: 96
      eth17: 102
      eth18: 106
      eth19: 110
      eth20: 114
      eth21: 17
      eth22: 21
      eth23: 25
      eth24: 29
      eth25: 118
      eth26: 122
      eth27: 126
      eth28: 130
      eth29: 1
      eth30: 5
      eth31: 9
      eth32: 13
      lxdbr0: 500
```
