#!/bin/bash

# Variables - https://pandoc.org/chunkedhtml-demo/6.2-variables.html

go run main.go work.yaml > README.md
pandoc -f markdown \
    --pdf-engine=pdflatex \
    --variable fontsize:10pt \
    --variable geometry:"top=0.75in, bottom=0.75in, left=0.75in, right=0.75in" \
    --variable linkcolor:blue \
    -o ChrisCampbellCV.pdf \
    README.md