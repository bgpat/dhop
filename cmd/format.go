package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	FORMAT_TYPE_DEFAULT = ""
	FORMAT_TYPE_BINARY  = "binary"
	FORMAT_TYPE_JSON    = "json"
	FORMAT_TYPE_HEX     = "hex"
	FORMAT_TYPE_BASE64  = "base64"
)

type formatType string

func (f *formatType) String() string {
	return string(*f)
}

func (f *formatType) Set(v string) error {
	switch formatType(v) {
	case FORMAT_TYPE_BINARY, FORMAT_TYPE_JSON, FORMAT_TYPE_HEX, FORMAT_TYPE_BASE64:
		*f = formatType(v)
	default:
		return fmt.Errorf("invalid format argument \"%s\"", v)
	}
	return nil
}

func (f *formatType) Type() string {
	return "{binary,json,hex,base64}"
}

func (f *formatType) Encode(src []byte) (string, error) {
	switch *f {
	case FORMAT_TYPE_BINARY:
		return string(src), nil
	case FORMAT_TYPE_JSON:
		bin := make([]uint, len(src))
		for i, n := range src {
			bin[i] = uint(n)
		}
		dst, err := json.Marshal(bin)
		if err != nil {
			return "", err
		}
		return string(dst), nil
	case FORMAT_TYPE_HEX:
		return strings.Join(strings.Split(fmt.Sprintf("% x", src), " "), separator), nil
	case FORMAT_TYPE_BASE64:
		return base64.StdEncoding.EncodeToString(src), nil
	}
	return "", fmt.Errorf("invalid input \"%s\"", src)
}

func (f *formatType) Decode(src []byte) ([]byte, error) {
	if *f == FORMAT_TYPE_BINARY && noTrimSpace {
		return src, nil
	}
	src = bytes.TrimSpace(src)
	switch *f {
	case FORMAT_TYPE_BINARY:
		return src, nil
	case FORMAT_TYPE_JSON:
		buf := []byte{}
		if err := json.Unmarshal(src, &buf); err != nil {
			return nil, err
		}
		return buf, nil
	case FORMAT_TYPE_HEX:
		var arr [][]byte
		if separator == "" {
			if len(src)%2 == 1 {
				return nil, fmt.Errorf("invalid input \"%s\"", src)

			}
			arr = make([][]byte, len(src)/2)
			for i := 0; i < len(src)/2; i++ {
				arr[i] = src[i*2 : (i+1)*2]
			}
		} else {
			arr = bytes.Split(src, []byte(separator))
		}
		buf := make([]byte, 0, len(src)/2)
		for _, s := range arr {
			n, err := strconv.ParseUint(string(s), 16, 8)
			if err != nil {
				return nil, err
			}
			buf = append(buf, byte(n))
		}
	case FORMAT_TYPE_BASE64:
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
		dst = make([]byte, len(src))
		n, err := base64.StdEncoding.Decode(dst, src)
		if err != nil {
			return nil, err
		}
		return dst[:n], nil
	}
	return nil, fmt.Errorf("invalid input \"%s\"", src)
}
