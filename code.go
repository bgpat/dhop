package dhop

type Code byte

func (c *Code) String() string {
	// https://www.iana.org/assignments/bootp-dhcp-parameters/bootp-dhcp-parameters.xhtml
	switch byte(*c) {
	case 0:
		return "Pad"
	case 1:
		return "Subnet Mask"
	case 2:
		return "Time Offset"
	case 3:
		return "Router"
	case 4:
		return "Time Server"
	case 5:
		return "Name Server"
	case 6:
		return "Domain Server"
	case 7:
		return "Log Server"
	case 8:
		return "Quotes Server"
	case 9:
		return "LPR Server"
	case 10:
		return "Impress Server"
	case 11:
		return "RLP Server"
	case 12:
		return "Hostname"
	case 13:
		return "Boot File Size"
	case 14:
		return "Merit Dump File"
	case 15:
		return "Domain Name"
	case 16:
		return "Swap Server"
	case 17:
		return "Root Path"
	case 18:
		return "Extension File"
	case 19:
		return "Forward On/Off"
	case 20:
		return "SrcRte On/Off"
	case 21:
		return "Policy Filter"
	case 22:
		return "Max DG Assembly"
	case 23:
		return "Default IP TTL"
	case 24:
		return "MTU Timeout"
	case 25:
		return "MTU Plateau"
	case 26:
		return "MTU Interface"
	case 27:
		return "MTU Subnet"
	case 28:
		return "Broadcast Address"
	case 29:
		return "Mask Discovery"
	case 30:
		return "Mask Supplier"
	case 31:
		return "Router Discovery"
	case 32:
		return "Router Request"
	case 33:
		return "Static Route"
	case 34:
		return "Trailers"
	case 35:
		return "ARP Timeout"
	case 36:
		return "Ethernet"
	case 37:
		return "Default TCP TTL"
	case 38:
		return "Keepalive Time"
	case 39:
		return "Keepalive Data"
	case 40:
		return "NIS Domain"
	case 41:
		return "NIS Servers"
	case 42:
		return "NTP Servers"
	case 43:
		return "Vendor Specific"
	case 44:
		return "NETBIOS Name Srv"
	case 45:
		return "NETBIOS Dist Srv"
	case 46:
		return "NETBIOS Node Type"
	case 47:
		return "NETBIOS Scope"
	case 48:
		return "X Window Font"
	case 49:
		return "X Window Manager"
	case 50:
		return "Address Request"
	case 51:
		return "Address Time"
	case 52:
		return "Overload"
	case 53:
		return "DHCP Msg Type"
	case 54:
		return "DHCP Server Id"
	case 55:
		return "Parameter List"
	case 56:
		return "DHCP Message"
	case 57:
		return "DHCP Max Msg Size"
	case 58:
		return "Renewal Time"
	case 59:
		return "Rebinding Time"
	case 60:
		return "Class Id"
	case 61:
		return "Client Id"
	case 62:
		return "NetWare/IP Domain"
	case 63:
		return "NetWare/IP Option"
	case 64:
		return "NIS-Domain-Name"
	case 65:
		return "NIS-Server-Addr"
	case 66:
		return "Server-Name"
	case 67:
		return "Bootfile-Name"
	case 68:
		return "Home-Agent-Addrs"
	case 69:
		return "SMTP-Server"
	case 70:
		return "POP3-Server"
	case 71:
		return "NNTP-Server"
	case 72:
		return "WWW-Server"
	case 73:
		return "Finger-Server"
	case 74:
		return "IRC-Server"
	case 75:
		return "StreetTalk-Server"
	case 76:
		return "STDA-Server"
	case 77:
		return "User-Class"
	case 78:
		return "Directory Agent"
	case 79:
		return "Service Scope"
	case 80:
		return "Rapid Commit"
	case 81:
		return "Client FQDN"
	case 82:
		return "Relay Agent Information"
	case 83:
		return "iSNS"
	case 85:
		return "NDS Servers"
	case 86:
		return "NDS Tree Name"
	case 87:
		return "NDS Context"
	case 88:
		return "BCMCS Controller Domain Name list"
	case 89:
		return "BCMCS Controller IPv4 address option"
	case 90:
		return "Authentication"
	case 91:
		return "client-last-transaction-time option"
	case 92:
		return "associated-ip option"
	case 93:
		return "Client System"
	case 94:
		return "Client NDI"
	case 95:
		return "LDAP"
	case 97:
		return "UUID/GUID"
	case 98:
		return "User-Auth"
	case 99:
		return "GEOCONF_CIVIC"
	case 100:
		return "PCode"
	case 101:
		return "TCode"
	case 112:
		return "Netinfo Address"
	case 113:
		return "Netinfo Tag"
	case 114:
		return "URL"
	case 116:
		return "Auto-Config"
	case 117:
		return "Name Service Search"
	case 118:
		return "Subnet Selection Option"
	case 119:
		return "Domain Search"
	case 120:
		return "SIP Servers DHCP Option"
	case 121:
		return "Classless Static Route Option"
	case 122:
		return "CCC"
	case 123:
		return "GeoConf Option"
	case 124:
		return "V-I Vendor Class"
	case 125:
		return "V-I Vendor-Specific Information"
	case 128:
		return "TFTP Server IP address (for IP Phone software load)"
	case 129:
		return "Call Server IP address"
	case 130:
		return "Discrimination string (to identify vendor)"
	case 131:
		return "Remote statistics server IP address"
	case 132:
		return "IEEE 802.1Q VLAN ID"
	case 133:
		return "IEEE 802.1D/p Layer 2 Priority"
	case 134:
		return "Diffserv Code Point (DSCP) for VoIP signalling and media streams"
	case 135:
		return "HTTP Proxy for phone-specific applications"
	case 136:
		return "OPTION_PANA_AGENT"
	case 137:
		return "OPTION_V4_LOST"
	case 138:
		return "OPTION_CAPWAP_AC_V4"
	case 139:
		return "OPTION-IPv4_Address-MoS"
	case 140:
		return "OPTION-IPv4_FQDN-MoS"
	case 141:
		return "SIP UA Configuration Service Domains"
	case 142:
		return "OPTION-IPv4_Address-ANDSF"
	case 143:
		return "OPTION_V4_ZEROTOUCH_REDIRECT (TEMPORARY - registered 2018-02-08, expires 2019-02-08)"
	case 144:
		return "GeoLoc"
	case 145:
		return "FORCERENEW_NONCE_CAPABLE"
	case 146:
		return "RDNSS Selection"
	case 150:
		return "TFTP server address"
	case 151:
		return "status-code"
	case 152:
		return "base-time"
	case 153:
		return "start-time-of-state"
	case 154:
		return "query-start-time"
	case 155:
		return "query-end-time"
	case 156:
		return "dhcp-state"
	case 157:
		return "data-source"
	case 158:
		return "OPTION_V4_PCP_SERVER"
	case 159:
		return "OPTION_V4_PORTPARAMS"
	case 160:
		return "DHCP Captive-Portal"
	case 161:
		return "OPTION_MUD_URL_V4 (TEMPORARY - registered 2016-11-17, extension registered 2017-10-02, expires 2018-11-17)"
	case 175:
		return "Etherboot (Tentatively Assigned - 2005-06-23)"
	case 176:
		return "IP Telephone (Tentatively Assigned - 2005-06-23)"
	case 177:
		return "PacketCable and CableHome (replaced by 122)"
	case 208:
		return "PXELINUX Magic"
	case 209:
		return "Configuration File"
	case 210:
		return "Path Prefix"
	case 211:
		return "Reboot Time"
	case 212:
		return "OPTION_6RD"
	case 213:
		return "OPTION_V4_ACCESS_DOMAIN"
	case 220:
		return "Subnet Allocation Option"
	case 221:
		return "Virtual Subnet Selection (VSS) Option"
	case 255:
		return "End"
	}
	if 224 <= *c && *c <= 254 {
		return "Reserved (Private Use)"
	}
	return "N/A"
}
