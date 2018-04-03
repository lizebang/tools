# v2ray

config.json download from `https://free-ss.site/`

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

.zshrc add

```config
# proxy
alias proxy='export https_proxy=http://127.0.0.1:1081;export http_proxy=https://127.0.0.1:1081;export socks5_proxy=socks5://127.0.0.1:1080'
alias unproxy='unset https_proxy http_proxy socks5_proxy'
```