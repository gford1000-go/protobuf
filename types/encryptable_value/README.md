## Encryptable_Value

The message `EncryptableValue` either contains a `Value` or an `EncryptedObject`, where its contents will decrypt and deserialise to a `Value`.

`NewEncryptableValue` provides a simple mechanism to create instances of `EncryptableValue`, using the provided `encryption.TokenKeyEncryptor` when encryption is required.

`NewEncryptableValueParser` likewise provides a simple mechanism to 
return the `Value` from an `EncryptableValue`, using the provided `encryption.TokenKeyDecryptor`.  This approach allows different behaviours (e.g. error on decryption failure, default value on decryption failure) to be implemented by calling logic, by providing different decryptors.
