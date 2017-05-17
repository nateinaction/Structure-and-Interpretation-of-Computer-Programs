package main

import (
  "fmt"
  "math"
)

func main() {
  a := makePoint(0,0)
  b := makePoint(2,0)
  c := makePoint(2,8)
  d := makePoint(0,8)

  z := makeSegment(a,b)
  midPoint := getMidPoint(z)
  length := getLength(z)

  fmt.Println(midPoint)
  fmt.Println(length)

  y := makeRectangle(a, b, c, d)
  area := getArea(y)
  perimeter := getPerimeter(y)

  fmt.Println(area)
  fmt.Println(perimeter)
}

func makeRectangle(a, b, c, d Point) Rectangle {
  return Rectangle{a, b, c, d}
}

func getArea(r Rectangle) float64 {
  l := makeSegment(r.a, r.b)
  w := makeSegment(r.b, r.c)
  length := getLength(l)
  width := getLength(w)
  return (length * width)
}

func getPerimeter(r Rectangle) float64 {
  a := makeSegment(r.a, r.b)
  b := makeSegment(r.b, r.c)
  c := makeSegment(r.c, r.d)
  d := makeSegment(r.d, r.a)
  lengthA := getLength(a)
  lengthB := getLength(b)
  lengthC := getLength(c)
  lengthD := getLength(d)
  return (lengthA + lengthB + lengthC + lengthD)
}

func getLength(z Line) float64 {
  x := math.Pow(float64(z.end.x - z.start.x), 2)
  y := math.Pow(float64(z.end.y - z.start.y), 2)
  return math.Sqrt(x + y)
}

func getMidPoint(z Line) Point {
  x := (z.start.x + z.end.x) / 2
  y := (z.start.y + z.end.y) / 2
  return Point{x, y}
}

func makeSegment(a, b Point) Line {
  return Line{a, b}
}

func startSegment(z Line) Point {
  return z.start
}

func endSegment(z Line) Point {
  return z.end
}

func makePoint(x, y int) Point {
  return Point{x, y}
}

type Rectangle struct {
  a, b, c, d Point
}

type Line struct {
  start, end Point
}

type Point struct {
  x, y int
}
