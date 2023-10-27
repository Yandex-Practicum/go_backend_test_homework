package main

import (
	"os"
	"testing"
)

var (
	requiredFiles = []string{"README.md", "main.go", "main_test.go", "go.mod"}
)

func TestCheckRequiredFiles(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		t.Errorf(err.Error())
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		t.Errorf("Please check permissions for the dir - %s ", dir)
	}
req:
	for _, reqFile := range requiredFiles {
		for _, file := range files {
			if reqFile == file.Name() {
				continue req
			}
		}
		t.Errorf("File doesn't exist - %s", reqFile)
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
