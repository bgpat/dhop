package dhop

type Option struct {
	OptionData
	Code Code
}

type OptionData interface {
	Encode() []byte
	Decode([]byte) error
	Marshal() []byte
	Unmarshal([]byte) error
}
