package dhcpop

func Decode(code byte, b []byte) (Option, error) {
	switch code {
	case 0:
		o := new(Padding)
		err := o.Decode(b)
		return o, err
	case 1, 16, 28, 32, 50, 54, 78, 95:
		o := new(IPv4)
		err := o.Decode(b)
		return o, err
	case 2:
		o := new(TimeOffset)
		err := o.Decode(b)
		return o, err
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 41, 42, 44, 45, 48, 49, 65, 68, 69, 70, 71, 72, 73, 74, 75, 76, 92, 112, 118, 138, 150: // []net.IP
		o := new(IPv4s)
		err := o.Decode(b)
		return o, err
	case 13, 22, 26, 57, 93:
		o := new(Size)
		err := o.Decode(b)
		return o, err
	case 19, 20, 27, 29, 30, 31, 34, 36, 39:
		o := new(Boolean)
		err := o.Decode(b)
		return o, err
	case 21, 33:
		o := new(IPv4Pair)
		err := o.Decode(b)
		return o, err
	case 23, 37, 46, 52, 53, 116:
		o := new(Byte)
		err := o.Decode(b)
		return o, err
	case 24, 35, 38, 51, 58, 59, 91:
		o := new(TimeDuration)
		err := o.Decode(b)
		return o, err
	case 25:
		o := new(Sizes)
		err := o.Decode(b)
		return o, err
	case 121, 249:
		o := new(Routes)
		err := o.Decode(b)
		return o, err
	case 119:
		o := new(DomainNames)
		err := o.Decode(b)
		return o, err
	case 255:
		o := new(Padding)
		err := o.Decode(b)
		return o, err
	default:
		o := new(String)
		err := o.Decode(b)
		return o, err
	}
}
