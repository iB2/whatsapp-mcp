"""
Purge messages and chat rows for JIDs listed in store/ignore-chats.json.

Run AFTER stopping the bridge (it holds a lock on the DB).

Usage:
    python cleanup-ignored-chats.py              # dry-run (default)
    python cleanup-ignored-chats.py --apply      # actually delete
"""
import json
import os
import shutil
import sqlite3
import sys

STORE_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), "store")
DB_PATH = os.path.join(STORE_DIR, "messages.db")
FILTER_PATH = os.path.join(STORE_DIR, "ignore-chats.json")

def main():
    dry_run = "--apply" not in sys.argv

    if dry_run:
        print("=== DRY RUN (pass --apply to execute) ===\n")

    with open(FILTER_PATH, encoding="utf-8") as f:
        data = json.load(f)
    ignored = data.get("ignore_chats", [])
    comments = data.get("_comments", {})

    if not ignored:
        print("No chats to ignore.")
        return

    conn = sqlite3.connect(DB_PATH)
    cur = conn.cursor()

    total_msgs = 0
    total_media = 0

    for jid in ignored:
        label = comments.get(jid, jid)
        cur.execute("SELECT COUNT(*) FROM messages WHERE chat_jid = ?", (jid,))
        count = cur.fetchone()[0]
        total_msgs += count

        media_dir = os.path.join(STORE_DIR, jid.replace(":", "_"))
        media_files = 0
        if os.path.isdir(media_dir):
            media_files = len(os.listdir(media_dir))
            total_media += media_files

        print(f"  {jid}  ({label})")
        print(f"    messages: {count:,}   media files: {media_files}")

        if not dry_run:
            cur.execute("DELETE FROM messages WHERE chat_jid = ?", (jid,))
            cur.execute("DELETE FROM chats WHERE jid = ?", (jid,))
            cur.execute("DELETE FROM calls WHERE chat_jid = ?", (jid,))
            if os.path.isdir(media_dir):
                shutil.rmtree(media_dir)

    if not dry_run:
        conn.commit()
        cur.execute("VACUUM")
        print(f"\nDeleted {total_msgs:,} messages, {total_media} media files. DB vacuumed.")
    else:
        print(f"\nWould delete {total_msgs:,} messages, {total_media} media files.")

    conn.close()

if __name__ == "__main__":
    main()
