def parse_commands(line):
    return list(map(lambda x: 1 if x == "R" else 0, line.strip()))


def parse_network(lines):
    network = dict()
    for line in lines:
        if len(line) <= 2:
            continue
        source, dest = line.split("=")
        source = source.strip()
        dest = dest.strip()[1:-1].split(", ")
        network[source] = dest

    return network


def parse_input(file_name):
    with open(file_name, "r") as file:
        commands = parse_commands(file.readline())
        network = parse_network(file.readlines())
    return commands, network


def gcd_pair(a, b):
    while True:
        r = a % b
        if r == 0:
            break
        a = b
        b = r
    return b


def lcm(numbers):
    current_number = numbers[0]
    for number in numbers[1:]:
        current_number = current_number * number / gcd_pair(current_number, number)
    return current_number


commands, network = parse_input("input.txt")
current_nodes = [node for node in network.keys() if node.endswith("A")]
visited_nodes = [dict() for _ in range(len(current_nodes))]
cycle_starts = [None] * len(current_nodes)
cycle_lengths = [None] * len(current_nodes)
z_steps = [[] for _ in range(len(current_nodes))]
step_number = 0

while True:
    command_index = step_number % len(commands)
    command = commands[command_index]
    current_nodes = [network[node][command] for node in current_nodes]
    step_number += 1

    for i, node in enumerate(current_nodes):
        if cycle_starts[i] is None and (node, command_index) in visited_nodes[i]:
            visited_step = visited_nodes[i][(node, command_index)]
            cycle_starts[i] = visited_step
            cycle_lengths[i] = step_number - visited_step
        else:
            visited_nodes[i][(node, command_index)] = step_number

        if node.endswith("Z"):
            z_steps[i].append(step_number)

    if all(start is not None for start in cycle_starts):
        break

print(step_number)
print(cycle_starts)
print(cycle_lengths)
print(z_steps)

for z_step in z_steps:
    assert len(z_step) == 1

print(lcm([z_step[0] for z_step in z_steps]))
