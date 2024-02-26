---
title: "pairtree(1) user manual | version 1.0.3 5291ea3"
author: "R. S. Doiel"
pubDate: 2024-02-26
---

# NAME

pairtree

# SYNOPSIS

pairtree [OPTIONS] encode|decode STRING

# DESCRIPTION

pairtree will encode or decode a string to/from a pairtree path.

# OPTIONS

-h, -help
: display help

-version
: display version

-license
: display license

-s SEP
: set separator to SEP

# EXAMPLES

Encode key 12345

~~~shell
pairtree encode 12345
~~~

Decode path 12/34/5

~~~shell
pairtree decode 12/34/5
~~~

