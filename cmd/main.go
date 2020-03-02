package main

import (
	"fmt"
	"github.com/zornh94/mongostore"
)

func main(){
	store := mongostore.NewStore()
	fmt.Printf("healthy: %t", store.Healthy())
}
