#!/usr/bin/env bash

bins=`pwd`/bin/.

mkdir -p $bins

for d in ./utils/*; do
  pushd $d
  go build -o $bins
  popd
done
