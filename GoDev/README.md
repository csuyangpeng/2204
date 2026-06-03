# Lite5GC

Lite 5g core development.

## How to setup your development environment?
1. Prepare your golang enviroment on your dev terminal.
download the [Golong](https://golang.google.cn/dl/) and install it. 
set the **GOROOT** and **GOPATH**.

2. Download the 5gc code on your src folder.
> git clone http://10.18.1.2:9999/develop/lite5gc.git

3. Swith to your devep branch and start code design.
> git checkout <your branch name>

## TroubleShooting
1. How to solve merge conflict?

> In yourbranch
``` 
> git fetch origin master:master
> git rebase master
```
Then will meet the CONFLICTS, revise all the conflict.
```
> git add .
> git rebase --continue
> git push origin yourbranch
```

## UPF build
Add at the end of go.mod
```
go mod init lite5gc
```
replace github.com/intel-go/nff-go v0.9.2 => ./3rdParty/github.com/intel-go/nff-go

Build NFF
```
sudo apt install libnuma-dev libpcap-dev liblua5.3-dev
cd lite5gc/3rdParty/github.com/intel-go/nff-go
make -j
```

Build upf
```
cd lite5gc/upf
make
```

git submodule update --init --recursive
