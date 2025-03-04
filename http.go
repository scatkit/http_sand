package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error){
  jsonPayload, err:= json.Marshal(data)
  if err != nil{
    return User{}, err
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
  if err != nil{
    return User{}, err
  }
  
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("X-Api-Key", apiKey)
  
  client := &http.Client{}
  res, err := client.Do(req)
  if err != nil{
    return User{}, err
  }

  defer res.Body.Close()
  var out User
  if err := json.NewDecoder(res.Body).Decode(&out); err != nil{
    return User{}, err
  }
  
  return out, nil
}

