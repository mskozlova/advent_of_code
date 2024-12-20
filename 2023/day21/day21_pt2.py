from collections import deque


TARGET_N_STEPS = 6


def parse_field(file_name):
    with open(file_name, "r") as file:
        return [line.strip() for line in file.readlines()]


def get_start_coords(field):
    for i, row in enumerate(field):
        for j, sym in enumerate(row):
            if sym == "S":
                return i, j
            

def get_available_steps(row, col, field):
    steps = [
        (row + 1, col),
        (row - 1, col),
        (row, col + 1),
        (row, col - 1),
    ]
    
    steps = list(filter(
        lambda c: c[0] >= 0 and c[0] < len(field) and c[1] >= 0 and c[1] < len(field[0]) and field[c[0]][c[1]] == ".",
        steps
    ))
    
    return steps


def count_gardens(start_row, start_col, field):
    visited = set()
    visited.add((start_row, start_col))
    
    queue = deque()
    queue.append((start_row, start_col, 0))
    
    n_reached_gardens = 0
    
    while len(queue) > 0:
        row, col, level = queue.popleft()
        
        if level % 2 == TARGET_N_STEPS % 2:
            n_reached_gardens += 1
        
        if level == TARGET_N_STEPS:
            continue
        
        next_steps = get_available_steps(row, col, field)
        
        for next_row, next_col in next_steps:
            if (next_row, next_col) not in visited:
                visited.add((next_row, next_col))
                queue.append((next_row, next_col, level + 1))
        
    return n_reached_gardens


field = parse_field("input.txt")
start_row, start_col = get_start_coords(field)
n_accessible_gardens = count_gardens(start_row, start_col, field)
print(n_accessible_gardens)