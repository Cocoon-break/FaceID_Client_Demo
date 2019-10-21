#!/bin/bash

#minions
echo 'request verify -----------------------------------------'
curl -F image_idcard=@a.jpg -F image_best=@a.jpg http://10.201.102.109:9000/verify

echo 'request extract_with_detect -----------------------------------------'
curl -F image=@a.jpg http://10.201.102.109:9000/extract_with_detect

echo 'request compare -----------------------------------------'
curl -X POST -H 'Content-Type:application/json' --data-binary @feat.json http://10.201.102.109:9000/compare

#ocr
echo 'request ocr-----------------------------------------'
curl -F img=@idcard_front.jpg -F legality=1 http://10.201.102.109:9002/faceid/v1/idcard

#FMP
echo 'request fmp-----------------------------------------'
curl -F img=@a.jpg -F rotate=1 http://10.201.102.247:9001/faceid/v1/liveness_cfg

#Group
echo 'request group init-----------------------------------------'
curl -F groupname=test http://10.201.102.247:9090/group/init

echo 'request group add-----------------------------------------'
feat='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'
curl -F groupname=test -F feature=$feat http://10.201.102.247:9090/group/add

echo 'request group search-----------------------------------------'
curl -F groupname=test -F feature=$feat http://10.201.102.247:9090/group/search

