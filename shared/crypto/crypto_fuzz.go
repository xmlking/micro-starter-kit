// +build gofuzz

package crypto

import "log"

func Fuzz(data []byte) int {
    encrypted_data, err := AesEncrypt(string(data), "12345678123456781234567812345678")
    if err != nil {
        log.Panic("tried encrypt %v got err %v", encrypted_data, err)
    }

    decrypted_data, err := AesDecrypt(encrypted_data, "12345678123456781234567812345678")
    if err != nil {
        log.Panic("tried to encrypt/decrypt %v got err %v", data, err)
    }

    if decrypted_data != string(data) {
        log.Panic("decrypt(encrypt(%v)) != %v", data, data)
    }

    return 0
}

