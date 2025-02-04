go build splitter.go
cat aphex.md | tail -n 61 | shuf | sort -n | grep -v "##" | grep . | ./splitter
