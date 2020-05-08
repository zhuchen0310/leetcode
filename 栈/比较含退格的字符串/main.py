class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        
        def parse_str(s):
            stack = []
            for _s in s:
                if _s == "#" and stack:
                    stack.pop()
                elif _s != "#":
                    stack.append(_s)
            return stack

        stack1 = parse_str(S)
        stack2 = parse_str(T)
        return ''.join(stack1) == ''.join(stack2)