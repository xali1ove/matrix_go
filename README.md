# matrix_go
## тестовое задание в Yabbi
**Имеется:** матрица m*n, состоящая из блоков.
Размер матрицы заранее неизвестен, каждый блок имеет какой-либо цвет.

**Задача**: найти наибольшую группу смежных блоков одного цвета.

**Вход:** высота, ширина, количество цветов.

**Вывод:** сгенерированная матрица и найденная группа, в виде по которой ее можно найти в сгенеренной матрице.

## Реализация:

**Склонировать проект:**
```bash
git clone https://github.com/xali1ove/matrix_go.git matrix
cd matrix
```
**Запуск проекта:**
```bash
go run main.go
```
***Вводные данные:***
```bash
Enter height, width, and number of colors:
5 4 4
```
***Вывод:***
```bash
Generated Matrix:
[Color3] [Color1] [Color3] [Color3] 
[Color2] [Color2] [Color2] [Color2] 
[Color2] [Color2] [Color3] [Color1] 
[Color3] [Color2] [Color4] [Color1] 
[Color1] [Color4] [Color4] [Color3] 
Largest Group:
(2, 1) (3, 1) (3, 2) (2, 2) (2, 3) (2, 4) (4, 2) 
Matrix with Largest Group:
                                    
[Color2] [Color2] [Color2] [Color2] 
[Color2] [Color2]                   
         [Color2]                   
                                    
```
***Тестирование:***
```bash
go test -bench test.go
PASS
ok      main.go 0.363s
```
                                    
