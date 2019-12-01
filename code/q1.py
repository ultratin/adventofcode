total_fuel = 0

file = open("input/input.txt", "r")


def calculate(input):
    res = (input//3) - 2
    if res <= 0:
        return 0
    return res + calculate(res)


for x in file:
    total_fuel += calculate(int(x))

# total_fuel = calculate(1969)
print(total_fuel)
