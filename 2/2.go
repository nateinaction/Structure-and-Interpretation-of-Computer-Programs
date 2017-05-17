package main

import (
  "fmt"
)

func main() {
  a := makePoint(0,0)
  b := makePoint(2,8)

  z := makeSegment(a,b)
  midPoint := midPointSegment(z)

  fmt.Println(midPoint)
}

func midPointSegment(z Line) Point {
  x := (z.start.x + z.end.x) / 2
  y := (z.start.y + z.end.y) / 2
  return Point{x, y}
}

func makeSegment(a Point, b Point) Line {
  return Line{a, b}
}

func startSegment(z Line) Point {
  return z.start
}

func endSegment(z Line) Point {
  return z.end
}

func makePoint(x int, y int) Point {
  return Point{x, y}
}

type Line struct {
  start, end Point
}

type Point struct {
  x, y int
}
