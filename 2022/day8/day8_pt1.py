trees = []

with open("input.txt", "r") as file:
    for line in file.readlines():
        trees.append(list(map(int, line.strip())))

visibility = [[0] * len(trees[0]) for _ in range(len(trees))]

for row in range(len(trees)):
    max_row = -1
    for col in range(len(trees[0])):
        if trees[row][col] > max_row:
            max_row = trees[row][col]
            visibility[row][col] = 1

    max_row = -1
    for col in range(len(trees[0]) - 1, -1, -1):
        if trees[row][col] > max_row:
            max_row = trees[row][col]
            visibility[row][col] = 1

for col in range(len(trees[0])):
    max_col = -1
    for row in range(len(trees)):
        if trees[row][col] > max_col:
            max_col = trees[row][col]
            visibility[row][col] = 1

    max_col = -1
    for row in range(len(trees) - 1, -1, -1):
        if trees[row][col] > max_col:
            max_col = trees[row][col]
            visibility[row][col] = 1

total_visible = 0
for row in visibility:
    total_visible += sum(row)

print(total_visible)
