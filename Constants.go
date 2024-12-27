package main

const (
	SUCCESS = 0
	ERROR   = -1
)

type GitRepository struct {
	WorkTree string
	GitDir   string
	Config   map[string]any
}
