#!/bin/bash
list=""
type=`echo $1 | cut -d ':' -f 1`
name=`echo $1 | cut -d ':' -f 2`

if [ "$type" == "local" ]
then
  list="backend game master gate api"
elif [ "$type" == "backend" ]
then
  list="backend"
elif [ "$type" == "cluster" ]
then
  list="game master gate api"
else
  echo "error init type [$type]"
  exit 1
fi

echo "init server by type [$type]"

if [ -n "$name" ]; then
  if [[ $list != *"$name"* ]]; then
    echo "container name [$name] not match"
    exit 1
  fi
else
  name=$list
fi

echo "container name [$name]"

cd `dirname $0`

#检查容器是否在运行
docker ps --format '{{.Names}}' | grep "^slot-*"
if [ $? -eq 0 ]; then
    # 容器正在运行，重启
    echo "Restarting container: ${name}"
    docker-compose restart ${name}
else
    # 容器存在但未运行，启动容器
    echo "Starting container: ${name}"
    docker-compose up -d ${name}
fi
