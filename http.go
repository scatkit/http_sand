package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id 
  
  req, err := http.NewRequest("DELETE", fullURL, nil)
  if err != nil{
    return err
  }
  
  req.Header.Set("X-Api-Key", apiKey)
  
  cl := &http.Client{}
  
  resp, err := cl.Do(req)
  if err != nil{
    return err
  }
  
  defer resp.Body.Close()
  
  if resp.StatusCode > 299{
    return fmt.Errorf("User not deleted: %d", resp.StatusCode)
  }
  
  return nil
}
