# OpenSSL安全

以下都在/root/secure下进行。

## 准备环境

### 生成密钥

```
openssl genpkey -algorithm RSA -out privatekey.pem -des-ede3-cbc
```

命令执行后输出私钥文件 privatekey.pem，默认输出为PEM格式，密钥长度为1024。

接下来使用文本工具直接打开该文件，可以看到如下内容:
```
-----BEGIN ENCRYPTED PRIVATE KEY-----
MIICxjBABgkqhkiG9w0BBQ0wMzAbBgkqhkiG9w0BBQwwDgQIfGaocfOBhP0CAggA
MBQGCCqGSIb3DQMHBAhAyPxFcMAr2ASCAoBx9AlpAlXcrzHIMJeEy/B6cJcq2CKV
53ubvqADXZX+DtLy38tqAerVa5w2SGi1OSiXpb2ANNnC0Hwww9KafVl6p40tkjXq
db6M2pPqntLbCQgCDvF6OUDRTjXNrAEflSKtFHuPV0nuXJLnIx/Tq1yiShx3T0RX
FzVkQT+U3+iv8bZaReB9gYwnoV8HKnAhibQycPuldWC1H4TbrY4AwmRZ2pi8Rh7Y
hR9+vnQC+pDeHHXRbJTpR+YDh4mzqPc1CXbVE85xCro4z/DKYWdZsP+oNBgv1jka
EPECdznXBw0+sBXjBP1P9zxWy14MO6ViiiwLkedk+yP0gkn833J4M3h4/AKhCDZv
ejfRQtlCB7Z30QEfhlmJr7aWyDhWlcUEo7S27hU9BaiV76pKOKM/jjeYyczCgUOt
3fw9FNsRlcXbLRf1ptlAjSBHsTo8aPdhJN5Ev3s/LpK6Fxye9XD94tf67LzZXwWb
QuTr/wYbAdeOANdUynj6kgxdaef6POAzVUQwsCusQa4bGBD0aqVzwlAcmF652qaY
5OFuWPh3Z1+zUixZGMKjxXWiwERW1oP7x4eGFIPPlGjaYOS4SdjP8BYHWjVrjwDq
UvWHBukKMnEuNdW8ByvU1mVMNQ/BjU5MFeLwZqzh1xVk/U+x6Tnynk2xjyRSqVds
aq1ipzsJ233evfKABTp35o6A6xrJ3nv/U/FWX3Z5xFOPkwEW77jNaecu5PYm3dU2
oa01EtSaGp14qkSZVeGQNrUV1B8Q+iuWVCpCHU1QgYkUKxbjm70MHjzleLy8Y3T+
8+5kFRGLUyBYT0UOAyMcm2wqfx7M9FrYVFUrGeWZBr1TeNnJKSp3HmqS
-----END ENCRYPTED PRIVATE KEY-----
```

密钥经过PEM编码后在文件头尾都添加了标签，用以说明当前的文件格式。从标签内容也可以看出密钥是加密的，因为有“ENCRYPTED”。

中间的密钥内容是经过BASE64编码的，这样方便私钥的传递，例如在网络上传输，数据复制粘贴。

例子只是PEM文件格式的其中一种，以下是平时可能会碰到的PEM私钥格式：

#### PKCS#8 私钥加密格式

```
-----BEGIN ENCRYPTED PRIVATE KEY-----
BASE64私钥内容
-----ENDENCRYPTED PRIVATE KEY-----
```

#### PKCS#8 私钥非加密格式

```
-----BEGIN PRIVATE KEY-----
BASE64私钥内容
-----END PRIVATEKEY-----
```

#### Openssl ASN格式

```
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info:DES-EDE3-CBC,4D5D1AF13367D726
BASE64私钥内容
-----END RSA PRIVATE KEY-----
```

除了以上几种，还有微软的PVK格式；以及DER编码格式，就是在使用PEM编码前的数据，由于没有密码保护，平时很少直接使用。

Openssl ASN格式在加密私钥数据时只能用MD5算法生成key，而且只迭代计算了1次。

所以从1.0.0开始Openssl把PKCS#8格式作为默认格式，可以为私钥文件提供更好的安全性和扩展性。

### 解析密钥

可以使用openssl的pkey命令，执行后得到privatekey.der文件，这个是没有加密的私钥文件，数据是ASN.1格式，并使用DER编码。

```
openssl pkey -in privatekey.pem -out privatekey.der
```

privatekey.der文件：
```
-----BEGIN PRIVATE KEY-----
MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAOIsx5HXEwftSDUt
jSWZC4QEGL09XNUgoRe8oHnv7Evq/AeAHjO0CM+1tlXAuhBaXVPSbltxe6d6AVMA
rHDG/xqYXocxk/C60GSX83yEUAmw7HsMzzXWk48UIqqxt6tweTTd+nZ2r+5iia0O
loE8EyacD2C4Q5XmzZ3ldtF7UGz7AgMBAAECgYEA1FUhzqLRdQGoZnQrLH7vtrJ5
3z86ZcPKOJSXO6KofeVPUmNO908vboCzevICO9P6BfVqMWLqqaH6mEfrm7zWRAdS
5bVeTHIU7mSTEi0/xY6/BlgxNkjLJ01qYLQOJMoUYZglmpgUlJ7BpTUtp71Q843G
4tuBrN8k26ApZ4TLX3ECQQD6x8w8Nwo+S1gPyzIZd9vQdzQwSSNbelw5pFuV5EdW
r10oA133bAzwEKrneo3LIGUt0lwG8/WrG0eL+hpIiCuXAkEA5uHhD/V+8CbyA0EF
BYDQlqrUpCeLT8ty4M+WH8iJOQxTb3ec2tclh5VtV0nELaXeL/T/9LkI/dFVFrfP
KfuGPQJBAMY98Q5cuCU8bDW3/DezzzlBMilFd8TVulf6vEeGeHpnEC4UU1DGwod3
tZJdB//d8P3C/+qjKb4ER4+4utRBiKsCQQC5Bjt+tpZLjmpUAT8s8dY3aB+QjcAu
6jPdxX3haqvNc/tJUOn61n/U8AL3+L2md05f0E9upvp1rdQqqJL0agVhAkAue9iR
Q4TiNbNVhlnrXwdTLevGVzbFyckydFjws7axbxkQC7JsRPzVAZZEu9kEteUSoSvI
kZSeY6kOwvdMy9iY
-----END PRIVATE KEY-----
```
### 导出公钥

可以使用openssl命令导出公钥文件pubkey.pem。

```
openssl rsa -in privatekey.pem -pubout -out public.pem
```

使用文本工具打开公钥文件，pem头尾格式和私钥类似的标签：

```
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDiLMeR1xMH7Ug1LY0lmQuEBBi9
PVzVIKEXvKB57+xL6vwHgB4ztAjPtbZVwLoQWl1T0m5bcXunegFTAKxwxv8amF6H
MZPwutBkl/N8hFAJsOx7DM811pOPFCKqsbercHk03fp2dq/uYomtDpaBPBMmnA9g
uEOV5s2d5XbRe1Bs+wIDAQAB
-----END PUBLIC KEY-----
```

### 查看密钥信息

```
openssl rsa -noout -text -in private.pem
```

```
root@egaotan-VirtualBox:~/secure# openssl rsa -noout -text -in privatekey.pem
Enter pass phrase for privatekey.pem:
Private-Key: (1024 bit)
modulus:
    00:e2:2c:c7:91:d7:13:07:ed:48:35:2d:8d:25:99:
    0b:84:04:18:bd:3d:5c:d5:20:a1:17:bc:a0:79:ef:
    ec:4b:ea:fc:07:80:1e:33:b4:08:cf:b5:b6:55:c0:
    ba:10:5a:5d:53:d2:6e:5b:71:7b:a7:7a:01:53:00:
    ac:70:c6:ff:1a:98:5e:87:31:93:f0:ba:d0:64:97:
    f3:7c:84:50:09:b0:ec:7b:0c:cf:35:d6:93:8f:14:
    22:aa:b1:b7:ab:70:79:34:dd:fa:76:76:af:ee:62:
    89:ad:0e:96:81:3c:13:26:9c:0f:60:b8:43:95:e6:
    cd:9d:e5:76:d1:7b:50:6c:fb
publicExponent: 65537 (0x10001)
privateExponent:
    00:d4:55:21:ce:a2:d1:75:01:a8:66:74:2b:2c:7e:
    ef:b6:b2:79:df:3f:3a:65:c3:ca:38:94:97:3b:a2:
    a8:7d:e5:4f:52:63:4e:f7:4f:2f:6e:80:b3:7a:f2:
    02:3b:d3:fa:05:f5:6a:31:62:ea:a9:a1:fa:98:47:
    eb:9b:bc:d6:44:07:52:e5:b5:5e:4c:72:14:ee:64:
    93:12:2d:3f:c5:8e:bf:06:58:31:36:48:cb:27:4d:
    6a:60:b4:0e:24:ca:14:61:98:25:9a:98:14:94:9e:
    c1:a5:35:2d:a7:bd:50:f3:8d:c6:e2:db:81:ac:df:
    24:db:a0:29:67:84:cb:5f:71
prime1:
    00:fa:c7:cc:3c:37:0a:3e:4b:58:0f:cb:32:19:77:
    db:d0:77:34:30:49:23:5b:7a:5c:39:a4:5b:95:e4:
    47:56:af:5d:28:03:5d:f7:6c:0c:f0:10:aa:e7:7a:
    8d:cb:20:65:2d:d2:5c:06:f3:f5:ab:1b:47:8b:fa:
    1a:48:88:2b:97
prime2:
    00:e6:e1:e1:0f:f5:7e:f0:26:f2:03:41:05:05:80:
    d0:96:aa:d4:a4:27:8b:4f:cb:72:e0:cf:96:1f:c8:
    89:39:0c:53:6f:77:9c:da:d7:25:87:95:6d:57:49:
    c4:2d:a5:de:2f:f4:ff:f4:b9:08:fd:d1:55:16:b7:
    cf:29:fb:86:3d
exponent1:
    00:c6:3d:f1:0e:5c:b8:25:3c:6c:35:b7:fc:37:b3:
    cf:39:41:32:29:45:77:c4:d5:ba:57:fa:bc:47:86:
    78:7a:67:10:2e:14:53:50:c6:c2:87:77:b5:92:5d:
    07:ff:dd:f0:fd:c2:ff:ea:a3:29:be:04:47:8f:b8:
    ba:d4:41:88:ab
exponent2:
    00:b9:06:3b:7e:b6:96:4b:8e:6a:54:01:3f:2c:f1:
    d6:37:68:1f:90:8d:c0:2e:ea:33:dd:c5:7d:e1:6a:
    ab:cd:73:fb:49:50:e9:fa:d6:7f:d4:f0:02:f7:f8:
    bd:a6:77:4e:5f:d0:4f:6e:a6:fa:75:ad:d4:2a:a8:
    92:f4:6a:05:61
coefficient:
    2e:7b:d8:91:43:84:e2:35:b3:55:86:59:eb:5f:07:
    53:2d:eb:c6:57:36:c5:c9:c9:32:74:58:f0:b3:b6:
    b1:6f:19:10:0b:b2:6c:44:fc:d5:01:96:44:bb:d9:
    04:b5:e5:12:a1:2b:c8:91:94:9e:63:a9:0e:c2:f7:
    4c:cb:d8:98
```

### 测试文件

创建测试文件test，文件内容"123456789"。

## 加解密

### 公钥加密

```
openssl rsautl -encrypt -in test -out test.enc -inkey privatekey.pem
```

或者

```
openssl rsautl -encrypt -in test -out test.enc -inkey public.pem -pubin
```

### 私钥解密

```
openssl rsautl -decrypt -in test.enc -out test.dec -inkey privatekey.pem
```

可以验证test.dec文件内容和test文件内容一致。

## 签名及验证

### 私钥签名

```
openssl rsautl -sign -in test -out test.sig -inkey privatekey.pem
```

### 公钥验证

```
openssl rsautl -verify -in test.sig -out test.vfy -inkey public.pem -pubin
```

可以验证test.vfy文件内容和test文件内容一致。