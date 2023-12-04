# only 12 red cubes, 13 green cubes, and 14 blue cubes?
max_target = {
    "red": 12,
    "blue": 14,
    "green": 13,
}


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
        "id": int(line[5 : line.rindex(":")]),
        "steps": [parse_step(step) for step in line[line.rindex(":") + 1 :].split(";")],
    }


def is_game_possible(game):
    for step in game["steps"]:
        for color, count in step.items():
            if count > max_target[color]:
                return False
    return True


sum_possible_ids = 0

with open("input.txt", "r") as file:
    for line in file.readlines():
        game = parse_game(line)
        if is_game_possible(game):
            sum_possible_ids += game["id"]

print(sum_possible_ids)
