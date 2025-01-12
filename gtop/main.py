"""Git repository and system information display tool with commit heatmap visualization."""

import os
import re
import platform
import socket
import subprocess
from datetime import datetime, timedelta

import distro
import psutil
from screeninfo import get_monitors
from rich import box
from rich.console import Console
from rich.table import Table
from rich.panel import Panel
from rich.text import Text
from rich.align import Align
import pandas as pd

def get_git_info():
    """Retrieve git repository info including branch, commit details, user info, and more."""
    try:
        branch = subprocess.check_output(['git', 'rev-parse', '--abbrev-ref', 'HEAD']).strip().decode('utf-8')
        commit_hash = subprocess.check_output(['git', 'rev-parse', 'HEAD']).strip().decode('utf-8')
        commit_count = subprocess.check_output(['git', 'rev-list', '--count', 'HEAD']).strip().decode('utf-8')
        user_name = subprocess.check_output(['git', 'config', 'user.name']).strip().decode('utf-8')
        user_email = subprocess.check_output(['git', 'config', 'user.email']).strip().decode('utf-8')
        remote_url = subprocess.check_output(['git', 'config', '--get', 'remote.origin.url']).strip().decode('utf-8')
        last_commit_message = subprocess.check_output(['git', 'log', '-1', '--pretty=%B']).strip().decode('utf-8')
        status = subprocess.check_output(['git', 'status', '--short']).strip().decode('utf-8')
        return branch, commit_hash, commit_count, user_name, user_email, remote_url, last_commit_message, status
    except subprocess.CalledProcessError:
        return None, None, None, None, None, None, None, None

def get_commit_dates():
    """Get list of commit dates from git log for the current branch."""
    try:
        # Get commit dates for the current branch
        commit_dates = subprocess.check_output(
            ['git', 'log', '--pretty=format:%cd', '--date=short']
        ).decode('utf-8').split('\n')
        return commit_dates
    except subprocess.CalledProcessError:
        return []

def plot_commit_heatmap(dates):
    """Generate and display a heatmap visualization of git commits over time."""
    # Convert dates to a DataFrame
    df = pd.DataFrame(dates, columns=['date'])
    df['date'] = pd.to_datetime(df['date'])
    df['count'] = 1

    # Group by date and count commits
    df = df.groupby('date').count().reset_index()

    # Create a date range for the last year
    today = datetime.today()
    one_year_ago = today - timedelta(days=365)
    all_dates = pd.date_range(start=one_year_ago, end=today)

    # Reindex the DataFrame to include all dates
    df = df.set_index('date').reindex(all_dates, fill_value=0).reset_index()
    df.columns = ['date', 'count']  # Ensure columns are correctly named

    # Determine the maximum number of commits in a single day for scaling
    max_commits = df['count'].max()

    # Create a table for the heatmap
    console = Console()
    table = Table(show_header=True, header_style="bold magenta")
    table.add_column("Weekday", justify="right")
    for month in pd.date_range(start=one_year_ago, end=today, freq='ME').strftime('%b'):
        table.add_column(month, justify="center")

    # Fill the table with commit data
    weekdays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
    for weekday in weekdays:
        row = [weekday]
        for month in pd.date_range(start=one_year_ago, end=today, freq='ME'):
            month_data = df[(df['date'].dt.month == month.month) & (df['date'].dt.day_name() == weekday)]
            commit_count = month_data['count'].sum() if not month_data.empty else 0

            # Calculate the intensity of the green color based on the commit count
            if commit_count > 0:
                intensity = int((commit_count / max_commits) * 255)
                color = f"rgb(0,{intensity},0)"
            else:
                color = "white"
            cell_text = Text(" ", style=f"on {color}")

            row.append(cell_text)
        table.add_row(*row)

    console.print(table)

def main():
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
    branch, commit_hash, commit_count, user_name, user_email, remote_url, last_commit_message, status = get_git_info()
    if branch and commit_hash and commit_count:
        table.add_row(Text('Git Branch:', style=label_color), Text(branch, style=value_color))
        table.add_row(Text('Git Commit Hash:', style=label_color), Text(commit_hash, style=value_color))
        table.add_row(Text('Git Commit Count:', style=label_color), Text(commit_count, style=value_color))
        table.add_row(Text('Git User Name:', style=label_color), Text(user_name, style=value_color))
        table.add_row(Text('Git User Email:', style=label_color), Text(user_email, style=value_color))
        table.add_row(Text('Git Remote URL:', style=label_color), Text(remote_url, style=value_color))
        table.add_row(Text('Last Commit Message:', style=label_color), Text(last_commit_message, style=value_color))
        table.add_row(Text('Git Status:', style=label_color), Text(status, style=value_color))

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

    commit_dates = get_commit_dates()
    plot_commit_heatmap(commit_dates)

if __name__ == "__main__":
    main()
