package models

type User struct {
  Id int
  Username string
}

type UserWithAddress struct {
  Id int
  Username string
  Address string
} 

type UserWithEmail struct {
  Id int
  Username string
  Email string
}
