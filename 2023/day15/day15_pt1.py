def parse_sequence(file_name):
    with open(file_name, "r") as file:
        return file.readline().strip().split(",")


def hash(string):
    current_value = 0
    for symbol in string:
        current_value += ord(symbol)
        current_value = (current_value * 17) % 256
    return current_value


sequence = parse_sequence("input.txt")
print(sum(hash(string) for string in sequence))
