package main

import (
	"net/http"
	"io/ioutil"
	"strings"
)

type Wiki struct {
	Name string
	Branch string
}

type WikiVersions struct {
    Wikis []Wiki
}


func ListWikis() []Wiki {
	resp, err := http.Get("http://noc.wikimedia.org/conf/wikiversions.dat.txt")
	if err != nil {
		// Do something?
	}
	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Wat
	}
	body := string(rawBody)

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
