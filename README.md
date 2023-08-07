# GoStack

```go
s := gostack.New[Type]()
```

---

## Usage

To create a new stack, use the New[Type] function
```go
s := gostack.New[int]()
```

To push an element onto the stack, use the push method
```go
s.Push(1)
s.Push(2)
s.Push(3)
```

To pop an element from the stack, use the pop method
```go
elem, err := s.Pop()
if err != nil {
    // handle error
}
// use elem
```

To peek at the top element of the stack without popping it, use the peek method
```go
elem, err := s.Peek()
if err != nil {
    // handle error
}
// use elem
```

To get the length of the stack, use the len method
```go
len := s.Len()
```

To clear the stack, use the clear method
```go
s.Clear()
```

## Example

check it out [here](example/ex.go)
