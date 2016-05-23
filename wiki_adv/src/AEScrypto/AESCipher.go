package AEScrypto

import (
    "io"
    "errors"
    "crypto/aes"
    "crypto/rand"
    "crypto/cipher"
    "encoding/base64"
)

// randomly generate a 32-byte key
func RandomKey() ([]byte, error) {
    b := make([]byte, 32)
    
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

// encrypt text with AES encryption
func AESEncrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    b := base64.StdEncoding.EncodeToString(text)
    encrypted := make([]byte, aes.BlockSize + len(b))
    iv := encrypted[:aes.BlockSize]

    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(encrypted[aes.BlockSize:], []byte(b))
    
    return encrypted, nil
}

// decrypt AES encrypted text
func AESDecrypt(key, text []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    if len(text) < aes.BlockSize {
        return nil, errors.New("encrypted message is too short")
    }

    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(text, text)
    
    data, err := base64.StdEncoding.DecodeString(string(text))
    if err != nil {
        return nil, err
    }

    return data, nil
}
