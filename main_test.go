package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	requiredFiles = []string{"README.md", "main.go", "main_test.go", "go.mod"}
)

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func TestCheckRequiredFiles(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Errorf("Please check permissions for the dir - %s ", dir)
	}

	for _, file := range files {
		fileName := file.Name()
		if !Contains(requiredFiles, fileName) {
			t.Errorf("File doesnt exist - %s", fileName)
		}
	}
}

func TestSumFunc(t *testing.T) {

	type args struct {
		a int
		b int
	}
	// описывает структуру тестовых данных и сами тесты
	tests := []struct {
		name   string // название теста
		args   args   // аргументы
		wanted int    // ожидаемое значение
	}{
		{
			name: "Test-2-positive",
			args: args{
				a: 1,
				b: 10,
			},
			wanted: 11,
		},
		{
			name: "Test-2-negative",
			args: args{
				a: -3,
				b: -10,
			},
			wanted: -13,
		},
		{
			name: "Test-2-zero",
			args: args{
				a: 0,
				b: 0,
			},
			wanted: 0,
		},
		{
			name: "Test-positive-negative-zero",
			args: args{
				a: -5,
				b: 5,
			},
			wanted: 0,
		},
		{
			name: "Test-positive-negative-correct",
			args: args{
				a: -2,
				b: 9,
			},
			wanted: 7,
		},
	}
	// вызываем тестируемую функцию для каждого тестового случая
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			if got != tt.wanted {
				t.Errorf("Add() = %v, want %v", got, tt.wanted)
			}
		})
	}
}
