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


commands, network = parse_input("input.txt")
current_node = "AAA"
step_number = 0

while current_node != "ZZZ":
    command = commands[step_number % len(commands)]
    current_node = network[current_node][command]
    step_number += 1

print(step_number)
