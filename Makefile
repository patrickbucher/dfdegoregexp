.PHONY: all clean

all: dfdegoregexp.zip dfdegoregexp.txt dfdegoregexp.man

dfdegoregexp.zip: go.mod go.sum *.go replexp/*.go godocfuncs/*.go manperf/*.go
	rm -f dfdegoregexp
	mkdir dfdegoregexp
	cp -r replexp dfdegoregexp/
	cp -r manperf dfdegoregexp/
	cp -r godocfuncs dfdegoregexp/
	cp go.mod go.sum *.go dfdegoregexp/
	zip $@ -r dfdegoregexp
	rm -rf dfdegoregexp

dfdegoregexp.txt: README.md
	pandoc -s -t plain $< -o $@

dfdegoregexp.man: README.md
	pandoc -V title=dfdegoregexp -V section=7 -s -t man $< -o $@

clean:
	rm -rf dfdegoregexp *.zip *.txt *.man
