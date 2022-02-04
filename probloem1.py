class Solution:
    def isPalindrome(self, x: int) -> bool:
        return str(x) == str(x)[::-1]
    
test_set = [1, 557933, 415541, 234567899998765432, 12345321, 989898989, 00000000000000]
solution = Solution()
for n in test_set:
    print(n, "palindrome?", solution.isPalindrome(n))