package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}


type SnippetModel struct {
  DB *sql.DB
}

func(m *SnippetModel) Insert(title, content string, expires int)(int, error) {
  stmt := `insert into snippets(title, content, created, expires)
  values(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := m.DB.Exec(stmt, title, content, expires)
  if err != nil {
    return 0, err
  }
  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }
  return int(id), err
}

func(m *SnippetModel) Get(id int)(*Snippet, error) {
  stmt := `select id, title, content, created, expires from snippets
  where expires > UTC_TIMESTAMP() and id = ?`
  row := m.DB.QueryRow(stmt, id)
  s := &Snippet{}
  // Notice that the arguments to row.Scan are *pointers* to the place you want to copy the data into,
  // and the number of arguments must be exactly the same as
  // the number of columns returned by your statement.
  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNoRecord
    }else {
      return nil, err
    }
  }
  return s, nil
}

// This will return the 10 most recently created snippet.
func(m *SnippetModel) Latest()([]*Snippet, error) {
  stmt := `select id, title, content, created, expires from snippets
  where expires > UTC_TIMESTAMP() order by id desc limit 10`
  rows, err := m.DB.Query(stmt)
  if err != nil {
    return nil, err
  }
  // This defer statement should come *after* you check for an error from the Query() method.
  // Otherwise, if Query() returns an error, you'll get a panic trying to close a nil resultset.
  defer rows.Close()

  snippets := []*Snippet{}
  for rows.Next() {
    s := &Snippet{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
      return nil, err
    }
    snippets = append(snippets, s)
  }
  // When the rows.Next() loop has finished we call rows.Err() to retrieve any 
  // error that was encountered during the iteration. It's important to 
  // call this - don't assume that a successful iteration was completed 
  // over the whole resultset
  if err = rows.Err(); err != nil {
    return nil, err
  }
  return snippets, nil
}
