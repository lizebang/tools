# ishadowx

ishadowx can help you get free ss accounts.

## Note

Please set `Shadowsocks` to `Manual Mode`.

## Setting proxy from terminal

Add to `$HOME/.bash_profile` or `$HOME/.zshrc`.

```config
alias proxy='export https_proxy=http://127.0.0.1:1087;export http_proxy=https://127.0.0.1:1087;export all_proxy=socks5://127.0.0.1:1086'
alias unproxy='unset https_proxy http_proxy all_proxy'
```

`Usage`

```shell
$proxy
--------
In Proxy
--------
$unproxy
```
