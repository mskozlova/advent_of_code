def parse_map(file_name):
    with open(file_name, "r") as file:
        return [line.strip() for line in file.readlines()]


def get_column_load(map, column_num):
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
        else:  # symbol == "#"
            last_occupied_row = row

    return load


map = parse_map("input.txt")
total_load = 0
for column_num in range(len(map[0])):
    total_load += get_column_load(map, column_num)

print(total_load)
