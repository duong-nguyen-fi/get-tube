#!/bin/bash

id=$1

#mkdir -p "$(pwd)"/tmp
sudo mkdir -p /files
sudo chown $(whoami):$(whoami) /files
rm /files/audio.wav
#docker run -v downloads:/home/cloud-user/files:Z --rm docker.io/nzuong/get-tube:latest ${id} -x --audio-format wav -P /home/cloud-user/files
docker run  --rm docker.io/nzuong/get-tube:latest ${id} --print filename
docker run --mount type=bind,source=/files,target=/files --rm docker.io/nzuong/get-tube:latest ${id} -x --audio-format wav -P /files -o audio.wav