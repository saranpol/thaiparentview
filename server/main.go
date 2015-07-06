package main

import (
    "fmt"
    "net/http"

    "time"
    "appengine"
    "appengine/datastore"
)

// User
type User struct {
    Name *datastore.Key // Text
    Username string
    Email string
    PasswordMD5 string
    Photo *datastore.Key // Image
}


// Content

type Review struct {
    Owner *datastore.Key // User 
    Type *datastore.Key // ReviewType
    Title *datastore.Key // Text
    Name *datastore.Key // Text
    TitleURL string
    Location *datastore.Key // Text
    Position appengine.GeoPoint
    Created time.Time
    CriteriaList []*datastore.Key // []Criteria
}

type Criteria struct {
    Type *datastore.Key // CriteriaType
    Star float32
    Description *datastore.Key // TextLong
}

type ReviewType struct {
    Name *datastore.Key // Text
}

type CriteriaType struct {
    Name *datastore.Key // Text
}

type GroupType struct {
    Name *datastore.Key // Text
}

type ProgramType struct {
    Name *datastore.Key // Text
}

type Text struct {
    TH string
    EN string
}

type TextLong struct {
    TH string
    EN string
}

type Image struct {
    Owner *datastore.Key // User
    Original appengine.BlobKey
    W50 appengine.BlobKey
    W100 appengine.BlobKey
    W200 appengine.BlobKey
    W400 appengine.BlobKey
    W600 appengine.BlobKey
    W800 appengine.BlobKey
    W1000 appengine.BlobKey
    W1200 appengine.BlobKey
}



func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h1>Thai Parent View</h1>")
}


func init() {
    http.HandleFunc("/", index)
}
