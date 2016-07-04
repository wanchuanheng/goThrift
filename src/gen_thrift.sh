#!/bin/bash

thrift_bin=/home/qthd/wanchuanheng/trunk/common/bin/thrift


for loop in `ls -l |grep -v ^d |grep -v ".sh"`
 do   
    filename=`echo $loop|grep ".thrift"|awk "{print $9}"`
    if [ -n "$filename" ];then
        $thrift_bin --gen go  -out ./  $filename
    fi
done
