import os
import re
import platform
import socket
import subprocess
from datetime import datetime

import distro
import psutil
from screeninfo import get_monitors
from rich import box
from rich.console import Console
from rich.table import Table
from rich.panel import Panel
from rich.text import Text
from rich.align import Align

def get_git_info():
    try:
        branch = subprocess.check_output(['git', 'rev-parse', '--abbrev-ref', 'HEAD']).strip().decode('utf-8')
        commit_hash = subprocess.check_output(['git', 'rev-parse', 'HEAD']).strip().decode('utf-8')
        commit_count = subprocess.check_output(['git', 'rev-list', '--count', 'HEAD']).strip().decode('utf-8')
        return branch, commit_hash, commit_count
    except subprocess.CalledProcessError:
        return None, None, None

if __name__ == "__main__":
    console = Console()

    table = Table(show_header=False, box=None, padding=(0, 1))

    label_color = 'bold cyan'
    value_color = 'white'

    # Example of adding meaningful data
    table.add_row(Text('OS:', style=label_color), Text(platform.system(), style=value_color))
    table.add_row(Text('Version:', style=label_color), Text(platform.version(), style=value_color))
    table.add_row(Text('Machine:', style=label_color), Text(platform.machine(), style=value_color))
    table.add_row(Text('Processor:', style=label_color), Text(platform.processor(), style=value_color))
    table.add_row(Text('IP Address:', style=label_color), Text(socket.gethostbyname(socket.gethostname()), style=value_color))

    # Add Git repository details
    branch, commit_hash, commit_count = get_git_info()
    if branch and commit_hash and commit_count:
        table.add_row(Text('Git Branch:', style=label_color), Text(branch, style=value_color))
        table.add_row(Text('Git Commit Hash:', style=label_color), Text(commit_hash, style=value_color))
        table.add_row(Text('Git Commit Count:', style=label_color), Text(commit_count, style=value_color))

    logo = """
   ____  _  _   
  / ___|(_)| |_ 
 | |  _ | || __|
 | |_| || || |_ 
  \____||_| \__|
                
    """

    logo_text = Text(logo, style="bold yellow")
    logo_panel = Panel.fit(logo_text, border_style="yellow", padding=(0, 2))

    layout = Table.grid(expand=True)
    layout.add_column(justify="left", ratio=1)
    layout.add_column(justify="left", ratio=4)
    layout.add_row(logo_panel, table)

    console.print(Align(layout))
