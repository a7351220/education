# education


1. 下载并解压 Go 1.15.2：

   ```bash
    wget https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
    sudo tar -xvf go1.15.2.linux-amd64.tar.gz
   ```

2. 将 `go` 目录的所有文件的所有权更改为 `root` 用户和组：

   ```bash
    sudo chown -R root:root ./go
   ```

3. 将 `go` 目录移动到 `/usr/local`：

   ```bash
   rm -rf /usr/local/go
    sudo mv go /usr/local
   ```

4. 创建一个符号链接，以便您可以在任何地方使用 `go` 命令：

   ```bash
    rm -rf /usr/bin/go
    sudo ln -s /usr/local/go/bin/go /usr/bin/go
   ```

5. 验证安装，检查 Go 版本：

   ```bash
    go version
   ```
要将 `GOPATH` 设置为 `/root/go`，您需要编辑您的shell配置文件（例如`.bashrc` 或 `.zshrc`）并将相应的环境变量设置添加到其中。以下是如何在Bash shell中设置 `GOPATH` 为 `/root/go` 的步骤：

1. 打开终端。

2. 编辑您的shell配置文件，通常是`~/.bashrc`，您可以使用文本编辑器打开它。例如，使用`vim`编辑器：

   ```
   vim ~/.bashrc
   ```

3. 在文件末尾添加以下行，使用i插入模式：

   ```
   export GOPATH=/root/go
   ```

   这将设置 `GOPATH` 环境变量为 `/root/go`。

4. 保存并关闭文件，按`Esc`跳出，然后按`:wq`退出保存。

5. 让新的配置生效，可以运行以下命令：

   ```
   source ~/.bashrc
   ```

现在，您的`GOPATH` 已经设置为 `/root/go`。您可以通过运行 `echo $GOPATH` 来验证它是否已正确设置。记住，任何时候修改环境变量后，需要在新终端窗口或会话中才能生效。
将`GOPATH`设置为`/root/go`,拉取项目：
```
cd $GOPATH/src && git clone https://github.com/a7351220/education.git
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

例如，运行以下命令以启动 VS Code 以超级用户身份：

bash
```
code --no-sandbox --user-data-dir=/path/to/alternate/directory
```
家目录下的一个子目录： 您可以在您的家目录下创建一个新的子目录，例如 ~/.vscode_user_data，然后将其用作参数：

```
code --no-sandbox --user-data-dir=~/.vscode_user_data
```
