# Self-Check Notes

Write your own answers under each question before looking anything up.

## Slices

- What is the difference between an array and a slice?
- What does a slice header contain?
- When does `append` reuse the old backing array?
- When is `copy(dst, src)` safer than direct assignment?
- What is the difference between `nil` slice and empty slice?

## Channels

- What is the difference between `chan T`, `<-chan T`, and `chan<- T`?
- When should you prefer an unbuffered channel?
- What does channel capacity actually control?
- What does closing a channel mean?
- Who should close a channel?

## Generics

- Is `T any` a "generic type"? If yes, where is the type parameter declared?
- Why is `any` just an alias for `interface{}`?
- When should you use a type parameter instead of `interface{}`?
- What is a constraint?
- Why do many infra packages still avoid generics in some APIs?

## Infra Design

- Can callers mutate internal memory through the API you expose?
- If data is shared across goroutines, what synchronizes access?
- Are zero values useful and safe?
- Is the API easy to misuse?
- What tests would catch a race, aliasing bug, or ordering bug?
