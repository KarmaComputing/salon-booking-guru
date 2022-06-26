#!/bin/bash

cat $1 | docker run --rm -i think/plantuml -tpng > "$(echo $1 | sed -r "s/(.*)\..*/\1/").png"
