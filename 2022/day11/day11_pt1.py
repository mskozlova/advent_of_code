from collections import defaultdict
from itertools import islice
from math import floor


class Monkey:
    def __init__(self, lines):
        lines = [line.strip() for line in lines]
        self.list = list(map(int, lines[1][16:].split(", ")))
        self.rule = self.__parse_rule(lines[2])
        self.test = int(lines[3][19:])
        self.if_true = int(lines[4][-1])
        self.if_false = int(lines[5][-1])
        self.throw_cnt = 0

    def __parse_rule(self, rule_line):
        operation = rule_line[21]
        second_number = None

        if not rule_line.endswith("old"):
            second_number = int(rule_line[23:])

        if operation == "*":
            if second_number is None:
                return lambda old: old * old
            else:
                return lambda old: old * second_number
        elif operation == "+":
            return lambda old: old + second_number

    def __throw(self, item):
        operation_res = self.rule(item)
        worry_lvl = int(floor(operation_res / 3))

        if worry_lvl % self.test == 0:
            return self.if_true, worry_lvl
        return self.if_false, worry_lvl

    def round(self):
        throws = defaultdict(list)
        for item in self.list:
            monkey_idx, worry_lvl = self.__throw(item)
            throws[monkey_idx].append(worry_lvl)
            self.throw_cnt += 1

        self.list = []

        return throws

    def receive(self, items):
        self.list.extend(items)


n = 7  # Or whatever chunk size you want
monkeys = []
with open("input.txt", "r") as f:
    for n_lines in iter(lambda: tuple(islice(f, n)), ()):
        monkeys.append(Monkey(n_lines))

for i in range(20):
    for j, monkey in enumerate(monkeys):
        print(f"Round {i}, monkey {j}")
        throws = monkey.round()
        for mid, items in throws.items():
            monkeys[mid].receive(items)

    for j, monkey in enumerate(monkeys):
        print(f"ROUND {i}, MONKEY {j}, ITEMS {monkey.list}")

for monkey in monkeys:
    print(monkey.throw_cnt)


top_monkeys = sorted(monkeys, key=lambda m: -m.throw_cnt)
print(top_monkeys[0].throw_cnt * top_monkeys[1].throw_cnt)
