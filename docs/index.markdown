---
# Feel free to add content and custom Front Matter to this file.
# To modify the layout, see https://jekyllrb.com/docs/themes/#overriding-theme-defaults

# layout: default
---

# Foyer: Simple Terminal Message Board

[Foyer](https://github.com/Kukanani/foyer) is a simple terminal-based message board service where any user can post a message to the message board and view messages others have posted. It is meant for use on an internal server where users are highly trusted (since all users have full access to remove or modify messages). Foyer is designed to be run when external users log into the system. In this way it acts as the "foyer" or entryway that users must pass through, seeing critical messages before they proceed to use the system.

Foyer is inspired by [`debops.sysnews`](https://docs.debops.org/en/stable-2.1/ansible/roles/sysnews/index.html), but is designed to be much simpler and not linked to Debian or DebOps in any way.

Foyer is written in Go and deployed using [GoReleaser](https://goreleaser.com).
