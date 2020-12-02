function proxyon(){
    export ALL_PROXY=socks5://127.0.0.1:3213
    export http_proxy=http://127.0.0.1:3213
    export https_proxy=https://127.0.0.1:3213
    git config --global http.proxy http://127.0.0.1:3213
    git config --global https.proxy http://127.0.0.1:3213
    echo -e "Proxy On"
}

function proxyoff(){
    unset ALL_PROXY
    unset http_proxy
    unset https_proxy
    git config --global --unset http.proxy
    git config --global --unset https.proxy
    echo -e "Proxy Off"
}