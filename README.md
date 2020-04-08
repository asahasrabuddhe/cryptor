# Cryptor

A tool to encrypt and decrypt text passed in the terminal or in a file. Cryptor implements AES-256-CBC encryption with PKCS padding. Encrypted text is stored as hex code and while decryption the cipher text is decode from hex to byte.

Features:
- [X] Encryptes plain text passed in the command line
- [X] Encrypts text of a file when filepath is set
- [X] Overwrites the file content in file encryption mode
- Generates key for encryption

### Usage
```
Usage:
  cryptor [command]

Available Commands:
  decrypt     Performs aes-256 decryption in cbc mode
  encrypt     Performs aes-256 encryption in cbc mode
  help        Help about any command

Flags:
  -h, --help   help for cryptor

Use "cryptor [command] --help" for more information about a command.
```

### Commands
- To encrypt plain text passed in the command
 ```
cryptor encrypt --key 579C73BDB1AA3045B87059358F69060E --text 'plain text to encrypt'
```
- To encrypt plain text passed in the command
```
./cryptor encrypt --key 579C73BDB1AA3045B87059358F69060E --filepath text.txt
```
- To decrypt plain text passed in the command
 ```
cryptor decrypt --key 579C73BDB1AA3045B87059358F69060E --text 08b9b8c2d22821c08e82cad3bdb70d06a9fd1bf84dfd7dad7f5f46ef9bf5c176e5
```
- To decrypt plain text passed in the command
```
./cryptor decrypt --key 579C73BDB1AA3045B87059358F69060E --filepath text.txt
```