file = open("input/q2.txt", "r")
ORIGINAL = []
input_array = []
for x in file:
    ORIGINAL = [int(x) for x in x.split(",")]
input_array = ORIGINAL.copy()


def process(array, opcode, idx1, idx2, idx3):
    if opcode == 1:
        array[idx3] = array[idx1] + array[idx2]
    elif opcode == 2:
        array[idx3] = array[idx1] * array[idx2]


outer = False
for i in range(100):
    for j in range(100):
        input_array[1] = i
        input_array[2] = j

        for k in range(0, len(input_array), 4):
            if input_array[k] == 99:
                break
            process(input_array, input_array[k], input_array[k+1],
                    input_array[k+2], input_array[k+3])

        if input_array[0] == 19690720:
            print(f"answer: {i} , {j}")
            outer = True
            break
        input_array = ORIGINAL.copy()
    if outer:
        break
