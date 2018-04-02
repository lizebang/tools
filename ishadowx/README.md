# ishadowx

ishadowx can help you get free ss accounts.

## Note

Please set `Shadowsocks` to `Manual Mode`.

## Setting proxy from terminal

Add to `$HOME/.bash_profile` or `$HOME/.zshrc`.

```config
# proxy
alias proxy='export https_proxy=http://127.0.0.1:1087;export http_proxy=https://127.0.0.1:1087;export socks5_proxy=socks5://127.0.0.1:1086'
alias unproxy='unset https_proxy http_proxy socks5_proxy'
```

`Usage`

```shell
$proxy
--------
In Proxy
--------
$unproxy
```
