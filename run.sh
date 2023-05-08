#!/bin/bash

go run main.go work.yaml |
    pandoc -f markdown \
    --pdf-engine=pdflatex \
    -V mainfont:"Helvetica Neue" \
    -V fontsize:10pt \
    -V geometry:"top=0.75in, bottom=0.75in, left=0.75in, right=0.75in" \
    -V linkcolor:blue \
    -o cv.pdf - 