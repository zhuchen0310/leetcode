class Solution:
    def removeOuterParentheses(self, S: str) -> str:
        ret = ""
        if not S:
            return ret 
        stack = [S[0]]
        start = 1
        # 从下标1开始遍历 “（”入栈 “）”出栈
        while start < len(S):
            s = S[start]
            if s == "(":
                stack.append(s)
                if len(stack) > 1:
                    ret += s 
            else:
                stack.pop()
                if len(stack) > 0:
                    ret += s 
            start += 1 
        return ret