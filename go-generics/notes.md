//TODO: setup github action for building Go!

# Generics in Go

## Type Parameters

Type parameters enable writing generic code
functions and type declarations now accetp type parametes

func name(p int) int{ ..}
func name[T any] (p T) int {...}
func name[T any] (p []T) int {}

any includes all types like empty interface {}
T is a common used name! nothing more nothing less

```go
var name  string = "bla bla"
func (s *Solar) Print() string{
    return fmt.Sprintf("5s - %v\n", name, *s)
}
func (w *Wind) Print() string{
    return fmt.Sprintf("5s - %v\n", name, *w)
}

func PrintGeneric[T any] (t T) string {
     return fmt.Sprintf("%s - %v\n", name, t)
}
```

calling generic functions
i := 43;
j := name(i)
k := name[int](i) // passing the type between square brackets

## Type Sets

- A type set includes types defined in an interface!

types and interfaces are connected throught the change of specifications!

- Interfaces before Generics: Method Sets => defined a method set of its functions

```go
type Output interface {
    Area() int
}
```

And the type implement the interface if the type contains the method defined in the interface!

> with generics, we stop looking into the method sets and look at types.!
So we build a type set instead of method set!

A type set includes types defined in an interface
A type set can include methods and types => this wasn't possible with method sets!

What can be in a type set ?

- Empty interface: all non-interface types
- non-empty interface: intersection of all elements:
  - method, for example, Read
  - non interface type, for example, int

with type sets we can restrict our generics to certain types!
for example our demo which uses Energy types {Solar, Wind}, we can restrict PrintGenerics to print only those two!

### Package constraints

package contains six common constraints [pre-declared types] for type parameters

- Complex type

```go
type Complex interface {
    ~float32 | ~float64
}
```

the `~` allows for subtypes !

- Float type
- Signed type [int types]
- Unsigned type
- Integer type [combines Signed and UnSigned ]

```go
type Integer interface {
    Signed | UnSigned
}
```

- Orderd type:=> for any type that support orders operators (<<=>=>)

```go
type Orderd interface {
    Integer | Float | ~string
}
```

## Type Inference

### Function argument type inference

There're some steps first  for the compiler to infere the type

1. `first pass`: ignore untypes constants
2. `second pass`: Unify restuls with untyped constants
3. `Inference`: Check if inference if possible or not .
If we tried to remove the type from

```go
busniess.PrintGeneric[business.Solar](solar2k)
```

to

```go
busniess.PrintGeneric(solar2k)
```

the compiler will infer the type!

### Constraint type inference

## Generics Best Practices

## Generics in Practice

Use generics to avoid boilerplate code
Narrow down type parameters
For which types would I write the function without generics

Use cases:

1. Data structures (LinkedLists, Stacks, Queues, etc);
2. Type-independents
