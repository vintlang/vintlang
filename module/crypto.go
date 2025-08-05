package module

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

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
	if len(args) != 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"crypto", "hashMD5",
			"1 string argument (data to hash)",
			fmt.Sprintf("%d arguments or wrong type", len(args)),
			`crypto.hashMD5("hello") -> "5d41402abc4b2a76b9719d911017c592"`,
		)
	}
	str := args[0].(*object.String).Value
	hash := md5.Sum([]byte(str))
	return &object.String{Value: hex.EncodeToString(hash[:])}
}

// hashSHA256 takes a string as input and returns the SHA-256 hash of that string.
// SHA-256 is a more secure cryptographic hash function than MD5.
func hashSHA256(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return ErrorMessage(
			"crypto", "hashSHA256",
			"1 string argument (data to hash)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.hashSHA256("hello") -> "2cf24dff4f..."`,
		)
	}

	if args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"crypto", "hashSHA256",
			"string argument for data to hash",
			string(args[0].Type()),
			`crypto.hashSHA256("hello") -> "2cf24dff4f..."`,
		)
	}

	// Get the string value to hash
	str := args[0].(*object.String).Value

	// Compute the SHA-256 hash of the string
	hash := sha256.New()
	hash.Write([]byte(str))

	// Return the SHA-256 hash as a hexadecimal string
	return &object.String{Value: hex.EncodeToString(hash.Sum(nil))}
}

// encryptAES takes data and a key, encrypts the data using the AES algorithm, and returns the encrypted data as a hexadecimal string.
func encryptAES(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"crypto", "encryptAES",
			"2 string arguments (data to encrypt, encryption key)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.encryptAES("secret data", "mykey123") -> returns encrypted hex`,
		)
	}

	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			paramName := []string{"data", "key"}[i]
			return ErrorMessage(
				"crypto", "encryptAES",
				fmt.Sprintf("string argument for %s", paramName),
				string(arg.Type()),
				`crypto.encryptAES("secret data", "mykey123") -> returns encrypted hex`,
			)
		}
	}

	// Get the data and key values
	data := args[0].(*object.String).Value
	key := args[1].(*object.String).Value

	// Encrypt the data using AES encryption
	encrypted, err := aesEncrypt([]byte(data), []byte(key))
	if err != nil {
		return &object.Error{
			Message: fmt.Sprintf("\033[1;31mError in crypto.encryptAES()\033[0m:\n"+
				"  Failed to encrypt data: %v\n"+
				"  Please ensure your key is valid for AES encryption.\n"+
				"  Usage: crypto.encryptAES(\"secret data\", \"mykey123\") -> returns encrypted hex\n",
				err),
		}
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
