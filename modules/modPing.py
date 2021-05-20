# modPing.py
import logging

from modules.mod import mod

class module(mod):
	def __init__(self, lgr: logging.Logger):
		super().__init__(lgr)

	async def run(self, *args):
		del args
		return "Pong!"
