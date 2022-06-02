DEPRECATED
==========

This Go package is now being maintained as part of [dataset](https://github.com/caltechlibrary/dataset).

Pairtree
========

This is a library for translate a UTF-8 string to/from a pairtree 
notation. This is typically used in storing things on disc (e.g. repository filesystems). This code is based on the specification found at https://confluence.ucop.edu/download/attachments/14254128/PairtreeSpec.pdf?version=2&modificationDate=1295552323000&api=v2 which is cited on the [OCFL](https://github.com/OCFL/spec/wiki) wiki.


Features
--------

- `Set()` will let you set the path separator
    - `Separator` is a readonly value of the file separator used by `Encode()` and `Decode()`
- `Encode()` will encode the provided string as a pairtree path
- `Decode()` will decode a pairtree path returning the unencoded string

Example
-------

```
    import (
        "fmt"
        "os"

        "github.com/caltechlibrary/pairtree"
    )

    func main() {
        key := "12mIEERD11"
        fmt.Printf("Key: %q\n", key)
        pairPath := pairtree.Encode(key)
        fmt.Printf("Endoded key %q -> %q\n", key, pairPath)
        key = Decode(pairPath)
        fmt.Printf("Decoded path %q -> %q\n", pairPath, key)
    }
```

