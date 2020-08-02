# go mod使用

## 如何在mod中使用私有仓库

因为程序中总是这样写：

```
import (
	"github.com/ontio/multi-chain-go-sdk"
	"github.com/ontio/multi-chain"
)
```

即使multi-chain-go-sdk和multi-chain这两个代码的仓库位于私有仓库中，如位于https://git.ont.io/Cross-Chain/multi-chain-go-sdk.git和https://git.ont.io/Cross-Chain/multi-chain.git这两个仓库中。

如果是这种情况，那么在以下go mod中会出现错误：
```
go mod init
go build
```
如：
```
root@egaotan-VirtualBox:~/gopath/src/github.com/ontio/ontology-tool# go mod init
go: creating new go.mod: module github.com/ontio/ontology-tool
root@egaotan-VirtualBox:~/gopath/src/github.com/ontio/ontology-tool# go build
go: finding github.com/ethereum/go-ethereum v1.9.14
go: finding github.com/alecthomas/log4go latest
go: downloading github.com/ethereum/go-ethereum v1.9.14
go: extracting github.com/ethereum/go-ethereum v1.9.14
go: finding github.com/FactomProject/basen latest
go: downloading github.com/VictoriaMetrics/fastcache v1.5.7
go: extracting github.com/VictoriaMetrics/fastcache v1.5.7
go: finding github.com/VictoriaMetrics/fastcache v1.5.7
build github.com/ontio/ontology-tool: cannot load github.com/ontio/multi-chain-go-sdk: git ls-remote -q https://github.com/ontio/multi-chain-go-sdk in /root/gopath/pkg/mod/cache/vcs/0d62aaca1954b5047cafba1cfbb19625b3d3f92ebca553557f0701c9f18a1d7c: exit status 128:
	fatal: could not read Username for 'https://github.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```
go mod会提示从github.com/ontio/multi-chain-go-sdk处load该仓库失败，当然会失败了，因为这是在https://git.ont.io/Cross-Chain/multi-chain-go-sdk.git仓库中，并不在github上。

go mod从程序的import中默认导入，使用指定了github.com上的仓库，但其实我们使用的是git.ont.io上的仓库，需要做配置使得go mod从git.ont.io拉取仓库。

配置go.mod：
```
require (
        github.com/ontio/multi-chain v0.0.0
        github.com/ontio/multi-chain-go-sdk v0.0.0-00010101000000-000000000000
)
replace (
        github.com/ontio/multi-chain => git.ont.io/Cross-Chain/multi-chain v0.0.0-20200520112631-c1a6e27571fa
        github.com/ontio/multi-chain-go-sdk => git.ont.io/Cross-Chain/multi-chain-go-sdk v0.0.0-20200507075820-9d050f4067c0
)
```
上面的配置中，指定了程序使用的所有的github.com的multi-chain和multi-chain-go-sdk被替换为git.ont.io/Cross-Chain的multi-chain和multi-chain-go-sdk。

很多时候，我们的仓库是还未tag的，这是可以这样指定：
```
replace (
        github.com/ontio/multi-chain => git.ont.io/Cross-Chain/multi-chain c1a6e27571fa
        github.com/ontio/multi-chain-go-sdk => git.ont.io/Cross-Chain/multi-chain-go-sdk 9d050f4067c0
)
```

有可能你的私有仓库需要用户名和密码来登录，如果是这样，那么会出现以下错误：
```
root@egaotan-VirtualBox:~/gopath/src/github.com/ontio/ontology-tool# go build
go: finding git.ont.io/Cross-Chain/multi-chain c1a6e27571
go: finding git.ont.io/Cross-Chain/multi-chain c1a6e27571
go: finding git.ont.io/Cross-Chain/multi-chain-go-sdk 9d050f4067
go: finding git.ont.io/Cross-Chain/multi-chain-go-sdk 9d050f4067
go: errors parsing go.mod:
/root/gopath/src/github.com/ontio/ontology-tool/go.mod:13: replace git.ont.io/Cross-Chain/multi-chain: version "c1a6e27571" invalid: git fetch -f origin refs/heads/*:refs/heads/* refs/tags/*:refs/tags/* in /root/gopath/pkg/mod/cache/vcs/ab420e34763f0c4fe852efeb7766e1bacc2735aa0d5f3bfb1ef7a6de77dc0804: exit status 128:
	fatal: could not read Username for 'https://git.ont.io': terminal prompts disabled
/root/gopath/src/github.com/ontio/ontology-tool/go.mod:14: replace git.ont.io/Cross-Chain/multi-chain-go-sdk: version "9d050f4067" invalid: git fetch -f https://git.ont.io/Cross-Chain/multi-chain-go-sdk.git refs/heads/*:refs/heads/* refs/tags/*:refs/tags/* in /root/gopath/pkg/mod/cache/vcs/0519e791604d79391b83ae56cb9de96ac6f423bf625b138bf78106b59c2a0d8d: exit status 128:
	fatal: could not read Username for 'https://git.ont.io': terminal prompts disabled
```

go mod在从私有仓库拉取代码的时候需要输入用户名和密码，但这里并没有开启输入终端，我们设置为开启：
```
export GIT_TERMINAL_PROMPT=1
```

这个时候，go mod拉取代码的时候，会提示用户输入用户名和密码。

也可以配置用户名和密码，这样每次拉取这个私有仓库的代码都不用再输入用户名和密码。
```
git config --global --add url."git@github.com:".insteadOf "https://github.com/"
```

即使这样，在拉取私有仓库的代码的时候，还是会给出以下错误信息：
```
root@egaotan-VirtualBox:~/gopath/src/github.com/ontio/ontology-tool# go build
go: finding git.ont.io/Cross-Chain/multi-chain c1a6e27571
go: finding git.ont.io/Cross-Chain/multi-chain c1a6e27571
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
go: finding git.ont.io/Cross-Chain/multi-chain-go-sdk 9d050f4067
go: finding git.ont.io/Cross-Chain/multi-chain-go-sdk 9d050f4067
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
Username for 'https://git.ont.io': tangaoyuan@onchain.com
Password for 'https://tangaoyuan@onchain.com@git.ont.io':
go: downloading git.ont.io/Cross-Chain/multi-chain-go-sdk v0.0.0-20200507075820-9d050f4067c0
go: downloading git.ont.io/Cross-Chain/multi-chain v0.0.0-20200520112631-c1a6e27571fa
verifying git.ont.io/Cross-Chain/multi-chain-go-sdk@v0.0.0-20200507075820-9d050f4067c0: git.ont.io/Cross-Chain/multi-chain-go-sdk@v0.0.0-20200507075820-9d050f4067c0: reading https://sum.golang.org/lookup/git.ont.io/!cross-!chain/multi-chain-go-sdk@v0.0.0-20200507075820-9d050f4067c0: 410 Gone
```

私有仓库的代码验证失败，那么我们设置为不验证了吧：
```
export GOSUMDB=off
```

另一种配置为本地的repository:

replace (
	github.com/ontio/multi-chain => C:\Users\DELL\go\src\github.com\ontio\multi-chain
	github.com/ontio/multi-chain-go-sdk => C:\Users\DELL\go\src\github.com\ontio\multi-chain-go-sdk
)

