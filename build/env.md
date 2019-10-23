
# zsh
```bash
sudo apt install zsh
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
vi ~/.zshrc
mh
source
erport
```

# protoc[https://github.com/protocolbuffers/protobuf/releases]


# rust
curl https://sh.rustup.rs -sSf | sh

# chocolatey[https://chocolatey.org]
```bash
ChocolateyInstall = xxxx/Chocolatey
cmd:
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
ps:
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
choco upgrade chocolatey
```


# minikube[https://kubernetes.io/docs/tasks/tools/install-kubectl/]
minikube的ssh有问题，不能发送esc
- https://github.com/kubernetes/minikube/releases/latest/download/minikube-installer.exe
```bash
win:
choco install minikube
unix:
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube
sudo install minikube /usr/local/bin

minikube start --registry-mirror=https://registry.docker-cn.com --vm-driver=hyperv --hyperv-virtual-switch "Default Switch"
```


# bazel[https://bazel.build,https://github.com/bazelbuild/bazel/releases]


# node
wget https://nodejs.org/dist/v12.3.1/node-v12.3.1.tar.gz
tar -xzvf
./configure
apt install g++
wget https://nodejs.org/dist/v12.3.1/node-v12.3.1-linux-x64.tar.xz

tar -Jxvf


# yarn
wget https://yarnpkg.com/latest.tar.gz

tar zvxf latest.tar.gz


# git
git config --global user.name ${username}
git config --global user.email ${email}
--global credential.helper store

# ssh

powershell ssh
```
Add-WindowsCapability -Online -Name OpenSSH-Client
  ssh root@IP  -p PORT -i .\.ssh\id_rsa
ssh -R [local port]:[remote host]:[remote port] [SSH hostname]
ssh  -fNg -L <本地端口>:<服务器数据库地址>  <用户名>@<服务器地址>
想让SSH连接一直连接，可以加上 -NTf 参数。
exit
```

# sublimet text
```json
{
    "hot_exit": false,
	"remember_open_files": false,
	"default_line_ending": "unix"
}
```