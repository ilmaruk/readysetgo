# readysetgo
[![Build Status](https://github.com/ilmaruk/readysetgo/actions/workflows/main.yml/badge.svg)](https://github.com/ilmaruk/readysetgo/actions/workflows/main.yml/badge.svg)
[![Go Coverage](https://github.com/ilmaruk/readysetgo/wiki/coverage.svg)](https://raw.githack.com/wiki/ilmaruk/readysetgo/coverage.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/ilmaruk/readysetgo)](https://goreportcard.com/report/github.com/ilmaruk/readysetgo)
[![GoDoc](https://godoc.org/github.com/ilmaruk/readysetgo?status.svg)](https://godoc.org/github.com/ilmaruk/readysetgo)

A simple set implementation in Go, that provides the following functionalities:

|Function|Description|
|-|-|
|Set.Add|Adds one or more items to the set|
|Set.Clear|Removes all the items from the set|
|Set.Copy|Returns a copy of the set|
|Set.Has|Tells if the set contains the specified item|
|Set.Items|Returns all the items in the set|
|[Set.Update](https://pkg.go.dev/github.com/ilmaruk/readysetgo#Set.Update)|Updates the set with the union of this set and others|
|Set.IsSubset|Returns whether another set contains this set or not|
|Set.IsDisjoint|Returns whether two sets have a intersection or not|
|Set.IsSuperset|Returns whether this set contains another set or not|
|Set.Remove|Removes the specified item from the set and returns true if it was found or false otherwise|
|Set.DifferenceUpdate|Removes the items in this set that are also included in other, specified set(s)|
|Set.IntersectionUpdate|Removes the items in this set that are not present in other, specified set(s)|
|[Union](https://pkg.go.dev/github.com/ilmaruk/readysetgo#Union)|Returns a set containing the union of sets|
|[Difference](https://pkg.go.dev/github.com/ilmaruk/readysetgo#Difference)|Returns a set containing the difference between a set and one or more other sets|
|[Intersection](https://pkg.go.dev/github.com/ilmaruk/readysetgo#Intersection)|Returns a set, that is the intersection of two or more other sets|

