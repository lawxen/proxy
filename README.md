```
Host github.com
    HostName github.com
    User git
    Proxycommand /usr/bin/ncat --proxy 127.0.0.1:3213 --proxy-type http %h %p

Host gitlab.com
    HostName gitlab.com
    User git
    Proxycommand /usr/bin/ncat --proxy 127.0.0.1:3213 --proxy-type http %h %p

```