# BST
 A simple implementation of basic binary search methods:
 * insert key and its value
 * find by key
 * delete by key
## Create tree and insert nodes
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
## Find value by key
```
key := "5"
if value, ok :=tree.Find(key); ok {
  fmt.Printf("%s:%v\n", key, value)
}
```
## Delete value by key
```
key := "5"
tree.Delete(key)
```
## Using channel iterator for walking tree (in-order way)
```
for node := range tree.ChannelIterator(){
  fmt.Printf("%s:%v\n", node.key, node.value)
}
```
## Using stateful iterator for walking tree (in-order way)
```
for it := tree.StatefulIterator(); it.Next(); {
  fmt.Printf("%s:%v\n", it.Node().Key, it.Node().Value)
}
```
