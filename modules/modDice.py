# modDice.py
import logging
import re
from random import randint

from modules.mod import mod

class module(mod):
	def __init__(self, lgr: logging.Logger):
		super().__init__(lgr)

	async def run(self, *args):
		errMsg: str = """Argument must be a number > 1 any may start with 'd'.
```
Examples:
  "?dice d20" rolls a D20
  "?dice 13" rolls a D13
```
"""

		# No arguments given; empty args: ([],)
		if not args[0]:
			return errMsg

		dMax = re.search(r"^(|d|D)\d+", str(args[0][0]))
		if not dMax:
			return errMsg

		roll: int = -1
		if dMax.group(0)[0] == 'd':
			roll = int(dMax.group(0)[1:])
		else:
			roll = int(dMax.group(0))

		if roll <= 1:
			return errMsg
		else:
			return f"Rolled: `{randint(1, roll)}`"
