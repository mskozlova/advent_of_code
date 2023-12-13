def parse_patterns(file_name):
    current_pattern = []
    with open(file_name, "r") as file:
        for line in file.readlines():
            if line.strip() == "":
                yield current_pattern
                current_pattern = []
            else:
                current_pattern.append(line.strip())
    yield current_pattern


def find_col_symmetry(pattern):
    total_cols = len(pattern[0])

    columns = [[pattern[i][j] for i in range(len(pattern))] for j in range(total_cols)]

    for symmetry_col in range(total_cols - 1):
        is_symmetrical = True
        for i in range(min(symmetry_col + 1, total_cols - symmetry_col - 1)):
            if columns[symmetry_col - i] != columns[symmetry_col + 1 + i]:
                is_symmetrical = False
                break

        if is_symmetrical:
            return symmetry_col

    return None


def find_row_symmetry(pattern):
    total_rows = len(pattern)

    for symmetry_row in range(total_rows - 1):
        is_symmetrical = True
        for i in range(min(symmetry_row + 1, total_rows - symmetry_row - 1)):
            if pattern[symmetry_row - i] != pattern[symmetry_row + 1 + i]:
                is_symmetrical = False
                break

        if is_symmetrical:
            return symmetry_row

    return None


patterns = parse_patterns("input.txt")
sum_pattern_notes = 0

for pattern in patterns:
    if find_col_symmetry(pattern) is not None:
        sum_pattern_notes += find_col_symmetry(pattern) + 1
    else:
        assert find_row_symmetry(pattern) is not None, "\n".join(pattern)
        sum_pattern_notes += 100 * (find_row_symmetry(pattern) + 1)

print(sum_pattern_notes)
