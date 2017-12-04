package dhcpop

import (
	"bytes"
	"net"
	"testing"
	"time"
)

var (
	filenameString = "filename.ext"
	filenameBytes  = []byte{102, 105, 108, 101, 110, 97, 109, 101, 46, 101, 120, 116}

	ip       = net.IPv4(192, 168, 100, 1)
	ipBytes  = []byte{192, 168, 100, 1}
	ipString = "192.168.100.1"

	ips       = []IPv4{IPv4(net.IPv4(192, 168, 100, 1)), IPv4(net.IPv4(192, 168, 100, 2))}
	ipsBytes  = []byte{192, 168, 100, 1, 192, 168, 100, 2}
	ipsString = "192.168.100.1,192.168.100.2"

	ipPair       = ips
	ipPairBytes  = ipsBytes
	ipPairString = "192.168.100.1 192.168.100.2"

	ipPairs = []IPv4Pair{
		IPv4Pair{ips[0], ips[1]},
		IPv4Pair{ips[1], ips[0]},
	}
	ipPairsBytes  = []byte{192, 168, 100, 1, 192, 168, 100, 2, 192, 168, 100, 2, 192, 168, 100, 1}
	ipPairsString = "192.168.100.1 192.168.100.2,192.168.100.2 192.168.100.1"

	route = Route{
		Source: net.IPNet{
			IP:   net.IPv4(192, 168, 100, 0),
			Mask: net.CIDRMask(24, 32),
		},
		Destination: net.IPv4(127, 0, 0, 1),
	}
	routeBytes  = []byte{24, 192, 168, 100, 127, 0, 0, 1}
	routeString = "192.168.100.0/24 127.0.0.1"

	routes = Routes{
		route,
		Route{
			Source: net.IPNet{
				IP:   net.IPv4(0, 0, 0, 0),
				Mask: net.CIDRMask(0, 32),
			},
			Destination: ip,
		},
	}
	routesBytes  = []byte{24, 192, 168, 100, 127, 0, 0, 1, 0, 192, 168, 100, 1}
	routesString = "192.168.100.0/24 127.0.0.1,0.0.0.0/0 192.168.100.1"

	domainName       = DomainName{"domain", "tld"}
	domainNameBytes  = []byte{6, 100, 111, 109, 97, 105, 110, 3, 116, 108, 100, 0}
	domainNameString = "domain.tld"

	domainNames = DomainNames{
		domainName,
		DomainName{"example", "com"},
	}
	domainNamesBytes  = []byte{6, 100, 111, 109, 97, 105, 110, 3, 116, 108, 100, 0, 7, 101, 120, 97, 109, 112, 108, 101, 3, 99, 111, 109, 0}
	domainNamesString = domainNameString + ",example.com"

	timeOffset       = time.Duration(0x123456) * time.Second
	timeOffsetBytes  = []byte{0, 18, 52, 86}
	timeOffsetString = "331h24m6s"

	timeDuration       = timeOffset
	timeDurationBytes  = timeOffsetBytes
	timeDurationString = timeOffsetString
)

func TestEncodeString(t *testing.T) {
	op := String(filenameString)
	if bytes.Compare(op.Encode(), filenameBytes) != 0 {
		t.Error()
	}
}

func TestDecodeString(t *testing.T) {
	op := new(String)
	if err := op.Decode(filenameBytes); err != nil {
		t.Error(err)
	}
	if string(*op) != filenameString {
		t.Error()
	}
}

func TestMarshalString(t *testing.T) {
	op := String(filenameString)
	if string(string(op.Marshal())) != filenameString {
		t.Error()
	}
}

func TestUnmarshalString(t *testing.T) {
	op := new(String)
	if err := op.Unmarshal(filenameBytes); err != nil {
		t.Error(err)
	}
	if string(*op) != filenameString {
		t.Error()
	}
}

func TestEncodeBooleanTrue(t *testing.T) {
	op := Boolean(true)
	if bytes.Compare(op.Encode(), []byte{1}) != 0 {
		t.Error()
	}
}

func TestEncodeBooleanFalse(t *testing.T) {
	op := Boolean(false)
	if bytes.Compare(op.Encode(), []byte{0}) != 0 {
		t.Error()
	}
}

func TestDecodeBooleanTrue(t *testing.T) {
	op := new(Boolean)
	if err := op.Decode([]byte{1}); err != nil {
		t.Error(err)
	}
	if !*op {
		t.Error()
	}
}

func TestDecodeBooleanFalse(t *testing.T) {
	op := new(Boolean)
	if err := op.Decode([]byte{0}); err != nil {
		t.Error(err)
	}
	if *op {
		t.Error()
	}
}

func TestMarshalBooleanTrue(t *testing.T) {
	op := Boolean(true)
	if string(op.Marshal()) != "true" {
		t.Error()
	}
}

func TestMarshalBooleanFalse(t *testing.T) {
	op := Boolean(false)
	if string(op.Marshal()) != "false" {
		t.Error()
	}
}

func TestUnmarshalBooleanTrue(t *testing.T) {
	op := new(Boolean)
	if err := op.Unmarshal([]byte("true")); err != nil {
		t.Error(err)
	}
	if !*op {
		t.Error()
	}
}

func TestUnmarshalBooleanFalse(t *testing.T) {
	op := new(Boolean)
	if err := op.Unmarshal([]byte("false")); err != nil {
		t.Error(err)
	}
	if *op {
		t.Error()
	}
}

func TestEncodeByte(t *testing.T) {
	op := Byte(123)
	if bytes.Compare(op.Encode(), []byte{123}) != 0 {
		t.Error()
	}
}

func TestDecodeByte(t *testing.T) {
	op := new(Byte)
	if err := op.Decode([]byte{123}); err != nil {
		t.Error(err)
	}
	if byte(*op) != 123 {
		t.Error()
	}
}

func TestMarshalByte(t *testing.T) {
	op := Byte(123)
	if string(op.Marshal()) != "123" {
		t.Error()
	}
}

func TestUnmarshalByte(t *testing.T) {
	op := new(Byte)
	if err := op.Unmarshal([]byte("123")); err != nil {
		t.Error(err)
	}
	if byte(*op) != 123 {
		t.Error()
	}
}

func TestEncodeSize(t *testing.T) {
	op := Size(0x1234)
	if bytes.Compare(op.Encode(), []byte{0x12, 0x34}) != 0 {
		t.Error()
	}
}

func TestDecodeSize(t *testing.T) {
	op := new(Size)
	if err := op.Decode([]byte{0x12, 0x34}); err != nil {
		t.Error(err)
	}
	if uint16(*op) != 0x1234 {
		t.Error()
	}
}

func TestMarshalSize(t *testing.T) {
	op := Size(1234)
	if string(op.Marshal()) != "1234" {
		t.Error()
	}
}

func TestUnmarshalSize(t *testing.T) {
	op := new(Size)
	if err := op.Unmarshal([]byte("1234")); err != nil {
		t.Error(err)
	}
	if uint16(*op) != 1234 {
		t.Error()
	}
}

func TestEncodeSizes(t *testing.T) {
	op := Sizes{0x1234, 0x5678}
	if bytes.Compare(op.Encode(), []byte{0x12, 0x34, 0x56, 0x78}) != 0 {
		t.Error()
	}
}

func TestDecodeSizes(t *testing.T) {
	op := new(Sizes)
	if err := op.Decode([]byte{0x12, 0x34, 0x56, 0x78}); err != nil {
		t.Error(err)
	}
	if uint16((*op)[0]) != 0x1234 {
		t.Error()
	}
	if uint16((*op)[1]) != 0x5678 {
		t.Error()
	}
}

func TestMarshalSizes(t *testing.T) {
	op := Sizes{1234, 5678}
	if string(op.Marshal()) != "1234,5678" {
		t.Error()
	}
}

func TestUnmarshalSizes(t *testing.T) {
	op := new(Sizes)
	if err := op.Unmarshal([]byte("1234, 5678")); err != nil {
		t.Error(err)
	}
	if uint16((*op)[0]) != 1234 {
		t.Error()
	}
	if uint16((*op)[1]) != 5678 {
		t.Error()
	}
}

func TestEncodeIPv4(t *testing.T) {
	op := IPv4(ip)
	if bytes.Compare(op.Encode(), ipBytes) != 0 {
		t.Error()
	}
}

func TestDecodeIPv4(t *testing.T) {
	op := new(IPv4)
	if err := op.Decode(ipBytes); err != nil {
		t.Error(err)
	}
	if !(net.IP(*op)).Equal(ip) {
		t.Error()
	}
}

func TestMarshalIPv4(t *testing.T) {
	op := IPv4(ip)
	if string(op.Marshal()) != ipString {
		t.Error()
	}
}

func TestUnmarshalIPv4(t *testing.T) {
	op := new(IPv4)
	if err := op.Unmarshal([]byte(ipString)); err != nil {
		t.Error(err)
	}
	if !(net.IP(*op)).Equal(ip) {
		t.Error()
	}
}

func TestEncodeIPv4s(t *testing.T) {
	op := IPv4s(ips)
	if bytes.Compare(op.Encode(), ipsBytes) != 0 {
		t.Error()
	}
}

func TestDecodeIPv4s(t *testing.T) {
	op := new(IPv4s)
	if err := op.Decode(ipsBytes); err != nil {
		t.Error(err)
	}
	for i, ip := range *op {
		if !net.IP(ip).Equal(net.IP(ips[i])) {
			t.Error()
		}
	}
}

func TestMarshalIPv4s(t *testing.T) {
	op := IPv4s(ips)
	if string(op.Marshal()) != ipsString {
		t.Error()
	}
}

func TestUnmarshalIPv4s(t *testing.T) {
	op := new(IPv4s)
	if err := op.Unmarshal([]byte(ipsString)); err != nil {
		t.Error(err)
	}
	for i, ip := range *op {
		if !net.IP(ip).Equal(net.IP(ips[i])) {
			t.Error()
		}
	}
}

func TestEncodeIPv4Pair(t *testing.T) {
	op := IPv4Pair{ipPair[0], ipPair[1]}
	if bytes.Compare(op.Encode(), ipPairBytes) != 0 {
		t.Error()
	}
}

func TestDecodeIPv4Pair(t *testing.T) {
	op := new(IPv4Pair)
	if err := op.Decode(ipPairBytes); err != nil {
		t.Error(err)
	}
	for i, ip := range *op {
		if !net.IP(ip).Equal(net.IP(ipPair[i])) {
			t.Error()
		}
	}
}

func TestMarshalIPv4Pair(t *testing.T) {
	op := IPv4Pair{ipPair[0], ipPair[1]}
	if string(op.Marshal()) != ipPairString {
		t.Error()
	}
}

func TestUnmarshalIPv4Pair(t *testing.T) {
	op := new(IPv4Pair)
	if err := op.Unmarshal([]byte(ipPairString)); err != nil {
		t.Error(err)
	}
	for i, ip := range *op {
		if !net.IP(ip).Equal(net.IP(ipPair[i])) {
			t.Error()
		}
	}
}

func TestEncodeIPv4Pairs(t *testing.T) {
	op := IPv4Pairs(ipPairs)
	if bytes.Compare(op.Encode(), ipPairsBytes) != 0 {
		t.Error()
	}
}

func TestDecodeIPv4Pairs(t *testing.T) {
	op := new(IPv4Pairs)
	if err := op.Decode(ipPairsBytes); err != nil {
		t.Error(err)
	}
	for i, p := range *op {
		if !net.IP(p[0]).Equal(net.IP(ipPairs[i][0])) {
			t.Error()
		}
		if !net.IP(p[1]).Equal(net.IP(ipPairs[i][1])) {
			t.Error()
		}
	}
}

func TestMarshalIPv4Pairs(t *testing.T) {
	op := IPv4Pairs(ipPairs)
	if string(op.Marshal()) != ipPairsString {
		t.Error()
	}
}

func TestUnmarshalIPv4Pairs(t *testing.T) {
	op := new(IPv4Pairs)
	if err := op.Unmarshal([]byte(ipPairsString)); err != nil {
		t.Error(err)
	}
	for i, p := range *op {
		if !net.IP(p[0]).Equal(net.IP(ipPairs[i][0])) {
			t.Error()
		}
		if !net.IP(p[1]).Equal(net.IP(ipPairs[i][1])) {
			t.Error()
		}
	}
}

func TestEncodeRoute(t *testing.T) {
	op := route
	if bytes.Compare(op.Encode(), routeBytes) != 0 {
		t.Error()
	}
}

func TestDecodeRoute(t *testing.T) {
	op := new(Route)
	if err := op.Decode(routeBytes); err != nil {
		t.Error(err)
	}
	if bytes.Compare(op.Encode(), route.Encode()) != 0 {
		t.Error()
	}
}

func TestMarshalRoute(t *testing.T) {
	op := route
	if string(op.Marshal()) != routeString {
		t.Error()
	}
}

func TestUnmarshalRoute(t *testing.T) {
	op := new(Route)
	if err := op.Unmarshal([]byte(routeString)); err != nil {
		t.Error(err)
	}
	if bytes.Compare(op.Encode(), route.Encode()) != 0 {
		t.Error()
	}
}

func TestEncodeRoutes(t *testing.T) {
	op := routes
	if bytes.Compare(op.Encode(), routesBytes) != 0 {
		t.Error()
	}
}

func TestDecodeRoutes(t *testing.T) {
	op := new(Routes)
	if err := op.Decode(routesBytes); err != nil {
		t.Error(err)
	}
	if bytes.Compare(op.Encode(), routes.Encode()) != 0 {
		t.Error()
	}
}

func TestMarshalRoutes(t *testing.T) {
	op := routes
	if string(op.Marshal()) != routesString {
		t.Log(op)
		t.Error()
	}
}

func TestUnmarshalRoutes(t *testing.T) {
	op := new(Routes)
	if err := op.Unmarshal([]byte(routesString)); err != nil {
		t.Error(err)
	}
	if bytes.Compare(op.Encode(), routes.Encode()) != 0 {
		t.Error()
	}
}

func TestEncodeDomainName(t *testing.T) {
	op := domainName
	if bytes.Compare(op.Encode(), domainNameBytes) != 0 {
		t.Error()
	}
}

func TestDecodeDomainName(t *testing.T) {
	op := new(DomainName)
	if err := op.Decode(domainNameBytes); err != nil {
		t.Error(err)
	}
	for i, s := range *op {
		if s != domainName[i] {
			t.Error()
		}
	}
}

func TestMarshalDomainName(t *testing.T) {
	op := domainName
	if string(op.Marshal()) != domainNameString {
		t.Error()
	}
}

func TestUnmarshalDomainName(t *testing.T) {
	op := new(DomainName)
	if err := op.Unmarshal([]byte(domainNameString)); err != nil {
		t.Error(err)
	}
	for i, s := range *op {
		if s != domainName[i] {
			t.Error()
		}
	}
}

func TestEncodeDomainNames(t *testing.T) {
	op := domainNames
	if bytes.Compare(op.Encode(), domainNamesBytes) != 0 {
		t.Error()
	}
}

func TestDecodeDomainNames(t *testing.T) {
	op := new(DomainNames)
	if err := op.Decode(domainNamesBytes); err != nil {
		t.Error(err)
	}
	for i, dn := range *op {
		if string(dn.Marshal()) != string(domainNames[i].Marshal()) {
			t.Error()
		}
	}
}

func TestMarshalDomainNames(t *testing.T) {
	op := domainNames
	if string(op.Marshal()) != domainNamesString {
		t.Error()
	}
}

func TestUnmarshalDomainNames(t *testing.T) {
	op := new(DomainNames)
	if err := op.Unmarshal([]byte(domainNamesString)); err != nil {
		t.Error(err)
	}
	for i, dn := range *op {
		if string(dn.Marshal()) != string(domainNames[i].Marshal()) {
			t.Error()
		}
	}
}

func TestEncodeTimeOffset(t *testing.T) {
	op := TimeOffset(timeOffset)
	if bytes.Compare(op.Encode(), timeOffsetBytes) != 0 {
		t.Error()
	}
}

func TestDecodeTimeOffset(t *testing.T) {
	op := new(TimeOffset)
	if err := op.Decode(timeOffsetBytes); err != nil {
		t.Error(err)
	}
	if *op != TimeOffset(timeOffset) {
		t.Error()
	}
}

func TestMarshalTimeOffset(t *testing.T) {
	op := TimeOffset(timeOffset)
	if string(op.Marshal()) != timeOffsetString {
		t.Error()
	}
}

func TestUnmarshalTimeOffset(t *testing.T) {
	op := new(TimeOffset)
	if err := op.Unmarshal([]byte(timeOffsetString)); err != nil {
		t.Error(err)
	}
	if *op != TimeOffset(timeOffset) {
		t.Error()
	}
}

func TestEncodeTimeDuration(t *testing.T) {
	op := TimeDuration(timeDuration)
	if bytes.Compare(op.Encode(), timeDurationBytes) != 0 {
		t.Error()
	}
}

func TestDecodeTimeDuration(t *testing.T) {
	op := new(TimeDuration)
	if err := op.Decode(timeDurationBytes); err != nil {
		t.Error(err)
	}
	if *op != TimeDuration(timeDuration) {
		t.Error()
	}
}

func TestMarshalTimeDuration(t *testing.T) {
	op := TimeDuration(timeDuration)
	if string(op.Marshal()) != timeDurationString {
		t.Error()
	}
}

func TestUnmarshalTimeDuration(t *testing.T) {
	op := new(TimeDuration)
	if err := op.Unmarshal([]byte(timeDurationString)); err != nil {
		t.Error(err)
	}
	if *op != TimeDuration(timeDuration) {
		t.Error()
	}
}

func TestEncodePadding(t *testing.T) {
	op := new(Padding)
	if len(op.Encode()) > 0 {
		t.Error()
	}
}

func TestDecodePadding(t *testing.T) {
	op := new(Padding)
	if err := op.Decode([]byte{}); err != nil {
		t.Error(err)
	}
}

func TestMarshalPadding(t *testing.T) {
	op := new(Padding)
	if string(op.Marshal()) != "" {
		t.Error()
	}
}

func TestUnmarshalPadding(t *testing.T) {
	op := new(Padding)
	if err := op.Unmarshal([]byte{}); err != nil {
		t.Error(err)
	}
}

func TestEncodeEnd(t *testing.T) {
	op := new(End)
	if len(op.Encode()) > 0 {
		t.Error()
	}
}

func TestDecodeEnd(t *testing.T) {
	op := new(End)
	if err := op.Decode([]byte{}); err != nil {
		t.Error(err)
	}
}

func TestMarshalEnd(t *testing.T) {
	op := new(End)
	if string(op.Marshal()) != "" {
		t.Error()
	}
}

func TestUnmarshalEnd(t *testing.T) {
	op := new(End)
	if err := op.Unmarshal([]byte{}); err != nil {
		t.Error(err)
	}
}
