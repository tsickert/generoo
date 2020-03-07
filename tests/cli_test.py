import unittest
import generoo.cli
import os

class TestCli(unittest.TestCase):
    def test_has_dot_generoo(self):
        self.assertFalse(generoo.has_dot_generoo(os.getcwd()))

if __name__ == '__main__':
    unittest.main()