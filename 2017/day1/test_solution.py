import pytest
from unittest.mock import patch
from .solution import *

class TestDay1:
    @pytest.mark.parametrize(
        "data, expect",
        [
            ("1122", 3),
            ("1111", 4),
            ("1234", 0),
            ("91212129", 9),
        ],
    )
    def test_part1(self, data, expect):
        with patch("inputs.string", return_value=data):
            assert part1() == expect

    @pytest.mark.parametrize(
        "data, expect",
        [
            ("1212", 6),
            ("1221", 0),
            ("123425", 4),
            ("123123", 12),
            ("12131415", 4),
        ],
    )
    def test_part2(self, data, expect):
        with patch("inputs.string", return_value=data):
            assert part2() == expect
