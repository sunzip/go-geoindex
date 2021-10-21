#! /bin/bash

mkdir -p ../map/static
mkdir -p ../map/static-gaode

# not copy
cp -r ../static ../map/
cp ../static-gaode/* ../map/static-gaode/

mv ../server ../map
mv ../server.exe ../map
cp ../index.html ../map
cp ../README.md ../map

rm ../map/static-gaode/nearest.html
rm ../map/static-gaode/heatmap.html
rm ../map/static-gaode/cluster.html