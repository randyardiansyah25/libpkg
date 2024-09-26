package rsatool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"

	"golang.org/x/crypto/ssh"
)

func GenerateRSAKey(saveTo string, bitSize int) (err error) {
	savePrivateFileTo := saveTo
	savePublicFileTo := saveTo + ".pub"

	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		return
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	publicKeyBytes := encodePublicKeyToPEM(&privateKey.PublicKey)

	err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		return err
	}

	err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		return err
	}
	return
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {

	privDER, er := x509.MarshalPKCS8PrivateKey(privateKey)
	if er != nil {
		return nil
	}

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

func encodePublicKeyToPEM(privatekey *rsa.PublicKey) []byte {
	//publicRsaKey := x509.MarshalPKCS1PublicKey(privatekey)
	publicRsaKey, _ := x509.MarshalPKIXPublicKey(privatekey)

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicRsaKey,
	}
	publicPEM := pem.EncodeToMemory(&publicKeyBlock)
	return publicPEM
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := os.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func EncryptUsingPEM(text string, publicKeyFile string) (cipherText string, err error) {
	buf, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return
	}

	return Encrypt(text, buf)

}

func Encrypt(text string, publicKey []byte) (cipherText string, err error) {
	pemBlock, _ := pem.Decode(publicKey)
	pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return
	}

	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), []byte(text))
	if err != nil {
		return
	}

	cipherText = base64.StdEncoding.EncodeToString(cipherBytes)
	return
}

func DecryptUsingPem(cipherText string, privateKeyFile string) (text string, err error) {
	buf, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return "", err
	}

	return Decrypt(cipherText, buf)
}

func Decrypt(cipherText string, privateKey []byte) (text string, err error) {
	pemBlock, _ := pem.Decode(privateKey)
	priv, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return
	}

	chiperBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return
	}
	textBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), chiperBytes)
	text = string(textBytes)
	return
}
