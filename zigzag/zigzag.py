import numpy as np


class Zigzag:
    s: str
    num_rows: int
    index: int = 0
    zigzag: np.ndarray

    def __init__(self, s: str, num_rows: int):
        self.s = s
        self.num_rows = num_rows

    def make_col(self):
        mod = self.index % self.num_rows
        thresh = self.num_rows - 2
        if mod <= thresh and mod != 0:
            appr_col = np.zeros(self.num_rows, dtype=str)
            appr_col[-1 - mod] = self.s[0]
            self.s = self.s[1:]
            self.index += 1
            return appr_col
        full_col = list(self.s[:self.num_rows])
        self.s = self.s[self.num_rows:]
        if len(full_col) < self.num_rows:
            self.s = ""
            full_col += ["" for _ in range(self.num_rows - len(full_col))]
        self.index = 1
        return np.array(full_col)

    def __call__(self) -> str:
        cols = []
        if len(self.s) == 1 or self.num_rows < 2:
            return self.s
        for _ in range(len(self.s)):
            cols.append(self.make_col())
            if self.s == "":
                break
        zigzag = np.array(cols).T
        return "".join(zigzag.flatten().tolist())
