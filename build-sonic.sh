go build splitter.go
cat sonicpieces.md | tail -n 49 | shuf | sort -n | grep -v "##" | grep . | ./splitter
