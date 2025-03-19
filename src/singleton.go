package main

import "sync"

type Database struct{}

var instance *Database
var once sync.Once

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{}
    })
    return instance
}