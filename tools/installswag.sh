
# Ubuntu 2204
GOPATH=$(go env GOPATH)
echo "install swag to $GOPATH"

if [ ! -f $GOPATH/bin/swag ]; then
    mkdir tmp
    pushd tmp
    wget https://github.com/swaggo/swag/releases/download/v1.16.4/swag_1.16.4_Linux_x86_64.tar.gz
    tar -zxvf swag_1.16.4_Linux_x86_64.tar.gz
    rm -f swag_1.16.4_Linux_x86_64.tar.gz
    mv swag $GOPATH/bin/swag
    echo "swag installed!"
    popd
    rm -rf tmp
else
    echo "swag has been installed before!"
fi