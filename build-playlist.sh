go build splitter.go
cat electronic-2025.md | tail -n 139 | shuf | sort -n | grep -v "##" | grep . | ./splitter
