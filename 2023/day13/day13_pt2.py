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


def find_symmetry(arrays):
    n_arrays = len(arrays)

    for symmetry_i in range(n_arrays - 1):
        n_smudges = 0
        for i in range(min(symmetry_i + 1, n_arrays - symmetry_i - 1)):
            for lhs_sym, rhs_sym in zip(
                arrays[symmetry_i - i], arrays[symmetry_i + 1 + i]
            ):
                if lhs_sym != rhs_sym:
                    n_smudges += 1

        if n_smudges == 1:
            return symmetry_i

    return None


patterns = parse_patterns("input.txt")
sum_pattern_notes = 0

for pattern in patterns:
    columns = [
        [pattern[i][j] for i in range(len(pattern))] for j in range(len(pattern[0]))
    ]
    if find_symmetry(columns) is not None:
        sum_pattern_notes += find_symmetry(columns) + 1
    else:
        assert find_symmetry(pattern) is not None, "\n".join(pattern)
        sum_pattern_notes += 100 * (find_symmetry(pattern) + 1)

print(sum_pattern_notes)
