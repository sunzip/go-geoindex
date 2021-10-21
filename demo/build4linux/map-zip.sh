#! /bin/bash

rm ../map-nearest.zip
rm ../map-cluster.zip
rm ../map-heatmap.zip

cp ../static-gaode/nearest.html ../map/static-gaode/
mkdir  ../map/static-gaode/效果/
cp -r ../static-gaode/效果/nearest ../map/static-gaode/效果/
zip -r ../map-nearest.zip ../map
rm ../map/static-gaode/nearest.html
rm -r ../map/static-gaode/效果/nearest

cp ../static-gaode/cluster.html ../map/static-gaode/
cp -r ../static-gaode/效果/cluster ../map/static-gaode/效果/
zip -r ../map-cluster.zip ../map
rm ../map/static-gaode/cluster.html
rm -r ../map/static-gaode/效果/cluster

cp ../static-gaode/heatmap.html ../map/static-gaode/
# cp -r ../static-gaode/效果/heatmap ../map/static-gaode/效果/
zip -r ../map-heatmap.zip ../map
rm ../map/static-gaode/heatmap.html
# rm -r ../map/static-gaode/效果/heatmap

rm -r ../map/static-gaode/效果