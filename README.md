# litekube-vpn
A VPN implementation over websockets. This is the client/server implementation of a layer-2 software switch able to route packets over websockets connections. The litekube-vpn is built on top of Linux's tun/tap device.

## Build and Install

Building litekube-vpn needs Go 1.1 or higher.

```shell
go mod tidy
go build -o litekube-vpn main.go
```

## Configuration

There are two config files to distinguish between client and server.

To start server execute the following command:

```shell
litekube-vpn --config server.yml
```

client:

```shell
litekube-vpn --config client.yml
```

### Download

Todo：release

### Network forwarding
On the server the IP forwarding is needed. First we need to be sure that IP forwarding is enabled.
Very often this is disabled by default. This is done by running the following command line as root:

```shell
sysctl -w net.ipv4.ip_forward=1
iptables -t nat -A POSTROUTING -j MASQUERADE
```

So, lets look at the iptables rules required for this to work.
```shell
# Allow TUN interface connections to VPN server
iptables -A INPUT -i tun0 -j ACCEPT

 # Allow TUN interface connections to be forwarded through other interfaces
iptables -A FORWARD -i tun0 -j ACCEPT

iptables -t nat -A POSTROUTING -o tun0 -j MASQUERADE

iptables -A FORWARD -i eth0 -o tun0 -j ACCEPT

iptables -A FORWARD -i tun0 -o eth0 -m state --state RELATED,ESTABLISHED -j ACCEPT
```

## litekube-vpn-explain

if you want to know more about this project, please look at [litekube-vpn-explain](./litekube-vpn-explain.md)
