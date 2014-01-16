package main

import (
	"fmt"
)

func main() {
    wikis := ListWikis()
    for _, wiki := range wikis {
        fmt.Println(wiki)
    }
}
