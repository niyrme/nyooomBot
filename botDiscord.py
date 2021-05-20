# modDiscord.py
import logging
import os
import sys

import discord
from dotenv import load_dotenv

from modules import mod
from modules import modDice
from modules import modHelp
from modules import modPing

hdlr: logging.StreamHandler = logging.StreamHandler(sys.stdout)
hdlr.setFormatter(logging.Formatter("[%(name)s] | %(levelname)s | %(message)s"))

logger: logging.Logger = logging.getLogger("BOT-DISCORD")
logger.setLevel(logging.INFO)
logger.addHandler(hdlr)

class Client(discord.Client):
	# Login
	async def on_ready(self):
		logger.info("Logged in.")
		self.command: str = '?'
		self.cmds: dict[str, mod.mod] = {
			f"{self.command}dice": modDice.module(logger),
			f"{self.command}help": modHelp.module(logger),
			f"{self.command}ping": modPing.module(logger),
		}

	# Message-event
	async def on_message(self, msg: discord.Message):
		# Do not reply to self
		if msg.author == cl.user:
			return

		# Check if msg is a command
		if not str(msg.content).startswith(self.command) or len(str(msg.content)) <= 4:
			return

		response: str = ""

		logger.info(f"New Command: '{msg.content}'")

		command, *commandArgs = str(msg.content).split(' ')
		if str(command) in self.cmds:
			response = await self.cmds[command].run(commandArgs)


		if response:
			await msg.channel.send(response)
			return
		else:
			await msg.channel.send(f"Unknown command: `{msg.content}`")


if __name__ == '__main__':
	load_dotenv()
	cl: Client = Client()
	cl.run(os.getenv("DISCORD_TOKEN"))
