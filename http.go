package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
  
  jsonPayload, err := json.Marshal(data)
  if err != nil{
    return User{}, nil
  }
  
  req, err := http.NewRequest("PUT", fullURL, bytes.NewReader(jsonPayload))
  if err != nil{
    return User{}, nil
  }
  
  req.Header.Set("X-Api-Key", apiKey )
  
  cl := &http.Client{}
  resp, err := cl.Do(req)
  if err != nil{
    return User{}, nil
  }
  
  defer resp.Body.Close()
  
  var out User
  if err := json.NewDecoder(resp.Body).Decode(&out); err != nil{
    return User{}, nil
  }
  
  return out, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id
  req, err := http.NewRequest("GET", fullURL, nil)
  if err != nil{
    return User{}, nil
  }
  
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Api-Key", apiKey )
  
  cl := &http.Client{}
  resp, err := cl.Do(req)
  if err != nil{
    return User{}, nil
  }
  defer resp.Body.Close()
  
  var out User
  if err := json.NewDecoder(resp.Body).Decode(&out); err != nil{
    return User{}, nil
  }
  
  return out, nil
}
