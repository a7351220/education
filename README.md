# education


1. 下载并解压 Go 1.15.2：

   ```bash
   $ wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
   $ sudo tar -xvf go1.15.2.linux-amd64.tar.gz
   ```

2. 将 `go` 目录的所有文件的所有权更改为 `root` 用户和组：

   ```bash
   $ sudo chown -R root:root ./go
   ```

3. 将 `go` 目录移动到 `/usr/local`：

   ```bash
   $ sudo mv go /usr/local
   ```

4. 创建一个符号链接，以便您可以在任何地方使用 `go` 命令：

   ```bash
   $ sudo ln -s /usr/local/go/bin/go /usr/bin/go
   ```

5. 验证安装，检查 Go 版本：

   ```bash
   $ go version
   ```

将`GOPATH`设置为`/root/go`,拉取项目：
```
cd $GOPATH/src && git clone https://github.com/sxguan/education.git
```
在`/etc/hosts`中添加：
```
127.0.0.1  orderer.example.com
127.0.0.1  peer0.org1.example.com
127.0.0.1  peer1.org1.example.com
```
添加依赖：
```
cd education && go mod tidy
```
运行项目：
```
./clean_docker.sh
```
在`127.0.0.1:9000`进行访问
