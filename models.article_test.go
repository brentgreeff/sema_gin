package main

import "testing"

func TestGetAllArticles(t *testing.T) {
  want := len(articleList)

  if got := len( getAllArticles() ); got != want {
    t.Errorf("len( getAllArticles() ) = %d, want %d", got, want)
  }

  for i, v := range getAllArticles() {
    if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
      t.Fail()
      break
    }
  }
}
