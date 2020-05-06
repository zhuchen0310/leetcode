class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []
        # 辅助栈 长度同stack 一致 栈顶始终为当前最小
        self.helper_stack = []


    def push(self, x: int) -> None:
        if not self.helper_stack or x < self.helper_stack[-1]:
            self.helper_stack.append(x)
        else:
            self.helper_stack.append(self.helper_stack[-1])
        self.stack.append(x)


    def pop(self) -> None:
        if self.stack:
            self.stack.pop()
            self.helper_stack.pop()


    def top(self) -> int:
        if self.stack:
            return self.stack[-1]


    def getMin(self) -> int:
        if self.helper_stack:
            return self.helper_stack[-1]



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()