package encrypt

// StringToByte converts a string to *[32]byte for use in encrypt or decrypt
func StringToByte(passphrase string) *[32]byte {
	var key [32]byte
	copy(key[:], passphrase)

	return &key
}
