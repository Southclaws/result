**Note: this was more of an experiment, for an actual real-world library that's well tested and used in production codebases, [check out opt](https://github.com/Southclaws/result).**

---

# Result Type for Go

A super simple generic result type for Go.

## Why?

The idiomatic ternary operation in Go is to declare a zero-value mutable variable and mutate it behind an if-else branch.

This is readable and simple, though it can make some code quite terse and overly explicit.

Take this for example:

```go
var u User
var err error

if someCondition {
    u, err = db.GetUserByID(id)
} else {
    u, err = db.GetUserByEmail(email)
}

if err != nil {
    return nil, err
}
```

It's quite explicit, compared to the equivalents in Rust and TypeScript:

```rust
let u = if some_condition {
    db.get_user_by_id(id)
} else {
    db.get_user_by_email(email)
};
```

```typescript
const u = await someCondition
    ? db.getUserByID(id)
    : db.getUserByEmail(email);
```

And Go also has no immutable types, which isn't a huge issue but sometimes code reads simpler if you can assume things are assigned once.

So, this library was written partly as an experiment to see how ergonomic it can be in a notoriously un-ergonomic language.

## Usage

### Wrapping Errors

See the package documentation for full details.

Given some function:

```go
func (u *Users) GetByID(id string) (*User, error)
```

```go
r := result.Wrap(u.GetByID(id))
```

This constructs a result type from the return value of the API function, this result type contains either the value or the error as a single entity.

```go
r.Valid() // if err == nil
r.Value() // return the user if r.Valid() == true
r.Error() // return the error if r.Valid() == false
```

Now this is fairly similar to what is idiomatic, there's an error check and a valid state. But, there are two other invalid states...

One issue with Go's idiomatic approach to return values and error handling is that there's no ergonomic way to express a discriminated union via the type system.

This means that, given the following function:

```go
func F() (T, error)
```

There are not *two* possible return states encoded in the type system, but **four!**

- T == nil, error == nil
- T != nil, error == nil
- T == nil, error != nil
- T != nil, error != nil

Of course, in practice this isn't a huge issue as most programmers can make the assumption that, given `err == nil` then `T` *must* be valid. And given `err != nil` then `T` *must* be undefined/invalid.

But type systems should be there to aid us in not needing to rely on *assumptions* to write safe predictable code. Result types in strongly typed languages (often functional languages) encode the possibility of *only* two states.

You're probably thinking "that's a great point, but this library doesn't solve that problem at all" and you'd be right, you can still end up with multiple invalid states with this result type but that's just because Go's generics are still very basic and cannot express a proper result type yet. But I hope it will do in the future!

One great example of a discriminated union result type is in TypeScript, because you can encode literal values as types and combine that with type unions in order to create types that can *only* be one value or the other, not any other combination of states such as the 4 states in the list above.

```typescript
type Result<T, E = Error> = {
    value: T;
    error: undefined;
} | {
    value: undefined;
    error: E;
}
```

This type makes it impossible to construct an instance where both `value` and `error` are defined. You either have one or the other. And the type system enforces that. Building code with such a concrete and fundamental concept baked into the core APIs leads to viral strictness to spread throughout the code (for better or worse.)

### Ternary Operations

So, going back to the beginning of this readme, the end goal is to facilitate an ergonomic ternary switch operation purely for performing conditional assignment once.

With the same example of the two example user APIs: `GetUserByID` and `GetUserByEmail` we can rewrite the initial ternary operation as:

```go
r := result.TernaryResult(
    someCondition,
    result.Wrap(db.GetUserByID(id)),
    result.Wrap(db.GetUserByEmail(email)),
)

if !r.Valid() {
    return nil, r.Error()
}
```

In my opinion (which is a dangerous thing to say while writing Go libraries!) this is cleaner and easier to read for certain people.

The downside here is that it's not as *simple* and *explicit* as the initial version, which is one of Go's benefits for onboarding newcomers to the language. Code is generally pretty easy to read because there aren't many concepts compared to other languages.

It also doesn't really solve the possible states problem outlined above. There's nothing in the type system that can prevent `r.Value()` and `r.Error()` both returning `nil`. This is simply not possible (as far as I can figure out) with the current implementation of Go 1.18 generics.

## Is this idiomatic?

Short answer: no.

Long answer: not sure yet. The idioms and common practices with generics are still evolving. I've been using optional types a ton at work and in my own code, those are a real life saver and a welcomed alternative to ~~hacking~~ using pointers to encode optionality into data structures.

I don't imagine this library wil catch on, but I do hope to see something similar used more widely in the near future. I know Go library authors are stuck with `(T, error)` for good and a combination of stubbornness, conforming to standards and keeping things simple will prevent these sorts of experimental progressive ideas. And that's probably a good thing tbh.
