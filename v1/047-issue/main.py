#!/usr/bin/env python3

"""
https://coderun.yandex.ru/problem/stalker/solutions/10010066-fd0a-3bbb-da29-c0ec8957a5c0?currentPage=1&groups=algorithm&groups=backend&pageSize=20&search=&tag=bfsВ городе Н при невыясненных обстоятельствах территория одного из заводов превратилась в аномальную зону.
Все подъезды к территории были перекрыты, а сама она получила название промзоны.
В промзоне находятся N зданий, некоторые из них соединены дорогами.
По любой дороге можно перемещаться в обоих направлениях.

Начинающий сталкер получил задание добраться до склада в промзоне.
Он нашел в электронном архиве несколько карт территории промзоны.
Так как карты составлялись разными людьми, то на каждой из них есть информация только о некоторых дорогах промзоны.
Одна и та же дорога может присутствовать на нескольких картах.

В пути сталкер может загружать из архива на мобильный телефон по одной карте.
При загрузке новой карты предыдущая в памяти телефона не сохраняется.
Сталкер может перемещаться лишь по дорогам, отмеченным на карте, загруженной на данный момент.
Каждая загрузка карты стоит 1 рубль. Для минимизации расходов сталкеру нужно выбрать такой маршрут,
чтобы как можно меньшее число раз загружать карты.
Сталкер может загружать одну и ту же карту несколько раз, при этом придется заплатить за каждую загрузку.
Изначально в памяти мобильного телефона нет никакой карты.

Требуется написать программу, которая вычисляет минимальную сумму расходов, необходимую сталкеру,
чтобы добраться от входа в промзону до склада.

Формат ввода
В первой строке входных данных содержатся два натуральных числа N и K (2 ≤ N ≤2000; 1 ≤ K ≤ 2000) — количество зданий
промзоны и количество карт соответственно.
Вход в промзону находится в здании с номером 1, а склад — в здании с номером N.

В последующих строках находится информация об имеющихся картах.
Первая строка описания i-ой карты содержит число rᵢ — количество дорог, обозначенных на i-ой карте.
Затем идут rᵢ строк, содержащие по два натуральных числа a и b (1 ≤ a, b ≤ N; a ≠ b),
означающих наличие на i-ой карте дороги, соединяющей здания a и b.
Суммарное количество дорог, обозначенных на всех картах, не превышает 300 000 (r₁ + r₂ + ... + rₖ ≤ 300_000).

Формат вывода
Выведите одно число — минимальную сумму расходов сталкера.
В случае, если до склада добраться невозможно, выведите число — 1.

Ограничения
Ограничение времени 2 с
Ограничение памяти 256 МБ
"""
import sys
from queue import Queue
from typing import TextIO

MAX_INT = (2 ** 64 - 1) // 2


class BFS:
    def __init__(self, capacity: int):
        self.adjacency = list[list[int]]()
        for i in range(capacity):
            self.adjacency.append(list[int]())
        self.color = ['white'] * capacity
        self.distance = [0] * capacity
        self.previous = [None] * capacity

    def add_edge(self, u: int, v: int) -> None:
        self.adjacency[u].append(v)
        self.adjacency[v].append(u)

    def bfs(self, s: int) -> None:
        planned = Queue()
        planned.put(s)
        self.color[s] = 'gray'
        self.distance[s] = 0
        while not planned.empty():
            u = planned.get()
            for _, v in enumerate(self.adjacency[u]):
                if self.color[v] == 'white':
                    self.distance[v] = self.distance[u] + 1
                    self.previous[v] = u
                    self.color[v] = 'gray'
                    planned.put(v)
            self.color[u] = 'black'
        pass

    def shortest_path(self, v: int) -> list[int]:
        path = list[int]()
        current_vertex = v
        while current_vertex is not None:
            path.append(current_vertex)
            current_vertex = self.previous[current_vertex]
        return path


class DFS:
    def __init__(self, capacity: int, start_component: int):
        self.adjacency = list[list[int]]()
        for i in range(capacity):
            self.adjacency.append(list[int]())
        self.color = ['white'] * capacity
        self.components = dict[int, bool]()
        self.map_component = dict[int, list[int]]()
        self.start_component = start_component
        self.vertices = dict[int, bool]()

    def add_edge(self, u: int, v: int) -> None:
        self.adjacency[u].append(v)
        self.adjacency[v].append(u)
        if self.vertices.get(u) is None:
            self.vertices[u] = True
        if self.vertices.get(v) is None:
            self.vertices[v] = True

    def all(self) -> None:
        for i in self.vertices.keys():
            if self.color[i] == 'white':
                self.components[self.start_component] = True
                self.dfs(i)
                self.start_component += 1
            pass
        pass

    def dfs(self, v: int) -> None:
        self.color[v] = 'gray'
        for _, w in enumerate(self.adjacency[v]):
            if self.color[w] == 'white':
                self.dfs(w)
        self.color[v] = 'black'
        if self.map_component.get(v) is None:
            self.map_component[v] = list[int]()
        self.map_component[v].append(self.start_component)


def read(name: str = 'input.txt', inp: TextIO = None) -> tuple[int, int, list[DFS], list[int], list[int], int]:
    reader = inp if inp is not None else open(name, 'r')
    result = list[DFS]()
    n, k = [int(i) for i in reader.readline().split()]
    start_component = 0
    start = list[int]()
    finish = list[int]()
    for _ in range(k):
        m = int(reader.readline())
        dfs = DFS(n + 1, start_component)
        is_start = False
        is_finish = False
        for _ in range(m):
            u, v = [int(i) for i in reader.readline().split()]
            dfs.add_edge(u, v)
            if u == 1 or v == 1:
                is_start = True
            if u == n or v == n:
                is_finish = True
        dfs.all()
        if is_start:
            start.extend(dfs.map_component[1])
        if is_finish:
            finish.extend(dfs.map_component[n])
        result.append(dfs)
        start_component = dfs.start_component
    reader.close()
    return n, k, result, start, finish, start_component


def write(result: int, name: str = 'output.txt', out: TextIO = None) -> None:
    writer = out if out is not None else open(name, 'w')
    writer.write("%d\n" % result)
    writer.close()
    pass


def solution(inp: tuple[int, int, list[DFS], list[int], list[int], int]) -> int:
    n, k, maps, start, finish, upperbound_component = inp
    print("start : ", start) # TODO
    print("finish: ", finish) # TODO
    for dfs in maps:
        print("components: %s" % dfs.map_component) # TODO
    capacity = (n + 1)
    middle = list[list[int]]()
    for i in range(capacity):
        middle.append(list[int]())
    for _, bfs in enumerate(maps):
        for key, value in bfs.map_component.items():
            for c in value:
                middle[key].append(c)
    bfs = BFS(upperbound_component)
    for v in middle:
        for i, _ in enumerate(v):
            for j in range(i + 1, len(v)):
                bfs.add_edge(v[i], v[j])
    minimum = MAX_INT
    for i in start:
        for j in finish:
            bfs.bfs(i)
            if i != j:
                shortest_path = bfs.shortest_path(j)
                if len(shortest_path) > 0:
                    minimum = min(minimum, len(shortest_path))
                    print("i: %d, j: %d, shortest_path: %s" % (i, j, shortest_path)) # TODO
    if minimum == MAX_INT:
        minimum = -1
    return minimum


def main():
    write(solution(read(inp=sys.stdin)), out=sys.stdout)
    pass


if __name__ == "__main__":
    main()
    pass
