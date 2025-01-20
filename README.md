# BGTables

# BGTables is an agent for configuring linux network stack via [BGP](https://en.wikipedia.org/wiki/Border_Gateway_Protocol)

```
    +=========================+
    |          BGP server     |
    +=========================+
                 | <- BGP Protocol
    +=========================+
    |        BGTables         |
    +=========================+
                 | <- netlink/netfilter
    +=========================+
    |   linux network stack   |
    +=========================+
```
