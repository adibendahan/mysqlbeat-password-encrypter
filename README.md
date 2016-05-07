# mysqlbeat-password-encrypter
password encrypter for [mysqlbeat](https://github.com/adibendahan/mysqlbeat)

##How to use
after building your own mysqlbeat with a unique secret, just paste the secret from your mysqlbeat to encrypt.go, build and run.

```
$ go build
$ ./mysqlbeat-password-encrypter 
  Enter password to encrypt: mysqlbeat_pass
  Encrypted password: 2321f38819cf693951e88f00cd82
```

## License
GNU General Public License v2
