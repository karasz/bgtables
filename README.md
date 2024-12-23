# GoBGPNet

# BGTables is an agent for configuring linux network stack via [GoBGP](https://github.com/osrg/gobgp)

```
    +=========================+
    |          GoBGP          |
    +=========================+
                 | <- gRPC API
    +=========================+
    |        BGTables         |
    +=========================+
                 | <- netlink/netfilter
    +=========================+
    |   linux network stack   |
    +=========================+
```
