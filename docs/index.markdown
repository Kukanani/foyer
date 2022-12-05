---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

# layout: default
---

# Welcome to the Foyer Documentation

[Foyer](https://github.com/Kukanani/foyer) is a simple terminal-based message board service where any user can post a message to the board and view messages others have posted. All users have full access to remove or modify messages. Therefore, `foyer` is meant for use on an internal server where all users are highly trusted. It is designed to be run when external users log into the system. In this way it acts as the "[foyer](https://en.wiktionary.org/wiki/foyer)" or entryway that users must pass through, seeing any critical messages before they proceed to use the system.

Foyer is inspired by [`debops.sysnews`](https://docs.debops.org/en/stable-2.1/ansible/roles/sysnews/index.html), but is designed to be much simpler and not linked to Debian or DebOps in any way.

Foyer is [MIT](https://choosealicense.com/licenses/mit/)-licensed, written in [Go](https://go.dev) and deployed using [GoReleaser](https://goreleaser.com).

## How Foyer Works

![A recording of using Foyer on a linux machine](foyer_example.gif)

Foyer stores messages as text files in `/opt/foyer`. The first line of each message is its title or "short text". Subsequent lines make up the message's full text. Users can add messages by creating text files in `/opt/foyer` directly. They can also use the subcommand `foyer add` to create a message using a simple terminal prompt.

Messages are kept in order and displayed based on their last modified date.

The `foyer` executable (without any subcommands) exits immediately if there are no messages to display. If there are any messages, `foyer` instead displays the messages and presents the user with a prompt. The prompt lets users view full details of any message(s), delete messages, or continue on past the foyer.
