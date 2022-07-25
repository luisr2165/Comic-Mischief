package main

import "fmt"

func main() {

var publisher string
publisher = "DizzyBooks Publishing Inc."

var writer string
writer = "Tracey Hatchet"

var artist string
artist = "Jewel Tampson"

var title string 
title = "Mr. GoToSleep"

var year uint
year = 1997

var pageNumber uint
pageNumber = 14

var grade float32
grade = 6.5

fmt.Println(title, "written by", writer, "drawn by", artist, "in the year", year, "from the publisher", publisher, "It has", pageNumber, "pages and its condition grade is", grade, ".")

title = "Epic Vol. 1"

writer = "Ryan N. Shawn"

artist = "Phoebe Paperclips"

publisher = "DizzyBooks Publishing Inc."

year = 2013

pageNumber = 160

grade = 9.0

fmt.Println(title, "written by", writer, "drawn by", artist, "in the year", year, "from the publisher.", publisher, "It has", pageNumber, "pages and its condition grade is", grade, ".")

}
