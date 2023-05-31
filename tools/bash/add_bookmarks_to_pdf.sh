#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# https://www.xmodulo.com/add-bookmarks-pdf-document-linux.html

gs -sDEVICE=pdfwrite -q -dBATCH -dNOPAUSE \
  -sOutputFile="Data Structures And Algorithmic Thinking With Go (Narasimha Karumanchi) 2021.pdf" \
  -dPDFSETTINGS=/prepress index.info \
  -f input.pdf

exit

apt-get install ghostscript

cat > index.info <<EOF
[/Page 3 /Title (Acknowledgements) /OUT pdfmark

[/Count 2 /Page 5 /Title (Preface) /OUT pdfmark
  [/Page 5 /Title (Dear Reader) /OUT pdfmark
  [/Page 6 /Title (Other Books by Narasimha Karumanchi) /OUT pdfmark

[/Page 7 /Title (Table of Contents) /OUT pdfmark

[/Page 15 /Title (Chapter 0 Organization of Chapters) /OUT pdfmark

[/Page 27 /Title (Chapter 1 Introduction) /OUT pdfmark

[/Page 50 /Title (Chapter 2 Recursion and Backtracking) /OUT pdfmark

[/Page 57 /Title (Chapter 3 Linked Lists) /OUT pdfmark

[/Page 111 /Title (Chapter 4 Stacks) /OUT pdfmark

[/Page 139 /Title (Chapter 5 Queues) /OUT pdfmark

[/Page 157 /Title (Chapter 6 Trees) /OUT pdfmark

[/Page 235 /Title (Chapter 7 Priority Queues and Heaps) /OUT pdfmark

[/Page 255 /Title (Chapter 8 Disjoint Sets ADT) /OUT pdfmark

[/Page 263 /Title (Chapter 9 Graph Algorithms) /OUT pdfmark

[/Page 313 /Title (Chapter 10 Sorting) /OUT pdfmark

[/Page 341 /Title (Chapter 11 Searching) /OUT pdfmark

[/Page 374 /Title (Chapter 12 Selection Algorithms[Medians]) /OUT pdfmark

[/Page 384 /Title (Chapter 13 Symbol Tables) /OUT pdfmark

[/Page 386 /Title (Chapter 14 Hashing) /OUT pdfmark

[/Page 407 /Title (Chapter 15 String Algorithms) /OUT pdfmark

[/Page 432 /Title (Chapter 16 Algorithms Design Techniques) /OUT pdfmark

[/Page 435 /Title (Chapter 17 Greedy Algorithms) /OUT pdfmark

[/Page 446 /Title (Chapter 18 Divide and Conquer Algorithms) /OUT pdfmark

[/Page 464 /Title (Chapter 19 Dynamic Programming) /OUT pdfmark

[/Page 503 /Title (Chapter 20 Complexity Classes) /OUT pdfmark

[/Page 509 /Title (Chapter 21 Miscellaneous Concepts) /OUT pdfmark

[/Page 520 /Title (References) /OUT pdfmark
EOF
