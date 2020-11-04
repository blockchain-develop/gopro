# OpenSSL证书

我们下面用OpensSSL生成一套证书系统，包括根证书和两级证书系统，验证证书，查看证书内容。

## 准备环境

以下均使用root用户，在目录/root/cert下进行。

## 生成根证书

### 生成根证书的密钥

```
openssl genrsa -des3 -out keys/RootCA.key 2048
```

openssl的命令手册请查阅[OpenSSL commands](https://www.openssl.org/docs/man1.1.1/man1/)。

上述命令解释：

+ genrsa: 生成一个RSA的密钥，openssl还支持其他的非对称加密算法，包括RSA、DSA、DH等
+ des3: 使用的RSA算法，支持的RSA算法有-aes128, -aes192, -aes256, -aria128, -aria192, -aria256, -camellia128, -camellia192, -camellia256, -des, -des3, -idea
+ out: 密钥输出文件
+ 2048: 密钥长度

在生成上述RootCA.key文件时，需要用户输入密码。

RooCA.key文件：
```
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,889DCDB7A65D4ED7

fBBNgLPb5kQ8ZCdkJqoqVBeXXxDjgIqBnSpqnYA5qSVqP1QhpZduG7ucMwp52vT1
b3tK3MWCzzB5OgO0mT3uuDFT9tw0ckB63UJWJGxQHO4IjAeu0/aoNEFQzK8PFYWX
6BOq6AAHELJ5ZpAQJpl+dIpOGE9k8wybesMXOKIZGmNdZWUFwdB6FWjiFkQGQ17m
+cWqNsnFJQp1EmmfO8Wd4tzZzXK06bDgXgDwT6sEjxS3+4hRSeGfxyaGj6NNRIH2
SobzPI22Oc2MWhF6MMYCXV4yVimG74Edj6OSGM7aVofKlWOhrjykAkCURZh21XB/
uVPAsLvh0K1dqTTNGNVMK0ObdTLeva2VGLztSkR76MJMOdCjTOxL0lQwxFdj1W8p
AG/FUDs5PTyalrtFiIm7XLaoD7JKxuPNUVO8q6lWx1sje54kD0p2NChWmDBzJIOy
FaYy4z0J+Uhg3dFf1YgAuVaBOfL3ciF7I0exZYCcfTBnvleM1lPTLtPI4S7uKJ0/
u6BoCSvaItk90tk9aCKDpbYeCVU1dj6jWlwBlIebZeWqfO3mPQBekLyrqLgG6MaG
AQXqpXLycEUIV9bF/XZLTVbyC9i2ZsAAaO3SgOXR+6CLR5y4SuAL0jtNYp99rUP2
xkFmCCHfcpi3f7vz6LtJ6Z3YCe2OQJw7AcBG+VKR4c0OiolCyr6HAoxaN15KV0Z2
6gclDKkAzhwGrv/Hszto/My9GPTa3XuI2zgwSKEWvGj9aDZ6jkmCoj0+FJQf0ywY
Q++e3Wt4M7lSjHY+sgT5F5wpxmyPd61H4/BbCrky0/18b4b0N7J0L+1efT3UKml5
tElGrUEqGeuZ1HGbhpQtpnWFzgukcVCRW1LT+CC8cUo4DVlAvvipFKr7EI7j1fPW
ZCXyiSSVAm4AhFgsLOsXfCLLH/qf0qULqZY3tNN7SHeyesmd1lXahPd0H06QbySr
as7DT4F6pOE5fVFDjspT4/8b/LHBAL8QDHK+gioS0LKLGofumifwImUSDr2TO/bB
EE47P6eWUP3RyEXG5I6rKdVtHnQh/g61Z4R0HfLUYvRd1ehEyravSu0CrdGtHcRE
rZyv6G5pkKmi4ij4QpuScOG8icF6xRBGZ48U4oNTK5mHCsG4F69apEuGmGlac4mP
+F6qFKfa/lT9IoLdYRKbJu/ygIUcNR2xBdjRGycqvvgej81LLsFn/s16xVXjfZD0
0/8kwHa7laGZZCTHhpYI/Q8wtdKcgHWyGmTDNOcttii0rAgkz0zuKx51KE23eEgf
v9qnS2wKJkgKbHSATxR2hbclKhiGHxOMNR3xLNlkPEhvwIJburKlm3DYdJyemB8z
JwjxWL5ilcvtFjSQ9bplJyDZ2qEXM25lm4K6K4pOwZs3z40AUJV6prmYlS8cR9uH
GLKl/V/dhstb1itLO611edKpNMuK6mgttUbANeGP2XiDsBCKxJBJQxnpoNVhT5ZJ
IxeVOQzuAsxtUsWOtZO2rETv9q3ayHMToDJ0vMZ/TX/+SdrISNYGbqaFjze6zqVo
DRgcp4qHcWi82ax28xo9d04xfi5DxKqg02LQu5GOi4di57ueNV0azg==
-----END RSA PRIVATE KEY-----
```

密钥文件经过编码后在文件头尾都添加了标签，用于说明当前文件的格式。

上面只是密钥文件格式的一种，以下时可能会遇到的密钥格式：

### PKCS#8密钥加密格式
```
-----BEGIN ENCRYPTED PRIVATE KEY-----  
BASE64私钥内容  
-----ENDENCRYPTED PRIVATE KEY-----  
```

### PKCS#8密钥非加密格式
```
-----BEGIN PRIVATE KEY-----  
BASE64私钥内容  
-----END PRIVATEKEY-----  
```

### OpenSSL ASN格式
```
-----BEGIN RSA PRIVATE KEY-----  
Proc-Type: 4,ENCRYPTED  
DEK-Info:DES-EDE3-CBC,4D5D1AF13367D726
BASE64私钥内容  
-----END RSA PRIVATE KEY-----  
```

除了上面几种，还有PVK、DER等格式。

OpenSSL ASN格式在加密密钥数据时只能用MD5算法生成Key，而且只迭代计算了1次。

从OpenSSL1.0.0开始PKCS#8密钥加密格式作为默认格式，为私钥提供更好的安全性和扩展性。

### 解析私钥文件
```
openssl asn1parse -i -in keys/RootCA_1.key
```

```
root@egaotan-VirtualBox:~/cert# openssl asn1parse -i -in keys/RootCA_1.key
    0:d=0  hl=4 l=1187 cons: SEQUENCE          
    4:d=1  hl=2 l=   1 prim:  INTEGER           :00
    7:d=1  hl=4 l= 257 prim:  INTEGER           :B3ACCB3D5D9DC9543175F77C73831D8D70C93103893876D7E09EBC41AA6A712493600648EA983F19387E1C16BCCAF90C4FDD03141EC5E87C5B099DD974D977530FAF11E72AE4E15F0FABFE74427A1ABA7C131B5D513333DB96975F1238D89C4BB8E99CCD5E3693541D5331E24B2CB31973B517F48F13D9D9557BAED7FA158E5599372C5BD8111E003825F86D24A714C7E35070EA39EC6009C0252F06671C1B3B4C12DCA923109DF1699B86A590781EF4878BADEC1CCEDD3285E4B0F6F7B492075D839154D5C639025923F9397222CE0E8089CA2E51ECCED83CCE4DF3F5A483720436022E534835641516A102A56AAA077B04CC7B6DACA65CA2E61008B674EAEF
  268:d=1  hl=2 l=   3 prim:  INTEGER           :010001
  273:d=1  hl=4 l= 256 prim:  INTEGER           :598A3F2EB1BCB45C74D9A9202820AB5595636E3DA8E0AEB91E837CA5ECDB66CF2772AF2F1F5A07F7F5A343344199954B81714EDB740A84B59458D285502A1FC1110BDA0C2DE02A40497B192FAF0152D8A484911A8C20CC331E750937F7845B3F91433A954DE23CC1AA6B8F6C8AB217FF131EA87583E80CA4B66910D529181CF9CC8E530A2C9112B92F2967BBBF08958E7D2599F14F596408836886B0A5EE174E4CFD76E6EB047EFC134BA54F032BF3630285629195FC68235D35E364044D7DE3988A94B5F894609E273445A08983459A2516D1C1F62D3CA80E5676479DAE07F1D24F55C170302CBC1E25C88D62BB6C9328F46A6648941E24B6842C9118CE3701
  533:d=1  hl=3 l= 129 prim:  INTEGER           :EE74661D84D24B28785A66B87BFAD7C91F2F4176AD19D63AF827969395A4619F9AA9007E49FABC8BDE379A7C511F674DE0103FDA9C34F80C995629A505DF4E0AA4578FA80B8693B2A7D5349291B21C369755B7CC353A791ACD591192125BBF6857EBB59938FBA464AD4EC15693F78D22DED9006117A9DED8A1D0A7FEBCE1288F
  665:d=1  hl=3 l= 129 prim:  INTEGER           :C0E5341A0566A5DD48B1E641B223907F12FB8981B63009EF145BF3DD9A53B8C2314E3D13D9D7B0EB7FDABF360126500148223097B0B68109EA3A0ED87BF6CBE5387B11E15C05F23B88A79DF49BBF1DC82685E1B77CDD14237F95CEA120B94FBFD97E202E177EF9AF264CD272925E44B1059596A7D77330EB46E68A545D9987A1
  797:d=1  hl=3 l= 129 prim:  INTEGER           :9F81CF5909A31E51690897D3EA63267B4BD7F18370516B23B121D230D115CF93D3021FE95A9556107DCA5615B8B2380938CBC5DAF748DB709A15AFEFCA98D07C7FE86E992B748F521FC097D53F8E941ED466AB76F75830F083D283D561606C0E6159B3F555CFC6ECED9D53E9153BAFF65679742E7E84CC23BF42734E0C8CD4F3
  929:d=1  hl=3 l= 128 prim:  INTEGER           :301DFCAD3103B6044909F78213C9C355AF8A87768FD7D3E28B95947386F7B372DD91C2B29CF8ACED51EB631F3992310AE0CF8687905136471EE274A993E2B061F180E1A8F1A79A1137B978317858683971429C57851230DFDA07A88F90F628EA967A5BD4A38FCD00DFF0F1DE1A9CE14DB732E56DE0D49F33517058B94ACFED81
 1060:d=1  hl=3 l= 128 prim:  INTEGER           :106A0A7EAB9A2934CA7DAF1C51A30BCB69F0480A7CEA3F9E389FA095B5EC3401E7FD6A66EFAA43A557A2ED8EA92CF6AD37A245459FAFB0C52E7AC0648085CDDB139549A072B39A95431FD31F8903338881A9748C0AAB05AEA0E4CFFFC40C3C81938C8DB27550DC6D7D7901AF9EE094A4C9122FC62909F67B033715A8F03ED167
```

从加密密钥文件中提却密钥:
```
openssl pkey -in RootCA.key -out RootCA.der
```
这一解密过程使用openssl的pkey命令完成，执行完成后得到RootCA.der文件，这是没有加密的私钥文件，数据是ASN.1格式，并使用DER编码。

然后再次分析密钥数据，由于输入是der格式，需要使用inform参数说明：
```
openssl asn1parse -in RootCA.der
```

```
root@egaotan-VirtualBox:~/cert/keys# openssl asn1parse -in RootCA.der
    0:d=0  hl=4 l=1213 cons: SEQUENCE          
    4:d=1  hl=2 l=   1 prim: INTEGER           :00
    7:d=1  hl=2 l=  13 cons: SEQUENCE          
    9:d=2  hl=2 l=   9 prim: OBJECT            :rsaEncryption
   20:d=2  hl=2 l=   0 prim: NULL              
   22:d=1  hl=4 l=1191 prim: OCTET STRING      [HEX DUMP]:308204A30201000282010100B3ACCB3D5D9DC9543175F77C73831D8D70C93103893876D7E09EBC41AA6A712493600648EA983F19387E1C16BCCAF90C4FDD03141EC5E87C5B099DD974D977530FAF11E72AE4E15F0FABFE74427A1ABA7C131B5D513333DB96975F1238D89C4BB8E99CCD5E3693541D5331E24B2CB31973B517F48F13D9D9557BAED7FA158E5599372C5BD8111E003825F86D24A714C7E35070EA39EC6009C0252F06671C1B3B4C12DCA923109DF1699B86A590781EF4878BADEC1CCEDD3285E4B0F6F7B492075D839154D5C639025923F9397222CE0E8089CA2E51ECCED83CCE4DF3F5A483720436022E534835641516A102A56AAA077B04CC7B6DACA65CA2E61008B674EAEF020301000102820100598A3F2EB1BCB45C74D9A9202820AB5595636E3DA8E0AEB91E837CA5ECDB66CF2772AF2F1F5A07F7F5A343344199954B81714EDB740A84B59458D285502A1FC1110BDA0C2DE02A40497B192FAF0152D8A484911A8C20CC331E750937F7845B3F91433A954DE23CC1AA6B8F6C8AB217FF131EA87583E80CA4B66910D529181CF9CC8E530A2C9112B92F2967BBBF08958E7D2599F14F596408836886B0A5EE174E4CFD76E6EB047EFC134BA54F032BF3630285629195FC68235D35E364044D7DE3988A94B5F894609E273445A08983459A2516D1C1F62D3CA80E5676479DAE07F1D24F55C170302CBC1E25C88D62BB6C9328F46A6648941E24B6842C9118CE370102818100EE74661D84D24B28785A66B87BFAD7C91F2F4176AD19D63AF827969395A4619F9AA9007E49FABC8BDE379A7C511F674DE0103FDA9C34F80C995629A505DF4E0AA4578FA80B8693B2A7D5349291B21C369755B7CC353A791ACD591192125BBF6857EBB59938FBA464AD4EC15693F78D22DED9006117A9DED8A1D0A7FEBCE1288F02818100C0E5341A0566A5DD48B1E641B223907F12FB8981B63009EF145BF3DD9A53B8C2314E3D13D9D7B0EB7FDABF360126500148223097B0B68109EA3A0ED87BF6CBE5387B11E15C05F23B88A79DF49BBF1DC82685E1B77CDD14237F95CEA120B94FBFD97E202E177EF9AF264CD272925E44B1059596A7D77330EB46E68A545D9987A1028181009F81CF5909A31E51690897D3EA63267B4BD7F18370516B23B121D230D115CF93D3021FE95A9556107DCA5615B8B2380938CBC5DAF748DB709A15AFEFCA98D07C7FE86E992B748F521FC097D53F8E941ED466AB76F75830F083D283D561606C0E6159B3F555CFC6ECED9D53E9153BAFF65679742E7E84CC23BF42734E0C8CD4F3028180301DFCAD3103B6044909F78213C9C355AF8A87768FD7D3E28B95947386F7B372DD91C2B29CF8ACED51EB631F3992310AE0CF8687905136471EE274A993E2B061F180E1A8F1A79A1137B978317858683971429C57851230DFDA07A88F90F628EA967A5BD4A38FCD00DFF0F1DE1A9CE14DB732E56DE0D49F33517058B94ACFED81028180106A0A7EAB9A2934CA7DAF1C51A30BCB69F0480A7CEA3F9E389FA095B5EC3401E7FD6A66EFAA43A557A2ED8EA92CF6AD37A245459FAFB0C52E7AC0648085CDDB139549A072B39A95431FD31F8903338881A9748C0AAB05AEA0E4CFFFC40C3C81938C8DB27550DC6D7D7901AF9EE094A4C9122FC62909F67B033715A8F03ED167
```

### 转换密钥格式

```
openssl rsa -in keys/RootCA.key -out keys/RootCA_1.key
```

RootCA_1.key文件：
```
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAs6zLPV2dyVQxdfd8c4MdjXDJMQOJOHbX4J68QapqcSSTYAZI
6pg/GTh+HBa8yvkMT90DFB7F6HxbCZ3ZdNl3Uw+vEecq5OFfD6v+dEJ6Grp8Extd
UTMz25aXXxI42JxLuOmczV42k1QdUzHiSyyzGXO1F/SPE9nZVXuu1/oVjlWZNyxb
2BEeADgl+G0kpxTH41Bw6jnsYAnAJS8GZxwbO0wS3KkjEJ3xaZuGpZB4HvSHi63s
HM7dMoXksPb3tJIHXYORVNXGOQJZI/k5ciLODoCJyi5R7M7YPM5N8/Wkg3IENgIu
U0g1ZBUWoQKlaqoHewTMe22splyi5hAItnTq7wIDAQABAoIBAFmKPy6xvLRcdNmp
ICggq1WVY249qOCuuR6DfKXs22bPJ3KvLx9aB/f1o0M0QZmVS4FxTtt0CoS1lFjS
hVAqH8ERC9oMLeAqQEl7GS+vAVLYpISRGowgzDMedQk394RbP5FDOpVN4jzBqmuP
bIqyF/8THqh1g+gMpLZpENUpGBz5zI5TCiyRErkvKWe7vwiVjn0lmfFPWWQIg2iG
sKXuF05M/Xbm6wR+/BNLpU8DK/NjAoVikZX8aCNdNeNkBE1945iKlLX4lGCeJzRF
oImDRZolFtHB9i08qA5Wdkedrgfx0k9VwXAwLLweJciNYrtskyj0amZIlB4ktoQs
kRjONwECgYEA7nRmHYTSSyh4Wma4e/rXyR8vQXatGdY6+CeWk5WkYZ+aqQB+Sfq8
i943mnxRH2dN4BA/2pw0+AyZVimlBd9OCqRXj6gLhpOyp9U0kpGyHDaXVbfMNTp5
Gs1ZEZISW79oV+u1mTj7pGStTsFWk/eNIt7ZAGEXqd7YodCn/rzhKI8CgYEAwOU0
GgVmpd1IseZBsiOQfxL7iYG2MAnvFFvz3ZpTuMIxTj0T2dew63/avzYBJlABSCIw
l7C2gQnqOg7Ye/bL5Th7EeFcBfI7iKed9Ju/HcgmheG3fN0UI3+VzqEguU+/2X4g
Lhd++a8mTNJykl5EsQWVlqfXczDrRuaKVF2Zh6ECgYEAn4HPWQmjHlFpCJfT6mMm
e0vX8YNwUWsjsSHSMNEVz5PTAh/pWpVWEH3KVhW4sjgJOMvF2vdI23CaFa/vypjQ
fH/obpkrdI9SH8CX1T+OlB7UZqt291gw8IPSg9VhYGwOYVmz9VXPxuztnVPpFTuv
9lZ5dC5+hMwjv0JzTgyM1PMCgYAwHfytMQO2BEkJ94ITycNVr4qHdo/X0+KLlZRz
hvezct2RwrKc+KztUetjHzmSMQrgz4aHkFE2Rx7idKmT4rBh8YDhqPGnmhE3uXgx
eFhoOXFCnFeFEjDf2geoj5D2KOqWelvUo4/NAN/w8d4anOFNtzLlbeDUnzNRcFi5
Ss/tgQKBgBBqCn6rmik0yn2vHFGjC8tp8EgKfOo/njifoJW17DQB5/1qZu+qQ6VX
ou2OqSz2rTeiRUWfr7DFLnrAZICFzdsTlUmgcrOalUMf0x+JAzOIgal0jAqrBa6g
5M//xAw8gZOMjbJ1UNxtfXkBr57glKTJEi/GKQn2ewM3FajwPtFn
-----END RSA PRIVATE KEY-----
```

### 生成根证书请求

```
openssl req -new -days 3650 -key keys/RootCA_1.key -out keys/RootCA.csr
```

上述命令解释：
+ req: PKCS#10的证书工具
+ new: 新生成一个证书请求
+ days: 证书有效期
+ key: 指定密钥文件
+ out: 输出的证书请求

```
root@egaotan-VirtualBox:~/cert# openssl req -new -days 3650 -key keys/RootCA_1.key -out keys/RootCA.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:cn
State or Province Name (full name) [Some-State]:shanghai
Locality Name (eg, city) []:shanghai
Organization Name (eg, company) [Internet Widgits Pty Ltd]:onchain
Organizational Unit Name (eg, section) []:chain
Common Name (e.g. server FQDN or YOUR name) []:egaotan
Email Address []:tgy_25@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:123456
An optional company name []:onchain
```

RootCA.csr文件：
```
-----BEGIN CERTIFICATE REQUEST-----
MIIC+zCCAeMCAQAwgYYxCzAJBgNVBAYTAmNuMREwDwYDVQQIDAhzaGFuZ2hhaTER
MA8GA1UEBwwIc2hhbmdoYWkxEDAOBgNVBAoMB29uY2hhaW4xDjAMBgNVBAsMBWNo
YWluMRAwDgYDVQQDDAdlZ2FvdGFuMR0wGwYJKoZIhvcNAQkBFg50Z3lfMjVAMTYz
LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALOsyz1dnclUMXX3
fHODHY1wyTEDiTh21+CevEGqanEkk2AGSOqYPxk4fhwWvMr5DE/dAxQexeh8Wwmd
2XTZd1MPrxHnKuThXw+r/nRCehq6fBMbXVEzM9uWl18SONicS7jpnM1eNpNUHVMx
4ksssxlztRf0jxPZ2VV7rtf6FY5VmTcsW9gRHgA4JfhtJKcUx+NQcOo57GAJwCUv
BmccGztMEtypIxCd8WmbhqWQeB70h4ut7BzO3TKF5LD297SSB12DkVTVxjkCWSP5
OXIizg6AicouUezO2DzOTfP1pINyBDYCLlNINWQVFqECpWqqB3sEzHttrKZcouYQ
CLZ06u8CAwEAAaAvMBUGCSqGSIb3DQEJBzEIDAYxMjM0NTYwFgYJKoZIhvcNAQkC
MQkMB29uY2hhaW4wDQYJKoZIhvcNAQELBQADggEBAH0d9XRz1CzvCMvf7iEnkpNV
SOdyft1tS8oDJbbCfS/WqPQoU6DXVw6yDonr+lcXaQGFNrKGdvnLfogIqRxTP2s0
ZxLN/OIR3oKGqD9iHSXVYRTdopDUBm4EFtOjfZlqfOrO1W6HSHTvEbxOVQsgFriy
7GPyNU9Qk5WCt3bPsjUz8FC2529Nb5l+JtVFTPbvhHVq6RDMEKy1UzfrQT6QtYm+
UYwpag+lCJ9NCgJ8iUe13rRp17bDN1pvcOO1I8eRHUi92HBauoTOh1BCkSL3dth9
xnEdQuLZM7poPhDPQilF9Gii8DQq3WAVBXk+q3+9VRkqWJ9y2om0aK43Ly4YnrU=
-----END CERTIFICATE REQUEST-----
```

### 生成根证书

```
openssl x509 -req -in keys/RootCA.csr -signkey keys/RootCA_1.key -out keys/RootCA.crt 
```

上述命令解释：
+ req: PKCS#10的证书工具
+ x509: 生成一个自签名证书而不是一个证书生成请求
+ in: 证书请求
+ signkey: 指定签名证书的密钥文件
+ out: 输出的证书文件

```
root@egaotan-VirtualBox:~/cert# openssl x509 -req -in keys/RootCA.csr -signkey keys/RootCA_1.key -out keys/RootCA.crt
Signature ok
subject=/C=cn/ST=shanghai/L=shanghai/O=onchain/OU=chain/CN=egaotan/emailAddress=tgy_25@163.com
Getting Private key
```

RootCA.crt文件：
```
-----BEGIN CERTIFICATE-----
MIIDijCCAnICCQDd3FlQmerq/zANBgkqhkiG9w0BAQsFADCBhjELMAkGA1UEBhMC
Y24xETAPBgNVBAgMCHNoYW5naGFpMREwDwYDVQQHDAhzaGFuZ2hhaTEQMA4GA1UE
CgwHb25jaGFpbjEOMAwGA1UECwwFY2hhaW4xEDAOBgNVBAMMB2VnYW90YW4xHTAb
BgkqhkiG9w0BCQEWDnRneV8yNUAxNjMuY29tMB4XDTIwMTEwNDA3MDExOFoXDTIw
MTIwNDA3MDExOFowgYYxCzAJBgNVBAYTAmNuMREwDwYDVQQIDAhzaGFuZ2hhaTER
MA8GA1UEBwwIc2hhbmdoYWkxEDAOBgNVBAoMB29uY2hhaW4xDjAMBgNVBAsMBWNo
YWluMRAwDgYDVQQDDAdlZ2FvdGFuMR0wGwYJKoZIhvcNAQkBFg50Z3lfMjVAMTYz
LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALOsyz1dnclUMXX3
fHODHY1wyTEDiTh21+CevEGqanEkk2AGSOqYPxk4fhwWvMr5DE/dAxQexeh8Wwmd
2XTZd1MPrxHnKuThXw+r/nRCehq6fBMbXVEzM9uWl18SONicS7jpnM1eNpNUHVMx
4ksssxlztRf0jxPZ2VV7rtf6FY5VmTcsW9gRHgA4JfhtJKcUx+NQcOo57GAJwCUv
BmccGztMEtypIxCd8WmbhqWQeB70h4ut7BzO3TKF5LD297SSB12DkVTVxjkCWSP5
OXIizg6AicouUezO2DzOTfP1pINyBDYCLlNINWQVFqECpWqqB3sEzHttrKZcouYQ
CLZ06u8CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEASa+Gd++Z63lEu9W0FLLp+23q
QrEu1AS/E2SZqKqzOReEKnEwuZmd84daZEk96MLkuJrgf7YDKhFmdvF1pgV0JsgC
HKiV0lM6Yd4zCeTD91yL5bWEoXpSlJRplkRboKWJHa6YLJCilOECzOtHJfbfH0hv
Dr5H7dTD2x66gkPRqqbtkTcpLsBJJ8i848pMwo8QDQIU+Y7/0p4HHFhfLBT4opUM
ZDVHcQPRMvH8BxtTPnORYsZtsYDnkai6P3Ehd2pwsArwO2osThWv3/AtzVINxSdt
kae1TehyLHV2WxA0rSK9IiAdALJZsv6KpbWYlQx9Mn0F0NClqu8QDjlkmP6exw==
-----END CERTIFICATE-----
```

可以从密钥直接生成一个自签名证书。

```
openssl req -new -x509 -days 3650 -key keys/RootCA.key -out keys/RootCA.crt
```

上述命令解释：
+ req: PKCS#10的证书工具
+ new: 新生成一个证书
+ x509: 生成一个自签名证书而不是一个证书生成请求
+ days: 证书有效期
+ key: 指定密钥文件
+ out: 签名的证书

```
root@egaotan-VirtualBox:~/cert# openssl req -new -x509 -days 3650 -key keys/RootCA.key -out keys/RootCA.crt
Enter pass phrase for keys/RootCA.key:
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:cn
State or Province Name (full name) [Some-State]:shanghai
Locality Name (eg, city) []:shanghai
Organization Name (eg, company) [Internet Widgits Pty Ltd]:onchain
Organizational Unit Name (eg, section) []:chain
Common Name (e.g. server FQDN or YOUR name) []:egaotan
Email Address []:tgy_25@163.com
```

### 文件格式解释

+ .key格式：私有的密钥
+ .csr格式：证书签名请求（证书请求文件），含有公钥信息，certificate signing request的缩写
+ .crt格式：证书文件，certificate的缩写
+ .crl格式：证书吊销列表，Certificate Revocation List的缩写
+ .pem格式：用于导出，导入证书时候的证书的格式，有证书开头，结尾的格式

### 根证书生成流程

+ 生成CA私钥（.key）
+ 生成CA证书请求（.csr）
+ 自签名得到根证书（.crt）（CA给自已颁发的证书）

## 生成二级证书
Root证书作为第一级证书，现在生成第二级证书。

### 先生成一个用户密钥

```
openssl genrsa -des3 -out user1/User1CA.key 2048
```

### 转换用户密钥格式

```
openssl rsa -in user1/User1CA.key -out user1/User1CA_1.key
```

### 生成证书请求

```
openssl req -new -days 3650 -key user1/User1CA_1.key -out user1/User1CA.csr
```

```
root@egaotan-VirtualBox:~/cert# openssl req -new -days 3650 -key user1/User1CA_1.key -out user1/User1CA.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:cn
State or Province Name (full name) [Some-State]:shanghai
Locality Name (eg, city) []:shanghai
Organization Name (eg, company) [Internet Widgits Pty Ltd]:onchain
Organizational Unit Name (eg, section) []:chain
Common Name (e.g. server FQDN or YOUR name) []:xgaotan
Email Address []:tgy_25@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:123456
An optional company name []:onchain
```

### 签名证书

```
openssl ca -in user1/User1CA.csr -out user1/User1CA.crt -cert keys/RootCA.crt -keyfile keys/RootCA.key
```

命令解释：
+ ca: 签名证书
+ in: 证书请求
+ out: 签名的证书
+ cert: 根证书文件
+ keyfile: 根证书的密钥

```
root@egaotan-VirtualBox:~/cert# openssl ca -in user1/User1CA.csr -out user1/User1CA.crt -cert keys/RootCA.crt -keyfile keys/RootCA.key
Using configuration from /usr/lib/ssl/openssl.cnf
Enter pass phrase for keys/RootCA.key:
Check that the request matches the signature
Signature ok
Certificate Details:
        Serial Number: 1 (0x1)
        Validity
            Not Before: Nov  4 07:22:04 2020 GMT
            Not After : Nov  4 07:22:04 2021 GMT
        Subject:
            countryName               = cn
            stateOrProvinceName       = shanghai
            organizationName          = onchain
            organizationalUnitName    = chain
            commonName                = egaotan
            emailAddress              = tgy_25@163.com
        X509v3 extensions:
            X509v3 Basic Constraints: 
                CA:FALSE
            Netscape Comment: 
                OpenSSL Generated Certificate
            X509v3 Subject Key Identifier: 
                45:A2:68:A4:D0:D9:41:54:67:62:12:0F:4C:9D:7C:EE:C2:63:F5:1A
            X509v3 Authority Key Identifier: 
                DirName:/C=cn/ST=shanghai/L=shanghai/O=onchain/OU=chain/CN=egaotan/emailAddress=tgy_25@163.com
                serial:DD:DC:59:50:99:EA:EA:FF

Certificate is to be certified until Nov  4 07:22:04 2021 GMT (365 days)
Sign the certificate? [y/n]:y


1 out of 1 certificate requests certified, commit? [y/n]y
Write out database with 1 new entries
Data Base Updated
root@egaotan-VirtualBox:~/cert# 
```

如果遇到以下错误
```
root@egaotan-VirtualBox:~/cert# openssl ca -extensions v3_ca -in user1/User1CA.csr -out user1/User1CA.crt -cert keys/RootCA.crt -keyfile keys/RootCA.key
Using configuration from /usr/lib/ssl/openssl.cnf
Enter pass phrase for keys/RootCA.key:
I am unable to access the ./demoCA/newcerts directory
./demoCA/newcerts: No such file or directory
```

解决办法：
```
mkdir -p ./demoCA/newcerts 
touch demoCA/index.txt 
touch demoCA/serial 
echo 01 > demoCA/serial
```

用户证书文件User1CA.crt:
```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 1 (0x1)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=cn, ST=shanghai, L=shanghai, O=onchain, OU=chain, CN=egaotan/emailAddress=tgy_25@163.com
        Validity
            Not Before: Nov  4 07:22:04 2020 GMT
            Not After : Nov  4 07:22:04 2021 GMT
        Subject: C=cn, ST=shanghai, O=onchain, OU=chain, CN=egaotan/emailAddress=tgy_25@163.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:f3:12:11:1a:91:8b:d8:1e:d5:2f:b8:f9:53:76:
                    f5:5b:fa:42:68:af:8b:f7:e8:a9:13:63:21:ff:9b:
                    77:30:1c:6a:d5:f9:c4:05:1b:3a:f6:61:47:ce:41:
                    2a:2d:ce:c4:d0:b0:ea:d6:30:95:65:17:50:f0:1b:
                    53:0a:cd:ce:e3:cb:cd:e2:df:fc:37:8a:8b:9a:34:
                    62:16:cc:bc:3b:f6:7a:04:69:1d:1c:68:56:c7:fa:
                    8d:25:93:7d:8d:77:2d:91:eb:49:83:72:b4:fb:b3:
                    32:ac:6d:2c:f5:35:ad:91:11:81:b2:df:c1:8b:0d:
                    ec:b2:27:2b:28:de:47:be:e4:a7:a9:09:f0:b2:02:
                    c8:09:05:79:84:d9:98:90:45:18:6a:4f:b8:73:c3:
                    d5:39:d6:c3:df:f7:e2:25:74:b3:80:d8:aa:7a:cd:
                    c2:48:1a:2b:c3:45:c9:ba:76:2a:f1:ec:cc:ed:6d:
                    51:e8:77:4e:10:5b:4e:31:08:21:ab:a9:dd:c7:fa:
                    30:d9:1a:48:a7:08:e8:b1:ea:76:23:36:23:90:27:
                    58:34:b7:a0:eb:a2:d4:56:f8:e1:44:18:47:8f:8b:
                    fe:ea:b7:f6:6b:3c:33:cc:8f:32:f8:7c:4c:a5:b1:
                    67:a5:22:0a:38:6c:04:ae:fc:18:f2:63:a1:c3:fb:
                    0a:d9
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints:
                CA:FALSE
            Netscape Comment:
                OpenSSL Generated Certificate
            X509v3 Subject Key Identifier:
                45:A2:68:A4:D0:D9:41:54:67:62:12:0F:4C:9D:7C:EE:C2:63:F5:1A
            X509v3 Authority Key Identifier:
                DirName:/C=cn/ST=shanghai/L=shanghai/O=onchain/OU=chain/CN=egaotan/emailAddress=tgy_25@163.com
                serial:DD:DC:59:50:99:EA:EA:FF

    Signature Algorithm: sha256WithRSAEncryption
         77:6f:b9:6c:5b:97:fa:d3:cc:cc:aa:0c:2a:dd:71:fa:db:09:
         53:0e:01:d9:54:93:5f:16:d6:90:c6:c4:37:0a:ec:73:99:c0:
         16:24:1e:3a:5b:a9:af:95:a0:c2:87:9e:2b:fe:0d:b0:4f:28:
         a9:e3:de:ff:b1:85:57:0f:11:8c:db:34:02:ea:51:2c:56:64:
         b8:1e:cb:78:b4:5b:84:33:6b:20:fe:0f:5c:0f:0b:2c:cf:15:
         25:dd:a4:09:63:05:4d:ce:45:09:ad:6e:b9:7e:1b:cc:2d:7c:
         2c:07:75:ec:b9:b2:fc:4b:f9:a3:ea:bb:0a:47:8f:92:af:9d:
         ca:ab:c2:a1:f2:ce:e8:9b:88:66:89:75:d4:67:55:ad:29:68:
         9b:a0:7a:98:21:51:c7:52:66:94:c8:c1:30:a2:66:45:1e:99:
         54:b0:eb:27:50:d4:8b:47:4f:eb:c1:93:53:b8:c0:2e:e4:55:
         3d:2d:40:19:ab:50:b4:a8:4a:eb:be:89:98:0f:49:af:b9:fd:
         4b:97:85:66:ed:c3:58:94:af:f9:c3:5a:f7:dc:01:3b:ba:6c:
         fe:16:63:da:65:a3:d7:44:48:f5:e2:2d:2b:ad:33:3f:89:ad:
         1e:84:65:b4:6f:11:0b:56:82:b7:c7:8c:fd:54:08:a3:a8:13:
         42:c4:22:a1
-----BEGIN CERTIFICATE-----
MIIEezCCA2OgAwIBAgIBATANBgkqhkiG9w0BAQsFADCBhjELMAkGA1UEBhMCY24x
ETAPBgNVBAgMCHNoYW5naGFpMREwDwYDVQQHDAhzaGFuZ2hhaTEQMA4GA1UECgwH
b25jaGFpbjEOMAwGA1UECwwFY2hhaW4xEDAOBgNVBAMMB2VnYW90YW4xHTAbBgkq
hkiG9w0BCQEWDnRneV8yNUAxNjMuY29tMB4XDTIwMTEwNDA3MjIwNFoXDTIxMTEw
NDA3MjIwNFowczELMAkGA1UEBhMCY24xETAPBgNVBAgMCHNoYW5naGFpMRAwDgYD
VQQKDAdvbmNoYWluMQ4wDAYDVQQLDAVjaGFpbjEQMA4GA1UEAwwHZWdhb3RhbjEd
MBsGCSqGSIb3DQEJARYOdGd5XzI1QDE2My5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDzEhEakYvYHtUvuPlTdvVb+kJor4v36KkTYyH/m3cwHGrV
+cQFGzr2YUfOQSotzsTQsOrWMJVlF1DwG1MKzc7jy83i3/w3iouaNGIWzLw79noE
aR0caFbH+o0lk32Ndy2R60mDcrT7szKsbSz1Na2REYGy38GLDeyyJyso3ke+5Kep
CfCyAsgJBXmE2ZiQRRhqT7hzw9U51sPf9+IldLOA2Kp6zcJIGivDRcm6dirx7Mzt
bVHod04QW04xCCGrqd3H+jDZGkinCOix6nYjNiOQJ1g0t6DrotRW+OFEGEePi/7q
t/ZrPDPMjzL4fEylsWelIgo4bASu/BjyY6HD+wrZAgMBAAGjggEEMIIBADAJBgNV
HRMEAjAAMCwGCWCGSAGG+EIBDQQfFh1PcGVuU1NMIEdlbmVyYXRlZCBDZXJ0aWZp
Y2F0ZTAdBgNVHQ4EFgQURaJopNDZQVRnYhIPTJ187sJj9RowgaUGA1UdIwSBnTCB
mqGBjKSBiTCBhjELMAkGA1UEBhMCY24xETAPBgNVBAgMCHNoYW5naGFpMREwDwYD
VQQHDAhzaGFuZ2hhaTEQMA4GA1UECgwHb25jaGFpbjEOMAwGA1UECwwFY2hhaW4x
EDAOBgNVBAMMB2VnYW90YW4xHTAbBgkqhkiG9w0BCQEWDnRneV8yNUAxNjMuY29t
ggkA3dxZUJnq6v8wDQYJKoZIhvcNAQELBQADggEBAHdvuWxbl/rTzMyqDCrdcfrb
CVMOAdlUk18W1pDGxDcK7HOZwBYkHjpbqa+VoMKHniv+DbBPKKnj3v+xhVcPEYzb
NALqUSxWZLgey3i0W4QzayD+D1wPCyzPFSXdpAljBU3ORQmtbrl+G8wtfCwHdey5
svxL+aPquwpHj5KvncqrwqHyzuibiGaJddRnVa0paJugepghUcdSZpTIwTCiZkUe
mVSw6ydQ1ItHT+vBk1O4wC7kVT0tQBmrULSoSuu+iZgPSa+5/UuXhWbtw1iUr/nD
WvfcATu6bP4WY9plo9dESPXiLSutMz+JrR6EZbRvEQtWgrfHjP1UCKOoE0LEIqE=
-----END CERTIFICATE-----
```

这份用户证书中包含了证书的明文信息。

### 用户证书生成步骤

+ 生成私钥（.key）
+ 生成证书请求（.csr）
+ 用CA根证书签名得到证书（.crt）

### 验证证书

```
openssl verify -verbose -CAfile keys/RootCA.crt user1/User1CA.crt
```

验证结果:
```
root@egaotan-VirtualBox:~/cert# openssl verify -verbose -CAfile keys/RootCA.crt user1/User1CA.crt
user1/User1CA.crt: OK
```

## 生成三级证书

### 生成用户密钥

```
openssl genrsa -des3 -out user3/User3CA.key 2048
```

### 转换密钥格式
```
openssl rsa -in user3/User3CA.key -out user3/User3CA_1.key
```
### 生成证书请求

```
openssl req -new -days 3650 -key user3/User3CA_1.key -out user3/User3CA.csr
```

```
root@egaotan-VirtualBox:~/cert# openssl req -new -days 3650 -key user3/User3CA_1.key -out user3/User3CA.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:cn
State or Province Name (full name) [Some-State]:shanghai
Locality Name (eg, city) []:shanghai
Organization Name (eg, company) [Internet Widgits Pty Ltd]:onchain
Organizational Unit Name (eg, section) []:chain
Common Name (e.g. server FQDN or YOUR name) []:agaotan
Email Address []:tgy_25@163.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:123456
An optional company name []:onchain
```

### 签名证书

```
openssl ca -in user3/User3CA.csr -out user3/User3CA.crt -cert user1/User1CA.crt -keyfile user1/User1CA_1.key
```

```
root@egaotan-VirtualBox:~/cert# openssl ca -in user3/User3CA.csr -out user3/User3CA.crt -cert user1/User1CA.crt -keyfile user1/User1CA_1.key
Using configuration from /usr/lib/ssl/openssl.cnf
Check that the request matches the signature
Signature ok
Certificate Details:
        Serial Number: 3 (0x3)
        Validity
            Not Before: Nov  4 07:59:52 2020 GMT
            Not After : Nov  4 07:59:52 2021 GMT
        Subject:
            countryName               = cn
            stateOrProvinceName       = shanghai
            organizationName          = onchain
            organizationalUnitName    = chain
            commonName                = agaotan
            emailAddress              = tgy_25@163.com
        X509v3 extensions:
            X509v3 Basic Constraints: 
                CA:FALSE
            Netscape Comment: 
                OpenSSL Generated Certificate
            X509v3 Subject Key Identifier: 
                56:98:BE:7B:D6:10:6A:77:C1:0B:A2:CE:A0:97:E8:9B:3B:7B:93:9C
            X509v3 Authority Key Identifier: 
                keyid:45:A2:68:A4:D0:D9:41:54:67:62:12:0F:4C:9D:7C:EE:C2:63:F5:1A

Certificate is to be certified until Nov  4 07:59:52 2021 GMT (365 days)
Sign the certificate? [y/n]:y


1 out of 1 certificate requests certified, commit? [y/n]y
Write out database with 1 new entries
Data Base Updated
```

签名的证书文件User3CA.crt:
```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 3 (0x3)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=cn, ST=shanghai, O=onchain, OU=chain, CN=xgaotan/emailAddress=tgy_25@163.com
        Validity
            Not Before: Nov  4 07:59:52 2020 GMT
            Not After : Nov  4 07:59:52 2021 GMT
        Subject: C=cn, ST=shanghai, O=onchain, OU=chain, CN=agaotan/emailAddress=tgy_25@163.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:d4:b8:f0:11:58:e0:d0:31:ef:26:88:d8:5d:28:
                    d5:61:96:75:24:77:d9:f4:e3:6e:8e:14:5b:b2:69:
                    c2:36:87:0f:4d:bc:69:c6:b7:06:5b:87:c0:42:97:
                    83:3d:4f:e8:b8:ee:0d:62:5b:97:4d:ab:c9:2c:70:
                    6f:86:0e:3c:bf:c0:49:5e:c9:8d:73:7e:94:e3:a1:
                    cb:31:2d:b9:d0:f6:9b:20:86:bd:05:61:71:d4:74:
                    df:74:19:83:26:78:e0:2a:6c:7a:62:63:28:87:1b:
                    36:49:81:a2:72:37:06:b1:b7:06:42:b3:3d:ad:2e:
                    1c:a1:0c:af:2a:bf:9f:0d:54:b7:0d:01:0d:6a:cb:
                    6d:b5:18:3c:c4:00:cd:26:0c:40:46:17:85:ea:f5:
                    b8:95:df:2c:bb:c1:78:ae:ef:95:b8:78:b2:b5:cb:
                    d2:4a:9b:9a:14:b5:26:e8:89:ee:81:2b:e7:6a:db:
                    95:40:78:0d:b4:33:9b:d2:23:8e:e2:b2:a1:5e:1d:
                    50:de:7e:45:f3:2a:04:a3:8a:6f:bf:c3:f8:85:fb:
                    75:1b:85:59:1b:63:07:a1:ca:85:8b:26:99:1c:e2:
                    36:40:40:5f:f3:4a:62:6d:73:85:ff:6a:d4:fb:56:
                    88:82:3d:49:f2:6f:ca:1d:25:de:d7:17:ac:07:a0:
                    d5:39
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints:
                CA:FALSE
            Netscape Comment:
                OpenSSL Generated Certificate
            X509v3 Subject Key Identifier:
                56:98:BE:7B:D6:10:6A:77:C1:0B:A2:CE:A0:97:E8:9B:3B:7B:93:9C
            X509v3 Authority Key Identifier:
                keyid:45:A2:68:A4:D0:D9:41:54:67:62:12:0F:4C:9D:7C:EE:C2:63:F5:1A

    Signature Algorithm: sha256WithRSAEncryption
         78:64:83:bf:6f:46:13:96:4a:1e:da:61:08:ea:04:37:3c:46:
         93:0f:10:74:e3:1e:a6:e8:12:b9:52:94:6a:6b:db:ea:1f:f3:
         fc:f6:4d:1a:b3:ca:51:dc:07:ce:d0:47:41:ff:f1:07:17:18:
         6d:ee:2a:3a:4d:51:26:9c:27:2f:1d:e2:70:4a:d7:ca:65:89:
         85:85:96:c3:e9:e4:73:a1:e1:a3:53:3a:e6:c5:bf:2f:bd:ee:
         ca:cf:b2:cf:43:b0:ff:ca:28:af:17:96:4a:5d:3f:97:f1:03:
         67:35:85:d5:6d:0d:f7:f3:4e:2e:a7:73:4f:d5:1e:1d:bd:bd:
         17:01:a8:5a:a7:bb:c9:b5:b4:47:b8:2f:27:c8:16:7d:8f:88:
         e3:c7:ce:c9:b6:37:58:f7:b0:35:9e:20:b8:cd:fe:d2:99:0b:
         7c:eb:d4:1e:e9:db:d0:82:b1:1c:d6:46:81:e6:95:1d:72:c9:
         2e:ad:c7:60:a1:3f:bf:ee:3e:fc:c5:77:9a:25:27:88:d1:3a:
         c3:a9:ba:76:1f:f0:ea:e8:ab:7d:07:9e:25:15:c4:8e:b3:81:
         26:e4:68:86:61:4c:11:06:0d:0d:71:45:b4:07:5f:f7:58:c4:
         62:6f:9e:6f:6d:fa:5d:cc:77:32:cd:bd:f3:8b:8b:70:3d:8c:
         ac:20:f3:52
-----BEGIN CERTIFICATE-----
MIID3DCCAsSgAwIBAgIBAzANBgkqhkiG9w0BAQsFADBzMQswCQYDVQQGEwJjbjER
MA8GA1UECAwIc2hhbmdoYWkxEDAOBgNVBAoMB29uY2hhaW4xDjAMBgNVBAsMBWNo
YWluMRAwDgYDVQQDDAd4Z2FvdGFuMR0wGwYJKoZIhvcNAQkBFg50Z3lfMjVAMTYz
LmNvbTAeFw0yMDExMDQwNzU5NTJaFw0yMTExMDQwNzU5NTJaMHMxCzAJBgNVBAYT
AmNuMREwDwYDVQQIDAhzaGFuZ2hhaTEQMA4GA1UECgwHb25jaGFpbjEOMAwGA1UE
CwwFY2hhaW4xEDAOBgNVBAMMB2FnYW90YW4xHTAbBgkqhkiG9w0BCQEWDnRneV8y
NUAxNjMuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1LjwEVjg
0DHvJojYXSjVYZZ1JHfZ9ONujhRbsmnCNocPTbxpxrcGW4fAQpeDPU/ouO4NYluX
TavJLHBvhg48v8BJXsmNc36U46HLMS250PabIIa9BWFx1HTfdBmDJnjgKmx6YmMo
hxs2SYGicjcGsbcGQrM9rS4coQyvKr+fDVS3DQENastttRg8xADNJgxARheF6vW4
ld8su8F4ru+VuHiytcvSSpuaFLUm6InugSvnatuVQHgNtDOb0iOO4rKhXh1Q3n5F
8yoEo4pvv8P4hft1G4VZG2MHocqFiyaZHOI2QEBf80pibXOF/2rU+1aIgj1J8m/K
HSXe1xesB6DVOQIDAQABo3sweTAJBgNVHRMEAjAAMCwGCWCGSAGG+EIBDQQfFh1P
cGVuU1NMIEdlbmVyYXRlZCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQUVpi+e9YQanfB
C6LOoJfomzt7k5wwHwYDVR0jBBgwFoAURaJopNDZQVRnYhIPTJ187sJj9RowDQYJ
KoZIhvcNAQELBQADggEBAHhkg79vRhOWSh7aYQjqBDc8RpMPEHTjHqboErlSlGpr
2+of8/z2TRqzylHcB87QR0H/8QcXGG3uKjpNUSacJy8d4nBK18pliYWFlsPp5HOh
4aNTOubFvy+97srPss9DsP/KKK8XlkpdP5fxA2c1hdVtDffzTi6nc0/VHh29vRcB
qFqnu8m1tEe4LyfIFn2PiOPHzsm2N1j3sDWeILjN/tKZC3zr1B7p29CCsRzWRoHm
lR1yyS6tx2ChP7/uPvzFd5olJ4jROsOpunYf8Oroq30HniUVxI6zgSbkaIZhTBEG
DQ1xRbQHX/dYxGJvnm9t+l3MdzLNvfOLi3A9jKwg81I=
-----END CERTIFICATE-----
```

### 验证证书

```
openssl verify -verbose -CAfile user1/User1CA.crt user3/User3CA.crt
```

没有指定根证书，验证失败:
```
root@egaotan-VirtualBox:~/cert# openssl verify -verbose -CAfile user1/User1CA.crt user3/User3CA.crt
user3/User3CA.crt: C = cn, ST = shanghai, O = onchain, OU = chain, CN = xgaotan, emailAddress = tgy_25@163.com
error 2 at 1 depth lookup:unable to get issuer certificate
```

指定根证书，验证成功：
```
root@egaotan-VirtualBox:~/cert# openssl verify -CAfile keys/RootCA.crt -untrusted user1/User1CA.crt user3/User3CA.crt
user3/User3CA.crt: C = cn, ST = shanghai, O = onchain, OU = chain, CN = xgaotan, emailAddress = tgy_25@163.com
error 24 at 1 depth lookup:invalid CA certificate
OK
```

## 参考
[证书之间的转换(crt pem key)](https://blog.csdn.net/qq_37049781/article/details/84837342)
[加密解密工具](http://tool.chacuo.net/cryptrsapubkey)
[OpenSSL中RSA私钥文件](https://www.cnblogs.com/jukan/p/5527922.html)