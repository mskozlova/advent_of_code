direction_mapping = {
    "R": "-",
    "L": "-",
    "U": "|",
    "D": "|",
    "UR": "F",
    "LD": "F",
    "UL": "7",
    "RD": "7",
    "DL": "J",
    "RU": "J",
    "DR": "L",
    "LU": "L",
}


def parse_trench_commands(file_name):
    commands = []
    with open(file_name, "r") as file:
        for line in file.readlines():
            _, _, hex_code = line.strip().split()
            duration_code, direction = hex_code[2:-2], hex_code[-2]
            duration = int(duration_code, 16)
            commands.append((direction, duration))
    return commands


def build_trench(commands):
    field_size = sum(command[1] for command in commands)
    print(f"Field size: {field_size}")
    field = [[] for _ in range(2 * field_size + 1)]

    row, col = field_size + 1, field_size + 1

    for i, (direction, duration) in enumerate(commands):
        for step in range(duration):
            if direction == "U":
                row -= 1
            elif direction == "D":
                row += 1
            elif direction == "R":
                col += 1
            else:
                col -= 1

            if step == duration - 1:
                next_direction = commands[(i + 1) % (len(commands))][0]
                if direction != next_direction:
                    direction += next_direction

            field[row].append((col, direction))

    field = [sorted(row) for row in filter(lambda row: len(row) > 0, field)]
    min_col = min(min(map(lambda x: x[0], row)) for row in field)
    return [[(x - min_col, direction_mapping[sym]) for x, sym in row] for row in field]


def generate_row(row):
    new_row = []
    for i, (coord, sym) in enumerate(row):
        if i == 0:
            new_row.append(sym)
        else:
            if row[i - 1][0] < coord - 1:
                new_row.extend(["."] * (coord - 1 - row[i - 1][0]))
            new_row.append(sym)
    return "".join(new_row)


def calculate_inside_area(row):
    total_area = 0
    inside = False
    group_start = None

    for sym in row:
        if sym != ".":
            total_area += 1

        if sym == "|":
            inside = not inside
        elif sym in ("F", "L"):
            group_start = sym
        elif sym in ("J", "7"):
            assert group_start is not None
            if (group_start == "F" and sym == "J") or (
                group_start == "L" and sym == "7"
            ):
                inside = not inside
            group_start = None
        elif sym == ".":
            if inside:
                total_area += 1

    return total_area


commands = parse_trench_commands("input.txt")
field = build_trench(commands)

inside_area = 0
for row in field:
    print(row)
    # inside_area += calculate_inside_area(generate_row(row))

# print(inside_area)
