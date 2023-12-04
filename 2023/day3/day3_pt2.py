from collections import defaultdict


def get_number_coords(line, line_num):
    numbers = []
    current_number_start = None

    for i, sym in enumerate(line):
        if sym.isdigit():
            if current_number_start is None:
                current_number_start = i

        elif current_number_start is not None:
            numbers.append(
                {
                    "row": line_num,
                    "col": current_number_start,
                    "length": i - current_number_start,
                    "number": int(line[current_number_start:i]),
                }
            )
            current_number_start = None

    if current_number_start is not None:
        numbers.append(
            {
                "row": line_num,
                "col": current_number_start,
                "length": len(line) - current_number_start,
                "number": int(line[current_number_start:]),
            }
        )

    return numbers


def get_adjucent_stars(row, col, length, number, lines):
    adj_coords = []
    adj_coords.extend(
        [(row - 1, col_num) for col_num in range(col - 1, col + length + 1)]
    )
    adj_coords.extend([(row, col - 1), (row, col + length)])
    adj_coords.extend(
        [(row + 1, col_num) for col_num in range(col - 1, col + length + 1)]
    )

    adj_coords = list(
        filter(
            lambda c: c[0] >= 0
            and c[1] >= 0
            and c[0] < len(lines)
            and c[1] < len(lines[0]),
            adj_coords,
        )
    )

    adj_stars = []
    for r, c in adj_coords:
        if lines[r][c] == "*":
            adj_stars.append((r, c, number))
    return adj_stars


def get_total_gear_ratio(adj_stars):
    stars = defaultdict(list)
    total_ratio = 0
    
    for row, col, number in adj_stars:
        stars[(row, col)].append(number)
    
    for numbers in stars.values():
        if len(numbers) == 2:
            total_ratio += (numbers[0] * numbers[1])
    
    return total_ratio


with open("input.txt", "r") as file:
    lines = [line.strip() for line in file.readlines()]

coords = []
for i, line in enumerate(lines):
    coords.extend(get_number_coords(line, i))

adj_stars = []
for number_info in coords:
    adj_stars.extend(get_adjucent_stars(
        number_info["row"],
        number_info["col"],
        number_info["length"],
        number_info["number"],
        lines
    ))

print(get_total_gear_ratio(adj_stars))
