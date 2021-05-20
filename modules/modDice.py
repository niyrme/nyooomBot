# modDice.py
import re
from random import randint

from modules.mod import mod

class module(mod):
	def __init__(self):
		super().__init__()
		self.description = "Roll a Dice of any size!"
		self.how = "`?dice {value}` (value mayb start with the letter 'd' e.g. d20)"

	async def run(self, *args):

		# No arguments given; empty args: ([],)
		if len(args[0]) != 0:
			return self.how

		dMax = re.search(r"^(|d|D)\d+", str(args[0][0]))
		if not dMax:
			return self.how

		roll: int = -1
		if dMax.group(0)[0] == 'd':
			roll = int(dMax.group(0)[1:])
		else:
			roll = int(dMax.group(0))

		if roll <= 1:
			return self.how
		else:
			return f"Rolled: `{randint(1, roll)}`"
