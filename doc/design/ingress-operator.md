# Ingress Operator Design

First, the Ingress Operator purchases capacity from the network.

```mermaid
flowchart LR
  IO(Ingress Operator) -- Purchases Capacity --> SC
  IP1(Infrastructure Provider 1) -- Registers Capacity --> SC
  IP2(Infrastructure Provider 2) -- Registers Capacity --> SC
  SC(Smart Contract)
```

Then, an Ingress Operator is deployed to Kubernetes via InfraKit.
```mermaid
flowchart LR
  subgraph KUB [Kubernetes]
    subgraph IO [Ingress Operator]
      direction TB
      Proxy(Proxy)
      Proxy -- Signs requests to send to providers <--> EP
      EP(Ethereum Private Key)
    end
    IK(InfraKit) -- Provisions --> IO
  end
  User -- API Requests --> IO
  IO -- Relayed requests --> IP1
  IO -- Relayed requests --> IP2
  IP1(Infrastructure Provider 1)
  IP2(Infrastructure Provider 2)
```

The Ingress Operator recieves API requests from a user, signs them with its private key
to prove entitlement, and relays them to the Infrastructure Providers it has purchased
capacity from.