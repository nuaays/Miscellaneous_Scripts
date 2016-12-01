


##官网
https://geth.ethereum.org/

##GitHub
https://github.com/ethereum/go-ethereum

##下载
https://geth.ethereum.org/downloads/

##简单试玩
* 0. compile geth or download geth binary
```
https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.5.4-b70acf3c.tar.gz 
```
* 1. create genesis.json
```
```
* 2. init
```
./geth --datadir datadir init genesis.json
```
* 3. exec geth
```
./geth --datadir datadir/ --rpc --rpcapi eth,admin,personal,miner,txpool,net --maxpeers 25 --networkid 912912 --port 30000 --rpcport 8000 --rpcaddr 0.0.0.0  console
```
* 4. create personal accout
```
> personal.newAccount('xxxxx')
"0x8d3b95ca14d7b6e75699c830142647c11bda007e"
```
* 5. see nodeinfo
```
admin.nodeInfo

{
  enode: "enode://1c29dd0ea646a1fa4c0228829d4888a94358677c1969f63c5b66a16ef842fe8073abb23a305f5a5cc848e1902329fd1557924b29acad5b417190b1e67cf0df3f@[::]:30000",
  id: "1c29dd0ea646a1fa4c0228829d4888a94358677c1969f63c5b66a16ef842fe8073abb23a305f5a5cc848e1902329fd1557924b29acad5b417190b1e67cf0df3f",
  ip: "::",
  listenAddr: "[::]:30000",
  name: "Geth/v1.5.5-unstable/linux/go1.7.1",
  ports: {
    discovery: 30000,
    listener: 30000
  },
  protocols: {
    eth: {
      difficulty: 131072,
      genesis: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
      head: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
      network: 912912
    }
  }
}
```

* 6. addPeer this node in other machines
```
admin.addPeer("enode://1c29dd0ea646a1fa4c0228829d4888a94358677c1969f63c5b66a16ef842fe8073abb23a305f5a5cc848e1902329fd1557924b29acad5b417190b1e67cf0df3f@${THIS_HOST_IP}:30000")
```
* 7. check peers
```
admin.peers
```
* 8. check network and net.peerCount
```
net.listening
net.peerCount
```
* 9. start to miner
```
miner
miner.start(1)
```
* 10. check eth block number
```
eth.blockNumber
```
