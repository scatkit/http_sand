package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
  resp, err := http.Get(url)
  if err != nil{
    return []User{}, err
  }
  defer resp.Body.Close()

  var out []User
  err = json.NewDecoder(resp.Body).Decode(&out)
  if err != nil{
    return []User{}, err
  }
  
  return out, nil
}
