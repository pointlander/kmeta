// Copyright 2025 The kmeta Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"

	"github.com/pointlander/kmeta/kmeans"
	"github.com/pointlander/kmeta/vector"
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

	{
		rng := rand.New(rand.NewSource(1))
		input := NewMatrix(4, len(iris))
		for _, item := range iris {
			for _, v := range item.Measures {
				input.Data = append(input.Data, float32(v))
			}
		}
		meta := make([][]float64, len(iris))
		for i := range meta {
			meta[i] = make([]float64, len(iris))
		}
		k := 3
		for i := 0; i < 33; i++ {
			project := NewMatrix(4, 8)
			for r := 0; r < project.Rows; r++ {
				for c := 0; c < project.Cols; c++ {
					project.Data = append(project.Data, float32(rng.NormFloat64()))
				}
			}
			for r := 0; r < project.Rows; r++ {
				row := project.Data[r*project.Cols : (r+1)*project.Cols]
				norm := sqrt(vector.Dot(row, row))
				for k := range row {
					row[k] /= norm
				}
			}
			l1 := project.MulT(input)
			output := make([][]float64, 150)
			for i := range output {
				measures := make([]float64, 8)
				for j := range measures {
					measures[j] = float64(l1.Data[i*8+j])
				}
				output[i] = measures
			}

			clusters, _, err := kmeans.Kmeans(int64(i+1), output, k, kmeans.SquaredEuclideanDistance, -1)
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
