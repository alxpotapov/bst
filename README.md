# BST
## Usage
```
go get github.com/alxpotapov/bst
```
## A simple implementation of basic binary search methods:
 * insert key and its value
 * find by key
 * delete by key
### Create tree and insert nodes
```
tree := &Tree{}

tree.Insert("2", 2)
tree.Insert("1", 1)
tree.Insert("7", 7)
tree.Insert("6", 6)
tree.Insert("8", 8)
tree.Insert("3", 3)
tree.Insert("3", 33)
```
### Find value by key
```
key := "5"
if value, ok :=tree.Find(key); ok {
  fmt.Printf("%s:%v\n", key, value)
}
```
### Delete value by key
```
key := "5"
tree.Delete(key)
```
## 3 iteration patterns in BST:
 * Callback iterator
 * Stateful iterator
 * Channel iterator
### Iterating via callback
```
tree.ForEach(func(node *Node){
  fmt.Printf("%s:%v\n", node.key, node.value)
})
```
### Iterating with Next()
```
for it := tree.StatefulIterator(); it.Next(); {
  fmt.Printf("%s:%v\n", it.Node().Key, it.Node().Value)
}
```
### Iterating with a channel
```
for node := range tree.ChannelIterator(){
  fmt.Printf("%s:%v\n", node.key, node.value)
}
```
