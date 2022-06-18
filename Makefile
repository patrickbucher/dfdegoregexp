.PHONY: clean

dfdegoregexp.zip: go.mod go.sum *.go replexp/*.go godocfuncs/*.go manperf/*.go
	rm -f dfdegoregexp
	mkdir dfdegoregexp
	cp -r replexp dfdegoregexp/
	cp -r manperf dfdegoregexp/
	cp -r godocfuncs dfdegoregexp/
	cp go.mod go.sum *.go dfdegoregexp/
	zip $@ -r dfdegoregexp
	rm -rf dfdegoregexp

clean:
	rm -rf dfdegoregexp
	rm -f dfdegoregexp.zip
