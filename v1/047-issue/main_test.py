#!/usr/bin/env python3

import main
import unittest


class TestPart3ProblemsC(unittest.TestCase):

    def test_sub_solution_c_case_1(self):
        self.assertEqual(3, main.solution(main.read("input.txt")))

    def test_sub_solution_c_case_2(self):
        self.assertEqual(-1, main.solution(main.read("input_test_02.txt")))

    def test_sub_solution_c_case_XX2(self):
        self.assertEqual(-1, main.solution(main.read("input_test_XX2.txt")))


if __name__ == "__main__":
    unittest.main()
    pass
