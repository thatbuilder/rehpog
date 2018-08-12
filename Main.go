package main

import (
"net/http"
"github.com/tmli3b3rm4n/router"
"thatBuilder/models"
)

type Action func(out http.ResponseWriter, req router.Request)

type Resource map[string] Action
type Users map[models.User]Action
type User map[string]User

var user = map[string]Resource{

}

func main() {

}

