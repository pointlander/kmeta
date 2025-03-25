// Copyright 2025 The kmeta Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/pointlander/kmeta/kmeans"
)

func main() {
	iris := Load()
	fmt.Println(len(iris))

	{
		input := make([][]float64, len(iris))
		for i, item := range iris {
			measures := make([]float64, 4)
			for j := range measures {
				measures[j] = item.Measures[j]
			}
			input[i] = measures
		}
		meta := make([][]float64, len(iris))
		for i := range meta {
			meta[i] = make([]float64, len(iris))
		}
		k := 3
		for i := 0; i < 100; i++ {
			clusters, _, err := kmeans.Kmeans(int64(i+1), input, k, kmeans.SquaredEuclideanDistance, -1)
			if err != nil {
				panic(err)
			}
			for i := 0; i < len(meta); i++ {
				target := clusters[i]
				for j, v := range clusters {
					if v == target {
						meta[i][j]++
					}
				}
			}
		}
		clusters, _, err := kmeans.Kmeans(1, meta, k, kmeans.SquaredEuclideanDistance, -1)
		if err != nil {
			panic(err)
		}
		a := make(map[string][3]int)
		for i := range iris {
			histogram := a[iris[i].Label]
			histogram[clusters[i]]++
			a[iris[i].Label] = histogram
		}
		for k, v := range a {
			fmt.Println(k, v)
		}
	}
}
