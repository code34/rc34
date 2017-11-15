// Copyright 2009 The Go Authors. All rights reserved
// RC4 packaged Rework by code34 
// nicolas_boiteux@yahoo.fr - 2017-2018
// Add padding <-> Interface with OO_CIPHER on ARMA3

package rc34

// A Cipher is an instance of RC4 using a particular key.
type Cipher struct {
	s    [256]uint8
	i, j uint8
}

// NewCipher creates and returns a new Cipher. The key argument should be the
// RC4 key, at least 1 byte and at most 256 bytes.
func NewCipher(key []byte) (*Cipher, error) {
	newkey := make([]byte, 256, 256)
	
	for i := range key {
        		newkey[i] = uint8(key[i])
	}
	for i := len(key);i < 256;i++{
		newkey[i] = uint8(1)
	}

	var c Cipher
	for i := 0; i < 256; i++ {
		c.s[i] = uint8(i)
	}

	var j uint8 = 0 // equivalent d un mod 256
	for i := 0; i < 256; i++ {
		j += c.s[i] +  newkey[i]
		c.s[i], c.s[j] = c.s[j], c.s[i]
	}
	return &c, nil
}

// Reset zeros the key data so that it will no longer appear in the
// process's memory.
func (key *Cipher) Reset() {
	for i := range key.s {
		key.s[i] = 0
	}
	key.i, key.j = 0, 0
}


// xorKeyStreamGeneric sets dst to the result of XORing src with the
// key stream. Dst and src may be the same slice but otherwise should
// not overlap.
//
// This is the pure Go version. rc4_{amd64,386,arm}* contain assembly
// implementations. This is here for tests and to prevent bitrot.

func (key *Cipher) XorKeyStreamGeneric(dst, src []byte) {
	i := uint8(0)
	j := uint8(0)
	for index, value := range src {
		i += 1
		j += key.s[i]
		key.s[i], key.s[j] = key.s[j], key.s[i]
		dst[index] = value ^ key.s[key.s[i]+key.s[j]]
	}
	key.i, key.j = i, j
}

