#!/usr/bin/env bash
echo '请输入要提交的注释信息:'

read msg

git pull origin master

git add .

git commit -a -m "$msg"

git push origin master



