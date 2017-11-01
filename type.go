package dhcpop

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Option interface {
	Encode() []byte
	Decode([]byte) error
	String() string
}

type String string

type Boolean bool

type Byte byte

type Size uint16

type Sizes []Size

type IPv4 net.IP

type IPv4s []IPv4

type IPv4Pair [2]IPv4

type IPv4Pairs []IPv4Pair

type Route struct {
	Source      net.IPNet
	Destination net.IP
}

type Routes []Route

type DomainName []string

type DomainNames []DomainName

type TimeOffset time.Duration

type TimeDuration time.Duration

type Padding struct{}

type End struct{}

func (o *String) Encode() []byte {
	return []byte(*o)
}

func (o *String) Decode(b []byte) error {
	*o = String(b)
	return nil
}

func (o *String) String() string {
	return string(*o)
}

func (o *Boolean) Encode() []byte {
	if *o == true {
		return []byte{1}
	}
	return []byte{0}
}

func (o *Boolean) Decode(b []byte) error {
	if err := validateSize(b, 1); err != nil {
		return err
	}
	if b[0] == 1 {
		*o = true
	} else if b[0] == 0 {
		*o = false
	} else {
		return &InvalidFormatError{}
	}
	return nil
}

func (o *Boolean) String() string {
	if *o {
		return "true"
	}
	return "false"
}

func (o *Byte) Encode() []byte {
	return []byte{byte(*o)}
}

func (o *Byte) Decode(b []byte) error {
	if err := validateSize(b, 1); err != nil {
		return err
	}
	*o = Byte(b[0])
	return nil
}

func (o *Byte) String() string {
	return strconv.Itoa(int(*o))
}

func (o *Size) Encode() []byte {
	return []byte{
		byte(*o >> 8),
		byte(*o),
	}
}

func (o *Size) Decode(b []byte) error {
	if err := validateSize(b, 2); err != nil {
		return err
	}
	*o = Size(int(b[0])<<8 | int(b[1]))
	return nil
}

func (o *Size) String() string {
	return strconv.Itoa(int(*o))
}

func (o *Sizes) Encode() []byte {
	b := make([]byte, len(*o)<<1)
	for i, s := range *o {
		b[i<<1] = byte(s >> 8)
		b[i<<1|1] = byte(s)
	}
	return b
}

func (o *Sizes) Decode(b []byte) error {
	if err := validateSizeFactor(b, 2); err != nil {
		return err
	}
	if err := validateMinimumSize(b, 2); err != nil {
		return err
	}
	l := len(b) >> 1
	*o = make([]Size, l)
	for i := 0; i < l; i++ {
		err := (*o)[i].Decode(b[i<<1 : (i+1)<<1])
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Sizes) String() string {
	s := make([]string, len(*o))
	for i, n := range *o {
		s[i] = n.String()
	}
	return strings.Join(s, ",")
}

func (o *IPv4) Encode() []byte {
	return []byte(net.IP(*o).To4())
}

func (o *IPv4) Decode(b []byte) error {
	if err := validateSize(b, 4); err != nil {
		return err
	}
	*o = IPv4(b)
	return nil
}

func (o *IPv4) String() string {
	return net.IP(*o).String()
}

func (o *IPv4s) Encode() []byte {
	b := make([]byte, 0, len(*o)*4)
	for _, ip := range *o {
		b = append(b, ip.Encode()...)
	}
	return b
}

func (o *IPv4s) Decode(b []byte) error {
	if err := validateSizeFactor(b, 4); err != nil {
		return err
	}
	if err := validateMinimumSize(b, 4); err != nil {
		return err
	}
	l := len(b) / 4
	*o = make([]IPv4, l)
	for i := 0; i < l; i++ {
		(*o)[i] = IPv4(b[i*4 : i*4+4])
	}
	return nil
}

func (o *IPv4s) String() string {
	s := make([]string, len(*o))
	for i, ip := range *o {
		s[i] = ip.String()
	}
	return strings.Join(s, ",")
}

func (o *IPv4Pair) Encode() []byte {
	b := make([]byte, 0, len(*o)*8)
	for _, ip := range *o {
		b = append(b, ip.Encode()...)
	}
	return b
}

func (o *IPv4Pair) Decode(b []byte) error {
	if err := validateMinimumSize(b, 8); err != nil {
		return err
	}
	*o = IPv4Pair{
		IPv4(b[:4]),
		IPv4(b[4:]),
	}
	return nil
}

func (o *IPv4Pair) String() string {
	s := make([]string, len(*o))
	for i, ip := range *o {
		s[i] = ip.String()
	}
	return strings.Join(s, " ")
}

func (o *IPv4Pairs) Encode() []byte {
	b := make([]byte, 0, len(*o)*8)
	for _, ip := range *o {
		b = append(b, ip.Encode()...)
	}
	return b
}

func (o *IPv4Pairs) Decode(b []byte) error {
	if err := validateSizeFactor(b, 8); err != nil {
		return err
	}
	if err := validateMinimumSize(b, 8); err != nil {
		return err
	}
	l := len(b) / 8
	*o = make([]IPv4Pair, l)
	for i := 0; i < l; i++ {
		err := (*o)[i].Decode(b[i*8 : i*8+8])
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *IPv4Pairs) String() string {
	s := make([]string, len(*o))
	for i, p := range *o {
		s[i] = p.String()
	}
	return strings.Join(s, ",")
}

func (o *Route) Encode() []byte {
	n, _ := o.Source.Mask.Size()
	l := (n + 7) >> 3
	b := make([]byte, 1, 5+l)
	b[0] = byte(n)
	b = append(b, o.Source.IP.To4()[:l]...)
	b = append(b, o.Destination.To4()...)
	return b
}

func (o *Route) Decode(b []byte) error {
	if err := validateMinimumSize(b, 5); err != nil {
		return err
	}
	n := int(b[0])
	l := (n + 7) >> 3
	if err := validateMinimumSize(b, 5+l); err != nil {
		return err
	}
	srcIP := make([]byte, 4)
	for i := 1; i <= l; i++ {
		srcIP[i-1] = b[i]
	}
	o.Source = net.IPNet{
		IP:   net.IPv4(srcIP[0], srcIP[1], srcIP[2], srcIP[3]),
		Mask: net.CIDRMask(n, 32),
	}
	o.Destination = net.IPv4(b[l+1], b[l+2], b[l+3], b[l+4])
	return nil
}

func (o *Route) String() string {
	return fmt.Sprintf("%s %s", o.Source.String(), o.Destination.String())
}

func (o *Routes) Encode() []byte {
	b := make([]byte, 0, 5)
	for _, r := range *o {
		b = append(b, r.Encode()...)
	}
	return b
}

func (o *Routes) Decode(b []byte) error {
	if err := validateMinimumSize(b, 5); err != nil {
		return err
	}
	*o = make([]Route, 0, 1)
	a := b[:]
	for len(a) > 0 {
		r := Route{}
		err := r.Decode(a)
		if err != nil {
			return err
		}
		*o = append(*o, r)
		n, _ := r.Source.Mask.Size()
		a = a[5+((n+7)>>3):]
	}
	return nil
}

func (o *Routes) String() string {
	s := make([]string, len(*o))
	for i, r := range *o {
		s[i] = r.String()
	}
	return strings.Join(s, ",")
}

func (o *DomainName) Encode() []byte {
	b := make([]byte, 0, 1)
	for _, s := range *o {
		b = append(b, byte(len(s)))
		b = append(b, s...)
	}
	b = append(b, 0)
	return b
}

func (o *DomainName) Decode(b []byte) error {
	if err := validateMinimumSize(b, 1); err != nil {
		return err
	}
	*o = make([]string, 0)
	a := b[:]
	for len(a) > 0 {
		l := a[0]
		a = a[1:]
		if l == 0 {
			break
		}
		*o = append(*o, string(a[:l]))
		a = a[l:]
	}
	return nil
}

func (o *DomainName) String() string {
	return strings.Join(*o, ".")
}

func (o *DomainNames) Encode() []byte {
	b := make([]byte, 0, 1)
	for _, dn := range *o {
		b = append(b, dn.Encode()...)
	}
	return b
}

func (o *DomainNames) Decode(b []byte) error {
	if err := validateMinimumSize(b, 1); err != nil {
		return err
	}
	*o = make([]DomainName, 0, 1)
	a := b[:]
	for len(a) > 0 {
		dn := DomainName{}
		if err := dn.Decode(a); err != nil {
			return err
		}
		*o = append(*o, dn)
		a = a[len(dn.Encode()):]
	}
	return nil
}

func (o *DomainNames) String() string {
	s := make([]string, len(*o))
	for i, dn := range *o {
		s[i] = dn.String()
	}
	return strings.Join(s, ",")
}

func (o *TimeOffset) Encode() []byte {
	t := int32(time.Duration(*o).Seconds())
	return []byte{
		byte(t >> 24),
		byte(t >> 16),
		byte(t >> 8),
		byte(t),
	}
}

func (o *TimeOffset) Decode(b []byte) error {
	if err := validateSize(b, 4); err != nil {
		return err
	}
	s := int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
	*o = TimeOffset(time.Duration(s) * time.Second)
	return nil
}

func (o *TimeOffset) String() string {
	return time.Duration(*o).String()
}

func (o *TimeDuration) Encode() []byte {
	t := uint32(time.Duration(*o).Seconds())
	return []byte{
		byte(t >> 24),
		byte(t >> 16),
		byte(t >> 8),
		byte(t),
	}
}

func (o *TimeDuration) Decode(b []byte) error {
	if err := validateSize(b, 4); err != nil {
		return err
	}
	s := int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
	*o = TimeDuration(time.Duration(s) * time.Second)
	return nil
}

func (o *TimeDuration) String() string {
	return time.Duration(*o).String()
}

func (o *Padding) Encode() []byte {
	return []byte{}
}

func (o *Padding) Decode(b []byte) error {
	return validateSize(b, 0)
}

func (o *Padding) String() string {
	return ""
}

func (o *End) Encode() []byte {
	return []byte{}
}

func (o *End) Decode(b []byte) error {
	return validateSize(b, 0)
}

func (o *End) String() string {
	return ""
}
