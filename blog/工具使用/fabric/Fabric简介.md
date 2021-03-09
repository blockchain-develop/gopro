# Fabric简介

## 1. 关键概念

### 1.1. 网络架构

#### 1.1.1. 交易

和其他区块链系统一样地，Fabric也是通过交易来实现与状态机的交互。交易是由具备写权限的账户发起，调用链码接口，修改对应状态的请求，所有节点执行该交易应该获得同样的结果，从而使整个状态机保持一致性。

和其他系统略有不同的是交易执行成功与否的标准。在以太坊中，节点按顺序执行区块中的交易，每执行一笔交易，即对账本做出更改，下一笔交易在新的状态下执行，比如小明有100块钱，第一笔交易扣掉小明10块，还有90块，第二笔扣掉20块，还有70块，两笔交易都成功落账。可是，在Fabric中，第二笔交易就不会成功，而会被标记为invalid。

这里就要提到[RWSet](https://hyperledger-fabric.readthedocs.io/zh_CN/release-1.4/readwrite.html)这个概念。每笔交易会有一个执行状态，称为读集，以及一个执行结果，称为写集，读写集是背书时候确定的。在交易执行的时候，读集对应的KVPair必须和当前状态相同，若相同，则可以按照写集修改状态。

经过背书，第一笔交易*TX1{<RSet: (K: 小明, V: 100)>, <WSet: (K: 小明, V: 90)>}*，第二笔交易*TX2{<RSet: (K: 小明, V: 100)>, <WSet: (K: 小明, V: 80)>}*，它们的读集是一样的，当执行第二笔交易的时候，发现TX2的读集和当前状态不同了，那这笔交易就会被标志为错误。

Fabric主要的特点就是读集不同，即执行的预期状态不同，则交易失败，对比以太坊等公链多了这个限制条件，更加严格。

#### 1.1.2. 通道

> Hyperledger Fabric 中的通道是两个或两个以上特定网络成员之间通信的专用“子网”，用于进行私有和机密的交易。通道由成员(组织)、每个成员的锚点节点、共享账本、链码应用程序和排序服务节点定义。网络上的每个交易都在一个通道上执行，在这个通道上，每一方都必须经过身份认证和授权才能在该通道上进行交易。

通道的定义如上，通道即channel，是由多个组织构成的子网络，具有自己的区块链账本，成员管理也是由channel自己负责的。channel可以分为两种，一种是系统通道，运行着Fabric排序服务，维护了所有可以创建通道的组织的列表，这个列表称之为”联盟“，其他通道的创建、更新都要经过系统通道，系统通道可以通过Raft等协议启动，是整个系统的共识部分，其他通道依赖于排序服务实现自己的一致性；另一种是应用通道，即由多个组织联合创建的通道，可以运行链码，执行某些业务。

通道的准入都是通过CA来控制的，在Fabric中CA可以看做账户，CA中包含了很多该实体的属性，比如角色、权限和公钥等，每个组织可以创建自己的私钥和对应的根CA，为组织的其他成员颁发CA，Fabric提供了FabricCA等一系列工具来完成这些工作。

通道的所有组织的根证书等信息都会存储在账本中，所以任意组织签发的证书都可以被其他组织的节点验证。

### 1.2. 节点角色

#### 1.2.1. Orderer

Orderer节点就是排序服务的节点，负责为所有通道打包区块，同时维护了联盟组织列表，只有这些组织可以创建通道，但是非联盟成员也可以**加入**通道。

#### 1.2.2. Peer

Peer节点负责维护一个通道的区块链账本，每个组织可以有多个Peer节点，Peer节点负责为该通道的交易进行背书，当交易满足背书策略的时候，就可以提交给排序服务，通道将获得新的区块，Fabric的区块是惰性产生的，不会产生空块。应用程序就是直接和Peer节点交互的。总之，Peer身份具备背书、提交交易和查询等权力。

无论是什么角色，都对应于一个CA证书，证书中包含了相对应的信息。

#### 1.2.3. Client

Client一般是应用程序需要具备的身份，具备提出交易、查询Peer等权限。管理员也会被标记为client，这些身份会被记录在证书的OU字段中。

### 1.3. 账户管理

#### 1.3.1. 证书体系

在Fabric中，证书就是账户，只有具备有效证书才可以向网络发送交易。

每个证书对应一个公私钥对，但是一个公私钥对可以不止一个证书。证书的唯一ID可以由其Subject和Issuer确定。

每个组织需要自行维护证书体系，需要处理组织内证书的签发、回收和记录。

组织的证书体系由一个组织根证书和组织成员证书组成。根证书往往是一个自签证书，用来给组织成员颁发证书，这个过程可以通过Fabric CA工具来实现。组织成员可以分为多种角色，比如管理员、Peer和Client等，可以在CA中添加各种属性，也可以是被根CA签发的中间CA，可以定义各种权限，比如撤销证书、颁发证书和定义属性的权限，FabricCA的可拓展性很强，尤其是可以自定义属性，这些属性可以在不同的业务中发挥权限控制的作用，比如在链码中检查调用者是否具备某个属性。

#### 1.3.2. MSP

面对复杂的证书体系和权限控制，Fabric将这些抽象成一个实体，称为Membership Provider，即MSP。MSP中包含了证明自己或他人权限的所有信息，具体实例如下：

```
Client@org1.example.com
├── msp
│   ├── admincerts // 管理员文件夹，存储着当前组织的管理员证书，这里没有
│   ├── cacerts // 存储当前组织的根证书
│   │   └── ca.org1.example.com-cert.pem
│   ├── config.yaml 
│   ├── keystore // 存储当前账户的私钥
│   │   └── d6a78efa7db939cd68e059a9593cbbb89bc68d14e7caebbbc38e4fb81e0fe579_sk
│   ├── signcerts // 当前账户的证书，由组织根证书颁发
│   │   └── cert.pem
│   └── tlscacerts // TLS的根证书
│       └── tlsca.org1.example.com-cert.pem
└── tls
    ├── ca.crt // TLS的根证书
    ├── client.crt // 本地TLS证书，由TLS根证书颁发
    └── client.key // TLS私钥
```

MSP的实体就是一些文件，包含证书和私钥等，都是PEM格式，这在Fabric中称为LocalMSP，是供本地账户使用的，Peer节点、Orderer节点都会维护一份，对于整个channel来说，还维护了所有组织的根证书等信息，这是存储在账本里的，供全通道所有组织使用。

#### 1.3.3. TLS

Fabric可以开启双向TLS保护网络内信息的隐私，比如证书等信息都不会泄露出去，毕竟证书中包含了很多权限和组织信息，这可能是很关键的商业信息。

在CS模式的TLS协议中，往往是仅有服务端有证书，客户端仅需要在验证服务端证书之后，用服务端公钥加密关键信息后发送给服务端，服务端用私钥解出该信息，双方都在本地生成同样的对称密钥，后续通信使用该对称密钥加密即可。而双向TLS则需要服务端验证客户端的TLS证书，后续流程与单向相同。

## 2. 交易流程

### 2.1. 交易背书

不像以太坊，把交易广播出去就不关用户的事情了，Fabric用户发出交易分两步，第一步交易背书，第二步提交交易到Orderer。

在这里用户指的是一个应用程序，它的CA具备写权限。假设通道有两个组织Org1和Org2，现在要调用链码CC，要发交易调用CC需要满足它的背书策略，这是链码要求的，而背书就是指满足背书策略的过程，举例来讲，就是要求某些组织的Peer为该交易签名。

应用程序首先向Org1和Org2的Peer节点发送自己签名的交易请求，Peer节点会验证应用程序是否有足够的权限，然后执行交易判断是否有效，此时是不写入账本的，如果交易请求通过检验，Peer会为其签名，返回给应用程序。

然后，应用程序收集Peer的返回，组装交易，发送给Orderer节点，Orderer节点会将该交易打包到当前通道的区块中。

区块分发给各个Peer节点，Peer执行交易，同时修改账本状态，此时交易也不一定会成功，因为RSet可能与当前状态不符，交易会被标记为invalid。

具体流程可[见](https://hyperledger-fabric.readthedocs.io/zh_CN/release-1.4/arch-deep-dive.html)

## 3. 链码开发

### 3.1. 背书策略

背书实际上就是组织对某个提案进行检查并签名，需要哪些组织、哪些角色签名，就是背书策略。

背书策略是针对链码来说的，在链码安装或者升级的时候会设置背书策略，这是对链码整体的背书策略，除此之外，还可以针对某些特殊的Key设置策略来覆盖整体背书策略，这称作基于状态的背书策略。

简单举例：`AND('Org1.member', 'Org2.member')`

这个就是背书策略的表达式，简单来看，就是要求该链码的调用要经过Org1和Org2任何成员的背书。详情可[见](https://hyperledger-fabric.readthedocs.io/zh_CN/release-1.4/endorsement-policies.html)

针对键级别的背书策略，需要在编写链码时以代码的形式预先写好，这里有具体的[接口](https://hyperledger-fabric.readthedocs.io/zh_CN/release-1.4/endorsement-policies.html#key-level-endorsement)。

### 3.2. 安装与初始化

在Fabric中链码都是通过docker容器的形式存在的。

安装的命令可以走一遍[教程](https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html)，走完就熟悉了。

安装实际上就是把代码发送到各个Peer节点，实例化会编译链码并为其启动一个容器，Peer节点与其通过gRPC通信。

### 3.3. 标准接口

链码编写比较简单，1.4提供的接口和2.0有较大差距，后面再研究一下。

链码实际上就是实现接口：

```go
// Chaincode interface must be implemented by all chaincodes. The fabric runs
// the transactions by calling these functions as specified.
type Chaincode interface {
	// Init is called during Instantiate transaction after the chaincode container
	// has been established for the first time, allowing the chaincode to
	// initialize its internal data
	Init(stub ChaincodeStubInterface) pb.Response

	// Invoke is called to update or query the ledger in a proposal transaction.
	// Updated state variables are not committed to the ledger until the
	// transaction is committed.
	Invoke(stub ChaincodeStubInterface) pb.Response
}
```

这里作者实现了几个[例子](https://github.com/polynetwork/fabric-contract)，这是[Poly](https://www.poly.network/)跨链协议中关于Fabric端的协议实现。不像solidity合约，单纯一本合约就是一个contract实体，就像是写了一个类或者结构体，在这里需要实现main入口，是完整的可执行程序。

### 3.4. 跨链码调用

在使用跨链码调用的时候，存在一些问题。

首先Fabric不存在合约账户这种概念，链码单纯就是逻辑、业务，当跨链码调用的时候，交易上下文除了参数外没有变化，原始的信息都可以从Proposal中手动解出来。

跨链码的时候，被调用的链码的日志不会写到区块中，具体还需要看一下代码。

## 4. 准入控制

### 4.1. 通道配置解析

在最初创建通道的时候，需要指定配置信息，这里面就有权限定义等内容。

配置交易会生成一个配置区块，配置区块是中除了配置交易不存在其他交易，第一个配置区块就是创世区块。更新配置流程是拉取配置、转换为人可读的格式、修改并提交审核。

这个是整个channel的配置结构，proto的定义如下：

```
message ConfigGroup {
    uint64 version = 1; ## 每次修改增长
    map<string,ConfigGroup> groups = 2; ## 包括Orderer信息和应用信息，比如channelMSP信息修改权限、
    map<string,ConfigValue> values = 3; ## 其中包括很多channel设置，比如hash算法等
    map<string,ConfigPolicy> policies = 4; ## 定义了很多权限策略，满足策略才可以修改对应配置
    string mod_policy = 5; ## 修改当前结构下内容的权限策略
}
```

在解读详细配置之前，首先要理解Fabric配置的修改策略，它是一种分层的策略配置，比如最上层的`ConfigGroup`有自己的`mod_policy`，这个策略对当前这一层适用，比如要修改values里面的hash算法，交易签名就必须满足`mod_policy`，而这个`mod_policy`被定义在`policies`中。

下面来解读一下具体配置文件：

```
{
  "channel_group": {
    "groups": {
      "Application": {},
      "Orderer": {}
    },
    "mod_policy": "Admins", 
    "policies": {
      "Admins": {
        "mod_policy": "Admins",
        "policy": {
          "type": 3,
          "value": {
            "rule": "MAJORITY",
            "sub_policy": "Admins"
          }
        },
        "version": "0"
      },
      "Readers": {
        "mod_policy": "Admins",
        "policy": {
          "type": 3,
          "value": {
            "rule": "ANY",
            "sub_policy": "Readers"
          }
        },
        "version": "0"
      },
      "Writers": {
        "mod_policy": "Admins",
        "policy": {
          "type": 3,
          "value": {
            "rule": "ANY",
            "sub_policy": "Writers"
          }
        },
        "version": "0"
      }
    },
    "values": {
      "BlockDataHashingStructure": {
        "mod_policy": "Admins",
        "value": {
          "width": 4294967295
        },
        "version": "0"
      },
      "Capabilities": {
        "mod_policy": "Admins",
        "value": {
          "capabilities": {
            "V1_4_3": {}
          }
        },
        "version": "0"
      },
      "Consortium": {
        "mod_policy": "Admins",
        "value": {
          "name": "SampleConsortium"
        },
        "version": "0"
      },
      "HashingAlgorithm": {
        "mod_policy": "Admins",
        "value": {
          "name": "SHA256"
        },
        "version": "0"
      },
      "OrdererAddresses": {
        "mod_policy": "/Channel/Orderer/Admins",
        "value": {
          "addresses": [
            "orderer.example.com:7050"
          ]
        },
        "version": "0"
      }
    },
    "version": "0"
  },
  "sequence": "3" ## 每次提交配置加1
}
```

首先，来看`policies`。策略`Admins`如下：

```
      "Admins": {
        "mod_policy": "Admins",
        "policy": {
          "type": 3,
          "value": {
            "rule": "MAJORITY",
            "sub_policy": "Admins"
          }
        },
        "version": "0"
      },
```

其中`policy`结构对应整个策略的内容：

- `type`：总共有四种类型，如下：

```
    enum PolicyType {
        UNKNOWN = 0; // Reserved to check for proper initialization
        SIGNATURE = 1;
        MSP = 2;
        IMPLICIT_META = 3;
    }
```

在实例中是3，代表`ImplicitMetaPolicy`，Fabric对其解释如下：

```protobuf
// ImplicitMetaPolicy is a policy type which depends on the hierarchical nature of the configuration
// It is implicit because the rule is generate implicitly based on the number of sub policies
// It is meta because it depends only on the result of other policies
// When evaluated, this policy iterates over all immediate child sub-groups, retrieves the policy
// of name sub_policy, evaluates the collection and applies the rule.
// For example, with 4 sub-groups, and a policy name of "foo", ImplicitMetaPolicy retrieves
// each sub-group, retrieves policy "foo" for each subgroup, evaluates it, and, in the case of ANY
// 1 satisfied is sufficient, ALL would require 4 signatures, and MAJORITY would require 3 signatures.
message ImplicitMetaPolicy {
    enum Rule {
        ANY = 0;      // Requires any of the sub-policies be satisfied, if no sub-policies exist, always returns true
        ALL = 1;      // Requires all of the sub-policies be satisfied
        MAJORITY = 2; // Requires a strict majority (greater than half) of the sub-policies be satisfied
    }
    string sub_policy = 1;
    Rule rule = 2;
}
```

它这个策略的意思是将会在找下一层的策略中找名叫`sub_policy`的策略来执行，`MAJORITY`意味着要有超过一半的子策略被满足，即`Application`和`Orderer`都必须满足。

- `value`：这部分对应于`ImplicitMetaPolicy`，`sub_policy`就是指下一层要满足的策略名字，rule则是三选一。

下面来看下第一个子组Application的内容：

```json
{
    "Application": {
        "groups": {
            "Org1MSP": {...},
            "Org2MSP": {...}
        },
        "mod_policy": "Admins",
        "policies": {
            "Admins": {
                "mod_policy": "Admins",
                "policy": {
                    "type": 3,
                    "value": {
                        "rule": "MAJORITY",
                        "sub_policy": "Admins"
                    }
                },
                "version": "0"
            },
            "Readers": {
                "mod_policy": "Admins",
                "policy": {
                    "type": 3,
                    "value": {
                        "rule": "ANY",
                        "sub_policy": "Readers"
                    }
                },
                "version": "0"
            },
            "Writers": {
                "mod_policy": "Admins",
                "policy": {
                    "type": 3,
                    "value": {
                        "rule": "ANY",
                        "sub_policy": "Writers"
                    }
                },
                "version": "0"
            }
        },
        "values": {
            "Capabilities": {
                "mod_policy": "Admins",
                "value": {
                    "capabilities": {
                        "V1_4_2": {}
                    }
                },
                "version": "0"
            }
        },
        "version": "1"
    }
}
```

可以看到`Application`这一层也是类似于`channel_group`的，有`groups`、`values`、`policies`、`versions`和`mod_policy`五部分组成，这里就能看出，所谓的分层，就是groups套groups，每层都会定义policy。

关键是`Application`下一层`Org1MSP`和`Org2MSP`的内容，这里展示`Org1MSP`：

```json
{
    "groups": {},
    "mod_policy": "Admins",
    "policies": {
        "Admins": {
            "mod_policy": "Admins",
            "policy": {
                "type": 1,
                "value": {
                    "identities": [{
                        "principal": {
                            "msp_identifier": "Org1MSP",
                            "role": "ADMIN"
                        },
                        "principal_classification": "ROLE"
                    }],
                    "rule": {
                        "n_out_of": {
                            "n": 1,
                            "rules": [{
                                "signed_by": 0
                            }]
                        }
                    },
                    "version": 0
                }
            },
            "version": "0"
        },
        "Readers": {
            "mod_policy": "Admins",
            "policy": {
                "type": 1,
                "value": {
                    "identities": [{
                            "principal": {
                                "msp_identifier": "Org1MSP",
                                "role": "ADMIN"
                            },
                            "principal_classification": "ROLE"
                        },
                        {
                            "principal": {
                                "msp_identifier": "Org1MSP",
                                "role": "PEER"
                            },
                            "principal_classification": "ROLE"
                        },
                        {
                            "principal": {
                                "msp_identifier": "Org1MSP",
                                "role": "CLIENT"
                            },
                            "principal_classification": "ROLE"
                        }
                    ],
                    "rule": {
                        "n_out_of": {
                            "n": 1,
                            "rules": [{
                                    "signed_by": 0
                                },
                                {
                                    "signed_by": 1
                                },
                                {
                                    "signed_by": 2
                                }
                            ]
                        }
                    },
                    "version": 0
                }
            },
            "version": "0"
        },
        "Writers": {
            "mod_policy": "Admins",
            "policy": {
                "type": 1,
                "value": {
                    "identities": [{
                            "principal": {
                                "msp_identifier": "Org1MSP",
                                "role": "ADMIN"
                            },
                            "principal_classification": "ROLE"
                        },
                        {
                            "principal": {
                                "msp_identifier": "Org1MSP",
                                "role": "CLIENT"
                            },
                            "principal_classification": "ROLE"
                        }
                    ],
                    "rule": {
                        "n_out_of": {
                            "n": 1,
                            "rules": [{
                                    "signed_by": 0
                                },
                                {
                                    "signed_by": 1
                                }
                            ]
                        }
                    },
                    "version": 0
                }
            },
            "version": "0"
        }
    },
    "values": {
        "AnchorPeers": {
            "mod_policy": "Admins",
            "value": {
                "anchor_peers": [{
                    "host": "peer0.org1.example.com",
                    "port": 7051
                }]
            },
            "version": "0"
        },
        "MSP": {
            "mod_policy": "Admins",
            "value": {
                "config": {
                    "admins": [],
                    "crypto_config": {
                        "identity_identifier_hash_function": "SHA256",
                        "signature_hash_family": "SHA2"
                    },
                    "fabric_node_ous": {
                        "admin_ou_identifier": {
                            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
                            "organizational_unit_identifier": "admin"
                        },
                        "client_ou_identifier": {
                            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
                            "organizational_unit_identifier": "client"
                        },
                        "enable": true,
                        "orderer_ou_identifier": {
                            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
                            "organizational_unit_identifier": "orderer"
                        },
                        "peer_ou_identifier": {
                            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
                            "organizational_unit_identifier": "peer"
                        }
                    },
                    "intermediate_certs": [],
                    "name": "Org1MSP",
                    "organizational_unit_identifiers": [],
                    "revocation_list": [],
                    "root_certs": [
                        "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
                    ],
                    "signing_identity": null,
                    "tls_intermediate_certs": [],
                    "tls_root_certs": [
                        "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXRENDQWY2Z0F3SUJBZ0lSQUtqL29zM1c5R2FaRzJCT1d0T0NxbUl3Q2dZSUtvWkl6ajBFQXdJd2RqRUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIekFkQmdOVkJBTVRGblJzCmMyTmhMbTl5WnpFdVpYaGhiWEJzWlM1amIyMHdIaGNOTWpBeE1UQTFNRGN3TXpBd1doY05NekF4TVRBek1EY3cKTXpBd1dqQjJNUXN3Q1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRQpCeE1OVTJGdUlFWnlZVzVqYVhOamJ6RVpNQmNHQTFVRUNoTVFiM0puTVM1bGVHRnRjR3hsTG1OdmJURWZNQjBHCkExVUVBeE1XZEd4elkyRXViM0puTVM1bGVHRnRjR3hsTG1OdmJUQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDkKQXdFSEEwSUFCTnlNTjVPaGFVb3NVUWtTVDBodllaOFNSeFJNTnJVZE1mdkkzLy9VcHUyTkJJWG4xNWxSWk9yOQp1akZzNUNFQXlBeGVTVE9neFNxOWloRDJXVXRLejF5amJUQnJNQTRHQTFVZER3RUIvd1FFQXdJQnBqQWRCZ05WCkhTVUVGakFVQmdnckJnRUZCUWNEQWdZSUt3WUJCUVVIQXdFd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBcEJnTlYKSFE0RUlnUWdWdWdaaDV5S3AvWnc4RGFXc2FleWhrZE1OT3A0aFFVbk1UVTJ1UnRNaHlRd0NnWUlLb1pJemowRQpBd0lEU0FBd1JRSWhBTXAvMDRncE5jZEZGSHhzMDhXVmNZbXZuU3kwYUVrdWFlWnc1Y2pLekRwNUFpQXRHcnpJCmR3ZmN2bmNtc0p2NnVCNEhabUtFU3A1ZUVLQ2tsbkNNTGZIeTBnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
                    ]
                },
                "type": 0
            },
            "version": "0"
        }
    },
    "version": "1"
}
```

先从权限策略开始。

这里的`policies`就是很具体的了，以下是`Admins`的内容：

```json
{
    "mod_policy": "Admins",
    "policy": {
        "type": 1,
        "value": {
            "identities": [{
                "principal": {
                    "msp_identifier": "Org1MSP",
                    "role": "ADMIN"
                },
                "principal_classification": "ROLE"
            }],
            "rule": {
                "n_out_of": {
                    "n": 1,
                    "rules": [{
                        "signed_by": 0
                    }]
                }
            },
            "version": 0
        }
    },
    "version": "0"
}
```

结构很清晰了，首先`policy`是一个通用的策略结构，`type`是1，即`SIGNATURE`，`value`对应如下结构：

```protobuf

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
message SignaturePolicyEnvelope {
    int32 version = 1;
    SignaturePolicy rule = 2;
    repeated MSPPrincipal identities = 3;
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
message SignaturePolicy {
    message NOutOf {
        int32 n = 1;
        repeated SignaturePolicy rules = 2;
    }
    oneof Type {
        int32 signed_by = 1;
        NOutOf n_out_of = 2;
    }
}
```

在`SignaturePolicyEnvelope`中，`identities`是指签名的实体列表，`rule`是指验证签名的规则集合，这个集合就是`SignaturePolicy`，它是以递归的形式构造的，可以看到它分为`n_out_of`和`signed_by`两种，当时`n_out_of`的时候，意味着，下面递归的子规则中满足n个就可以，继续递归直到`signed_by`，`signed_by`的意思是这个子规则要求`identities`中的第几个ID签名。

所以策略`Admins`的意思是必须要有一个`Org1MSP`的`ADMIN`角色的签名。

至此权限策略定义就明了了。从第一层`channel_group`的`ImplicitMetaPolicy`开始递归验证到`Org1MSP`这一层的`SignaturePolicyEnvelope`，最终完成验证。

下面介绍一下`values`这一部分。

```json
{
    "AnchorPeers": {
        "mod_policy": "Admins",
        "value": {
            "anchor_peers": [{
                "host": "peer0.org1.example.com",
                "port": 7051
            }]
        },
        "version": "0"
    },
    "MSP": {...},
    "version": "0"
    }
}
```

第一部分`AnchorPeers`是组织1的锚节点列表，锚节点可以和其他组织的锚节点通过P2P通信，基于gossip协议去同步交易等。

第二部分`MSP`，还是直接看protobuf的结构体：

```protobuf
// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
message MSPConfig {
    // Type holds the type of the MSP; the default one would
    // be of type FABRIC implementing an X.509 based provider
    int32 type = 1;

    // Config is MSP dependent configuration info
    bytes config = 2;
}
```

MSP就是前面提到的channelMSP。

`config`内容如下：

```
{
    "admins": [],
    "crypto_config": {
        "identity_identifier_hash_function": "SHA256",
        "signature_hash_family": "SHA2"
    },
    "fabric_node_ous": {
        "admin_ou_identifier": {
            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
            "organizational_unit_identifier": "admin"
        },
        "client_ou_identifier": {
            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
            "organizational_unit_identifier": "client"
        },
        "enable": true,
        "orderer_ou_identifier": {
            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
            "organizational_unit_identifier": "orderer"
        },
        "peer_ou_identifier": {
            "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
            "organizational_unit_identifier": "peer"
        }
    },
    "intermediate_certs": [],
    "name": "Org1MSP",
    "organizational_unit_identifiers": [],
    "revocation_list": [],
    "root_certs": [
        "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNVVENDQWZpZ0F3SUJBZ0lSQU5TMEM5Nkdpb1U1ZWNiMUpUVi9PYmt3Q2dZSUtvWkl6ajBFQXdJd2N6RUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIREFhQmdOVkJBTVRFMk5oCkxtOXlaekV1WlhoaGJYQnNaUzVqYjIwd0hoY05NakF4TVRBMU1EY3dNekF3V2hjTk16QXhNVEF6TURjd016QXcKV2pCek1Rc3dDUVlEVlFRR0V3SlZVekVUTUJFR0ExVUVDQk1LUTJGc2FXWnZjbTVwWVRFV01CUUdBMVVFQnhNTgpVMkZ1SUVaeVlXNWphWE5qYnpFWk1CY0dBMVVFQ2hNUWIzSm5NUzVsZUdGdGNHeGxMbU52YlRFY01Cb0dBMVVFCkF4TVRZMkV1YjNKbk1TNWxlR0Z0Y0d4bExtTnZiVEJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUEKQktvUy85YUxFMW1NdExPclNsdCtESDlTVTNKM2VmUnczTkZsU1JMMXh2dUZ1WkcvanQvZEdXRnZwa3lXZEdOZwpGYS9xcDBTcm1zSjhnSXZuVWhRMTlmU2piVEJyTUE0R0ExVWREd0VCL3dRRUF3SUJwakFkQmdOVkhTVUVGakFVCkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3RXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QXBCZ05WSFE0RUlnUWcKblFvR0JLRk9uYzNUcW84emE4am1qdHFkdXBhdW5NU0ZTSm9TUUgrM0MzRXdDZ1lJS29aSXpqMEVBd0lEUndBdwpSQUlnY2dOOUdUdk85NDZNN2dwbmhJY1RYdXplcDAxdTYxQlZlOXhleEw3K1lEY0NJRWpPR2ZxZnpURkRQMWFaClBvdThUbVoyZmtjYnVZWVNhcHdLRFE3blZtYmoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
    ],
    "signing_identity": null,
    "tls_intermediate_certs": [],
    "tls_root_certs": [
        "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNXRENDQWY2Z0F3SUJBZ0lSQUtqL29zM1c5R2FaRzJCT1d0T0NxbUl3Q2dZSUtvWkl6ajBFQXdJd2RqRUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhHVEFYQmdOVkJBb1RFRzl5WnpFdVpYaGhiWEJzWlM1amIyMHhIekFkQmdOVkJBTVRGblJzCmMyTmhMbTl5WnpFdVpYaGhiWEJzWlM1amIyMHdIaGNOTWpBeE1UQTFNRGN3TXpBd1doY05NekF4TVRBek1EY3cKTXpBd1dqQjJNUXN3Q1FZRFZRUUdFd0pWVXpFVE1CRUdBMVVFQ0JNS1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRQpCeE1OVTJGdUlFWnlZVzVqYVhOamJ6RVpNQmNHQTFVRUNoTVFiM0puTVM1bGVHRnRjR3hsTG1OdmJURWZNQjBHCkExVUVBeE1XZEd4elkyRXViM0puTVM1bGVHRnRjR3hsTG1OdmJUQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDkKQXdFSEEwSUFCTnlNTjVPaGFVb3NVUWtTVDBodllaOFNSeFJNTnJVZE1mdkkzLy9VcHUyTkJJWG4xNWxSWk9yOQp1akZzNUNFQXlBeGVTVE9neFNxOWloRDJXVXRLejF5amJUQnJNQTRHQTFVZER3RUIvd1FFQXdJQnBqQWRCZ05WCkhTVUVGakFVQmdnckJnRUZCUWNEQWdZSUt3WUJCUVVIQXdFd0R3WURWUjBUQVFIL0JBVXdBd0VCL3pBcEJnTlYKSFE0RUlnUWdWdWdaaDV5S3AvWnc4RGFXc2FleWhrZE1OT3A0aFFVbk1UVTJ1UnRNaHlRd0NnWUlLb1pJemowRQpBd0lEU0FBd1JRSWhBTXAvMDRncE5jZEZGSHhzMDhXVmNZbXZuU3kwYUVrdWFlWnc1Y2pLekRwNUFpQXRHcnpJCmR3ZmN2bmNtc0p2NnVCNEhabUtFU3A1ZUVLQ2tsbkNNTGZIeTBnPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
    ]
}
```

proto结构体如下：

```protobuf
// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
message FabricMSPConfig {
    // Name holds the identifier of the MSP; MSP identifier
    // is chosen by the application that governs this MSP.
    // For example, and assuming the default implementation of MSP,
    // that is X.509-based and considers a single Issuer,
    // this can refer to the Subject OU field or the Issuer OU field.
    string name = 1;

    // List of root certificates trusted by this MSP
    // they are used upon certificate validation (see
    // comment for IntermediateCerts below)
    repeated bytes root_certs = 2;

    // List of intermediate certificates trusted by this MSP;
    // they are used upon certificate validation as follows:
    // validation attempts to build a path from the certificate
    // to be validated (which is at one end of the path) and
    // one of the certs in the RootCerts field (which is at
    // the other end of the path). If the path is longer than
    // 2, certificates in the middle are searched within the
    // IntermediateCerts pool
    repeated bytes intermediate_certs = 3;

    // Identity denoting the administrator of this MSP
    repeated bytes admins = 4;

    // Identity revocation list
    repeated bytes revocation_list = 5;

    // SigningIdentity holds information on the signing identity
    // this peer is to use, and which is to be imported by the
    // MSP defined before
    SigningIdentityInfo signing_identity = 6;

    // OrganizationalUnitIdentifiers holds one or more
    // fabric organizational unit identifiers that belong to
    // this MSP configuration
    repeated FabricOUIdentifier organizational_unit_identifiers = 7;

    // FabricCryptoConfig contains the configuration parameters
    // for the cryptographic algorithms used by this MSP
    FabricCryptoConfig crypto_config = 8;

    // List of TLS root certificates trusted by this MSP.
    // They are returned by GetTLSRootCerts.
    repeated bytes tls_root_certs = 9;

    // List of TLS intermediate certificates trusted by this MSP;
    // They are returned by GetTLSIntermediateCerts.
    repeated bytes tls_intermediate_certs = 10;

    // fabric_node_ous contains the configuration to distinguish clients from peers from orderers
    // based on the OUs.
    FabricNodeOUs fabric_node_ous = 11;
}
```

LocalMSP和channelMSP用的同一个结构体。注释中清楚解释了每个字段的含义，这里主要说一下`fabric_node_ous`。可以看到`fabric_node_ous`实际上指定了组织最基础的四个角色`client`、`peer`、`admin`和`orderer`，`certificate`设置了这些角色的证书的发行者，`organizational_unit_identifier`则指定了他们的名字，这个与他们CA的OU中必须要一致。

```protobuf
// FabricNodeOUs contains configuration to tell apart clients from peers from orderers
// based on OUs. If NodeOUs recognition is enabled then an msp identity
// that does not contain any of the specified OU will be considered invalid.
message FabricNodeOUs {
    // If true then an msp identity that does not contain any of the specified OU will be considered invalid.
    bool   enable = 1;

    // OU Identifier of the clients
    FabricOUIdentifier client_ou_identifier = 2;

    // OU Identifier of the peers
    FabricOUIdentifier peer_ou_identifier = 3;

    // OU Identifier of the admins
    FabricOUIdentifier admin_ou_identifier = 4;

    // OU Identifier of the orderers
    FabricOUIdentifier orderer_ou_identifier = 5;
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
message FabricOUIdentifier {

    // Certificate represents the second certificate in a certification chain.
    // (Notice that the first certificate in a certification chain is supposed
    // to be the certificate of an identity).
    // It must correspond to the certificate of root or intermediate CA
    // recognized by the MSP this message belongs to.
    // Starting from this certificate, a certification chain is computed
    // and bound to the OrganizationUnitIdentifier specified
    bytes certificate = 1;

    // OrganizationUnitIdentifier defines the organizational unit under the
    // MSP identified with MSPIdentifier
    string organizational_unit_identifier = 2;
}
```

类似地，`Orderer`组中记录了各种策略和配置，比如排序通道的共识类型，这里是solo。

### 4.2. 访问控制列表（ACL）

在[ACL](https://hyperledger-fabric.readthedocs.io/zh_CN/release-1.4/access_control.html)这部分可以更细粒度控制权限，包括调用链码的某个方法也可以加权限，设置和上面很类似。

添加自己的策略可以通过：初始化通道之前，修改configtx.yaml；或者，通过更新通道的方式，即获取配置区块，修改配置，发送更新配置交易。