package main

import "gorm.io/gorm"

type Game struct {
    gorm.Model
    Name        string
    Description string
    Price       int
    ImageURL    string
}