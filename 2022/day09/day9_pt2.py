def move_head(rope, direction):
    if direction == "U":
        rope[0][0] += 1
    elif direction == "D":
        rope[0][0] -= 1
    elif direction == "L":
        rope[0][1] -= 1
    else:
        rope[0][1] += 1


def make_move(rope, knot):
    if (
        abs(rope[knot - 1][0] - rope[knot][0]) <= 1
        and abs(rope[knot - 1][1] - rope[knot][1]) <= 1
    ):
        return

    if rope[knot - 1][0] == rope[knot][0] and rope[knot - 1][1] > rope[knot][1]:
        rope[knot][1] += 1
        return
    if rope[knot - 1][0] == rope[knot][0] and rope[knot - 1][1] < rope[knot][1]:
        rope[knot][1] -= 1
        return
    if rope[knot - 1][1] == rope[knot][1] and rope[knot - 1][0] > rope[knot][0]:
        rope[knot][0] += 1
        return
    if rope[knot - 1][1] == rope[knot][1] and rope[knot - 1][0] < rope[knot][0]:
        rope[knot][0] -= 1
        return

    if rope[knot - 1][0] > rope[knot][0] and rope[knot - 1][1] > rope[knot][1]:
        rope[knot][0] += 1
        rope[knot][1] += 1
        return
    if rope[knot - 1][0] > rope[knot][0] and rope[knot - 1][1] < rope[knot][1]:
        rope[knot][0] += 1
        rope[knot][1] -= 1
        return
    if rope[knot - 1][0] < rope[knot][0] and rope[knot - 1][1] > rope[knot][1]:
        rope[knot][0] -= 1
        rope[knot][1] += 1
        return
    if rope[knot - 1][0] < rope[knot][0] and rope[knot - 1][1] < rope[knot][1]:
        rope[knot][0] -= 1
        rope[knot][1] -= 1
        return


def draw_positions(rope, max_coord):
    for i in range(max_coord):
        row = []
        for j in range(max_coord):
            for knot in range(len(rope)):
                if max_coord - i - 1 == rope[knot][0] and rope[knot][1] == j:
                    if knot == 0:
                        row.append("H")
                    elif knot == 9:
                        row.append("T")
                    else:
                        row.append(str(knot))
                else:
                    row.append(".")
        print("".join(row))
    print()


positions = set()
positions.add((0, 0))

rope = [[0, 0] for _ in range(10)]

with open("input.txt", "r") as file:
    for line in file.readlines():
        direction, num = line.strip().split(" ")
        num = int(num)

        print(direction, num)
        for _ in range(num):
            move_head(rope, direction)
            for knot in range(1, len(rope)):
                make_move(rope, knot)
            positions.add(tuple(rope[-1]))

print(len(positions))
