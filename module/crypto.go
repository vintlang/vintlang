package module

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"

	// "errors"
	"github.com/vintlang/vintlang/object"
)

// CryptoFunctions is a map that holds the available functions in the Crypto module.
var CryptoFunctions = map[string]object.ModuleFunction{
	"hashMD5":       hashMD5,
	"hashSHA256":    hashSHA256,
	"encryptAES":    encryptAES,
	"decryptAES":    decryptAES,
	"generateRSA":   generateRSAKeyPair,
	"encryptRSA":    encryptRSA,
	"decryptRSA":    decryptRSA,
	"signRSA":       signRSA,
	"verifyRSA":     verifyRSA,
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
		return ErrorMessage(
			"crypto", "encryptAES",
			"valid data and key for AES encryption",
			fmt.Sprintf("error: %v", err),
			`crypto.encryptAES("secret data", "mykey123") -> returns encrypted hex`,
		)
	}

	// Return the encrypted data as a hexadecimal string
	return &object.String{Value: hex.EncodeToString(encrypted)}
}

// decryptAES takes encrypted data and a key, decrypts the data using the AES algorithm, and returns the decrypted plaintext.
func decryptAES(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"crypto", "decryptAES",
			"2 string arguments (data to decrypt, decryption key)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.decryptAES("encrypted data", "mykey123") -> returns decrypted plaintext`,
		)
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

// generateRSAKeyPair generates an RSA key pair with the specified bit size.
// Returns a dictionary with "private" and "public" keys in PEM format.
func generateRSAKeyPair(args []object.Object, defs map[string]object.Object) object.Object {
	// Default key size is 2048 bits
	keySize := 2048
	
	if len(args) > 1 {
		return ErrorMessage(
			"crypto", "generateRSA",
			"0 or 1 integer argument (key size in bits, defaults to 2048)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.generateRSA() -> {"private": "...", "public": "..."}`,
		)
	}
	
	if len(args) == 1 {
		if args[0].Type() != object.INTEGER_OBJ {
			return ErrorMessage(
				"crypto", "generateRSA",
				"integer argument for key size",
				string(args[0].Type()),
				`crypto.generateRSA(2048) -> {"private": "...", "public": "..."}`,
			)
		}
		keySize = int(args[0].(*object.Integer).Value)
		if keySize < 1024 || keySize > 4096 {
			return ErrorMessage(
				"crypto", "generateRSA",
				"key size between 1024 and 4096 bits",
				fmt.Sprintf("%d bits", keySize),
				`crypto.generateRSA(2048) -> {"private": "...", "public": "..."}`,
			)
		}
	}

	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to generate RSA key pair: %v", err)}
	}

	// Encode private key to PEM format
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// Encode public key to PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to marshal public key: %v", err)}
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	// Return as dictionary
	pairs := make(map[object.HashKey]object.DictPair)
	
	privateKeyStr := &object.String{Value: string(privateKeyPEM)}
	publicKeyStr := &object.String{Value: string(publicKeyPEM)}
	
	privateKeyKey := &object.String{Value: "private"}
	publicKeyKey := &object.String{Value: "public"}
	
	pairs[privateKeyKey.HashKey()] = object.DictPair{
		Key:   privateKeyKey,
		Value: privateKeyStr,
	}
	pairs[publicKeyKey.HashKey()] = object.DictPair{
		Key:   publicKeyKey,
		Value: publicKeyStr,
	}

	return &object.Dict{Pairs: pairs}
}

// encryptRSA encrypts data using RSA public key encryption.
func encryptRSA(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"crypto", "encryptRSA",
			"2 string arguments (data to encrypt, public key in PEM format)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.encryptRSA("hello", publicKey) -> "encrypted_hex_string"`,
		)
	}

	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			paramName := []string{"data", "public key"}[i]
			return ErrorMessage(
				"crypto", "encryptRSA",
				fmt.Sprintf("string argument for %s", paramName),
				string(arg.Type()),
				`crypto.encryptRSA("hello", publicKey) -> "encrypted_hex_string"`,
			)
		}
	}

	data := args[0].(*object.String).Value
	publicKeyPEM := args[1].(*object.String).Value

	// Parse public key
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return &object.Error{Message: "failed to decode PEM block containing public key"}
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to parse public key: %v", err)}
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return &object.Error{Message: "not an RSA public key"}
	}

	// Encrypt data
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPublicKey, []byte(data))
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to encrypt data: %v", err)}
	}

	return &object.String{Value: hex.EncodeToString(encryptedData)}
}

// decryptRSA decrypts data using RSA private key decryption.
func decryptRSA(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"crypto", "decryptRSA",
			"2 string arguments (encrypted data in hex, private key in PEM format)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.decryptRSA("encrypted_hex", privateKey) -> "decrypted_string"`,
		)
	}

	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			paramName := []string{"encrypted data", "private key"}[i]
			return ErrorMessage(
				"crypto", "decryptRSA",
				fmt.Sprintf("string argument for %s", paramName),
				string(arg.Type()),
				`crypto.decryptRSA("encrypted_hex", privateKey) -> "decrypted_string"`,
			)
		}
	}

	encryptedHex := args[0].(*object.String).Value
	privateKeyPEM := args[1].(*object.String).Value

	// Decode hex data
	encryptedData, err := hex.DecodeString(encryptedHex)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to decode hex data: %v", err)}
	}

	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return &object.Error{Message: "failed to decode PEM block containing private key"}
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to parse private key: %v", err)}
	}

	// Decrypt data
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to decrypt data: %v", err)}
	}

	return &object.String{Value: string(decryptedData)}
}

// signRSA creates a digital signature using RSA private key.
func signRSA(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return ErrorMessage(
			"crypto", "signRSA",
			"2 string arguments (data to sign, private key in PEM format)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.signRSA("hello", privateKey) -> "signature_hex_string"`,
		)
	}

	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			paramName := []string{"data", "private key"}[i]
			return ErrorMessage(
				"crypto", "signRSA",
				fmt.Sprintf("string argument for %s", paramName),
				string(arg.Type()),
				`crypto.signRSA("hello", privateKey) -> "signature_hex_string"`,
			)
		}
	}

	data := args[0].(*object.String).Value
	privateKeyPEM := args[1].(*object.String).Value

	// Parse private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return &object.Error{Message: "failed to decode PEM block containing private key"}
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to parse private key: %v", err)}
	}

	// Hash the data
	hashed := sha256.Sum256([]byte(data))

	// Create signature
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, 0, hashed[:])
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to create signature: %v", err)}
	}

	return &object.String{Value: hex.EncodeToString(signature)}
}

// verifyRSA verifies a digital signature using RSA public key.
func verifyRSA(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return ErrorMessage(
			"crypto", "verifyRSA",
			"3 string arguments (original data, signature in hex, public key in PEM format)",
			fmt.Sprintf("%d arguments", len(args)),
			`crypto.verifyRSA("hello", signature, publicKey) -> true/false`,
		)
	}

	for i, arg := range args {
		if arg.Type() != object.STRING_OBJ {
			paramName := []string{"data", "signature", "public key"}[i]
			return ErrorMessage(
				"crypto", "verifyRSA",
				fmt.Sprintf("string argument for %s", paramName),
				string(arg.Type()),
				`crypto.verifyRSA("hello", signature, publicKey) -> true/false`,
			)
		}
	}

	data := args[0].(*object.String).Value
	signatureHex := args[1].(*object.String).Value
	publicKeyPEM := args[2].(*object.String).Value

	// Decode signature
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to decode signature: %v", err)}
	}

	// Parse public key
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return &object.Error{Message: "failed to decode PEM block containing public key"}
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("failed to parse public key: %v", err)}
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return &object.Error{Message: "not an RSA public key"}
	}

	// Hash the data
	hashed := sha256.Sum256([]byte(data))

	// Verify signature
	err = rsa.VerifyPKCS1v15(rsaPublicKey, 0, hashed[:], signature)
	if err != nil {
		return &object.Boolean{Value: false}
	}

	return &object.Boolean{Value: true}
}
