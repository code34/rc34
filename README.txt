	/*
	Copyright 2009 The Go Authors. All rights reserved
	2017-2018 - Hardened by code34 nicolas_boiteux@yahoo.fr 

	Package RC34

	RC4 reference:
	https://fr.wikipedia.org/wiki/RC4

	==================================================

	import "rc34"

	Generate a new key from the passphrase
	newkey, err := rc34.NewCipher([]byte)

	Cipher / Uncipher the stripe
	newkey.XorKeyStreamGeneric(dest []byte, src []byte)

	Reset the key in memory
	newkey.Reset()

	*/