package models

type Runner struct {
  Bib string
  Race *Race
  Name string
  Category *Category
  Result *Result
}
