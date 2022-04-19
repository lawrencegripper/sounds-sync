#!/bin/bash

DIR="$(dirname "${BASH_SOURCE[0]}")"
cd $DIR
for f in ../responses/*.json; do
    echo $f
    base_name=$(basename ${f} .json)
    GENERATED_FILE=../generated/$base_name.go
    json2struct -f $f > $GENERATED_FILE
    sed -i "s/JSONToStruct/$base_name/g" $GENERATED_FILE
    echo "Generated $GENERATED_FILE"
done
