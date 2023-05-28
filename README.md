# matrix_go
## тестовое задание в Yabbi
**Имеется:** матрица m*n, состоящая из блоков.
Размер матрицы заранее неизвестен, каждый блок имеет какой-либо цвет.

**Задача**: найти наибольшую группу смежных блоков одного цвета.

**Вход:** высота, ширина, количество цветов.

**Вывод:** сгенерированная матрица и найденная группа, в виде по которой ее можно найти в сгенеренной матрице.

## Реализация:

Склонировать проект:
```bash
git clone https://github.com/xali1ove/matrix_go.git matrix
cd matrix
```
Запуск проекта:
```bash
go run main.go
```
Пример запроса:

Enter height, width, and number of colors:

6 4 3

Generated Matrix:

[Color2] [Color3] [Color2] [Color3]
[Color2] [Color3] [Color2] [Color3]
[Color2] [Color3] [Color3] [Color1]
[Color2] [Color3] [Color2] [Color1]
[Color3] [Color2] [Color1] [Color1]
[Color2] [Color1] [Color2] [Color1]

Largest Group:
(0, 1) (1, 1) (2, 1) (3, 1) (2, 2)

Matrix with Largest Group:
[Color3]                   
[Color3]                   
[Color3] [Color3]          
[Color3]                   
                                    
