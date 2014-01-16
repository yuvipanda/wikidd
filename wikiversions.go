package main

import (
	"strings"
)

type Wiki struct {
	Name string
	Branch string
}


func ListWikis() []Wiki {
    body := GetHttpResource("http://noc.wikimedia.org/conf/wikiversions.dat.txt")

    // Handle extra newline at the end of the file
	wikisData := strings.Split(strings.TrimRight(body, "\n"), "\n")
    wikis := make([]Wiki, len(wikisData))

	for i, wiki := range wikisData {
		wikiData := strings.Split(wiki, " ");
		wiki := Wiki{wikiData[0], strings.Replace(wikiData[1], "php-", "", 1)}
        wikis[i] = wiki
	}

    return wikis
}
