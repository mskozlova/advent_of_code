import math


def parse_races(filename):
    with open(filename, "r") as file:
        lines = file.readlines()
        assert len(lines) == 2

    target_time = int("".join(lines[0][5:].strip().split()))
    record = int("".join(lines[1][9:].strip().split()))

    return target_time, record


def get_quadratic_roots(a=1, b=1, c=1):
    return (
        (-b - math.sqrt(b**2 - 4 * a * c)) / (2 * a),
        (-b + math.sqrt(b**2 - 4 * a * c)) / (2 * a),
    )


target_time, record = parse_races("input.txt")
root1, root2 = get_quadratic_roots(b=-target_time, c=record)
print(int(math.floor(root2)) + 1 - int(math.ceil(root1)))
