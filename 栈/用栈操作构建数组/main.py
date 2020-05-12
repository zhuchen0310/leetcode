class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        push = "Push"
        pop = "Pop"
        ret = []
        if not target:
            for i in range(n+1):
                ret.append(push)
                ret.append(pop) 
            return           
        stack = []
        start = 0
        for i in range(1, n+1):
            if len(stack) == len(target):
                break
            stack.append(i)
            ret.append(push)
            if target[start] != i:
                stack.pop()
                ret.append(pop)
            else:
                start += 1
        return ret