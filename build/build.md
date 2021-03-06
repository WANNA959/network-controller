# Build Document

## simple compile

Building network-controller needs Go 1.1 or higher.

```shell
go mod tidy
# network-controller build
cd ./cmd/network-controller
go build -o network-controller main.go

# ncadm build
cd ./cmd/ncadm
go build -o ncadm main.go
```

## container-based compile

> Ref：https://github.com/Litekube/LiteKube/tree/main/build

We still recommend that you use container-based compilation, and we provide a [Dockerfile](./Dockerfile) to help you build your image. You can follow the rules for writing DockerFile and compile them to your schema, which may require a little knowledge of container. The current version of the image, at least, provides a native go-compilation environment and a cross-compilation environment for the `armv7l` architecture. You can run by `Docker`as follow:

```shell
# assum you start your work in /mywork/

# download project first
cd /mywork
git clone https://github.com/Litekube/network-controller.git

# build docker image
cd /mywork/network-controller/build/

# pre-download & COPY or use proxy
# pre-download & COPY: https://github.com/WANNA959/network-controller/blob/main/build/Dockerfile
wget http://releases.linaro.org/components/toolchain/binaries/7.5-2019.12/arm-linux-gnueabihf/gcc-linaro-7.5.0-2019.12-i686_arm-linux-gnueabihf.tar.xz 
docker build -t litekube/centos-go:v1 .

# if you need proxy, you can use proxy of your host-device and run: https://github.com/Litekube/LiteKube/blob/main/build/Dockerfile
docker build --network=host -t  litekube/centos-go:v1 .

# run container
chmox +x /mywork/network-controller/build/build.sh
docker run -v /mywork/networker-controller:/LiteKube --name=compile-litekube litekube/centos-go:v1 /LiteKube/build/build.sh

# demo
docker run -v /root/go_project/network-controller:/LiteKube --name=compile-litekube litekube/centos-go:v1 /LiteKube/build/build.sh
```

then you can view binary in `/mywork/network-container/build/outputs/`
