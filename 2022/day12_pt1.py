from collections import deque

with open("input.txt", "r") as file:
    heightmap = [line.strip() for line in file.readlines()]

def find_sym(heightmap, sym):
    for i, line in enumerate(heightmap):
        for j, map_sym in enumerate(line):
            if map_sym == sym:
                return i, j


def is_neighbor_height(h1, h2):
    h1 = "a" if h1 == "S" else h1
    h1 = "z" if h1 == "E" else h1
    h2 = "a" if h2 == "S" else h2
    h2 = "z" if h2 == "E" else h2
    return ord(h1) - ord(h2) >= -1

def get_neighbors(heightmap, visited, point):
    potential = [
        [point[0] - 1, point[1]],
        [point[0] + 1, point[1]],
        [point[0], point[1] + 1],
        [point[0], point[1] - 1],
    ]
    neighbors = []
    for p in potential:
        if tuple(p) not in visited:
            if p[0] >= 0 and p[0] < len(heightmap):
                if p[1] >= 0 and p[1] < len(heightmap[0]):
                    if is_neighbor_height(heightmap[point[0]][point[1]], heightmap[p[0]][p[1]]):
                        neighbors.append(p)
                        visited.add(tuple(p))

    return neighbors


start_coords = find_sym(heightmap, "S")
end_coords = find_sym(heightmap, "E")

print(start_coords, end_coords)

queue = deque()
visited = set()

queue.append((0, start_coords))
visited.add(start_coords)

while len(queue) > 0:
    step, point = queue.popleft()
    print(step, point)

    if tuple(point) == end_coords:
        print(step)
        break

    neighbors = get_neighbors(heightmap, visited, point)

    for n in neighbors:
        queue.append((step + 1, n))
