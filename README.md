# Go Infra Playground

This workspace is a small, test-driven practice ground for building confidence
with core Go features that appear frequently in infrastructure codebases such as
Kubernetes and etcd.

## Why This Exists

If AI has been filling in too much code for you, the fix is not "stop using AI".
The fix is to create a place where you can repeatedly write the parts yourself,
get fast feedback, and learn to explain the code in your own words.

This playground is organized around that idea.

## How To Use It

1. Open one exercise package at a time.
2. Read the package comment and the `TODO` markers.
3. Implement the missing pieces in `exercise.go`.
4. Run the tests for only that package first.
5. When it passes, explain to yourself:
   - Why does this API use a slice and not an array?
   - Why is this channel buffered or unbuffered?
   - What does the type parameter mean?
   - What memory is shared and what memory is copied?

Example commands:

```bash
go test ./01_slices_copy -v
go test ./02_channels -v
go test ./03_generics -v
go test ./04_queue -v
go test ./05_context_loop -v
go test ./06_worker_pool -v
go test ./07_json_codec -v
go test ./08_retry_backoff -v
go test ./09_singleflight -v
go test ./10_coalescer -v
go test ./11_select_patterns -v
go test ./12_interfaces -v
go test ./13_options -v
go test ./14_slice_semantics -v
go test ./15_slice_ownership -v
go test ./16_array_vs_slice -v
go test ./17_buffer_reuse -v
go test ./18_semaphore -v
go test ./19_errgroup_lite -v
go test ./20_conflict_retry -v
go test ./21_workqueue -v
go test ./22_fake_clock_retry -v
go test ./23_map_ownership -v
go test ./24_subslice_leak -v
go test ./25_custom_json -v
go test ./26_context_tree -v
go test ./27_pipeline_shutdown -v
```

## Recommended Order

1. `01_slices_copy`
2. `02_channels`
3. `03_generics`
4. `04_queue`
5. `05_context_loop`
6. `06_worker_pool`
7. `07_json_codec`
8. `08_retry_backoff`
9. `09_singleflight`
10. `10_coalescer`
11. `11_select_patterns`
12. `12_interfaces`
13. `13_options`
14. `14_slice_semantics`
15. `15_slice_ownership`
16. `16_array_vs_slice`
17. `17_buffer_reuse`
18. `18_semaphore`
19. `19_errgroup_lite`
20. `20_conflict_retry`
21. `21_workqueue`
22. `22_fake_clock_retry`
23. `23_map_ownership`
24. `24_subslice_leak`
25. `25_custom_json`
26. `26_context_tree`
27. `27_pipeline_shutdown`

`00_notes` contains short concept notes that you can update in your own words.

## Practical Patterns

These later exercises are intentionally closer to the code shapes that show up
in Kubernetes, controller loops, and config handling:

- `05_context_loop`: context cancellation, stop signals, and periodic loops.
- `06_worker_pool`: bounded parallelism, first-error return, ordered results.
- `07_json_codec`: strict JSON decoding for config and API payloads.
- `08_retry_backoff`: bounded retries, exponential backoff, cancellation.
- `09_singleflight`: duplicate suppression for concurrent loads.
- `10_coalescer`: event coalescing to avoid signal storms.
- `11_select_patterns`: `select`, pipeline stop signals, fan-in, leak avoidance.
- `12_interfaces`: small interfaces and function adapters like `HandlerFunc`.
- `13_options`: functional options and constructor-style API design.
- `14_slice_semantics`: slice passing, in-place mutation, append, clone, mixed struct fields.
- `15_slice_ownership`: defensive copy on input/output and internal slice ownership.
- `16_array_vs_slice`: direct comparison of array value semantics and slice sharing.
- `17_buffer_reuse`: reusable buffers, returned bytes, and accidental aliasing.
- `18_semaphore`: bounded concurrency with acquire/release and cancellation.
- `19_errgroup_lite`: sibling goroutines, first-error cancellation, coordinated wait.
- `20_conflict_retry`: optimistic concurrency retry on conflict only.
- `21_workqueue`: deduplicated key queue, done/reenqueue, shutdown behavior.
- `22_fake_clock_retry`: injected time dependency for deterministic retry tests.
- `23_map_ownership`: defensive copy for maps on both input and output.
- `24_subslice_leak`: detached small windows to avoid aliasing and large-array retention.
- `25_custom_json`: enum and duration custom JSON encoding/decoding.
- `26_context_tree`: parent-child cancellation and timeout scoping.
- `27_pipeline_shutdown`: multi-stage pipeline shutdown and channel closing discipline.

## Rules For Yourself

- Do not ask AI for the final code immediately.
- First write your own explanation in comments or on paper.
- If stuck, ask for a hint, not a solution.
- After passing tests, refactor once.
- For every function you write, ask:
  - What are the zero values?
  - Who owns the memory?
  - Is this API exposing too much mutability?
  - What happens under contention?

## Stretch Goal

When these feel comfortable, reimplement one small etcd or k8s helper from
memory and compare your version with the upstream code.
