package dhop

type Option struct {
	OptionData
	Code byte
}

type OptionData interface {
	Encode() []byte
	Decode([]byte) error
	Marshal() []byte
	Unmarshal([]byte) error
}
