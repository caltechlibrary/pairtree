---
title: "pairtree(1) user manual | version 1.0.4 9e3dac9"
author: "R. S. Doiel"
pubDate: 2025-07-29
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

