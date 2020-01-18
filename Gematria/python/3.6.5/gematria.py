import inflect, string, numpy, re
test_num = 0
largest_num = 0

for test_num in range (0,1000):
    print(test_num)

    p = inflect.engine()
    num = p.number_to_words(test_num) #turn inputted number into words
    num = re.sub(r'[^a-zA-Z]|and', '',num) #takes out "and" | all characters not 

    res = []
    res[:0] = num 

    size = len(res) #get the size of the string list
    asc = [] #initialize asc (represents the letters is ascii values minus 96)

    i = 0
    sum = 0
    while i < size: #assign array of letters to ascii values
        asc.append(ord(res[i])-96)
        sum+=ord(res[i])-96
        i += 1
    print(asc)

    print(sum)
    if sum > test_num:
        largest_num = test_num
    print(largest_num)

    test_num +=1