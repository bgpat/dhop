package main

import (
	"fmt"
	"strconv"
	"strings"
)

type codeRange struct {
	From byte
	To   byte
}

type codeRanges []codeRange

func (r *codeRange) String() string {
	count := r.To - r.From
	from := strconv.Itoa(int(r.From))
	to := strconv.Itoa(int(r.To))
	switch count {
	case 0:
		return from
	case 1:
		return from + "," + to
	default:
		return from + "-" + to
	}
}

func (r *codeRanges) String() string {
	s := make([]string, len(*r))
	for i, codes := range *r {
		s[i] = codes.String()
	}
	return strings.Join(s, ",")
}

func (r *codeRange) Set(s string) error {
	a := strings.SplitN(s, "-", 2)
	switch len(a) {
	case 0:
		return fmt.Errorf("invalid range argument \"%s\"", s)
	case 1:
		v, err := strconv.ParseUint(a[0], 10, 8)
		if err != nil {
			return err
		}
		r.From = byte(v)
		r.To = byte(v)
	case 2:
		v, err := strconv.ParseUint(a[0], 10, 8)
		if err != nil {
			return err
		}
		r.From = byte(v)
		v, err = strconv.ParseUint(a[1], 10, 8)
		if err != nil {
			return err
		}
		r.To = byte(v)
		if r.From > r.To {
			return fmt.Errorf("invalid range argument \"%s\"", s)
		}
	}
	return nil
}

func (r *codeRanges) Set(s string) error {
	a := strings.Split(strings.TrimSpace(s), ",")
	if len(a) == 0 {
		return fmt.Errorf("invalid range argument \"%s\"", s)
	}
	*r = codeRanges{}
	for _, s := range a {
		v := codeRange{}
		if err := v.Set(s); err != nil {
			return err
		}
		*r = append(*r, v)
	}
	return nil
}

func (r *codeRanges) Type() string {
	return "range"
}

func (r *codeRange) Slice() []byte {
	count := int(r.To-r.From) + 1
	s := make([]byte, count)
	for i := 0; i < count; i++ {
		s[i] = r.From + byte(i)
	}
	return s
}

func (r *codeRanges) Slice() []byte {
	tmp := make(map[byte]struct{})
	for _, codes := range *r {
		for _, code := range codes.Slice() {
			tmp[code] = struct{}{}
		}
	}
	s := make([]byte, 0, 256)
	for i := byte(0); ; i++ {
		if _, ok := tmp[i]; ok {
			s = append(s, i)
		}
		if i == 255 {
			break
		}
	}
	return s
}
