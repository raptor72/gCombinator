# gCombinator
Concurrent solution of a combinatorial problem


Дано слово - **hppn**. Необходимо предоставить все возможные комбинации с данным словом, в которых между его буквами будут представлены гласные буквы английского алфавита или пробел. Условно маску требуемого решения можно выразить как \*h\*p\*p\*n\*.

Сочетанием является как наличие символа, так и его отсутствие, т.е. в итоговом списке могут быть следующие варианты:

 -  h p p n
 - ahppn
 - ahappn
 - yhypypyny

Список всех полученных комбинаций нужно записать в текстовый файл.

Для исполнения задачи склонируйте репозиторий и выполните:

```bash
go run ./main.go
```

Результат будет записан в файл `result.txt` в директории проекта.

Выполнение кода будет происходить конкурентно, нагрузка будет распределяться  между горутинами, в результате чего скорость выполнения будет значительно превосходить аналогичные по алгоритму однопоточные решения:

```bash
time go run ./main.go
real    0m18,925s
user    3m30,180s
sys     0m1,375s
```

Итоговое число полученых комбинаций - 16806:

```bash
cat ./result.txt | wc -l
16806
```
