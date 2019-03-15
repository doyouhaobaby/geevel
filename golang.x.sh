#比如先切换到 $GOPATH 的 src 目录，cd $GOPATH/src，然后按需要下载：
#https://www.jianshu.com/p/6fe61053c8aa?utm_campaign=maleskine&utm_content=note&utm_medium=seo_notes&utm_source=recommendation

cd $GOPATH/src 

git clone --depth 1 https://github.com/golang/tools.git golang.org/x/tools
git clone --depth 1 https://github.com/golang/lint.git golang.org/x/lint
git clone --depth 1 https://github.com/golang/net.git golang.org/x/net
git clone --depth 1 https://github.com/golang/sys.git golang.org/x/sys
git clone --depth 1 https://github.com/golang/crypto.git golang.org/x/crypto
git clone --depth 1 https://github.com/golang/text.git golang.org/x/text
git clone --depth 1 https://github.com/golang/image.git golang.org/x/image
git clone --depth 1 https://github.com/golang/oauth2.git golang.org/x/oauth2
