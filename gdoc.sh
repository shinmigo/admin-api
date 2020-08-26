#!/bin/bash

#------生成文档，本地需要安装swagger
#下载地址：https://github.com/go-swagger/go-swagger/releases

echo "正在生成文档..."
swagger serve --no-open -F=swagger ./docs/swagger.yaml > ./docs/buffer.txt 2>&1 &
sleep 2
port=$(cut -d ":" -f 5 ./docs/buffer.txt | cut -d "/" -f 1)
rm -f ./docs/buffer.txt
curl -so ./static/swagger/swagger.json http://localhost:${port}/swagger.json

pkill -9 swagger
echo "文档成功完毕..."