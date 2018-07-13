package cipher

import (
	"strings"
)

const alph = 26

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type MyShift struct {
	shift int
}

type MyVigenere struct {
	key string
}

func NewCaesar() Cipher {
	return MyShift{3}
}

func NewShift(shift int) Cipher {
	if shift < 26 && shift > -26 && shift != 0 {
		return MyShift{shift}
	}
	return nil
}

func NewVigenere(key string) Cipher {
	if Validation(key) {
		return MyVigenere{key}
	}
	return nil
}

func Validation(key string) bool {
	if len(key) < 1 {
		return false
	}
	check := false
	for _, val := range key {
		if val < 'a' || val > 'z' {
			return false
		}
		if val != 'a' {
			check = true
		}
	}
	if check {
		return true
	}
	return false
}

func Cheking_key_length(str, key string) string {
	new_key := make([]byte, len(str))

	j := 0
	for i := range new_key {
		if j == len(key) {
			j = 0
		}
		new_key[i] = key[j]
		j++
	}
	return string(new_key)
}

func Clean_input_str(str string) string {
	str_1 := ""
	if len(str) < 1 {
		return str_1
	}
	str_2 := strings.ToLower(str)
	for i := range str_2 {
		if str_2[i] >= 'a' && str_2[i] <= 'z' {
			str_1 += string(str_2[i])
		}
	}
	return str_1
}

func cycle(char uint8, shift int) uint8 {
	char += uint8(shift)
	if char > 'z' {
		char -= alph
	} else if char < 'a' {
		char += alph
	}
	return char
}

func (shift MyShift) Encode(str string) string {
	str_1 := Clean_input_str(str)
	enc_str := ""
	for i := range str_1 {
		enc_str += string(cycle(str_1[i], shift.shift))
	}
	return enc_str
}

func (shift MyShift) Decode(str string) string {
	str_1 := Clean_input_str(str)
	dec_str := ""
	for i := range str_1 {
		dec_str += string(cycle(str_1[i], -shift.shift))
	}
	return dec_str
}

func (vig MyVigenere) Encode(str string) string {
	if len(str) < 1 {
		return ""
	}
	str_1 := Clean_input_str(str)
	new_key := Cheking_key_length(str_1, vig.key)
	enc_str := ""
	var shift int

	for i := range str_1 {
		shift = int(new_key[i] - byte(97))
		enc_str += string(cycle(str_1[i], shift))
	}
	return enc_str
}

func (vig MyVigenere) Decode(str string) string {
	if len(str) < 1 {
		return ""
	}
	str_1 := Clean_input_str(str)
	new_key := Cheking_key_length(str_1, vig.key)
	enc_str := ""
	var shift int

	for i := range str_1 {
		shift = int(new_key[i] - byte(97))
		enc_str += string(cycle(str_1[i], -shift))
	}
	return enc_str
}
