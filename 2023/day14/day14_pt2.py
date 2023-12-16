def parse_map(file_name):
    with open(file_name, "r") as file:
        return [list(line.strip()) for line in file.readlines()]


def move_column_north(map, column_num):
    new_column = []
    total_rows = len(map)
    last_occupied_row = -1
    load = 0

    for row in range(total_rows):
        symbol = map[row][column_num]

        if symbol == ".":
            pass
        elif symbol == "O":
            last_occupied_row += 1
            load += total_rows - last_occupied_row
            new_column.append("O")
        else:  # symbol == "#"
            while len(new_column) < row:
                new_column.append(".")

            new_column.append("#")
            last_occupied_row = row

    while len(new_column) < total_rows:
        new_column.append(".")

    for row_num in range(total_rows):
        map[row_num][column_num] = new_column[row_num]

    return load


def move_column_south(map, column_num):
    new_column = []
    total_rows = len(map)
    last_occupied_row = total_rows
    load = 0

    for row_num in range(total_rows - 1, -1, -1):
        symbol = map[row_num][column_num]

        if symbol == ".":
            pass
        elif symbol == "O":
            last_occupied_row -= 1
            load += total_rows - last_occupied_row
            new_column.append("O")
        else:  # symbol == "#"
            while len(new_column) < total_rows - row_num - 1:
                new_column.append(".")

            new_column.append("#")
            last_occupied_row = row_num

    while len(new_column) < total_rows:
        new_column.append(".")

    new_column = new_column[::-1]

    for row_num in range(total_rows):
        map[row_num][column_num] = new_column[row_num]

    return load


def move_row_west(map, row_num):
    new_row = []
    total_cols = len(map[0])
    last_occupied_col = -1
    load = 0

    for col_num in range(total_cols):
        symbol = map[row_num][col_num]

        if symbol == ".":
            pass
        elif symbol == "O":
            last_occupied_col += 1
            load += total_cols - last_occupied_col
            new_row.append("O")
        else:  # symbol == "#"
            while len(new_row) < col_num:
                new_row.append(".")

            new_row.append("#")
            last_occupied_col = col_num

    while len(new_row) < total_cols:
        new_row.append(".")

    for col_num in range(total_cols):
        map[row_num][col_num] = new_row[col_num]

    return load


def move_row_east(map, row_num):
    new_row = []
    total_cols = len(map[0])
    last_occupied_col = total_cols
    load = 0

    for col_num in range(total_cols - 1, -1, -1):
        symbol = map[row_num][col_num]

        if symbol == ".":
            pass
        elif symbol == "O":
            last_occupied_col -= 1
            load += total_cols - last_occupied_col
            new_row.append("O")
        else:  # symbol == "#"
            while len(new_row) < total_cols - col_num - 1:
                new_row.append(".")

            new_row.append("#")
            last_occupied_col = col_num

    while len(new_row) < total_cols:
        new_row.append(".")

    new_row = new_row[::-1]

    for col_num in range(total_cols):
        map[row_num][col_num] = new_row[col_num]

    return load


def calculate_load(map):
    total_load = 0
    for row in range(len(map)):
        load = len(map) - row
        n_boulders = sum(int(sym == "O") for sym in map[row])
        total_load += load * n_boulders
    return total_load


def print_map(map):
    for row in map:
        print("".join(row))
    print()


def rotate_cycle(map):
    for col in range(len(map[0])):
        move_column_north(map, col)

    for row in range(len(map[0])):
        move_row_west(map, row)

    for col in range(len(map)):
        move_column_south(map, col)

    for row in range(len(map[0])):
        move_row_east(map, row)


def collapse_map(map):
    return "".join("".join(row) for row in map)


map = parse_map("input.txt")
map_last_occurence = dict()
total_steps = 1000000000
current_step = 1
has_skipped = False

while current_step < total_steps:
    print(f"current step: {current_step}")
    rotate_cycle(map)

    if collapse_map(map) not in map_last_occurence:
        map_last_occurence[collapse_map(map)] = current_step
        current_step += 1
    elif has_skipped:
        current_step += 1
        continue
    else:
        cycle_length = current_step - map_last_occurence[collapse_map(map)]
        print(f"cycle length: {cycle_length}")
        current_step += (total_steps - current_step) // cycle_length * cycle_length
        has_skipped = True

print(calculate_load(map))
