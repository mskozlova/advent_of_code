all_colors = ["red", "green", "blue"]

def parse_step(step):
    colors = map(lambda x: x.strip(), step.split(","))
    parsed = dict()
    
    for color_info in colors:
        count, color = color_info.split(" ")
        parsed[color] = int(count)
    
    return parsed


def parse_game(line):
    # example: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    return {
        "id": int(line[5:line.rindex(":")]),
        "steps": [
            parse_step(step)
            for step in line[line.rindex(":") + 1:].split(";")
        ]
    }


def get_min_power(game):
    power = 1
    for color in all_colors:
        power *= max([step.get(color, 0) for step in game["steps"]])
    return power


sum_power = 0

with open("input.txt", "r") as file:
    for line in file.readlines():
        game = parse_game(line)
        sum_power += get_min_power(game)

print(sum_power)