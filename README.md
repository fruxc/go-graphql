# Go Lang GraphQL Server

_👀 Easily create a new GoLang and GraphQL Server Using gqlgen_

## Run

```bash
Go run ./server.go
```

## Queries

### Mutation Query

```bash
mutation AddBike {
  createBike(input: {name:"BMW", isNewBike: true}){
    _id
    name
    isNewBike
  }
}
```

![Mutation](https://github.com/fruxc/go-graphql/blob/master/assets/mutation.png?raw=true "Mutation Query")

### FindOne Query

```bash
query findBike{
  bike(_id: "your_id") {
    name
    isNewBike
  }
}
```

![FindOne](https://github.com/fruxc/go-graphql/blob/master/assets/findOne.png?raw=true "FindOne Query")

### FindAll Query

```bash
query findAll{
  bikes{
    name
    isNewBike
  }
}
```

![FindAll](https://github.com/fruxc/go-graphql/blob/master/assets/findAll.png?raw=true "FindAll Query")
