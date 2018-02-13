package dhop

func Decode(code byte, b []byte) (Option, error) {
	var (
		o   OptionData
		err error
	)
	switch code {
	case 0:
		o = new(Padding)
		err = o.Decode(b)
	case 1, 16, 28, 32, 50, 54, 78, 95:
		o = new(IPv4)
		err = o.Decode(b)
	case 2:
		o = new(TimeOffset)
		err = o.Decode(b)
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 41, 42, 44, 45, 48, 49, 65,
		68, 69, 70, 71, 72, 73, 74, 75, 76, 92, 112, 118, 138, 150:
		o = new(IPv4s)
		err = o.Decode(b)
	case 13, 22, 26, 57, 93:
		o = new(Size)
		err = o.Decode(b)
	case 19, 20, 27, 29, 30, 31, 34, 36, 39:
		o = new(Boolean)
		err = o.Decode(b)
	case 21, 33:
		o = new(IPv4Pair)
		err = o.Decode(b)
	case 23, 37, 46, 52, 53, 116:
		o = new(Byte)
		err = o.Decode(b)
	case 24, 35, 38, 51, 58, 59, 91:
		o = new(TimeDuration)
		err = o.Decode(b)
	case 25:
		o = new(Sizes)
		err = o.Decode(b)
	case 121, 249:
		o = new(Routes)
		err = o.Decode(b)
	case 119:
		o = new(DomainNames)
		err = o.Decode(b)
	case 255:
		o = new(End)
		err = o.Decode(b)
	default:
		o = new(String)
		err = o.Decode(b)
	}
	return Option{
		OptionData: o,
		Code:       Code(code),
	}, err
}

func Unmarshal(code byte, b []byte) (Option, error) {
	var (
		o   OptionData
		err error
	)
	switch code {
	case 0:
		o = new(Padding)
		err = o.Unmarshal(b)
	case 1, 16, 28, 32, 50, 54, 78, 95:
		o = new(IPv4)
		err = o.Unmarshal(b)
	case 2:
		o = new(TimeOffset)
		err = o.Unmarshal(b)
	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 41, 42, 44, 45, 48, 49, 65,
		68, 69, 70, 71, 72, 73, 74, 75, 76, 92, 112, 118, 138, 150:
		o = new(IPv4s)
		err = o.Unmarshal(b)
	case 13, 22, 26, 57, 93:
		o = new(Size)
		err = o.Unmarshal(b)
	case 19, 20, 27, 29, 30, 31, 34, 36, 39:
		o = new(Boolean)
		err = o.Unmarshal(b)
	case 21, 33:
		o = new(IPv4Pair)
		err = o.Unmarshal(b)
	case 23, 37, 46, 52, 53, 116:
		o = new(Byte)
		err = o.Unmarshal(b)
	case 24, 35, 38, 51, 58, 59, 91:
		o = new(TimeDuration)
		err = o.Unmarshal(b)
	case 25:
		o = new(Sizes)
		err = o.Unmarshal(b)
	case 121, 249:
		o = new(Routes)
		err = o.Unmarshal(b)
	case 119:
		o = new(DomainNames)
		err = o.Unmarshal(b)
	case 255:
		o = new(End)
		err = o.Unmarshal(b)
	default:
		o = new(String)
		err = o.Unmarshal(b)
	}
	return Option{
		OptionData: o,
		Code:       Code(code),
	}, err
}
