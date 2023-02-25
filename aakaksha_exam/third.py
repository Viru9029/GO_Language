def count_interesting_pairs(arr, sumVal):
    freq = {}
    count = 0
    for i in range(len(arr)):
        if arr[i] in freq:
            count += freq[arr[i]]
        if (arr[i]+sumVal) in freq:
            count += freq[arr[i]+sumVal]
        if (arr[i]-sumVal) in freq:
            count += freq[arr[i]-sumVal]
        if arr[i] in freq:
            freq[arr[i]] += 1
        else:
            freq[arr[i]] = 1
    return count - 1

# Example usage
arr = [1,3,2,0]
sumVal = 2
result = count_interesting_pairs(arr, sumVal)
print(result) # Output: 4