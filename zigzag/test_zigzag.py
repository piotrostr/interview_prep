import pytest

from zigzag import Zigzag


@pytest.mark.parametrize(
    "s, num_rows, expected",
    [
        ("PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"),
        ("PAYPALISHIRING", 4, "PINALSIGYAHRPI"),
        ("A", 1, "A"),
    ],
)
def test_zigzag(s: str, num_rows: int, expected: str):
    assert Zigzag(s, num_rows)() == expected
