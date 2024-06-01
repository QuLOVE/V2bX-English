# V2bX-English

[![](https://img.shields.io/badge/TgChat-%E4%BA%A4%E6%B5%81%E7%BE%A4-blue)](https://t.me/YuzukiProjects)

A multi-core V2board node server, modified from XrayR, supporting V2ay, Trojan, and Shadowsocks protocols.

**Note: This project requires V2board version >= 1.7.0**

## Features

* Perpetually open source and free.
* Supports multiple protocols: Vmess/Vless, Trojan, Shadowsocks, and Hysteria.
* Supports new features such as Vless and XTLS.
* Single instance can handle multiple nodes, eliminating the need for repeated startup.
* Supports online IP limiting.
* Supports limiting the number of TCP connections.
* Supports port-level and user-level speed limiting.
* Simple and clear configuration.
* Automatically restarts instances when the configuration is modified.
* Supports multiple cores, easy to extend.
* Supports conditional compilation, allowing compilation of only the required cores.

## Functions

| Function        | v2ray | trojan | shadowsocks | hysteria |
|-----------|-------|--------|-------------|----------|
| Automatic TLS certificate application | √     | √      | √           | √        |
| Automatic TLS certificate renewal | √     | √      | √           | √        |
| Online user statistics    | √     | √      | √           | √        |
| Auditing Rules      | √     | √      | √           | √         |
| Custom DNS    | √     | √      | √           | √        |
| Online IP Limit   | √     | √      | √           | √        |
| Connection Limit     | √     | √      | √           | √         |
| Cross-node IP limit  |      |       |            |          |
| User-based Speed Limit    | √     | √      | √           | √         |
| Dynamic Speed Limit (Untested) | √     | √      | √           | √         |

## TODO

- [ ] Reimplement dynamic speed limit
- [ ] Reimplement online IP synchronization (cross-node online IP restriction)
- [ ] Improve documentation

## Installation

### One-click installation

```
wget -N wget -N https://raw.githubusercontents.com/QuLOVE/V2bX-script/master/install.sh && bash install.sh
```

### Manual installation

[Manual installation tutorial (outdated, to be updated)](https://yuzuki-1.gitbook.io/v2bx-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## Build
``` bash
# Specify the kernel to compile through the -tags option, optional: xray, sing
go build -o V2bX -ldflags '-s -w' -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD} -tags "xray sing"
```

## Configuration file and detailed usage tutorial

[Detailed usage tutorial](https://yuzuki-1.gitbook.io/v2bx-doc/)

## Disclaimer

* This project is for my personal use, so I cannot guarantee backward compatibility.
* Due to my limited ability, I cannot guarantee the availability of all functions. If you encounter any problems, please report them in Issues.
* I am not responsible for any consequences caused by anyone using this project.
* I am quite fickle, so this project may change the project structure or refactor the code on a large scale as ideas or thoughts change. If you cannot accept this, please do not use it.

## Thanks

* [V2bX fork of wyx2685](https://github.com/wyx2685/V2bX)
* [Project X](https://github.com/XTLS/)
* [V2Fly](https://github.com/v2fly)
* [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
* [Air-Universe](https://github.com/crossfw/Air-Universe)
* [XrayR](https://github.com/XrayR/XrayR)
* [sing-box](https://github.com/SagerNet/sing-box)

## Stars growth record

[![Stargazers over time](https://starchart.cc/InazumaV/V2bX.svg)](https://starchart.cc/InazumaV/V2bX)
