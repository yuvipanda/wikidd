package main

import (
    "net/http"
    "io/ioutil"
)

func GetHttpResource(URL string) string {
    resp, err := http.Get(URL)
    if err != nil {
        // Do something?
    }
    defer resp.Body.Close()

    rawBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // Fine, fine, I'll do something
    }

    return string(rawBody)
}

