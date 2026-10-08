class MinStack:

    def __init__(self):
        self.data=[]     
        self.minData=[]  

    def push(self, val: int) -> None:
        self.data.append(val)
        if len(self.minData) == 0:
            self.minData.append(val)
        else:
            self.minData.append(min(val, self.minData[-1]))

    def pop(self) -> None:
        self.data.pop()
        self.minData.pop()

    def top(self) -> int:
        return self.data[-1]

    def getMin(self) -> int:
        return self.minData[-1]