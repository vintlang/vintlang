package module

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	// "errors"
	"github.com/vintlang/vintlang/object"
)

// CryptoFunctions is a map that holds the available functions in the Crypto module.
var CryptoFunctions = map[string]object.ModuleFunction{
	"hashMD5":    hashMD5,
	"hashSHA256": hashSHA256,
	"encryptAES": encryptAES,
	"decryptAES": decryptAES,
}

// hashMD5 takes a string as input and returns the MD5 hash of that string.
// The MD5 hash is commonly used for checksums or for detecting duplicate data.
func hashMD5(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to hash"}
	}

	// Get the string value to hash
	str := args[0].Inspect()

	// Compute the MD5 hash of the string
	hash := md5.New()
	hash.Write([]byte(str))

	// Return the MD5 hash as a hexadecimal string
	return &object.String{Value: hex.EncodeToString(hash.Sum(nil))}
}

// hashSHA256 takes a string as input and returns the SHA-256 hash of that string.
// SHA-256 is a more secure cryptographic hash function than MD5.
func hashSHA256(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "We need one argument: the string to hash"}
	}

	// Get the string value to hash
	str := args[0].Inspect()

	// Compute the SHA-256 hash of the string
	hash := sha256.New()
	hash.Write([]byte(str))

	// Return the SHA-256 hash as a hexadecimal string
	return &object.String{Value: hex.EncodeToString(hash.Sum(nil))}
}

// encryptAES takes data and a key, encrypts the data using the AES algorithm, and returns the encrypted data as a hexadecimal string.
func encryptAES(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the data to encrypt and the key"}
	}

	// Get the data and key values
	data := args[0].Inspect()
	key := args[1].Inspect()

	// Encrypt the data using AES encryption
	encrypted, err := aesEncrypt([]byte(data), []byte(key))
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	// Return the encrypted data as a hexadecimal string
	return &object.String{Value: hex.EncodeToString(encrypted)}
}

// decryptAES takes encrypted data and a key, decrypts the data using the AES algorithm, and returns the decrypted plaintext.
func decryptAES(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "We need two arguments: the data to decrypt and the key"}
	}

	// Get the encrypted data and key values
	data := args[0].Inspect()
	key := args[1].Inspect()

	// Decrypt the data using AES decryption
	decrypted, err := aesDecrypt([]byte(data), []byte(key))
	if err != nil {
		return &object.Error{Message: err.Error()}
	}

	// Return the decrypted data as a string
	return &object.String{Value: string(decrypted)}
}

// aesEncrypt is a helper function that performs AES encryption on the given data using the provided key.
// It returns the encrypted data or an error if encryption fails.
func aesEncrypt(data, key []byte) ([]byte, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Prepare a buffer for the encrypted data
	ciphertext := make([]byte, len(data))

	// AES encryption uses an initialization vector (IV). In this example, we use a zero IV.
	iv := make([]byte, aes.BlockSize)

	// Create an AES encryption stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the data using the AES stream
	stream.XORKeyStream(ciphertext, data)

	return ciphertext, nil
}

// aesDecrypt is a helper function that performs AES decryption on the given encrypted data using the provided key.
// It returns the decrypted plaintext or an error if decryption fails.
func aesDecrypt(data, key []byte) ([]byte, error) {
	// Create a new AES cipher block using the provided key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Prepare a buffer for the decrypted data
	plaintext := make([]byte, len(data))

	// AES decryption uses an initialization vector (IV). In this example, we use a zero IV.
	iv := make([]byte, aes.BlockSize)

	// Create an AES decryption stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt the data using the AES stream
	stream.XORKeyStream(plaintext, data)

	return plaintext, nil
}


