trees = []

with open("input.txt", "r") as file:
    for line in file.readlines():
        trees.append(list(map(int, line.strip())))

for row in trees:
    print(row)
print()

print(len(trees), len(trees[0]))


def count_visibility(x, y, direction):
    if direction == "up":
        for i in range(1, max(len(trees), len(trees[0])) + 1):
            print(direction, x, y, i)
            if x - i >= 0:
                if trees[x - i][y] < trees[x][y]:
                    continue
                else:
                    return i
            else:
                return i - 1
    elif direction == "down":
        for i in range(1, max(len(trees), len(trees[0])) + 1):
            print(direction, x, y, i)
            if x + i < len(trees):
                if trees[x + i][y] < trees[x][y]:
                    continue
                else:
                    return i
            else:
                return i - 1
    elif direction == "right":
        for i in range(1, max(len(trees), len(trees[0])) + 1):
            print(direction, x, y, i)
            if y + i < len(trees[0]):
                if trees[x][y + i] < trees[x][y]:
                    continue
                else:
                    return i
            else:
                return i - 1
    elif direction == "left":
        for i in range(1, max(len(trees), len(trees[0])) + 1):
            print(direction, x, y, i)
            if y - i >= 0:
                if trees[x][y - i] < trees[x][y]:
                    continue
                else:
                    return i
            else:
                return i - 1


max_visibility = 0
visibilities = []

for x in range(len(trees)):
    row = []
    for y in range(len(trees[0])):
        current_visibility = (
            count_visibility(x, y, "up")
            * count_visibility(x, y, "down")
            * count_visibility(x, y, "left")
            * count_visibility(x, y, "right")
        )
        max_visibility = max(max_visibility, current_visibility)
        row.append(
            [
                count_visibility(x, y, "up"),
                count_visibility(x, y, "down"),
                count_visibility(x, y, "left"),
                count_visibility(x, y, "right"),
            ]
        )

    visibilities.append(row)

for row in visibilities:
    print(row)
print()

print(max_visibility)
