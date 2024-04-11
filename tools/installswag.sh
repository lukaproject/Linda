
GOPATH=$(go env GOPATH)
echo "install swag to $GOPATH"

if [ ! -f $GOPATH/bin/swag ]; then
    wget https://github.com/swaggo/swag/releases/download/v1.16.3/swag_1.16.3_Linux_amd64.tar.gz
    tar -zxvf swag_1.16.3_Linux_amd64.tar.gz
    rm -f swag_1.16.3_Linux_amd64.tar.gz
    mv swag $GOPATH/bin/swag
    echo "swag installed!"
else
    echo "swag has been installed before!"
fi