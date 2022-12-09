def make_move(x_head, y_head, x_tail, y_tail, direction):
    if direction == "U":
        x_head += 1
    elif direction == "D":
        x_head -= 1
    elif direction == "L":
        y_head -= 1
    else:
        y_head += 1

    if abs(x_head - x_tail) <= 1 and abs(y_head - y_tail) <= 1:
        return x_head, y_head, x_tail, y_tail

    if x_head == x_tail and y_head > y_tail:
        y_tail += 1
        return x_head, y_head, x_tail, y_tail
    if x_head == x_tail and y_head < y_tail:
        y_tail -= 1
        return x_head, y_head, x_tail, y_tail
    if y_head == y_tail and x_head > x_tail:
        x_tail += 1
        return x_head, y_head, x_tail, y_tail
    if y_head == y_tail and x_head < x_tail:
        x_tail -= 1
        return x_head, y_head, x_tail, y_tail

    if x_head > x_tail and y_head > y_tail:
        x_tail += 1
        y_tail += 1
        return x_head, y_head, x_tail, y_tail
    if x_head > x_tail and y_head < y_tail:
        x_tail += 1
        y_tail -= 1
        return x_head, y_head, x_tail, y_tail
    if x_head < x_tail and y_head > y_tail:
        x_tail -= 1
        y_tail += 1
        return x_head, y_head, x_tail, y_tail
    if x_head < x_tail and y_head < y_tail:
        x_tail -= 1
        y_tail -= 1
        return x_head, y_head, x_tail, y_tail



def draw_positions(x_head, y_head, x_tail, y_tail, max_coord):
    for i in range(max_coord):
        row = []
        for j in range(max_coord):
            if max_coord - i - 1 == x_head and y_head == j:
                row.append("H")
            elif max_coord - i - 1 == x_tail and y_tail == j:
                row.append("T")
            else:
                row.append(".")
        print("".join(row))
    print()


positions = set()
positions.add((0, 0))

x_head, y_head, x_tail, y_tail = 0, 0, 0, 0

with open("input.txt", "r") as file:
    for line in file.readlines():
        direction, num = line.strip().split(" ")
        num = int(num)

        print(direction, num)
        for _ in range(num):
            x_head, y_head, x_tail, y_tail = make_move(x_head, y_head, x_tail, y_tail, direction)
            positions.add((x_tail, y_tail))

print(len(positions))
