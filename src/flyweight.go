package main

import "fmt"

type TreeType struct {
    color string
    texture string
}

type TreeFactory struct {
    treeTypes map[string]*TreeType
}

func (tf *TreeFactory) GetTreeType(color, texture string) *TreeType {
    key := fmt.Sprintf("%s_%s", color, texture)
    if tt, ok := tf.treeTypes[key]; ok {
        return tt
    }
    tt := &TreeType{color: color, texture: texture}
    tf.treeTypes[key] = tt
    return tt
}