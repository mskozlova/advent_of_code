import re


def get_stacks(lines):
    number_of_stacks = max(map(int, lines[-1].split()))
    stacks = [[] for _ in range(number_of_stacks)]
    
    print(number_of_stacks)

    for line in lines[:-1]:
        try:
            for i in range(number_of_stacks):
                symbol_pos = i * 4 + 1
                if line[symbol_pos].isalpha():
                    stacks[i].append(line[symbol_pos])
        except IndexError:
            continue

    return [stack[::-1] for stack in stacks]


def parse_move(line):
    number = int(re.findall("move ([\d]+)", line)[0])
    move = (
        int(re.findall("from ([\d]+)", line)[0]) - 1,
        int(re.findall("to ([\d]+)", line)[0]) - 1 
    )
    return number, move


def apply_move(stacks, move):
    box = stacks[move[0]].pop()
    stacks[move[1]].append(box)


def get_upper_layer(stacks):
    return "".join(stack[-1] for stack in stacks)


with open("input.txt", "r") as file:
    stack_lines = []
    stack_finished = False
    stacks = None

    for line in file.readlines():
        if line.strip() == "":
            stack_finished = True
            stacks = get_stacks(stack_lines)
            print(stacks)
        elif not stack_finished:
            stack_lines.append(line)
        else:
            number, move = parse_move(line.strip())
            for _ in range(number):
                apply_move(stacks, move)

print(get_upper_layer(stacks))

