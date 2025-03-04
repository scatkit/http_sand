package main
import "fmt"

func fetchTasks(baseURL, availability string) []Issue {
  
  var limit int
  switch availability{
  case "Low":
    limit = 1
  case "Medium":
    limit = 3
  default:
    limit = 5
  }

  query := fmt.Sprintf("?sort=estimate&limit=%v", limit)
	fullURL := baseURL + query
  fmt.Println(fullURL)
	return getIssues(fullURL)
}
