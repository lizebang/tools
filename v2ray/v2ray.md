# v2ray

## Installation

V2Ray: https://github.com/v2ray/v2ray-core

Mac: https://github.com/v2ray/homebrew-v2ray

Windows: https://github.com/2dust/v2rayN/

Android: https://github.com/V2Ray-Android/Actinium

## Configuration

config.json download from `http://free-ss.tk/`

config.json add http

```json
  "inboundDetour": [{
    "protocol": "http",
    "port": 1081,
    "settings": {
      "udp": true
    }
  }],
```
add to `.zshrc`/`.bash_profile`

```config
# proxy
alias proxy='export http_proxy=http://127.0.0.1:1081;export https_proxy=https://127.0.0.1:1081;export socks5_proxy=socks5://127.0.0.1:1080'
alias unproxy='unset http_proxy https_proxy socks5_proxy'
```
