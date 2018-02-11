#! /bin/bash
set -euo pipefail

# make sure the standalone makefont command is available
go get github.com/jung-kurt/gofpdf/makefont

# generate fonts from all of the .ttf files under fonts
ENC=$GOPATH/src/github.com/jung-kurt/gofpdf/font/cp1252.map
find fonts -name "*.ttf" | while read line
do
    makefont --embed --enc=$ENC --dst=fonts/ $line
done